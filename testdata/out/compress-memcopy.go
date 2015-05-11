package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _in_a_test_asset = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x57\x28\x4e\xcc\x2d\xc8\x49\x55\x48\xcb\xcc\x49\xe5\x02\x04\x00\x00\xff\xff\x8a\x82\x8c\x85\x0f\x00\x00\x00")

func in_a_test_asset_bytes() ([]byte, error) {
	return bindataRead(
		_in_a_test_asset,
		"in/a/test.asset",
	)
}

func in_a_test_asset() (*asset, error) {
	bytes, err := in_a_test_asset_bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "in/a/test.asset", size: 15, mode: os.FileMode(420), modTime: time.Unix(1431385279, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _in_b_test_asset = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x57\x28\x4e\xcc\x2d\xc8\x49\x55\x48\xcb\xcc\x49\xe5\x02\x04\x00\x00\xff\xff\x8a\x82\x8c\x85\x0f\x00\x00\x00")

func in_b_test_asset_bytes() ([]byte, error) {
	return bindataRead(
		_in_b_test_asset,
		"in/b/test.asset",
	)
}

func in_b_test_asset() (*asset, error) {
	bytes, err := in_b_test_asset_bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "in/b/test.asset", size: 15, mode: os.FileMode(420), modTime: time.Unix(1430781941, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _in_c_test_asset = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x57\x28\x4e\xcc\x2d\xc8\x49\x55\x48\xcb\xcc\x49\xe5\x02\x04\x00\x00\xff\xff\x8a\x82\x8c\x85\x0f\x00\x00\x00")

func in_c_test_asset_bytes() ([]byte, error) {
	return bindataRead(
		_in_c_test_asset,
		"in/c/test.asset",
	)
}

func in_c_test_asset() (*asset, error) {
	bytes, err := in_c_test_asset_bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "in/c/test.asset", size: 15, mode: os.FileMode(420), modTime: time.Unix(1430781941, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _in_test_asset = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x57\x28\x4e\xcc\x2d\xc8\x49\x55\x48\xcb\xcc\x49\xe5\x02\x04\x00\x00\xff\xff\x8a\x82\x8c\x85\x0f\x00\x00\x00")

func in_test_asset_bytes() ([]byte, error) {
	return bindataRead(
		_in_test_asset,
		"in/test.asset",
	)
}

func in_test_asset() (*asset, error) {
	bytes, err := in_test_asset_bytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "in/test.asset", size: 15, mode: os.FileMode(420), modTime: time.Unix(1430781941, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"in/a/test.asset": in_a_test_asset,
	"in/b/test.asset": in_b_test_asset,
	"in/c/test.asset": in_c_test_asset,
	"in/test.asset": in_test_asset,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"in": &bintree{nil, map[string]*bintree{
		"a": &bintree{nil, map[string]*bintree{
			"test.asset": &bintree{in_a_test_asset, map[string]*bintree{
			}},
		}},
		"b": &bintree{nil, map[string]*bintree{
			"test.asset": &bintree{in_b_test_asset, map[string]*bintree{
			}},
		}},
		"c": &bintree{nil, map[string]*bintree{
			"test.asset": &bintree{in_c_test_asset, map[string]*bintree{
			}},
		}},
		"test.asset": &bintree{in_test_asset, map[string]*bintree{
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        // File
        if err != nil {
                return RestoreAsset(dir, name)
        }
        // Dir
        for _, child := range children {
                err = RestoreAssets(dir, path.Join(name, child))
                if err != nil {
                        return err
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

