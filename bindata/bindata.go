package bindata

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

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _static_index_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x92\x31\x53\xc4\x20\x10\x85\xeb\xbb\x5f\xb1\xd2\x1f\x24\x6a\xe1\x44\x72\x8d\x5a\x6b\x61\x63\x89\xb0\x06\xce\x10\x22\x6c\xf4\xf2\xef\x25\xc4\xc2\x51\xc7\x46\xab\xb7\x2c\xc3\xf7\xe6\xbd\x41\x9e\x5c\xdf\x5e\xdd\x3f\xdc\xdd\x80\x25\xdf\xef\xb7\x72\x95\x8d\xb4\xa8\x4c\xd6\x8d\x24\x47\x3d\xee\xa5\x58\x75\xd9\x78\x24\x05\xda\xaa\x98\x90\x5a\x36\xd1\xd3\xee\x82\x81\x28\x57\x49\x47\x37\x12\xa4\xa8\x5b\x66\x89\xc6\x46\x08\x1d\x0c\xf2\xc3\xcb\x84\x71\xe6\x3a\x78\xb1\x8e\xbb\x9a\xd7\x35\xaf\xb8\x77\x03\x3f\x24\x96\xf9\xeb\xd3\x5f\x29\xd6\x75\x76\xf1\xa5\x54\x48\x89\x82\x7e\x16\xcb\xb2\x4c\x7f\xe4\xf8\x60\xa6\x1e\x93\xc0\xe3\x18\x22\xb9\xa1\xfb\xca\x93\xe2\xa3\x13\xf9\x18\xcc\x5c\x0c\x8c\x7b\x05\x67\x5a\xa6\xc3\x40\xca\x0d\x18\x6b\x06\x89\xe6\x1e\x5b\x96\x83\xed\xde\x9c\x21\xdb\xc0\x59\x5d\x8d\xc7\x4b\xb0\x98\x6d\xa9\x81\xf3\xaa\x1c\xbd\x8a\x9d\x1b\x1a\xa8\x40\x4d\x14\x16\xa7\x4c\xfb\x99\x7a\xfa\x1f\xd4\xcf\x65\x78\xf5\xbd\x76\x29\xd6\x58\x39\x66\xf9\x02\xdb\xf7\x00\x00\x00\xff\xff\xde\x27\x8d\x1c\x1b\x02\x00\x00")

func static_index_html_bytes() ([]byte, error) {
	return bindata_read(
		_static_index_html,
		"static/index.html",
	)
}

func static_index_html() (*asset, error) {
	bytes, err := static_index_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "static/index.html", size: 539, mode: os.FileMode(420), modTime: time.Unix(1421573304, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _static_main_js = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x55\x5d\x6b\xe3\x38\x14\x7d\xb6\x7f\x85\xd6\x94\xc6\x66\xb3\xce\x57\x0b\x8b\x4b\x28\xdd\x2e\x74\x18\x98\xb6\x90\xc0\x3c\x94\x3c\x28\xf6\x75\x62\xaa\x48\x46\x92\x9b\xa4\x43\xfe\xfb\x5c\x49\xb6\x1b\x4f\x29\xb4\xc3\x30\x4f\xf3\x12\x1c\xdd\x73\xee\xe7\x91\xee\x13\x95\x24\x5d\x53\xa9\x47\x17\x7e\xfb\x3d\xbe\xf0\xfd\x93\x30\xaf\x78\xaa\x0b\xc1\xc3\x88\x7c\xf3\x3d\x63\xdc\x91\x29\xe1\xb0\x25\xff\x53\x0d\x61\x84\x20\xef\x53\xb1\x5a\x5b\x8a\x8a\x15\xe8\xbb\xd2\xe0\x55\x88\x70\x6f\xc5\xc4\x92\xb2\xc4\x50\x3d\x4f\x17\x1b\x78\x16\x1c\xee\xf2\x1c\x61\x09\xd9\xc5\x2b\xd0\xf3\xce\x61\x18\x21\xf0\xe0\x7b\x87\x08\xdd\x9e\x18\xfb\xe7\xd9\xdd\x6d\xd8\x1b\x64\xb0\xac\x56\x03\x17\x64\x90\x51\x4d\x2f\x53\xca\xd8\x92\xa6\x8f\xd3\xcb\x5e\x9f\xb4\x59\x1a\x93\xcd\xd4\x73\xf5\xd4\xb9\x1e\x65\x38\xd3\x22\x7d\xbc\x36\xdf\x36\x43\x87\xab\x33\xf4\x24\xf0\x0c\xe4\x5c\x24\xa4\x97\x0a\xae\x69\xc1\x41\x8e\x7a\x7d\x6b\x7b\x16\x62\x33\xdf\x97\x80\xb6\x5d\xcf\x9c\x1c\xfa\xae\x2a\xcd\xa0\xe1\x6b\xd8\xa1\xaf\xde\xcd\x35\x29\x69\xa5\x40\x1d\xe1\xf6\x57\xbb\x42\xb5\xb8\x63\x52\xc3\xba\xa5\x5c\x28\xc0\xb8\x99\xe3\x99\x46\x34\x6c\x95\x4a\x61\xea\x95\x0d\x09\x38\x5d\x32\xc8\x12\x92\x53\xa6\xe0\x05\x28\x29\x5f\xc1\x0c\x18\xa4\x5a\xb4\xe0\x65\xa5\x35\x8e\x24\x21\x0f\x4d\x44\x57\x87\x8b\x56\xd7\xe7\xa5\xa2\xe2\x98\xc7\x79\xbf\x93\xd5\x79\x93\x4c\x9f\xbc\x83\x3c\x19\x76\xd9\x93\xe1\x5b\xf4\x4d\xc1\x2b\x0d\x3f\xd0\x47\x5d\xf6\x68\xf3\x06\x19\x67\xdf\xeb\x42\xaf\xf0\xc4\x61\x17\xce\xa0\x6c\x0f\x4c\x87\x26\x47\x6d\x04\x59\xc0\x4b\x1f\x38\xdd\xa0\xb7\xa0\x1d\x57\xe0\xa8\x46\x44\x09\x31\xbf\xf1\x4d\x7a\x6f\x2d\xce\xd0\x44\x97\x40\xeb\xf0\x5a\x08\xa6\x8b\xb2\x1d\xe5\x13\x65\x15\xcc\xaa\x3c\x2f\x76\x08\xe4\x9d\x41\x2e\x8c\xb6\xf1\xc2\xd4\x9a\x1b\xff\x02\x6d\x8e\x3f\xae\xcd\x2f\xb0\x11\x72\x4f\xb0\x83\x22\xc5\x1b\x9c\x7d\x48\xa2\xff\xed\x35\xfc\x11\xe7\x6f\x15\xe7\x55\x33\xa8\xd7\xe2\xb4\xd3\x68\xed\x3f\x27\xd1\xe5\x6b\x85\xfa\x56\xa5\xbe\xd7\x3c\xaa\x64\xab\x2a\xc9\xdc\x02\xb0\x1b\x80\xa1\x72\xb7\x05\xcf\xc4\x36\xb6\xb1\x11\x63\x54\x2d\x41\x57\x92\x93\x30\x64\x71\x29\x05\x4a\x59\x20\x70\x3a\x25\xc1\x5a\xeb\x52\x25\x41\x44\x2e\x49\xb0\x55\x2a\x19\x0c\x02\x92\x98\x4f\xf3\x15\x91\xbf\x09\x8b\xd7\x42\x69\x53\x30\xfe\x09\xad\x03\x21\x35\xf9\x6b\x4a\xfe\x1d\x46\xe4\xf4\x94\x1c\x9d\x9c\x9d\x4d\x22\xeb\x29\x09\x2c\xd3\x9e\xa3\x37\xeb\x28\x78\xbd\x28\xfe\xc9\x01\x7b\x87\xf9\x1d\xb0\xa4\xad\xaa\x2f\xdd\x57\x58\xce\xf0\xae\xe1\xc6\xa9\x8b\x33\xf7\x72\xab\x62\xc1\x45\x09\x1c\x41\x6d\xf1\x75\xdd\xd6\xb6\x01\xa5\xe8\x0a\x3a\x66\x78\xd2\x0e\x61\x5b\x63\x22\xa2\xd9\x6c\xad\xb8\xa4\x52\x81\xb1\xc7\x76\x2b\x99\x16\x79\x45\x4e\xc2\xe3\x87\xc5\x54\x34\xac\xf9\xf5\xca\x8a\x9d\x0e\x1e\x86\x8b\x98\x66\xd9\xbd\x28\xb8\x0e\x1f\x2c\x67\xae\xfa\x9d\x57\x69\xd1\x27\x5a\x56\xe0\x3c\x1f\xda\x07\x63\xfc\x0e\x0f\x5d\xe9\x1c\x3b\x32\xdb\xf7\xc2\x47\x15\x7c\x0f\x00\x00\xff\xff\x46\xc4\x57\x98\x15\x08\x00\x00")

func static_main_js_bytes() ([]byte, error) {
	return bindata_read(
		_static_main_js,
		"static/main.js",
	)
}

func static_main_js() (*asset, error) {
	bytes, err := static_main_js_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "static/main.js", size: 2069, mode: os.FileMode(420), modTime: time.Unix(1421574661, 0)}
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
	"static/index.html": static_index_html,
	"static/main.js": static_main_js,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"static": &_bintree_t{nil, map[string]*_bintree_t{
		"index.html": &_bintree_t{static_index_html, map[string]*_bintree_t{
		}},
		"main.js": &_bintree_t{static_main_js, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

