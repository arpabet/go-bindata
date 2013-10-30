// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package bindata

import (
	"fmt"
	"os"
	"path/filepath"
)

// Config defines a set of options for the asset conversion.
type Config struct {
	// Name of the package to use. Defaults to 'main'.
	Package string

	// Tags specify a set of optional build tags, which should be
	// included in the generated output. The tags are appended to a
	// `// +build` line in the beginning of the output file
	// and must follow the build tags syntax specified by the go tool.
	Tags string

	// Input defines the directory path, containing all asset files.
	// This may contain sub directories, which will be included in the
	// conversion.
	Input string

	// Output defines the output file for the generated code.
	// If left empty, this defaults to 'bindata.go' in the current
	// working directory.
	Output string

	/*
		Prefix defines a path prefix which should be stripped from all
		file names when generating the keys in the table of contents.
		For example, running without the `-prefix` flag, we get:

			$ go-bindata /path/to/templates
			go_bindata["/path/to/templates/foo.html"] = _path_to_templates_foo_html

		Running with the `-prefix` flag, we get:

			$ go-bindata -prefix "/path/to/" /path/to/templates/foo.html
			go_bindata["templates/foo.html"] = templates_foo_html
	*/
	Prefix string

	/*
		NoMemCopy will alter the way the output file is generated.

		It will employ a hack that allows us to read the file data directly from
		the compiled program's `.rodata` section. This ensures that when we call
		call our generated function, we omit unnecessary mem copies.

		The downside of this, is that it requires dependencies on the `reflect` and
		`unsafe` packages. These may be restricted on platforms like AppEngine and
		thus prevent you from using this mode.

		Another disadvantage is that the byte slice we create, is strictly read-only.
		For most use-cases this is not a problem, but if you ever try to alter the
		returned byte slice, a runtime panic is thrown. Use this mode only on target
		platforms where memory constraints are an issue.

		The default behaviour is to use the old code generation method. This
		prevents the two previously mentioned issues, but will employ at least one
		extra memcopy and thus increase memory requirements.

		For instance, consider the following two examples:

		This would be the default mode, using an extra memcopy but gives a safe
		implementation without dependencies on `reflect` and `unsafe`:

		func myfile() []byte {
			return []byte{0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a}
		}

		Here is the same functionality, but uses the `.rodata` hack.
		The byte slice returned from this example can not be written to without
		generating a runtime error.

		var _myfile = "\x89\x50\x4e\x47\x0d\x0a\x1a"

		func myfile() []byte {
			var empty [0]byte
			sx := (*reflect.StringHeader)(unsafe.Pointer(&_myfile))
			b := empty[:]
			bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
			bx.Data = sx.Data
			bx.Len = len(_myfile)
			bx.Cap = bx.Len
			return b
		}
	*/
	NoMemCopy bool

	/*
	   NoCompress means the assets are /not/ GZIP compressed before being turned
	   into Go code. The generated function will automatically unzip
	   the file data when called. Defaults to true.
	*/
	NoCompress bool

	// Perform a debug build. This generates an asset file, which
	// loads the asset contents directly from disk at their original
	// location, instead of embedding the contents in the code.
	//
	// This is mostly useful if you anticipate that the assets are
	// going to change during your development cycle. You will always
	// want your code to access the latest version of the asset.
	// Only in release mode, will the assets actually be embedded
	// in the code. The default behaviour is Release mode.
	Debug bool

	// Recursively process all assets in the input directory and its
	// sub directories. This defaults to false, so only files in the
	// input directory itself are read.
	Recursive bool
}

// NewConfig returns a default configuration struct.
func NewConfig() *Config {
	c := new(Config)
	c.Package = "main"
	c.NoMemCopy = false
	c.NoCompress = false
	c.Debug = false
	c.Recursive = false
	return c
}

// validate ensures the config has sane values.
// Part of which means checking if certain file/directory paths exist.
func (c *Config) validate() error {
	if len(c.Package) == 0 {
		return fmt.Errorf("Missing package name")
	}

	stat, err := os.Lstat(c.Input)
	if err != nil {
		return fmt.Errorf("Input path: %v", err)
	}

	if !stat.IsDir() {
		return fmt.Errorf("Input path is not a directory.")
	}

	if len(c.Output) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("Unable to determine current working directory.")
		}

		c.Output = filepath.Join(cwd, "bindata.go")
	}

	stat, err = os.Lstat(c.Output)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("Output path: %v", err)
	}

	if stat != nil && stat.IsDir() {
		return fmt.Errorf("Output path is a directory.")
	}

	return nil
}
