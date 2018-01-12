// Code generated by go-bindata.
// sources:
// templates/html/034.html
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

var _templatesHtml034Html = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xdf\x6f\xd3\x3e\x10\x7f\xfe\xf6\xaf\xb8\xef\x3d\x6d\xd2\xb2\x0c\x6d\x0f\xa0\xd9\x95\xa6\xd2\x69\x48\x43\x43\x68\x93\xe0\xd1\x89\xaf\x8d\x37\xc7\x36\xf6\x25\x6d\x85\xf8\xdf\x51\x9c\x32\x52\x31\x10\xbc\xb4\xbe\xf3\xe7\x97\x2f\x71\xc4\xff\x6f\xef\x16\xf7\x9f\x3f\x2c\xa1\xe1\xd6\xce\x67\x62\xf8\x03\xab\xdc\x5a\x22\x39\x9c\xcf\x00\x44\x43\x4a\xcf\x67\xff\x01\x08\x36\x6c\x69\xfe\xfe\xea\xe3\x02\xce\xce\x2f\x80\x3d\x54\xbe\x73\xda\xb8\x35\x54\x7e\x0b\xec\xbd\x15\xe5\x08\xca\xf8\x96\x58\x41\xc3\x1c\x0a\xfa\xd2\x99\x5e\xe2\xc2\x3b\x26\xc7\xc5\xfd\x2e\x10\x42\x3d\x56\x12\x99\xb6\x5c\x0e\xce\x97\x50\x37\x2a\x26\x62\xf9\x70\x7f\x5d\xbc\xc6\x89\x8c\x53\x2d\x49\x8c\xb4\xa2\x18\x29\x4e\xc8\x3e\x9a\xb5\x71\xf8\x1b\xc7\x4f\xc5\xc3\x55\xb1\xf0\x6d\x50\x6c\x2a\x3b\x35\x7d\xb7\x94\x6f\x10\xca\x5f\x2c\x54\x08\x96\x8a\xd6\x57\xc6\x52\xb1\xa1\xaa\x50\x21\x14\xb5\x0a\xea\x90\xbe\xa3\xf4\xd7\xec\xc4\x8a\xbb\x54\x54\x2a\x16\x89\x77\x07\x32\x95\x55\xf5\xd3\x4b\x42\x37\xca\xe9\x86\xac\xbe\x8e\x86\x9c\xb6\xbb\xe9\xb8\x62\x47\x2f\x51\x7a\x43\x9b\xe0\x23\x4f\xa0\x1b\xa3\xb9\x91\x9a\x7a\x53\x53\x91\x8b\x13\x30\xce\xb0\x51\xb6\x48\xb5\xb2\x24\x5f\x9d\x40\xab\xb6\xa6\xed\xda\x49\xc3\xb8\xc3\x46\x97\x28\xe6\x6a\x18\x82\x74\x3e\xbb\xcf\x00\x86\x00\xd6\xb8\x27\x88\x64\x25\xe6\xb3\xa5\x86\x88\x11\x78\x17\x68\xff\x60\xeb\x94\x10\x9a\x48\x2b\x89\xc3\xba\x34\x4e\xd3\xf6\x34\x77\xf7\x67\xf8\x67\x89\xb3\xf3\x8b\xa9\x40\x16\x49\x75\x34\x81\xa7\xac\x47\xd5\xab\xb1\x8b\x90\x62\x2d\x71\xd2\xc9\x12\x8f\x09\xe7\xa2\x1c\x1b\x3f\x84\xfe\xac\x93\x51\xb0\x31\x4e\xfb\xcd\xa9\xd2\x7a\xd9\x93\xe3\x5b\x93\x98\x1c\xc5\x23\xb4\x5e\x69\x3c\x81\x55\xe7\x6a\x36\xde\xc1\x50\x1f\xd1\x80\x39\xfe\x9a\x99\xb5\x77\xc9\x5b\x3a\xb5\x7e\x7d\x84\x37\xcb\xdb\xdb\x3b\x3c\xce\x1b\xdf\x8e\x2f\x07\xf3\xe7\x30\x00\xa2\x1c\x2f\x1e\x80\xa8\xbc\xde\x0d\x0b\x80\xfc\x23\x56\x3e\xb6\x63\x0d\x20\x8c\x0b\xdd\x34\x2e\xee\xdf\x86\xb3\xf3\x0b\x84\x60\x55\x4d\x8d\xb7\x9a\xa2\xc4\xa5\x63\x8a\xa0\xe0\xf9\x02\xf7\xca\x76\x04\x0d\x45\xc2\x71\x2d\x31\x8f\x74\xaf\x5c\x75\xcc\xde\xed\xa5\x53\x57\xb5\x86\x71\xbe\xf0\xae\xa7\xc8\xa2\x1c\x77\x47\xb0\x28\x7f\x46\xca\xd1\xc7\xc4\xa2\x1c\xbf\x2a\xdf\x03\x00\x00\xff\xff\xbf\x7b\xc3\xdf\x66\x04\x00\x00")

func templatesHtml034HtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtml034Html,
		"templates/html/034.html",
	)
}

func templatesHtml034Html() (*asset, error) {
	bytes, err := templatesHtml034HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/034.html", size: 1126, mode: os.FileMode(420), modTime: time.Unix(1515772829, 0)}
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
	"templates/html/034.html": templatesHtml034Html,
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
			"034.html": &bintree{templatesHtml034Html, map[string]*bintree{}},
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

