// Package swagger Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// v1alpha/tracker/tracker.swagger.json
// v1alpha/extractor/extractor.swagger.json
// v1alpha/schema/schema.swagger.json
// v1alpha/deps/deps.swagger.json
// v1alpha/store/store.swagger.json
package swagger

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}


type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _v1alphaTrackerTrackerSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x6f\xa4\x36\x10\x7f\xdf\x4f\x61\xb9\x7d\x8c\x42\x9a\xbe\xe5\xad\x6a\x55\xa9\x52\xef\xa5\xc9\x5b\x15\x55\x0e\xcc\x12\x5f\xc1\x76\xec\x21\xd5\xf6\xb4\xdf\xbd\x02\x43\xf8\xb7\xb0\x0b\x06\x12\x52\x56\x3a\x1d\xb7\xf6\xcc\xcd\x9f\xdf\xcf\x33\x9e\xe5\xdb\x8e\x10\x6a\xfe\x61\x61\x08\x9a\xde\x11\x7a\x7b\x7d\x43\xaf\xd2\xef\xb8\xd8\x4b\x7a\x47\xd2\x75\x42\x28\x72\x8c\x20\x5d\x7f\xfd\x81\x45\xea\x99\x79\xa8\x99\xff\x37\xe8\xe2\xef\x6b\xa5\x25\xca\x4c\x92\x10\xfa\x0a\xda\x70\x29\xb2\xfd\xf6\x91\x08\x89\xc4\x00\xd2\x1d\x21\xc7\x4c\xbf\xf1\x9f\x21\x06\x43\xef\xc8\x9f\x56\xe8\x19\x51\x15\x0a\xd2\x67\x93\xee\x7d\xcc\xf6\xfa\x52\x98\xa4\xb6\x99\x29\x15\x71\x9f\x21\x97\xc2\xfb\x6a\xa4\x28\xf7\x2a\x2d\x83\xc4\xbf\x70\x2f\xc3\x67\x53\x3a\xe9\x15\xce\x85\x9a\xa9\x67\xef\x5b\xc4\x44\x98\xb0\x10\x8e\x5e\x00\x0a\x44\x00\xc2\xe7\x50\xee\x27\x84\x86\x80\x95\x7f\x12\x42\xa5\x02\x9d\xfd\x4f\xbf\x05\xa9\xf7\xbf\x73\x83\xbf\x54\x65\xaf\xca\xad\x1a\x8c\x92\xc2\xd4\x14\x66\x0b\xb7\x37\x37\x8d\xaf\x08\xa1\x01\x18\x5f\x73\x85\x79\x58\x7f\x22\x26\xf1\x7d\x30\x66\x9f\x44\xa4\xd0\x74\x5d\x51\x9f\x09\x65\x31\x66\x2d\x65\x84\xd0\xef\x35\xec\x53\x3d\xdf\x79\x01\xec\xb9\xe0\xa9\x5e\x53\x24\xb3\x69\xf5\x1f\xb9\x7e\x5a\xd3\x72\xdc\x9d\x7a\x3e\x56\x3c\x54\x4c\xb3\x18\x10\x74\x99\x0c\xfb\x69\xf8\x26\x58\x9c\x61\xab\x88\x77\xd3\x0d\x9e\xb9\x9c\x26\xab\xb9\xa2\xe1\x25\xe1\x1a\xd2\x58\xa3\x4e\xa0\xb1\x8a\x07\x95\xe9\x35\xa8\xb9\x08\xab\xd6\x1f\xaf\xce\x5b\x23\x75\xc8\x04\xff\x37\xcb\xe6\x69\x8b\x5e\x12\xd0\x87\x1e\x93\xf6\x2c\x32\xd3\xda\x14\xcb\x20\x89\x3a\xe2\x33\xb1\x35\x6f\xcf\x8f\x95\x8c\x22\x0b\x9b\xb9\xa4\x6f\x48\x39\xdc\x83\x7e\xe5\x7e\x05\x26\x8f\xbb\xaa\xb2\xdc\xc1\x0b\x69\xe6\xa1\x54\x32\x92\xe1\xc1\x85\x6f\x0f\x85\x8e\x8d\x77\x64\xe3\xdd\xe7\xe2\x5d\x81\xed\x79\x58\xe7\x21\x07\x6b\xbf\x33\xf9\x1e\xac\xa6\x55\x52\xd0\xda\xbe\x11\x71\x23\x62\xbe\x7f\x09\x22\xe2\xe8\x1e\x13\x57\xd8\x61\xe2\x56\xe7\x36\x7a\x15\xfb\x97\xe9\x2f\xd1\xbd\xbb\xc4\x15\xf7\x96\x1b\xe3\x36\xc6\xbd\xed\x5f\xa6\xa0\x4d\xd5\x57\xe2\xea\xbb\x4a\xdc\x7a\xca\x8d\x82\xb3\x51\xd0\xfa\x32\xb8\x7f\x5c\x0f\x8d\xbe\x64\x0e\x2e\x45\x1e\xd5\x49\x1c\x17\x60\x70\x81\x10\x82\x6e\x4a\xef\xa5\x8e\x19\xe6\x1b\x7e\xbc\x1d\x0a\x63\x5f\x26\x02\x3f\x84\xb1\x17\xa3\xdc\x26\x73\x24\xc6\xbd\x98\x09\x16\x0e\x2f\x25\x5f\x72\xb1\xf5\x40\xde\x1a\xbc\x14\xe6\x13\x1d\x7d\xac\xb3\xd0\x0d\x25\x46\x26\xda\x87\xa1\x20\xb9\xcf\xa4\x56\x74\x9b\xce\x0d\xfe\x18\x5d\xc5\xf2\x35\x7c\xeb\x2b\xec\x67\x1e\x2e\x99\x9c\x0d\x9f\xb6\xab\xb0\xec\xd9\xba\x8a\xb6\xad\x2b\xec\x2a\x6c\x32\x47\x62\xdc\xa2\xa2\x8a\x74\x25\x4d\x3f\xd4\x1f\x32\x89\x55\x60\x3d\x33\x75\x29\x98\x3f\xc9\xa0\x05\x0f\x8b\x9c\x53\x2b\xfd\xb7\xce\x91\xfe\x16\xbc\x7e\x49\xc0\xe0\x25\xfe\x4e\x83\xad\xb7\x97\x5c\x2a\x26\x95\xaf\x99\x04\xa0\x4c\x39\xd4\xb4\xdd\x5d\x0c\x02\x7f\xe5\x51\xad\x4f\x29\xb8\x22\x9f\xbe\x82\x5f\x72\x90\x2a\x9d\x02\x10\x79\x03\x5f\x65\x4d\x6e\xa0\xae\xab\x6e\x54\x73\x6b\x0e\x06\x21\x1e\x23\x59\xab\xbc\x23\xe4\xf3\x2a\x39\x42\xb2\x7c\xcb\x68\xb0\x68\xc7\xfb\x3c\x35\x79\xa6\x35\xab\xa3\x94\x72\x84\xb8\xb9\xbf\x13\x82\xf9\xc1\x52\xcf\xf6\xe9\x13\xed\x78\xf2\x68\xb2\x98\xb7\xb2\xe6\xbd\x80\x91\x07\xf9\x2f\x5f\x0a\x83\x9a\x71\x81\xa3\xe0\xe5\x4b\x35\x45\xa4\x2f\x69\x81\xfa\x82\x69\xc9\xf6\x6e\xc1\x1c\xcf\xb2\xa1\x58\xef\x8b\x41\x93\x70\x8b\x86\x60\xe1\xe3\xa2\x27\x0e\xf7\xad\x8b\xe1\xd0\x38\xa4\x37\x65\x27\x43\xce\xbd\xf7\xe3\x60\xdb\x82\x27\x5c\xee\xc5\xb8\x43\xee\xb2\xf7\x2e\x1c\x02\x81\xdc\x36\x2c\xb3\x44\xa0\x68\xab\x2a\xbf\x49\x38\x3a\x8f\x93\x66\x1f\xd7\x96\xfb\xd6\xaf\x23\x9f\x3f\xf3\xcd\xf1\x9e\x83\xc7\xed\xe9\xff\x1c\x3e\xe7\x06\xe7\x75\x64\xa4\xd3\xf5\x31\xbe\x83\xcf\xaa\xa7\x10\x9d\xba\x65\x76\xdf\x31\xab\x65\xc6\x5e\x7b\x27\x57\x3b\x73\x82\x6a\x05\x7e\x5c\x62\x1a\x93\x90\xff\x4b\x62\xda\x13\xae\x19\x12\x93\x77\x1c\x2e\x89\x99\xa4\x3a\xcc\xec\x6c\xfd\x98\x18\xed\x74\xfd\x98\x71\x39\x15\x5b\x4d\xff\xd9\x4c\x15\xf7\x84\x01\x2d\xe8\x10\x46\x5e\xe2\xb7\x7b\x7b\xda\xfa\xe5\xe3\xac\x9d\xcd\x5c\xd5\xbc\x9e\x22\x8a\x7d\x7e\xd7\x67\x35\x1f\xcc\xef\x62\x3e\x33\x1b\x69\x7a\x87\x42\xc3\xc9\x53\xeb\x4b\x1c\x7b\xa8\x15\xb4\x90\xf5\xb9\xa6\x8b\xbf\xa9\xa2\xf4\xe6\xd6\xe1\xf3\x93\x94\x11\x30\xd1\x55\x61\x8a\xe5\x33\x56\x9f\x1e\x0d\x39\x98\xfd\xee\x43\xb8\x9f\x57\x3e\x1e\xea\xc4\xa2\xf3\x25\x6c\xd0\x89\x59\x0c\xfb\x96\xa9\x3b\xbb\xf4\xcf\x71\xf7\x5f\x00\x00\x00\xff\xff\xf4\x75\xc9\x06\x03\x3a\x00\x00")

func v1alphaTrackerTrackerSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_v1alphaTrackerTrackerSwaggerJson,
		"v1alpha/tracker/tracker.swagger.json",
	)
}

func v1alphaTrackerTrackerSwaggerJson() (*asset, error) {
	bytes, err := v1alphaTrackerTrackerSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1alpha/tracker/tracker.swagger.json", size: 14851, mode: os.FileMode(420), modTime: time.Unix(1574222637, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v1alphaExtractorExtractorSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x56\x4d\x6f\xdb\x30\x0c\xbd\xfb\x57\x08\xda\x8e\x45\xd3\xf5\x98\xdb\xd0\x75\xc0\x0e\x05\x8a\x5d\x87\x1c\x58\x89\xb1\x55\xd8\x92\x2a\xd2\xdd\xb2\x21\xff\x7d\x90\x62\xd7\x1f\x8d\x8b\xc4\x19\x16\x0c\xcb\x49\xb1\xf8\x28\xbe\xa7\x67\xd2\xbf\x32\x21\x24\x7d\x87\x3c\xc7\x20\x97\x42\x5e\x5f\x5e\xc9\x8b\xf8\xcc\xd8\xb5\x93\x4b\x11\xf7\x85\x90\x6c\xb8\xc4\xb8\xff\xfc\x01\x4a\x5f\xc0\x02\x7f\x70\x00\xc5\x2e\x74\xab\x4b\x1f\x1c\xbb\x84\x16\x42\x3e\x63\x20\xe3\x6c\xc2\xec\x96\xc2\x3a\x16\x84\x2c\x33\x21\xb6\xe9\x0c\x52\x05\x56\x48\x72\x29\xbe\xed\x40\x05\xb3\x6f\x13\xc4\x35\xc5\xd8\x55\x8a\x55\xce\x52\x3d\x08\x06\xef\x4b\xa3\x80\x8d\xb3\x8b\x47\x72\xb6\x8b\xf5\xc1\xe9\x5a\x1d\x18\x0b\x5c\x50\x47\x74\xd1\x12\xd4\xe8\xd1\x6a\xb4\xca\x20\xb5\x1c\x5f\xc2\x22\xce\x51\xff\xbf\x10\xd2\x79\x0c\xe9\x84\x2f\x3a\xb2\xbe\x6d\x30\x17\x5d\x44\x40\xf2\xce\x12\xd2\x00\x28\x84\xbc\xbe\xba\x1a\x3d\x12\x42\x6a\x24\x15\x8c\xe7\x46\xc5\x8f\x82\x6a\xa5\x90\x68\x5d\x97\xa2\xcd\x74\xd9\x4b\x9f\x40\x49\x52\x78\x95\x4c\x08\xf9\x3e\xe0\x3a\xe6\x79\xb7\xd0\xb8\x36\xd6\xc4\xbc\xd4\xdd\x5e\x53\xee\xd7\x26\xb1\x1c\xc0\xb7\xd9\xbe\xf5\xb6\x47\xcd\x43\x80\x0a\x19\x43\x27\xfa\xee\x37\x22\x65\xa1\x4a\x3e\x7a\x70\x7a\x33\xae\xdd\xd8\xa9\x9d\x80\x4f\xb5\x09\x18\x75\xe5\x50\xe3\x1f\xe7\xfc\x54\x23\xf1\x21\x94\x57\x3d\xca\x0c\xf9\x98\xac\xfc\xd4\xba\x66\x73\xdb\x1e\xd2\xa5\x5d\x65\xfd\x74\x8d\x7a\x13\x8e\xab\x80\x55\x71\x94\xdf\xee\x12\xe2\x1f\x71\x5b\x2a\xf6\xbf\xf2\x5a\xc3\xf8\x2c\x4e\x7b\x69\xb8\xbd\xca\xba\x96\xa7\xd1\x53\x97\xad\xef\x39\xde\xf8\x24\xa0\x7b\x78\xc4\x5e\x27\x8b\xfd\xd5\x63\x60\x33\xb2\x96\x74\x21\x07\x6b\x7e\x42\x63\xa2\x81\xe9\xda\x5c\xc4\xc1\xd8\x5c\xee\xbd\xd8\xca\xe9\x3a\x8d\x99\xa3\x91\xcd\x88\xb9\x71\x96\x38\x80\xb1\x3c\x27\x09\x29\xe7\x5f\xbf\x2d\x2d\x12\x42\x80\xa1\x59\xa4\x61\xac\xc6\xf1\x6f\x9c\x35\xf0\xf3\xde\x5e\x30\xbc\x8a\x3b\xb0\x90\x63\x85\x96\x3f\x9b\x81\x2c\x47\x5f\x4c\x09\x36\xaf\x21\x9f\x25\x2d\x6d\x88\xb1\x9a\x83\x3c\xbb\x1d\xe6\x40\xfb\x5d\xf8\x74\x2b\xec\x6f\x0a\xa3\x37\xee\x08\x83\x4c\x8d\xae\x13\xbc\x41\x18\xfb\x69\x6c\x1f\x33\xd4\x5a\x9b\x12\x6f\x9c\x65\xb4\x3c\xa9\xd6\xa8\x90\xb4\x07\x5a\x27\x2d\xa0\xbc\xdf\x5f\xd6\x9b\x15\x4c\x0d\x86\x3a\x94\x07\x93\x38\x50\xde\x66\x42\x9d\xa0\x6f\x35\x78\x8b\xff\x92\xa7\x46\xad\x63\x8e\xc3\x06\x03\xeb\x5c\xfe\x1a\x7e\x9d\xcf\x96\x6c\x7e\x47\x9e\xf8\x62\x39\xc9\x0f\xac\x0a\xd4\xf7\x67\x63\x16\xbf\x06\xb2\x6d\xf6\x3b\x00\x00\xff\xff\xbe\x6f\x55\xf4\xfa\x0d\x00\x00")

func v1alphaExtractorExtractorSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_v1alphaExtractorExtractorSwaggerJson,
		"v1alpha/extractor/extractor.swagger.json",
	)
}

func v1alphaExtractorExtractorSwaggerJson() (*asset, error) {
	bytes, err := v1alphaExtractorExtractorSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1alpha/extractor/extractor.swagger.json", size: 3578, mode: os.FileMode(420), modTime: time.Unix(1574222638, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v1alphaSchemaSchemaSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xce\x31\xae\xc3\x20\x0c\x06\xe0\x9d\x53\x58\x9e\x9f\x92\xd7\x8e\xb9\x4a\xd5\x01\x11\x12\xa8\x52\x8c\xb0\x93\x0e\x51\xee\x5e\x19\x14\x55\xdd\x3a\x61\xd9\xdf\xff\x8b\xdd\x00\x20\xbf\xec\x3c\xfb\x82\x03\xe0\xb5\xfb\xc7\x3f\xdd\xc5\x34\x11\x0e\xa0\x77\x00\x94\x28\x8b\xd7\xfb\x76\xb1\x4b\x0e\xb6\x67\x17\xfc\xf3\x7c\xba\x5c\x48\xa8\xe6\x00\x70\xf3\x85\x23\xa5\xaa\xdb\x08\x89\x04\xd8\x0b\x1a\x80\xa3\xb6\xd7\x9c\x67\x1c\xe0\xd6\x42\x41\x24\x9f\x05\x3a\xb3\xda\x7b\xb5\x8e\x12\xaf\x5f\xd8\xe6\xbc\x44\x67\x25\x52\xea\x1f\x4c\xe9\x63\x73\xa1\x71\x75\x3f\x5a\x2b\x41\xe1\xde\xbe\x34\xfa\x29\xa6\xa8\xae\x2d\xcd\x61\xde\x01\x00\x00\xff\xff\x1c\xe4\x75\xba\x1d\x01\x00\x00")

func v1alphaSchemaSchemaSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_v1alphaSchemaSchemaSwaggerJson,
		"v1alpha/schema/schema.swagger.json",
	)
}

func v1alphaSchemaSchemaSwaggerJson() (*asset, error) {
	bytes, err := v1alphaSchemaSchemaSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1alpha/schema/schema.swagger.json", size: 285, mode: os.FileMode(420), modTime: time.Unix(1574222638, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v1alphaDepsDepsSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xce\xc1\x0a\xc2\x30\x0c\x06\xe0\x7b\x9f\x22\xe4\x2c\x9b\x7a\xdc\xab\x88\x87\xd2\x65\x6b\x65\x36\xa5\xc9\xe6\x61\xec\xdd\xa5\x2d\x22\xde\xbc\x94\x90\x7c\xff\x4f\x77\x03\x80\xf2\xb2\xf3\x4c\x19\x07\xc0\x6b\x77\xc6\x53\xd9\x85\x38\x31\x0e\x50\xee\x00\xa8\x41\x17\x2a\xf7\xed\x62\x97\xe4\x6d\x3f\x52\x92\xfa\x74\x29\xb3\x72\xcd\x00\xe0\x46\x59\x02\xc7\x2a\xdb\x08\x91\x15\x84\x14\x0d\xc0\x51\x9b\xc5\x79\x7a\x92\xe0\x00\xb7\x16\xf2\xaa\xe9\x53\x50\x66\x29\xf6\x5e\xad\xe3\x28\xeb\x0f\xb6\x29\x2d\xc1\x59\x0d\x1c\xfb\x87\x70\xfc\xda\x94\x79\x5c\xdd\x9f\xd6\xaa\x2f\x70\x6f\x5f\x1a\x69\x0a\x31\x14\xd7\x96\xe6\x30\xef\x00\x00\x00\xff\xff\x08\xa8\x20\x09\x19\x01\x00\x00")

func v1alphaDepsDepsSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_v1alphaDepsDepsSwaggerJson,
		"v1alpha/deps/deps.swagger.json",
	)
}

func v1alphaDepsDepsSwaggerJson() (*asset, error) {
	bytes, err := v1alphaDepsDepsSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1alpha/deps/deps.swagger.json", size: 281, mode: os.FileMode(420), modTime: time.Unix(1574222638, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v1alphaStoreStoreSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x4d\x8f\xda\x30\x10\xbd\xe7\x57\x58\x6e\x8f\x08\x0a\x47\x6e\x95\xa0\x55\xab\xaa\x45\xb4\x52\x0f\x2b\x0e\x26\x99\x24\x66\x13\x8f\x65\x4f\x58\x21\x94\xff\xbe\xb2\xb3\xf9\x00\xc2\x6e\x56\xb0\x97\xc8\x9a\x79\xef\x79\xde\x8c\x33\xc7\x80\x31\x6e\x9f\x44\x92\x80\xe1\x73\xc6\x67\xe3\x2f\x7c\xe4\x62\x52\xc5\xc8\xe7\xcc\xe5\x19\xe3\x24\x29\x03\x97\xdf\x4f\x45\xa6\x53\x31\xb1\x84\x06\xaa\xef\x58\x1b\x24\xf4\x2c\xc6\xf8\x1e\x8c\x95\xa8\x3c\xb6\x3a\x32\x85\xc4\x2c\x10\x0f\x18\x2b\xbd\xb6\x0d\x53\xc8\xc1\xf2\x39\x7b\xa8\x48\x29\x91\xae\x05\xdc\xd9\x3a\xec\xc6\x63\x43\x54\xb6\x38\x01\x0b\xad\x33\x19\x0a\x92\xa8\x26\x3b\x8b\xaa\xc5\x6a\x83\x51\x11\x0e\xc4\x0a\x4a\x1d\xf0\x58\x95\x14\x41\x2c\x95\x74\x38\xdb\xba\xf6\xf6\x16\x90\x01\xc1\x1a\xac\x46\x65\xa1\x49\xba\xa6\x1c\xb4\xef\x09\x6e\x77\x10\x7a\x7b\x2f\x06\x6b\xea\x37\xa9\xa2\x01\xc4\x51\x1d\xd7\x06\x35\x18\x92\x60\x3b\x68\x5f\xab\x34\xa7\xa1\x8e\x88\x30\x46\x1c\x1a\x0d\x9f\x92\x04\xf9\x39\x9e\x31\xfe\xd9\x40\xec\x18\x9f\x26\x1d\xb7\xd5\x0c\xbf\x1b\xa1\xd3\x1f\x04\xf9\x4a\x48\xc3\x3b\xb4\x32\x38\x3f\x95\x97\x3e\x1b\xf6\x2d\x26\x93\x5a\xe4\x5f\x45\xec\x35\x6b\xc9\x48\x95\xb4\x05\x96\xad\x6f\xfe\x38\x7d\x83\x75\xd2\xa3\x18\x4d\x2e\xc8\x65\xb7\x07\x82\x2b\x8a\xb3\x7b\x2b\x82\x0a\x31\x72\xcc\x33\xdd\x21\xa3\x59\xd6\xdc\x5e\xe5\xa6\x7d\x0b\x41\xe2\x2e\x65\x0f\x98\xf7\xf2\xd2\xce\xb5\xbb\x38\xa8\x22\x6f\x7e\x4c\x1f\x59\x7f\xfd\xdf\x29\x85\xff\xfc\xfb\xe7\x77\x7d\xfd\xa6\xa1\x45\x10\x8b\x22\xf3\xf5\x39\xfc\x2b\xc5\xf8\xa7\x7b\xc3\x03\x84\x28\xb9\x78\x77\x43\xe6\xd2\x3f\x0f\x85\xd1\xad\x6a\xd7\x07\xf0\x4b\x5a\xba\xc7\x62\xe9\x5b\x14\x1f\xb4\x58\xde\xbd\x54\x56\xc5\x10\x8b\x41\xcd\x2e\x83\x32\x78\x0e\x00\x00\xff\xff\x9b\x51\xf8\x1c\xcd\x06\x00\x00")

func v1alphaStoreStoreSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_v1alphaStoreStoreSwaggerJson,
		"v1alpha/store/store.swagger.json",
	)
}

func v1alphaStoreStoreSwaggerJson() (*asset, error) {
	bytes, err := v1alphaStoreStoreSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1alpha/store/store.swagger.json", size: 1741, mode: os.FileMode(420), modTime: time.Unix(1574222638, 0)}
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
	"v1alpha/tracker/tracker.swagger.json":     v1alphaTrackerTrackerSwaggerJson,
	"v1alpha/extractor/extractor.swagger.json": v1alphaExtractorExtractorSwaggerJson,
	"v1alpha/schema/schema.swagger.json":       v1alphaSchemaSchemaSwaggerJson,
	"v1alpha/deps/deps.swagger.json":           v1alphaDepsDepsSwaggerJson,
	"v1alpha/store/store.swagger.json":         v1alphaStoreStoreSwaggerJson,
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
	"v1alpha": &bintree{nil, map[string]*bintree{
		"deps": &bintree{nil, map[string]*bintree{
			"deps.swagger.json": &bintree{v1alphaDepsDepsSwaggerJson, map[string]*bintree{}},
		}},
		"extractor": &bintree{nil, map[string]*bintree{
			"extractor.swagger.json": &bintree{v1alphaExtractorExtractorSwaggerJson, map[string]*bintree{}},
		}},
		"schema": &bintree{nil, map[string]*bintree{
			"schema.swagger.json": &bintree{v1alphaSchemaSchemaSwaggerJson, map[string]*bintree{}},
		}},
		"store": &bintree{nil, map[string]*bintree{
			"store.swagger.json": &bintree{v1alphaStoreStoreSwaggerJson, map[string]*bintree{}},
		}},
		"tracker": &bintree{nil, map[string]*bintree{
			"tracker.swagger.json": &bintree{v1alphaTrackerTrackerSwaggerJson, map[string]*bintree{}},
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
