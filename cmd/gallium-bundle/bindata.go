// Code generated by go-bindata.
// sources:
// info.plist.tpl
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	name    string
	size    int64
	mode    os.FileMode
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

var _infoPlistTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x93\x5f\x4f\xdb\x30\x14\xc5\x9f\xc9\xa7\xc8\xf2\x5e\x3b\x66\x20\xb6\xc9\x04\xd1\x84\xa2\x6a\xfd\x13\x2d\x61\x82\xa7\xc9\x73\x4c\x7b\x35\xc7\xb6\x6c\x87\x82\x10\xdf\x7d\x4a\x5b\xba\x2d\x69\x1e\xf6\x16\xeb\x9e\xdf\xc9\xf5\xbd\x3e\xf4\xea\xb9\x96\xe1\x93\xb0\x0e\xb4\xba\x8c\x08\x8a\xa3\x50\x28\xae\x2b\x50\xab\xcb\xe8\xae\x9c\x8c\x3e\x45\x57\x49\x40\x3f\x64\xcb\xb4\x7c\xc8\x6f\x42\x23\xc1\xf9\x30\xbf\x1b\xcf\xa6\x69\x18\x8d\x30\xbe\x36\x46\x0a\x8c\xb3\x32\x0b\xf3\xd9\xb4\x28\x43\x82\x62\x8c\x6f\x16\x51\x18\xad\xbd\x37\x5f\x30\xde\x6c\x36\x88\xb5\x2a\xc4\x75\xdd\x0a\x1d\xce\xad\x36\xc2\xfa\x97\x19\x38\x3f\x22\x28\x46\x95\xaf\xa2\x24\xa0\x3b\xf7\x7f\xda\x49\x02\x5a\x01\xf7\x49\x70\x42\x7f\x89\x97\x64\xdc\x80\xac\xe6\x8c\xaf\x41\x89\x65\xb1\x3d\x51\xdc\x16\x82\x13\xea\xbc\x05\xb5\x4a\xc8\xd9\x84\x9c\xc7\x9f\x29\xde\x9f\xf7\x64\x3a\x19\x37\xaa\x92\x62\x5a\x09\xe5\xe1\x11\x84\xed\x82\xaf\xaf\xa8\x2b\x79\x7b\x1b\xb2\x59\xb0\x5a\x0c\x1a\xb4\xc5\x61\xb4\x68\x8c\xd1\xd6\x8b\x2a\x97\xcc\x3f\x6a\x5b\xbb\x83\x11\xb3\x96\xb5\x1f\x07\xcb\x39\xe3\xcb\xe2\xfe\x6f\x27\xfc\xae\xd9\x7a\x66\x65\xaa\x6b\x03\xb2\x7f\x1b\xae\xeb\x3f\x63\xdf\x2a\x1c\x92\xf2\xa9\x46\x5c\x32\xb5\x42\xe4\x47\xdc\xed\x2f\x2b\xdf\x1b\x3a\x3a\xd7\x8b\x94\xc4\xf1\xe9\x30\xf4\x7d\xb7\xb6\x2e\x76\x3b\xef\x23\x45\xf6\xf5\xf8\xea\xce\xd3\xb3\x8f\x47\xe5\xc7\xc6\x5d\x33\xae\xdd\x33\x89\x11\x21\x7d\xe6\x9e\xeb\xaa\x47\xc4\x17\xa7\x43\xd2\xff\xb8\xf2\xa2\x98\x33\x50\x0b\xf8\x39\x01\xd9\xfb\x45\x5b\x9a\x0b\xd5\xf4\xa1\xdc\x82\xe2\x60\x98\x4c\x25\x73\xae\xcb\x8d\xbf\x3d\xb4\x51\x02\xce\xfc\x76\x88\x5d\x7a\xff\x68\xdc\x75\xe3\x75\xcd\x3c\xf0\x5b\xcb\xcc\x1a\xb8\x2b\x36\xe0\xdb\x38\xac\x0e\x8e\xde\x36\x02\x27\x01\xc5\xbb\xe0\x50\xbc\x8d\x55\x12\xfc\x0e\x00\x00\xff\xff\xce\xbb\x95\x4c\xed\x03\x00\x00")

func infoPlistTplBytes() ([]byte, error) {
	return bindataRead(
		_infoPlistTpl,
		"info.plist.tpl",
	)
}

func infoPlistTpl() (*asset, error) {
	bytes, err := infoPlistTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "info.plist.tpl", size: 1005, mode: os.FileMode(420), modTime: time.Unix(1473867136, 0)}
	a := &asset{bytes: bytes, info: info}
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
	if err != nil {
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
	"info.plist.tpl": infoPlistTpl,
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
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"info.plist.tpl": &bintree{infoPlistTpl, map[string]*bintree{}},
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
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
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
		err = RestoreAssets(dir, filepath.Join(name, child))
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

