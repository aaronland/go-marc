// Code generated by go-bindata.
// sources:
// templates/html/marc-034.html
// DO NOT EDIT!

package html

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

var _templatesHtmlMarc034Html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x4d\x6f\xdb\x38\x10\x3d\xaf\x7f\xc5\x2c\xcf\xe1\x2a\x8b\x04\x8b\x2d\x4a\x0a\x08\xdc\x04\xed\xa1\x68\x51\x24\x40\x7b\x1c\x51\xe3\x88\x09\x45\xb2\xe4\x48\xb6\xfb\xeb\x0b\x49\xb6\xaa\xb8\x29\xfa\x71\xb1\x39\x0f\xf3\xde\xe3\x3c\x8a\x54\x7f\xbf\x7a\xb7\xbe\xfd\xf4\xfe\x1a\x1a\x6e\x5d\xb9\x52\xc3\x1f\x38\xf4\xf7\x5a\x90\x17\xe5\x0a\x40\x35\x84\x75\xb9\xfa\x0b\x40\xb1\x65\x47\xe5\xdb\xab\x0f\x6b\x38\xbf\xb8\x04\x0e\x50\x85\xce\xd7\xd6\xdf\x43\x15\x76\xc0\x21\x38\x55\x4c\x4d\x63\x7f\x4b\x8c\xd0\x30\x47\x49\x9f\x3b\xdb\x6b\xb1\x0e\x9e\xc9\xb3\xbc\xdd\x47\x12\x60\xa6\x4a\x0b\xa6\x1d\x17\x83\xf3\x4b\x30\x0d\xa6\x4c\xac\xef\x6e\x6f\xe4\xff\x62\x21\xe3\xb1\x25\x2d\x12\x6d\x28\x25\x4a\x0b\x72\x48\xf6\xde\x7a\xf1\x03\xc7\x8f\xf2\xee\x4a\xae\x43\x1b\x91\x6d\xe5\x96\xa6\x6f\xae\xf5\x0b\x01\xc5\x77\x16\x18\xa3\x23\xd9\x86\xca\x3a\x92\x5b\xaa\x24\xc6\x28\x0d\x46\x7c\x4a\xdf\x53\xfe\x65\x76\x66\xe4\x2e\xcb\x0a\x93\xcc\xbc\x7f\x22\x53\x39\x34\x8f\xcf\x09\xbd\x46\x5f\x37\xe4\xea\x9b\x64\xc9\xd7\x6e\xbf\x8c\x2b\x75\xf4\x1c\xa5\xb7\xb4\x8d\x21\xf1\xa2\x75\x6b\x6b\x6e\x74\x4d\xbd\x35\x24\xc7\xe2\x0c\xac\xb7\x6c\xd1\xc9\x6c\xd0\x91\xfe\xf7\x0c\x5a\xdc\xd9\xb6\x6b\x17\x80\xf5\x4f\x81\x2e\x53\x1a\xab\x21\x04\xed\xc3\xe8\x3e\xda\x67\x93\x6c\x64\xe0\x7d\xa4\xc3\x39\x3e\x60\x8f\x13\x2a\x20\x27\xa3\xc5\x02\x29\x1c\xe1\xc6\x11\xff\xf3\x90\x45\xa9\x8a\x09\x1c\x3e\xb2\x41\xca\x59\xff\x08\x89\x9c\x16\x63\x48\xb9\x21\x62\xb1\x54\x36\x39\x0b\x68\x12\x6d\xb4\x18\xd6\xb3\xd8\x88\x1f\x37\xf4\xdb\x2a\x2d\x26\x23\xcf\x2f\x2e\x67\x99\x3f\x18\x6b\xd6\x38\x9d\x4b\x15\xd3\xed\x01\x50\x55\xa8\xf7\x25\x00\xc0\x6a\x35\xfc\xaa\xda\xf6\x60\x6b\x2d\x06\xae\x00\xe3\x30\x67\x2d\x52\xd8\x8e\x97\x0e\x40\x6d\x42\x6a\xa7\x25\x80\xb2\x3e\x76\xcb\xed\x88\xc3\x91\x1f\x8d\xc5\x2c\x35\x55\xd1\xa1\xa1\x26\xb8\x9a\x92\x16\xd7\x9e\x29\x01\xc2\x7c\x71\x7b\x74\x1d\x41\x43\x89\xc4\xb4\xd6\x42\x40\xb6\x5f\x48\x8b\xff\x2e\xc7\x0c\x0e\xb6\x55\xc7\x1c\xfc\xa8\x9d\xbb\xaa\xb5\x73\x92\x87\xaa\x34\xc1\xf7\x94\xf8\xf4\x29\x50\xc5\xc4\x3c\x8c\x52\x7c\x9b\x45\x15\xb5\xed\xcb\x93\x08\x12\xe5\xce\x71\x3e\x4e\x7e\x84\xab\x2a\xec\x86\x5b\xb6\xcc\xe6\xc0\x3f\x49\x30\x2e\x9b\x20\xa2\x27\xf7\x7c\x6b\xc2\xed\xcf\x5a\xe7\x0d\xaa\x62\x3c\xb2\x95\x2a\xa6\xb7\xf1\x6b\x00\x00\x00\xff\xff\x53\x41\xd5\x1f\x2c\x05\x00\x00")

func templatesHtmlMarc034HtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlMarc034Html,
		"templates/html/marc-034.html",
	)
}

func templatesHtmlMarc034Html() (*asset, error) {
	bytes, err := templatesHtmlMarc034HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/marc-034.html", size: 1324, mode: os.FileMode(420), modTime: time.Unix(1515786433, 0)}
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
	"templates/html/marc-034.html": templatesHtmlMarc034Html,
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
	"templates": &bintree{nil, map[string]*bintree{
		"html": &bintree{nil, map[string]*bintree{
			"marc-034.html": &bintree{templatesHtmlMarc034Html, map[string]*bintree{}},
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

