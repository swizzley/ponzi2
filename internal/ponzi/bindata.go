// Code generated by go-bindata.
// sources:
// data/addButton.ply
// data/refreshButton.ply
// data/removeButton.ply
// data/roundedCornerNE.ply
// data/roundedCornerNW.ply
// data/roundedCornerSE.ply
// data/roundedCornerSW.ply
// DO NOT EDIT!

package ponzi

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

var _addbuttonPly = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x91\xdd\x6e\x9c\x30\x10\x85\xef\xe7\x29\xce\x5d\x5a\xa9\x8b\xfc\x13\xef\xb2\xbd\x4c\x1f\x24\xf2\xe2\x21\x41\x62\xcd\xca\x98\x6e\xe8\xd3\x57\xfc\x04\x10\x69\xa9\x48\x91\x90\x47\xc7\x73\x3e\xce\x0c\xb7\xb2\xa5\xbc\x0a\x57\x1b\x61\xeb\xac\x28\x20\x13\x41\x59\x75\xbd\xb2\x8f\xf8\x11\xd8\x46\x76\xb8\xb4\x78\x2a\xd9\x3b\x0e\x50\xc9\x29\xc5\x97\xba\xb9\x40\x7c\xc5\x01\xf7\xfb\x3d\xb9\x0c\x57\x49\x15\x5e\xbe\xa1\xae\x9a\x90\x31\xf2\xa2\xe4\xef\x78\xb0\xce\x3d\x35\x31\x56\x7e\x68\x7a\x20\x2e\xb9\x27\xff\xe4\x10\xf9\x0d\x52\xd1\x2d\x54\x37\x0e\xb1\x45\x5e\x56\x36\xe2\x6d\x2d\xb4\x6b\xe1\xd7\x5a\xf0\x1f\x3c\xfe\x83\xc9\x2f\x5c\x4d\xf6\x6a\x03\x02\xbb\xb5\xf4\x12\x98\xfd\x5a\xbc\x94\x0d\x4f\xb1\x73\x9b\x31\xa4\x98\x7b\xca\xa2\x8e\x63\x63\x53\x4c\x73\x3d\x17\xde\x15\x19\xd7\xc4\xde\x3d\xbf\xb2\x75\x1c\x48\x24\xd2\x08\x21\x04\x0e\x53\x25\x12\x21\xde\xa5\xb1\x9a\x0a\xf9\x5e\xa8\xc7\x23\x94\x31\xdd\x49\x0b\xe7\x3e\x44\x3a\x22\x52\x3a\x7c\x9a\x31\xc7\x38\x2c\x67\x39\xae\x1c\x7f\xa0\xcd\x10\x23\x07\x88\x91\xf4\x69\x46\x17\xa4\x7b\xf5\x2a\xc8\x7f\x4c\x33\x7d\x7f\x1f\xe3\x3c\x32\xce\x0b\xc6\xee\x20\xa7\x11\x72\xa2\x8d\x1c\xdb\x4b\x55\xe3\x52\x15\x6d\xc5\xd8\x5e\xea\x1c\x63\x69\xfd\xeb\x8f\xf9\xd7\x24\x13\xc1\x9c\xbb\x67\x37\x41\x43\x40\x42\xf5\xa7\xc6\x23\x69\x18\x1c\x31\xe8\x06\x9a\x34\x52\x08\x9c\x49\x43\x76\x8d\x52\x76\x15\x04\xd2\xbe\x43\xc1\x90\x86\xea\x2e\x44\x6f\x55\x38\xd2\xef\x00\x00\x00\xff\xff\x66\xb5\x46\x72\xea\x04\x00\x00")

func addbuttonPlyBytes() ([]byte, error) {
	return bindataRead(
		_addbuttonPly,
		"addButton.ply",
	)
}

func addbuttonPly() (*asset, error) {
	bytes, err := addbuttonPlyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "addButton.ply", size: 1258, mode: os.FileMode(438), modTime: time.Unix(1504839318, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _refreshbuttonPly = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\xdb\x6e\xdb\x4a\x0c\x7c\xdf\xaf\xe0\x5b\xce\x01\x6a\x81\xb7\xbd\xf5\x31\xfd\x90\x40\xb6\xd7\x89\x00\x45\x0a\x56\x52\x13\xf7\xeb\x0b\xad\x94\x34\x70\x0c\xd4\x50\x0d\x18\xa6\x67\x39\x23\x0e\xb9\xe2\x4b\x7b\x36\xa7\x3e\x3f\xd7\x23\xd4\xc3\xa1\x69\x80\x2a\x34\x87\xfe\xf9\x39\x75\x23\xfc\xc8\xa9\x1e\xd3\x11\xf6\x67\xb8\x6f\x53\x77\x4c\x19\xb8\xf2\x01\xfe\x1b\xa6\x3d\xe0\xff\xb0\x83\xd7\xd7\xd7\x6a\xbf\x1c\x55\x7d\x7e\xfc\x06\x43\x3f\xe5\x43\x82\x53\xd3\xa6\xef\x70\x97\xd3\x29\xa7\xe1\xe9\x7e\x1a\xc7\xbe\x5b\x12\xef\x4c\x6a\x53\x51\xff\x99\xf2\x98\xde\x80\xa2\x79\xc9\xfd\x4b\xca\xe3\x19\x4e\x6d\x5f\x8f\xf0\x76\x09\x9c\x2f\x81\x5f\x97\x40\xf7\x85\xd3\x7d\x21\x75\x9f\x58\xd3\xe1\xa9\xce\x90\xd3\xf1\x12\x7a\xcc\x29\x75\x97\xe0\xbe\x9d\xd2\x47\xd9\xa7\xfa\x90\x80\xfc\x9f\x9c\xb6\x19\xc6\x35\x71\x6a\x3e\x7c\x3d\x34\xdd\xb1\x39\xa4\xc1\xa4\xee\xf8\xf0\x94\xea\x63\xca\x66\x87\x95\x44\x46\x8d\x80\x95\x62\x74\xa8\x80\x15\x96\x0f\xec\xae\x44\xf4\x1e\xb0\xb5\xef\x5f\x53\x4e\x15\x1d\x60\x25\xa2\xc1\xfb\x5b\x35\x34\x2c\x1a\x1a\x3e\x6b\x58\x2f\x51\xe4\xe6\x3a\x78\xad\x83\x8b\x19\xe1\x28\x76\x49\x94\xa0\x7a\x95\xfc\x17\x3b\xbb\xb9\x15\xaa\x91\x61\x89\x9c\xb8\x8d\x32\x2c\x1e\x91\xe6\x44\x56\x1b\x54\xb6\xc8\x94\xb9\x78\xf5\xa5\x18\xf2\xe8\xc3\x36\x15\x51\xcf\xab\x11\x52\x96\xb0\x69\xce\xac\x12\x6d\xb1\xc1\x56\x68\xe9\xd0\x86\xbe\x58\xe7\x22\xda\x52\x81\x45\x47\xdb\x0c\xcd\x90\x2d\x5c\xeb\xa3\xc6\x4d\x2a\xb3\x11\x8e\xce\x12\xcc\xc3\xb2\x44\xb4\xa9\x2d\xe2\x88\x74\xa6\x2a\x47\x17\xf8\x66\x0d\x5a\x35\xa8\xd8\xa1\x68\xcb\x7d\x15\x8d\xc4\x7e\xe3\x5d\xb1\x5e\x2c\xce\x2f\x22\x33\xe9\xc6\x21\x07\xd4\x30\xbf\x88\x4c\x8a\xe2\x36\x69\xd8\xa0\x0e\x97\xfb\x4e\x6c\xc5\xfe\xdb\x56\xf1\x5e\x39\xc4\x9b\x35\x64\xd5\x90\xcf\x1a\xc4\x82\x78\xb5\x8e\x6b\x8b\xc9\xaf\x8b\xc9\x1b\x01\x04\x02\x36\x02\x02\x0a\xd6\x08\x38\xf0\x10\xca\xff\x08\x6a\x04\x08\xc1\x96\x00\x81\x08\xc8\x08\x30\x10\x10\x2f\x47\x24\x85\x14\x41\x80\x68\x81\x5c\xa1\xc7\x39\x1b\x17\x24\x00\xc9\x1c\x29\x90\x05\x5f\x9e\x41\xae\x04\xa4\xe0\x81\xdc\x22\xca\xf3\x96\x97\xf2\x0b\x14\xcc\xef\x00\x00\x00\xff\xff\xbc\x8a\x8c\x1c\x2a\x07\x00\x00")

func refreshbuttonPlyBytes() ([]byte, error) {
	return bindataRead(
		_refreshbuttonPly,
		"refreshButton.ply",
	)
}

func refreshbuttonPly() (*asset, error) {
	bytes, err := refreshbuttonPlyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "refreshButton.ply", size: 1834, mode: os.FileMode(438), modTime: time.Unix(1507493357, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _removebuttonPly = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\xdd\x8e\x9b\x30\x14\x84\xef\xcf\x53\xcc\xdd\xb6\x52\x41\xfe\x59\x27\xa4\x97\xdb\x07\x59\x11\x38\xec\x22\x81\x1d\x19\xb3\x09\x7d\xfa\x8a\x9f\x92\x94\x24\x08\x15\x09\xd9\x1a\x9f\xf9\x74\x66\x4e\x55\x47\x85\xf3\x75\x1a\x90\x36\x59\x59\x42\xc6\x82\x32\x57\xd7\x6c\x03\x7e\x79\x4e\x03\xe7\x38\x76\x78\xab\xd8\xe6\xec\xa1\xe2\x7d\x82\x6f\x4d\x7b\x84\xf8\x8e\x08\xe7\xf3\x39\x3e\x8e\x4f\xb1\xf3\x1f\x3f\xd0\xb8\xd6\x67\x8c\xa2\xac\xf8\x27\x5e\x3c\xd7\xee\x8b\xdf\xda\x10\x9c\x1d\xe7\x5e\x88\x2b\x1e\xe0\x5f\xec\x03\x5f\x20\x15\x9d\xbc\x3b\xb1\x0f\x1d\x8a\xca\xa5\x01\x97\xa5\xd0\x2d\x85\xdf\x4b\xc1\xde\x79\xec\x9d\xc9\xde\xb8\xda\xec\x33\xf5\xf0\x9c\x2f\xa5\x0f\xcf\x6c\x97\xe2\xb1\x6a\x79\x5e\xbb\x48\x33\x86\x14\xd7\x99\xaa\x6c\xc2\x34\xd8\x96\x73\xae\xf7\xd2\xe6\x65\xc6\x0d\xb1\xcd\xdf\x3f\x39\xcd\xd9\x93\x88\xc5\xf0\x21\x12\xb1\x92\x4a\x6a\x85\x59\xba\xbf\xc8\xbf\x17\xf5\xba\x83\x32\xa6\x3f\x69\x8b\xf1\x11\x21\x99\x08\x09\xdd\x8c\xfd\xe7\x0e\x91\x88\x8d\x16\x5a\x0f\x39\xb4\x4c\xe4\x21\x79\x60\x8d\x1e\x40\x8c\x1c\x21\x46\xd2\xd5\x7a\xc5\x6d\x83\xf4\x9b\xf4\xbf\x1e\x36\x79\x9a\x21\xda\x9e\x66\xce\xb0\xdc\x63\x9d\x71\x98\x18\x07\xba\x8d\xb0\x2c\x64\x9d\xb1\x9f\x18\x7b\x5a\x2b\x35\x5a\x6f\x55\x4d\xad\x2a\x5a\x2b\x75\x1d\xf2\xcf\x26\x4f\x0b\xd9\x8e\x78\xda\xc7\x56\x84\x86\x80\x84\x1a\x4e\x8d\x57\xd2\x30\xd8\x61\xd4\x0d\x34\x69\x24\x10\x38\x90\x86\xec\x07\xa5\xec\x6f\x10\x48\x86\x09\x05\x43\x1a\xaa\x7f\x10\x83\x55\x61\x47\x7f\x02\x00\x00\xff\xff\x91\x82\xfd\x33\xec\x04\x00\x00")

func removebuttonPlyBytes() ([]byte, error) {
	return bindataRead(
		_removebuttonPly,
		"removeButton.ply",
	)
}

func removebuttonPly() (*asset, error) {
	bytes, err := removebuttonPlyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "removeButton.ply", size: 1260, mode: os.FileMode(438), modTime: time.Unix(1504846062, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _roundedcornernePly = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x8f\xdd\x4a\xc4\x30\x10\x85\xef\x07\xe6\x1d\xe6\x05\x36\xe4\xaf\xae\x5e\x8b\xb7\xbe\x82\xc4\x66\x76\xb7\xd0\x26\x65\x4c\xa5\xf5\xe9\x85\xa2\xb5\x10\xbd\xa8\x85\xc2\xe1\x9b\x73\x4e\x66\xc6\x7e\x41\xb8\x64\x19\x42\xa1\xf0\xd6\x76\x1d\x19\xa5\x11\xda\x3c\x0c\x9c\x0a\x49\x9e\x52\xe4\xf8\x98\x25\xb1\x3c\x3f\xa9\xd5\xce\x3d\xaf\xc3\x77\x96\xc2\x33\x35\x08\xa3\xe4\x91\xa5\x2c\x74\xe9\x73\x28\x34\x57\x64\xa9\xc8\x47\x45\x52\x1d\x4b\x75\x2e\xed\x83\x53\x7b\x0b\x42\xc2\xb1\x62\x57\x61\x4e\x15\x7d\xed\x27\xfe\xd9\x9f\xe3\x95\xc9\xef\x4c\xdd\x76\x94\xf9\x95\x5a\x04\x4e\xf1\xe5\xc6\x21\xb2\x20\x18\xa5\xd7\x8f\x4e\x9b\xd2\x7f\x8b\xcd\x63\x9b\xe6\xfb\x47\xd0\xea\xde\x9f\xcf\xcd\x03\x9d\xb4\xb2\xce\xdf\x39\xf7\x9f\x0e\x6f\xbc\x35\x9e\x76\xe2\x60\xc5\xfe\xf5\xaf\x85\x8e\x77\x98\x6a\x7e\xfc\x14\x32\x08\x86\x2c\x82\x25\x87\xe0\xc8\x7f\x06\x00\x00\xff\xff\x0b\x4c\xbe\x91\xa1\x02\x00\x00")

func roundedcornernePlyBytes() ([]byte, error) {
	return bindataRead(
		_roundedcornernePly,
		"roundedCornerNE.ply",
	)
}

func roundedcornernePly() (*asset, error) {
	bytes, err := roundedcornernePlyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "roundedCornerNE.ply", size: 673, mode: os.FileMode(438), modTime: time.Unix(1504845949, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _roundedcornernwPly = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\xcf\x4e\xc4\x20\x18\xc4\xef\x3c\xc5\xf7\x02\x25\xfc\xab\xab\x67\xef\x5e\x3d\x1a\x2c\xb3\xbb\x4d\x5a\x68\x3e\xa9\xd9\xfa\xf4\xc6\x8d\x62\x43\xf5\xd0\x25\x21\x99\xfc\x98\x19\x3e\x98\x86\x45\x1c\x13\x8f\x3e\x93\x7f\xeb\xfa\x9e\xb4\x54\xa2\x4b\xe3\x88\x98\x89\xd3\x1c\x03\xc2\x63\xe2\x08\x7e\x7a\x96\x5f\x66\x0c\xb8\x9e\xbd\x83\x33\x2e\xd4\x8a\x89\xd3\x04\xce\x0b\x1d\x87\xe4\x33\x5d\x6a\xb0\xd4\xe0\xa3\x06\x71\x93\x89\x9b\x50\x5c\xa5\xe6\xee\xec\x99\x18\xa1\x46\x27\x06\x62\x0d\x5f\x87\x19\x65\x6a\x84\x13\xc8\xfd\x5a\xfa\xf2\x12\xfd\x17\x34\x02\x31\xbc\x9c\xe1\x03\x58\x68\xa9\xae\x8b\x8a\x50\xff\x8b\xe2\x31\x6d\xfb\xb3\x85\x92\xc6\xba\x3b\x6b\x49\xc9\x7b\x77\x38\xb4\x0f\xbb\x1b\x1a\x25\x9d\x76\x46\x3b\x5a\x89\xdd\x15\xdf\x97\x37\xab\x79\x76\x76\x14\xda\xdc\xfe\x19\xa4\x85\x26\x23\x0c\x59\x61\xc9\x7d\x06\x00\x00\xff\xff\x5e\x8a\x14\x69\x89\x02\x00\x00")

func roundedcornernwPlyBytes() ([]byte, error) {
	return bindataRead(
		_roundedcornernwPly,
		"roundedCornerNW.ply",
	)
}

func roundedcornernwPly() (*asset, error) {
	bytes, err := roundedcornernwPlyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "roundedCornerNW.ply", size: 649, mode: os.FileMode(438), modTime: time.Unix(1504845956, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _roundedcornersePly = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x8f\x6f\x4a\xc4\x30\x10\xc5\xbf\x0f\xcc\x1d\xe6\x02\x0d\xf9\x57\x57\x3f\x8b\x27\xf0\x00\x12\x9b\xd9\xdd\x42\x9b\x94\x31\x95\xad\xa7\x17\x8b\x76\x17\xa2\x1f\x6a\x20\x30\xfc\xe6\xbd\x97\x97\x69\x58\x10\x8e\x59\xc6\x50\x28\xbc\x75\x7d\x4f\x46\x69\x84\x2e\x8f\x23\xa7\x42\x92\xe7\x14\x39\x3e\x66\x49\x2c\xcf\x4f\x6a\x95\xf3\xc0\xeb\xf2\x9d\xa5\xf0\x85\x5a\x84\x49\xf2\xc4\x52\x16\x3a\x0e\x39\x14\xba\x54\x64\xa9\xc8\x47\x45\x52\x6d\x4b\xb5\x2f\xdd\x1a\xe7\xee\x1c\x84\x84\x63\xc5\x4e\xc2\x9c\x2a\xfa\x3a\xcc\x7c\xed\xcf\xf1\xc4\xe4\x6f\x44\xfd\xf6\x29\xf3\x2b\xb5\x08\x9c\xe2\xcb\x99\x43\x64\x41\x68\x8c\xd2\xeb\xa1\xeb\xa4\xff\x1e\x36\x8d\x6d\xdb\x9f\x8b\xd0\x68\x65\x9d\xbf\x73\x8e\x1a\xad\xee\xfd\xe1\xd0\x3e\xec\x0f\xd1\xca\x1b\x6f\x8d\xff\xca\x58\x27\xf7\x9f\x8c\xed\xf5\xef\x42\xbb\x23\x4c\xb5\xde\xdf\x82\x0c\x82\x21\x8b\x60\xc9\x21\x38\xf2\x9f\x01\x00\x00\xff\xff\x38\xf0\x26\xca\xa2\x02\x00\x00")

func roundedcornersePlyBytes() ([]byte, error) {
	return bindataRead(
		_roundedcornersePly,
		"roundedCornerSE.ply",
	)
}

func roundedcornersePly() (*asset, error) {
	bytes, err := roundedcornersePlyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "roundedCornerSE.ply", size: 674, mode: os.FileMode(438), modTime: time.Unix(1504845965, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _roundedcornerswPly = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x4d\x6a\xc3\x30\x10\x85\xf7\x03\x73\x87\xb9\x80\x8d\xfe\xdc\xb4\xeb\x1e\xa1\x8b\x2e\x8b\x6a\x4d\x12\x83\x2d\x99\xa9\x5c\xe2\x9e\xbe\x60\xa8\x62\x50\xb3\x48\x04\x82\xc7\xa7\xf7\x9e\x46\x9a\xc7\x15\xe1\x98\x64\xf2\x99\xfc\x57\x3f\x0c\xa4\x5b\x85\xd0\xa7\x69\xe2\x98\x49\xd2\x12\x03\x87\xd7\x24\x91\xe5\xed\xbd\xdd\xec\x3c\xf2\x76\xf8\xcd\x92\xf9\x42\x1d\xc2\x2c\x69\x66\xc9\x2b\x1d\xc7\xe4\x33\x5d\x2a\xb2\x56\xe4\xa7\x22\xb1\x8e\xc5\x3a\x17\xf7\xc1\xa5\x3f\x7b\x21\xe1\x50\xb1\x93\x30\xc7\x8a\x7e\x8e\x0b\x5f\xe7\xe7\x70\x62\x72\x3b\xd3\x50\x1e\xa5\xff\xa5\x06\x81\x63\xf8\x38\xb3\x0f\x2c\x08\x8d\x6e\xd5\xb6\xa8\x08\x75\x5b\x14\x8f\xe9\xba\xbf\x8d\xd0\xa8\xf6\xd9\x1d\x0e\xdd\x0b\xa9\xd6\x58\xf7\x64\xed\x43\x1d\x4e\x3b\xa3\x1d\x5d\xd5\xdd\x25\xe5\xfa\xfd\x44\x77\x76\x14\xda\x3c\xfe\x21\x8a\x34\x82\x26\x83\x60\xc8\x22\x58\x72\xbf\x01\x00\x00\xff\xff\x15\x20\x87\xfc\xa3\x02\x00\x00")

func roundedcornerswPlyBytes() ([]byte, error) {
	return bindataRead(
		_roundedcornerswPly,
		"roundedCornerSW.ply",
	)
}

func roundedcornerswPly() (*asset, error) {
	bytes, err := roundedcornerswPlyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "roundedCornerSW.ply", size: 675, mode: os.FileMode(438), modTime: time.Unix(1504845973, 0)}
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
	"addButton.ply": addbuttonPly,
	"refreshButton.ply": refreshbuttonPly,
	"removeButton.ply": removebuttonPly,
	"roundedCornerNE.ply": roundedcornernePly,
	"roundedCornerNW.ply": roundedcornernwPly,
	"roundedCornerSE.ply": roundedcornersePly,
	"roundedCornerSW.ply": roundedcornerswPly,
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
	"addButton.ply": &bintree{addbuttonPly, map[string]*bintree{}},
	"refreshButton.ply": &bintree{refreshbuttonPly, map[string]*bintree{}},
	"removeButton.ply": &bintree{removebuttonPly, map[string]*bintree{}},
	"roundedCornerNE.ply": &bintree{roundedcornernePly, map[string]*bintree{}},
	"roundedCornerNW.ply": &bintree{roundedcornernwPly, map[string]*bintree{}},
	"roundedCornerSE.ply": &bintree{roundedcornersePly, map[string]*bintree{}},
	"roundedCornerSW.ply": &bintree{roundedcornerswPly, map[string]*bintree{}},
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

