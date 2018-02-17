// Code generated by go-bindata.
// sources:
// .DS_Store
// .gitignore
// README.md
// api/handlers.go
// api/router.go
// logic/.DS_Store
// logic/xxx.go
// main.go
// sdk/rpc.go
// sdk/rpc_test.go
// types/types.go

package tpl

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

	"github.com/henrylee2cn/ant/cmd/ant/info"
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

var _Ds_store = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\x4f\x4a\xc3\x40\x14\xc6\xbf\x99\x44\x08\xb8\x99\xa5\xcb\x5c\xa1\x37\x08\x25\x9e\x20\x3b\x57\x4a\x22\xb1\x98\x36\xc1\xb6\x82\x0b\x21\x47\xf0\x74\x9e\x47\x86\xf9\x8a\xb1\x89\xd9\x89\x6d\xf9\x7e\x10\x7e\x09\x6f\xfe\x3c\xb2\x98\x79\x33\x00\xcc\x72\x5f\x2d\x00\x07\x20\x41\xb0\xf1\x2f\x13\x24\x7c\x46\xd8\x41\x03\xe3\xc7\x78\xd9\x34\xed\xa6\x0e\x5f\x67\x82\xcf\x3d\xc2\x03\x3a\xac\x5e\x8b\x1f\xf9\x5f\xa1\x41\x8b\x1a\x2b\x94\x47\x91\x08\x5b\x54\x78\x1e\xb5\xdf\xe1\x0d\x1d\x1e\xb1\x6d\xea\x45\x51\xb6\xeb\x6e\x30\xcf\x20\xba\x6e\xf3\xbc\xda\xef\x4a\xe0\xf3\xe3\xee\x7d\x1c\xad\x66\xa2\xdd\xd3\xdc\xc8\x47\x19\x09\x21\x84\x10\x53\x70\x8f\x48\xae\xff\x3b\x11\x21\xc4\xc9\xe1\xd7\x87\x94\xce\xe8\x3e\xd8\x30\x6e\xe9\x78\xd0\xc7\xd1\x29\x9d\xd1\x7d\xb0\x61\x3b\x4b\xc7\x74\x42\x3b\x3a\xa5\x33\xba\x0f\xe6\xa2\x65\x78\xf8\x30\x9c\xf9\x70\x78\x31\x8e\x4e\xe9\xec\x6f\xfe\x8d\x10\xe7\x4e\x14\xe4\xfc\xfe\x7f\xfb\xfb\xf9\x5f\x08\x71\xc1\x98\x38\x2f\xf2\xe5\xcc\xa5\x91\x65\x21\x70\x7f\xe8\xc0\x42\x00\x13\x45\x80\x0d\x77\x6a\x37\xf8\x8e\xab\x10\x10\xe2\xc4\xf8\x0a\x00\x00\xff\xff\x44\x4a\x2b\xfa\x04\x18\x00\x00")

func Ds_storeBytes() ([]byte, error) {
	return bindataRead(
		_Ds_store,
		".DS_Store",
	)
}

func Ds_store() (*asset, error) {
	bytes, err := Ds_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1517885954, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _Gitignore = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8e\xc1\x0a\xc2\x30\x10\x44\xef\xf3\x29\x7b\x18\x50\x50\xfc\x17\x91\xd0\xa6\x69\xac\xd8\x6e\x4c\x52\xa9\x3d\xf8\xed\x92\xa6\x5e\x66\x77\x1e\xcb\xec\x08\x15\xc2\x06\xc2\xa4\x30\xda\x3e\x60\xb2\x4b\x19\xc2\xeb\xe9\x7c\x79\xbf\x6e\xd8\x27\x75\x2e\xd4\x7a\x3d\xd0\x6b\xdd\x8e\xb4\x30\xd6\xab\xe9\x5c\x3f\x4f\x7f\xe3\x35\x7f\x82\x4b\xe5\x6a\xf3\x6e\x09\x1a\x33\xa5\x26\x8f\xcd\x30\xd5\x00\xb7\xb8\xaa\x5f\x08\xf7\xa7\x21\x6a\x0f\x61\x6c\x22\x84\xeb\x10\x20\xf4\x6b\xe1\xa9\x83\xb0\x1d\x8b\xda\xde\x17\x12\xf2\x02\xe1\x53\x3d\x64\xd2\xfb\x1c\xb6\x82\xbf\x00\x00\x00\xff\xff\xbf\xaa\x91\x77\xcf\x00\x00\x00")

func GitignoreBytes() ([]byte, error) {
	return bindataRead(
		_Gitignore,
		".gitignore",
	)
}

func Gitignore() (*asset, error) {
	bytes, err := GitignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".gitignore", size: 207, mode: os.FileMode(420), modTime: time.Unix(1517885954, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _readmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xa8\xae\x0e\x08\xf2\xf7\x8a\xf7\x73\xf4\x75\xad\xad\xe5\x02\x04\x00\x00\xff\xff\x40\x0d\xb5\xec\x10\x00\x00\x00")

func readmeMdBytes() ([]byte, error) {
	return bindataRead(
		_readmeMd,
		"README.md",
	)
}

func readmeMd() (*asset, error) {
	bytes, err := readmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "README.md", size: 16, mode: os.FileMode(420), modTime: time.Unix(1518840422, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _apiHandlersGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\x2c\xc8\xe4\x02\x04\x00\x00\xff\xff\x0c\x0c\x0a\x62\x0c\x00\x00\x00")

func apiHandlersGoBytes() ([]byte, error) {
	return bindataRead(
		_apiHandlersGo,
		"api/handlers.go",
	)
}

func apiHandlersGo() (*asset, error) {
	bytes, err := apiHandlersGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api/handlers.go", size: 12, mode: os.FileMode(420), modTime: time.Unix(1518840223, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _apiRouterGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8b\x41\x0e\x02\x21\x0c\x45\xd7\xf6\x14\xcd\xac\x66\x5c\x0c\x89\x47\xf1\x06\x88\x15\x88\x40\x9b\x52\x16\xc6\x78\x77\x83\xba\x70\xf9\xdf\x7f\x4f\x7c\xb8\xfb\x48\xe8\x25\x03\xe4\x2a\xac\x86\x2b\x1c\x4c\x70\x89\xd9\xd2\xb8\xec\x81\xab\x4b\xd4\xf4\x51\x88\x4e\xa1\x39\xa3\x42\x53\x5b\x60\x03\x70\x0e\xc3\xe8\xc6\xf5\xcc\xc3\x08\x95\x62\xee\x46\xda\x7f\x14\x93\x6f\xd7\x32\xb7\x31\xea\x54\x74\x87\xdb\x68\xe1\xbf\x5a\xbf\x07\x1e\x4d\xf6\x0f\xd0\x0d\x9f\xf0\x82\x77\x00\x00\x00\xff\xff\xd3\xb5\x7a\x55\x9b\x00\x00\x00")

func apiRouterGoBytes() ([]byte, error) {
	return bindataRead(
		_apiRouterGo,
		"api/router.go",
	)
}

func apiRouterGo() (*asset, error) {
	bytes, err := apiRouterGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api/router.go", size: 155, mode: os.FileMode(420), modTime: time.Unix(1517885954, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _logicDs_store = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\xd8\x31\x0a\x02\x31\x10\x85\xe1\x37\x31\x45\xc0\x26\xa5\x65\x1a\x0f\xe0\x0d\xc2\xb2\x9e\xc0\x0b\x58\x78\x05\xfb\x1c\x5d\x96\x79\x60\x60\xd5\x4e\x8c\xcb\xfb\x40\xfe\x05\x37\x2a\x16\x31\x23\x00\x9b\xee\xb7\x13\x90\x01\x24\x78\x71\xc4\x4b\x89\x8f\x95\xd0\x5d\x1b\x5f\x43\x44\x44\x44\xc6\x66\x9e\xb4\xff\xf5\x07\x11\x91\xe1\x2c\xfb\x43\x61\x2b\xdb\xbc\xc6\xe7\x03\x1b\xbb\x35\x99\x2d\x6c\x65\x9b\xd7\x78\x5f\x60\x23\x9b\xd8\xcc\x16\xb6\xb2\xcd\xcb\x4d\xcb\x38\x7c\x18\xdf\xd9\x38\xa1\x18\xa7\x10\x2b\x6c\xfd\xce\x77\x23\xf2\xef\x76\x9e\xbc\xfc\xfe\x9f\xdf\xcf\xff\x22\xb2\x61\x16\xe7\xcb\x3c\x3d\x07\x82\xf5\x0d\x00\xae\xdd\xf5\xa7\x43\x40\xf0\x3f\x0b\x0f\xdd\x5a\x1d\x04\x44\x06\xf3\x08\x00\x00\xff\xff\x6a\x00\x88\x6d\x04\x18\x00\x00")

func logicDs_storeBytes() ([]byte, error) {
	return bindataRead(
		_logicDs_store,
		"logic/.DS_Store",
	)
}

func logicDs_store() (*asset, error) {
	bytes, err := logicDs_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logic/.DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1517885954, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _logicXxxGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\xc9\x4f\xcf\x4c\xe6\xe2\xca\xcc\x2d\xc8\x2f\x2a\x51\xd0\xe0\xe2\x54\xaa\xae\x0e\x08\xf2\xf7\x8a\x0f\x70\x0c\xf1\xa8\xad\xd5\x2f\xa9\x2c\x48\x2d\x56\xe2\xe2\x2c\x29\x50\x50\x4a\xcf\x2c\xc9\x28\x4d\xd2\x4b\xce\xcf\xd5\xcf\x48\xcd\x2b\xaa\xcc\x49\x4d\x35\x4a\xce\xd3\x2f\x49\xcd\x49\x05\xe9\x56\xe2\xd2\xe4\xe2\x02\x04\x00\x00\xff\xff\xfc\x3c\x50\xb5\x58\x00\x00\x00")

func logicXxxGoBytes() ([]byte, error) {
	return bindataRead(
		_logicXxxGo,
		"logic/xxx.go",
	)
}

func logicXxxGo() (*asset, error) {
	bytes, err := logicXxxGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "logic/xxx.go", size: 88, mode: os.FileMode(420), modTime: time.Unix(1518840241, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _mainGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcd\xb1\x6a\xc3\x30\x10\xc6\xf1\x59\xf7\x14\x87\x26\x09\x8c\x1d\xba\xc5\xd0\x21\x94\x42\x28\xa5\x0d\x49\xf7\x72\x76\x2e\x8e\x68\x2c\x99\xd3\xd9\xa5\x18\xbf\x7b\x49\x3c\x7e\xdf\xf0\xfb\x0f\xd4\xfe\x50\xc7\xd8\x53\x88\x00\xa1\x1f\x92\x28\x3a\x30\xb6\x0b\x7a\x1d\x9b\xb2\x4d\x7d\x75\xe5\x28\x7f\x37\xe6\xa7\x36\x56\x14\xd5\x82\xb1\xf3\x7c\x38\x7e\xbe\x7d\x1f\x76\x5f\xfb\x65\xa9\x68\x08\x16\x3c\xc0\x65\x8c\xed\x43\x72\x1e\x67\x30\x59\x26\xac\x9f\x91\xa2\x96\x1f\xfc\x7b\x62\x99\x58\xdc\x7d\x9d\x64\x7a\x49\xf1\x12\xba\x19\x8c\x79\x0f\x59\x39\xee\xce\x67\xe1\x9c\x6b\x44\xb4\xf5\x76\xb3\xdd\xd8\x02\x8c\x79\x8d\xd4\xdc\x78\xcf\x24\xda\x30\x69\x8d\x2a\x23\x17\x60\x16\x0f\x86\x86\x50\x1e\xd3\xa8\xec\x6c\x25\x29\xa9\x2d\x30\xcb\xb4\x5e\xe2\xbc\x7f\xf4\xcb\x55\x77\x1e\x16\xf8\x0f\x00\x00\xff\xff\xda\x19\xaf\x73\xeb\x00\x00\x00")

func mainGoBytes() ([]byte, error) {
	return bindataRead(
		_mainGo,
		"main.go",
	)
}

func mainGo() (*asset, error) {
	bytes, err := mainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main.go", size: 235, mode: os.FileMode(420), modTime: time.Unix(1517915085, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sdkRpcGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x4e\xc9\xe6\x02\x04\x00\x00\xff\xff\x36\xfa\x03\xb1\x0c\x00\x00\x00")

func sdkRpcGoBytes() ([]byte, error) {
	return bindataRead(
		_sdkRpcGo,
		"sdk/rpc.go",
	)
}

func sdkRpcGo() (*asset, error) {
	bytes, err := sdkRpcGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sdk/rpc.go", size: 12, mode: os.FileMode(420), modTime: time.Unix(1517885954, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sdkRpc_testGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x4e\xc9\xe6\x02\x04\x00\x00\xff\xff\x36\xfa\x03\xb1\x0c\x00\x00\x00")

func sdkRpc_testGoBytes() ([]byte, error) {
	return bindataRead(
		_sdkRpc_testGo,
		"sdk/rpc_test.go",
	)
}

func sdkRpc_testGo() (*asset, error) {
	bytes, err := sdkRpc_testGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sdk/rpc_test.go", size: 12, mode: os.FileMode(420), modTime: time.Unix(1517885954, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typesTypesGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\xa9\x2c\x48\x2d\xe6\x02\x04\x00\x00\xff\xff\xa8\xd5\x8c\x30\x0e\x00\x00\x00")

func typesTypesGoBytes() ([]byte, error) {
	return bindataRead(
		_typesTypesGo,
		"types/types.go",
	)
}

func typesTypesGo() (*asset, error) {
	bytes, err := typesTypesGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "types/types.go", size: 14, mode: os.FileMode(420), modTime: time.Unix(1517885954, 0)}
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
	".DS_Store":       Ds_store,
	".gitignore":      Gitignore,
	"README.md":       readmeMd,
	"api/handlers.go": apiHandlersGo,
	"api/router.go":   apiRouterGo,
	"logic/.DS_Store": logicDs_store,
	"logic/xxx.go":    logicXxxGo,
	"main.go":         mainGo,
	"sdk/rpc.go":      sdkRpcGo,
	"sdk/rpc_test.go": sdkRpc_testGo,
	"types/types.go":  typesTypesGo,
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
	".DS_Store":  &bintree{Ds_store, map[string]*bintree{}},
	".gitignore": &bintree{Gitignore, map[string]*bintree{}},
	"README.md":  &bintree{readmeMd, map[string]*bintree{}},
	"api": &bintree{nil, map[string]*bintree{
		"handlers.go": &bintree{apiHandlersGo, map[string]*bintree{}},
		"router.go":   &bintree{apiRouterGo, map[string]*bintree{}},
	}},
	"logic": &bintree{nil, map[string]*bintree{
		".DS_Store": &bintree{logicDs_store, map[string]*bintree{}},
		"xxx.go":    &bintree{logicXxxGo, map[string]*bintree{}},
	}},
	"main.go": &bintree{mainGo, map[string]*bintree{}},
	"sdk": &bintree{nil, map[string]*bintree{
		"rpc.go":      &bintree{sdkRpcGo, map[string]*bintree{}},
		"rpc_test.go": &bintree{sdkRpc_testGo, map[string]*bintree{}},
	}},
	"types": &bintree{nil, map[string]*bintree{
		"types.go": &bintree{typesTypesGo, map[string]*bintree{}},
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
	data = bytes.Replace(data, projNameTpl, projNameBytes, -1)
	data = bytes.Replace(data, projPathTpl, projPathBytes, -1)
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

var (
	projNameBytes []byte
	projPathBytes []byte
	projNameTpl   = []byte("{{PROJ_NAME}}")
	projPathTpl   = []byte("{{PROJ_PATH}}")
)

// Create creates base files.
func Create() {
	projNameBytes = []byte(info.ProjName())
	projPathBytes = []byte(info.ProjPath())
	RestoreAssets("./", "")
}
