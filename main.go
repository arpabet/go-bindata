// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package main

import (
	"os"
	"flag"
	"fmt"
	"path"
	"strings"
	"runtime"
)

const (
	APP_NAME    = "bindata"
	APP_VERSION = "0.2"
)

func main() {
	in := flag.String("i", "", "Path to the input file.")
	out := flag.String("o", "", "Optional path to the output file.")
	pkgname := flag.String("p", "", "Optional name of the package to generate.")
	funcname := flag.String("f", "", "Optional name of the function to generate.")
	version := flag.Bool("v", false, "Display version information.")

	flag.Parse()

	if *version {
		fmt.Fprintf(os.Stdout, "%s v%s (Go runtime %s)\n",
			APP_NAME, APP_VERSION, runtime.Version())
		return
	}

	if len(*in) == 0 {
		fmt.Fprintln(os.Stderr, "[e] No input file specified.")
		os.Exit(1)
	}

	if len(*out) == 0 {
		// Ensure we create our own output filename that does not already exist.
		dir, file := path.Split(*in)

		*out = path.Join(dir, file) + ".go"
		if _, err := os.Lstat(*out); err == nil {
			// File already exists. Pad name with a sequential number until we
			// find a name that is available.
			count := 0
			for {
				f := path.Join(dir, fmt.Sprintf("%s.%d.go", file, count))
				if _, err := os.Lstat(f); err != nil {
					*out = f
					break
				}

				count++
			}
		}

		fmt.Fprintf(os.Stderr, "[w] No output file specified. Using '%s'.\n", *out)
	}

	if len(*pkgname) == 0 {
		fmt.Fprintln(os.Stderr, "[w] No package name specified. Using 'main'.")
		*pkgname = "main"
	}

	if len(*funcname) == 0 {
		_, file := path.Split(*in)
		file = strings.ToLower(file)
		file = strings.Replace(file, " ", "_", -1)
		file = strings.Replace(file, ".", "_", -1)
		file = strings.Replace(file, "-", "_", -1)
		fmt.Fprintf(os.Stderr, "[w] No function name specified. Using '%s'.\n", file)
		*funcname = file
	}

	// Read the input file, transform it into a gzip compressed data stream and
	// write it out as a go source file.
	if err := translate(*in, *out, *pkgname, *funcname); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %s\n", err)
		return
	}

	// If gofmt exists on the system, use it to format the generated source file.
	if err := gofmt(*out); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %s\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, "[i] Done.")
}
