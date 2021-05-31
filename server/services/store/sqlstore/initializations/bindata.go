// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/templates.json
package initializations

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

var _templatesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x7c\xdf\x8e\x5c\xc7\x71\xfe\xab\x9c\xdf\xe8\xe2\x07\x04\x5b\x48\x77\x75\x75\x77\x35\xef\x28\xd1\x96\x09\xc8\x12\x6d\xd2\x71\x82\x58\x30\xfa\x4f\xf5\x72\xac\xd9\x99\xc5\xcc\x59\xd2\x0b\x81\x57\xb9\xb5\x63\xc7\x30\x60\x24\x32\x22\x23\x46\x02\xdf\xc5\x41\x72\x11\x24\xc8\x8d\x1f\xc5\x2f\x10\x3d\x42\xd0\xb3\x14\x39\xbb\x7b\x86\x6c\xee\x8a\xe4\xda\x09\xa4\x8b\xe1\x9c\x99\x3e\x67\xaa\xbe\xfa\xea\xab\xea\xea\xfd\x74\xf6\x48\xd6\x9b\xf9\x6a\x39\xbb\xa5\x0f\x66\x25\x8e\x32\xbb\xa5\x9d\x62\x83\x56\x05\x85\x5a\x1f\xcc\xd2\x62\x95\x3f\xd9\xcc\x6e\xfd\xe5\xa7\xb3\x79\x99\xdd\x9a\x61\x4a\xbe\x64\x53\x20\x1b\x17\x81\x44\x29\x60\x51\x15\xa2\x2b\x46\xc7\xac\xac\xc9\x7e\x76\x30\x3b\x8e\x6b\x59\x8e\x77\xdb\x37\x66\x07\xb3\xf5\x6a\x75\xf6\xba\xf3\xdb\x9b\xfc\x50\x8e\xe2\xf6\xa1\xc6\xd3\x63\x99\xdd\x9a\xa5\x55\x5c\x97\xd9\xc1\x6c\x9c\x8f\x8b\xf6\xef\x7b\xeb\xd5\x0f\x24\x8f\xc3\x83\xb8\xf9\x64\x33\x3b\x98\xd5\xb9\x2c\xca\x66\x76\xeb\xd3\x59\x8e\xeb\x72\x6f\xbd\x3a\x96\xf5\x38\x97\x9d\x07\x8f\xc1\x63\xc9\x3e\x82\xad\x94\x81\x6c\x41\x60\x45\x04\x9c\x91\xb3\x0b\x5e\xfb\xaa\x67\x07\xb3\x65\x3c\x6a\xeb\xdf\x1f\xe3\x78\xd2\x16\x5e\x1d\x8f\xf3\xd5\xf2\x6c\x9d\xbc\x5a\xac\xd6\xb3\x5b\xb3\xe3\xf5\xea\xf8\xbd\xf6\xfa\x8e\xd4\x78\xb2\x18\x67\x07\x67\xb7\x20\xf2\x92\x2b\x69\x28\xd5\x16\x20\x14\x0d\x1c\x62\x02\xaf\x49\x9c\xb7\x12\x9d\x6f\xb7\x78\x14\x17\x27\xed\x1e\x1f\xca\x0f\xc7\xe1\x3b\xc7\xb3\x27\x07\x13\x4b\xff\x85\x2c\x16\xab\xc7\x5f\xae\x5c\x8a\x4f\x26\xfa\x00\x92\xb0\x00\x05\x63\x81\x83\x0d\x80\x41\xd9\x10\x4b\xb0\xde\xc4\x9d\x95\xef\x2e\x87\x7b\xeb\xd5\xe1\x5a\x36\x9b\xe9\xd5\xdf\x5f\x8b\x2c\x77\x16\x8f\x09\x53\x81\x90\xbd\x03\x2a\x12\x80\x55\x65\xb0\xda\x8b\x76\xb5\x68\xb6\xbb\x8f\xfd\xde\xea\xe8\x78\x21\xa3\x94\xdf\xfd\xea\x77\xbf\xfa\xe2\xf3\xbf\xfd\xd1\xf4\x2d\xde\x5d\xaf\x1e\x3f\xbb\x85\xb3\x62\x73\x88\x06\x4c\x8a\x11\xa8\x6a\x0f\xac\x5d\x06\x6c\x37\x8e\xde\x14\x29\x61\xe7\x16\xb7\xd7\xf9\xe1\xfc\x91\x94\xd9\x93\x8f\x9f\x21\x60\x23\x0b\xc9\xe3\xf6\x5e\x67\x4f\x6d\x8a\x63\x4c\x15\x44\x79\x02\x0a\x25\x00\x97\x6a\xc1\x1b\x54\x01\x75\xc6\x82\xe6\xb9\x3f\xef\xad\xe7\xab\xf5\x7c\x3c\x7d\x99\x47\xbf\x2d\xe5\x99\x59\x4c\xaa\xc9\xaa\x0a\xd5\xba\x00\x94\x62\x01\x8e\x26\x42\x29\xda\x66\x53\x9d\xd2\x6a\xd7\x2c\xdf\x98\x1f\x3e\x1c\xbe\xf8\xfc\xe7\xff\xd8\xe3\x4f\xf6\xd5\x06\xcf\x04\x89\x6d\x00\xca\x48\xc0\x92\x04\xb4\xcf\xde\x39\x51\xac\x4b\xd9\x59\xfb\x9b\x52\xe6\x27\x47\xfb\x5c\x19\x4f\xbf\x5c\x36\x70\xb4\xde\xa1\x07\x55\x9b\x27\xbd\x2e\xc0\x56\x15\x08\xba\x1a\x29\xa1\x16\xd4\x66\x67\xd9\x0f\x56\x8f\x5f\x64\x61\x8c\xb6\x44\x83\x0a\xbc\xb1\x19\x48\x05\x03\xec\xd9\x37\x7b\x88\xb6\xb9\x54\xd9\x1a\xeb\xa9\x85\xef\xc4\x51\x86\xf7\xd6\x12\xc7\xed\xbb\xcf\xad\xfc\xfc\x06\xf9\xec\xea\x83\xf9\x91\x6c\xef\x5b\x64\x93\xd7\xf3\xed\x07\xcf\x78\x62\x9e\xb7\xaf\xbe\xf8\xfc\xaf\xff\xb9\xfd\x6b\xf3\x40\x8e\x8e\x17\x5b\x56\x1a\xd7\x27\xf2\xe4\xe0\xe9\x0a\xb7\xc7\x2d\x4d\x79\x87\xda\x3b\x6d\x0d\x1e\xcc\x4e\x8e\xcb\xb9\x0b\x88\x8c\x14\x5c\xbb\x47\xc3\x6a\xbb\xa0\x9e\xfd\x30\x53\xa3\x45\x25\x09\x8c\xca\x05\x88\x2d\x02\x63\x88\x90\xa2\x49\xd9\x0b\xb3\xa0\xdb\xcf\x61\x9d\xdf\x7e\x39\x87\x7d\x53\x64\x9c\x2f\x0f\x87\x0f\x57\xa3\xf4\x71\x98\xcf\xa8\x51\x3c\x43\x30\x64\x81\xb2\x53\xc0\x3a\x59\x50\x49\x89\xf1\x59\xc8\x99\xfa\xdc\x23\x0f\xda\x6d\x5f\x82\xf7\xf3\xb0\x34\xba\x44\xab\x72\x84\xaa\x63\x00\x2a\xa8\x81\x9d\x71\xa0\x1b\x69\xb3\xcf\x3a\xe2\x2e\x7e\x6e\x97\xe1\x1b\xab\xbc\x27\xfc\xdb\x47\xbe\x8c\x24\xa9\xce\x68\x9f\x41\xb2\xd6\x40\x5a\xb5\x48\x72\x09\x24\x92\xd3\x06\x55\x35\x01\x77\x96\xbd\x3f\xc6\x65\x39\xd9\xc3\x8b\xf7\x4e\xd6\xc7\x8b\x67\x2b\x7b\xa5\x2a\x9b\xca\xe0\x22\x69\x20\x97\x0b\xb0\x17\x04\x6b\x44\x15\x95\x34\xe6\x6d\x3e\xf9\x72\xe5\xef\x8a\x7c\xb2\x38\x1d\xee\x9f\x2e\xf3\x8b\x80\xaf\x4d\x41\x13\x28\x82\x24\x2b\x40\x15\xb1\x65\x09\x04\x67\xb5\x95\x4c\x3a\x53\xa4\x9d\x54\x71\x72\x74\x14\xd7\xa7\x7b\x30\x3f\xca\x0f\xc7\x97\x80\xfd\x17\x3f\xfb\xef\x7f\xff\x49\x0f\xde\xbd\xf6\xda\xb9\xe0\xdc\x45\xbc\x7b\xed\x8d\x33\x81\xf5\x34\xde\x5d\x12\x13\x72\xd6\x50\xb1\x05\xb2\x4f\x15\xb8\x50\x80\xea\x84\x22\xab\x5c\xd9\xe3\x7e\xbc\x77\x7e\xbb\x23\x67\xcb\x7a\xb3\x5a\xc6\xc5\xf0\xfe\x2a\x2e\x3a\x93\x76\x75\x35\xa7\xc4\x90\xa3\x75\x40\xc9\x1b\x60\xd3\x42\xcf\x93\xf1\x29\x44\xe7\x7c\x79\xd5\xa4\xbd\x43\xf1\xa9\x5a\x4c\x55\x1c\xc4\xdc\x34\x41\x20\x06\x46\x5d\xdb\x3d\x62\xd4\xd9\x27\x45\xbb\x69\xf5\xc1\x6a\xb8\xb3\xea\xa1\x77\xef\xb3\x35\x21\x56\x30\x2a\x64\xa0\x92\x34\xb0\xc1\x00\x05\x95\x54\x2f\x41\x62\xde\xa5\xf7\x3b\xab\xf9\xf2\xb0\x23\x51\x07\x4e\x45\x62\xa3\xf7\x9c\x05\x28\x17\x01\x2e\xd9\x01\x6a\x8c\xa5\x18\x27\xee\x9c\x0a\xb8\xb3\x5a\xca\x70\x96\xa0\x5f\x90\x46\x83\x47\x5b\x34\x41\xb1\x91\x81\x58\x2c\x70\x11\x0d\xae\x72\x56\x8a\x62\x70\xac\x9e\x5b\xf8\xbd\x38\xca\xe1\x6a\xfd\xd2\x34\x7a\x3e\x4a\x0d\x92\x8d\x06\x0b\x54\xc7\x0c\xe4\x4c\x02\xf6\xb5\xa5\x3c\xaf\x09\xb3\x4e\x26\xec\x46\xe9\x07\xf3\x2a\xc3\xfd\x4f\xe6\x8b\x45\x8f\x7a\x61\x95\x64\xab\x25\x6a\x13\x77\x14\x90\x81\x89\x03\x70\x52\xa8\x59\x55\x0a\x96\x76\x16\xff\xfa\x7c\x19\x97\x59\xa6\x17\xfe\x68\x1d\x97\x87\xcf\x1e\xbb\xd6\x64\x6a\xb0\x1a\x12\xf9\x0a\xa4\xdb\x63\x57\x5d\x00\x0d\x3b\xe7\x3c\xb2\x52\xbc\x2b\x00\x24\x2e\xc6\x87\x2f\xb4\xb5\x4b\x1a\x29\x24\x48\x59\x33\x90\xad\x19\x98\x83\x40\x92\x4c\x5c\xb3\xb0\x92\x1d\xfa\xbe\x73\x22\x43\x4b\xaa\xaf\x28\x42\x43\x54\x41\x89\x31\x90\xb4\x56\x40\xe8\x18\x38\xa8\x00\xda\x60\xb6\x4a\x61\x0e\x4a\x76\x9e\xfa\x5b\x7a\xda\x14\x17\x56\x55\x91\x31\x78\x5f\xa1\x05\x0c\x90\xf5\x09\x98\x74\x02\xc1\xe4\xab\x77\x35\x59\xb4\xbb\xab\x62\xd7\xaa\x4e\xbc\x36\x41\x08\xac\xb1\x0d\x18\x29\x01\x67\xa5\xc1\x27\x85\x21\x5a\x9f\x58\xed\x02\xfa\x5b\xa6\x6b\xd5\x62\x8d\xd7\xd9\x19\x70\xae\x3d\x2b\x35\x0b\x78\xc3\x90\xa9\x64\x4a\x12\x89\x69\x37\x8b\x7d\x8b\xa6\x7c\xf6\x02\xc6\xfe\xfd\x67\xbf\xed\x26\x6c\x8b\x9a\xd1\xab\xcb\x84\x6d\xdd\xf6\xff\x69\xc2\x0e\x86\x6a\xc0\xe4\x21\xd4\xa2\x81\x52\x53\x5e\x35\x6b\x50\x44\x31\x95\x28\x39\xe7\xb0\x9f\xb0\x3b\xbf\xfd\x0a\x84\xdd\x5f\x65\x15\xef\x7d\x13\x42\xc0\x1e\x19\x48\x15\x0d\xec\xa3\x03\x1b\x48\xb9\x94\x52\xaa\x49\xbd\x2a\x61\xdf\x9b\x2f\x3f\x79\x46\x25\x24\x29\x64\xb4\x50\x6c\xf3\xed\x56\xf5\xdb\x20\x50\x3d\xd5\x56\x4e\x1a\x25\xfe\x2a\x8c\x5d\x8c\x8f\x4e\x57\x82\x6a\x0c\x36\xc6\x6e\x04\x82\x05\xac\xa8\x18\x74\x8c\xa8\xe4\x2a\x8c\x5d\x62\x8a\xa5\x84\x04\xb1\x54\x0d\x54\x42\x05\xf6\xaa\x55\x9e\x36\x66\x97\xa5\x3a\x85\xbd\x8c\xfd\x22\x44\xfe\xf2\xe7\x9d\x88\x74\xa8\x82\xb1\x56\x5f\x92\x10\x0e\xb5\x09\xd6\xe2\x1e\x44\x16\x5f\xd9\x78\x2e\xc0\x7a\xab\x3c\x19\x81\xa9\x38\x88\xb5\x26\x49\xbe\x16\x5f\xe2\x7e\x44\x76\x7e\xfb\xe5\x88\xfc\xf6\x2a\x96\xa3\x78\xdc\x05\x45\xab\xb4\xf6\xc5\x22\xa4\x9c\x3d\x90\xb7\x0a\x18\x63\x04\x36\x3a\x1a\xab\x33\x51\x7c\x65\x28\x5e\x60\x1a\xce\xd6\xfa\xea\x02\x24\x25\x05\xc8\x49\x06\x36\xd1\x00\x4b\x8d\x29\x14\xf2\x52\x77\x59\xf1\xc3\xd5\x38\xdc\x1f\xe3\xba\x55\x48\x1d\x98\x94\xec\x0a\xa6\x6c\xa1\x54\x4c\x40\xd5\x7b\x60\xf2\x01\xc4\x86\x2c\x49\x99\x10\xc8\x5d\xb9\xe8\x67\x0a\xde\xb9\x14\xc1\xba\x68\x81\xb8\x05\x2a\xbb\x0a\xa8\x9d\xc7\x4a\xda\x9a\x60\x27\x8a\xfe\x97\xea\x09\x54\x5e\xfb\x58\x0c\x58\xdf\x14\x39\x69\x07\x6c\x6a\xd3\x13\xda\x98\x5a\xab\x29\x5a\x5f\xa7\x44\x21\x24\x89\x56\x0c\x84\x18\x35\x90\xf2\x16\x38\xdb\x0c\x4a\x27\xa2\xe4\x9a\x06\xda\x4d\xf7\x5f\x3b\x9e\xe7\xe1\xf7\x9f\xfd\xb6\xc3\x22\x4e\x24\x06\x97\x5b\x39\xee\x34\x50\x68\x69\xc3\x5a\x6a\xaa\xd3\x1a\x2f\x5c\x3d\xed\xa6\xfb\xc6\x89\xad\xde\xff\xcd\xf4\xd2\x3b\x2a\x53\xd7\x92\x92\xd5\x16\xa4\x14\x04\x8a\xd5\x02\xab\x9a\x81\x8c\x97\x82\xa8\x75\xa4\xdd\xe6\xc7\xbb\x27\x87\xc3\x17\x9f\xff\xf4\xef\x5f\x64\x65\xa7\x02\xdb\x4a\x0e\x8c\x34\x2b\xb3\x93\x26\x5c\x0d\x60\x0a\xec\x89\x94\x44\x9d\x77\xb0\x7d\xbc\x9e\x2f\xc7\x57\xc4\x76\x56\xda\x79\x97\x5b\x5d\xbb\xcd\xf8\x86\x80\x93\x58\xc8\x59\x10\xd9\x97\x18\x5d\xde\x2d\xda\xb6\xf7\x18\xfa\xd4\x84\x14\x8a\xd6\x90\x02\x72\xaa\xc5\x8d\x4e\xc0\x6c\x18\x30\x5b\x27\x2c\x62\x43\x95\xcb\x6b\xf7\x69\x0a\x4d\x1c\xd0\xb0\x02\x1d\x1b\x93\x63\x29\xc0\xca\x10\x44\x95\x25\x1a\x6c\x25\xba\xbb\xbc\xb6\x79\x91\xb5\xab\xaf\x26\x16\x42\x48\x46\x47\xa0\x98\x11\x58\x57\x05\xc8\xd2\x0a\x20\x9b\x0c\xc9\xb5\x5a\x4d\x39\xb1\xe4\x12\x33\x18\x4e\x02\x54\x8c\x03\xf6\x1a\x9b\x62\xb1\x9c\x33\x47\x96\x5d\x84\xdc\xd3\xdd\x8d\x26\x71\xd1\x57\x0c\x1e\x88\x9a\xb9\x3d\x1b\xe0\x98\x0c\x98\x58\x6d\x8a\x0e\x73\xd4\xbb\x4d\xac\x7b\x7b\x6c\xbc\xdb\x64\xca\x0e\xb5\x47\x89\x60\x4b\xf4\x40\xa5\x91\x9e\x66\x07\xc6\xa3\xf3\x85\x45\x42\xdc\x4d\xc1\xf7\x26\x4d\xfb\xe2\x82\xf8\x3f\xba\xb3\x19\x5a\xb4\x8a\xe9\x72\x36\x33\xa8\x50\x91\x9a\xce\x66\x4a\xf9\x1c\xb3\x73\x40\x31\x59\xa0\xd4\x52\x83\x4d\x08\xa8\x42\x09\x62\x6d\xcc\x2a\x9f\xcf\x66\x9d\x8d\xeb\xeb\x37\xba\xf3\xf9\x84\x77\x5f\xc6\x71\xbe\x3c\xdc\x0c\xdf\xf9\xf3\x73\x49\xef\x79\xa7\xec\xb3\x33\x5b\x1d\xef\x64\xc0\x4f\xfb\x9a\xdd\xdd\x6d\xe5\xae\x56\x6b\x67\x53\xf3\xc9\x54\x17\x2f\x04\xab\x69\xaa\x8b\xa7\xc8\x06\x15\xa6\x9d\xa8\xb3\xb1\xae\x36\xed\x87\x1c\x81\x92\x16\x60\xd7\x9e\xaa\x92\x84\x42\xd1\xeb\x6d\x26\xb8\x01\x4e\x6c\xa9\x62\xca\x7b\x13\x10\xbf\xa2\x23\x3b\x77\x1e\x5e\xb7\x23\x91\x94\x45\x9a\xd0\x96\x48\x5a\x2b\x34\xfb\x1c\x59\xc8\x50\x52\xe0\x6d\x74\x40\xde\x14\x60\x34\x02\xc1\x72\xcc\x81\x63\x75\xae\xbc\x25\x47\x3e\x9a\xcb\xe3\x1d\x47\xbe\x7b\x3a\xec\xb0\xfb\x79\x09\xfa\xd1\xba\xc8\xfa\xac\xf5\x97\x57\x8b\x93\xa3\xe5\x77\xe7\x65\x7c\xd8\x2e\x3f\x69\x1f\x5d\x8c\xed\xea\xa7\x4f\x5f\x3d\xed\x11\x36\x4f\xc7\xa7\x04\x18\x97\x4d\x14\xce\x0e\xd7\xab\x93\xe3\x77\x4f\xef\xbe\xc2\x16\xc7\xc3\x79\x29\xb2\xfc\x68\xcb\xa4\x77\xcb\xd3\xa5\x37\xab\xf5\xf8\xd1\x6e\x3f\xb2\xfd\x94\x07\xe7\x55\xf5\xa3\xf9\x66\x9e\x16\x72\xe1\x9b\x4f\xdf\x7d\x2a\xa9\x4f\xcf\xde\xef\x7b\x96\x8f\xf7\xb6\xe8\xcd\x54\x8b\xde\x1a\xb7\xaf\xde\x70\xd5\xd8\x62\x92\x01\x95\x8a\x07\x8a\x6a\xdb\x19\x20\x20\xcb\x2a\x60\x36\xae\xb2\xb9\x19\xc1\x7d\x27\x8e\x31\xc5\x8d\x0c\xf7\xcf\x3e\x39\xc9\xd2\x3f\xfb\xaf\xab\x52\xf4\x57\x1b\xd9\x5d\x7b\x5a\x93\x14\xdd\xea\x05\xe3\x27\xaa\xc6\x40\xde\x39\x9a\xf6\x62\x0a\x45\xd7\xdc\x88\x59\xe9\x26\xd6\xab\x03\xd6\x39\x41\x74\x2a\x29\x6d\x42\x4e\x7c\x43\x22\xfb\xf6\x62\x5f\x97\x63\x7f\x5c\xf7\x6d\x90\xdd\xd2\x3e\x1c\xcc\xbe\xff\xfd\xa7\x77\x42\x56\x07\x9d\xbe\xd7\x88\xbd\x8e\xd5\x5a\xbd\x12\xcb\xbc\x32\x65\x8c\x31\x6d\xfb\xb9\xaf\x42\x19\x9d\x3b\xee\x9d\x2c\xd7\x65\xec\x29\x02\x42\x74\x14\xdc\x14\x01\x19\xa7\x8d\xc5\x69\xe8\x66\x95\x8b\x64\xcd\x80\x1c\x0c\x50\x90\x00\x4c\x92\xc1\xa2\x2d\x49\x45\x4b\xc5\xa4\x9b\x01\xdd\x77\x4f\x87\x67\xcd\x8b\xd7\x9e\x92\x3a\x7d\xfa\x26\x52\xd2\xb4\xaf\x89\x7c\x98\x28\x07\x10\x2d\x7b\xf2\xd3\xbe\xae\x2e\x86\xa2\x25\x81\x33\x4e\x80\x6a\xb6\xc0\x46\x7b\xc8\x8c\x31\x78\xc3\xda\x79\x7b\x33\x92\xcd\xed\x7b\x77\x87\x0f\xe2\xa9\xac\xa7\xd3\xcc\x8f\xfe\xed\x66\xa4\x99\xab\x0b\x48\xe5\x94\x9b\xf4\x9f\x76\x9a\x70\x4f\x39\xd7\xb2\x4b\x40\xe7\x21\x6d\xc5\x02\x65\x0f\x6c\x10\x01\x91\x44\xe5\x8a\xe2\x1d\x9f\xf7\x5f\xe7\x1e\xfe\xf5\xf7\xfc\x2f\xf8\xef\xc3\xd5\xa3\x01\xcf\xc7\xe9\x6a\x39\xca\x72\xfc\x32\x54\x67\x5e\x54\x2c\x51\xb5\xaa\x94\x35\x50\x29\x11\x58\x15\x07\xd5\x5a\x65\xab\x64\x2e\x2e\xcd\x3e\xde\x1d\x95\xf8\xcf\x4b\x3e\xef\xda\xbf\xbe\x35\xdb\x76\xbf\x86\xc5\xfc\xf0\xe1\xf8\xff\x66\x07\x7d\xc3\x05\x9d\x9b\xf9\x97\x7d\xeb\xb5\x47\x34\xe8\xec\x05\xdf\x6e\x67\xcd\x58\x19\xb7\xa7\xca\x4b\xac\x55\xd6\x21\x80\xc4\xa6\x58\x9c\x62\xe0\xa4\x0c\xa0\x8a\x92\x02\x72\xd5\xdb\x0d\x85\xb7\xe1\xdb\xcb\x3c\x3c\x9e\xb5\x33\x5f\x3b\x0b\x77\xce\x81\xbc\x91\xc2\xa0\x0b\x6c\x97\xb9\x7a\x3b\xcb\xe0\x95\xba\x54\xf5\x7b\xcd\x8a\x15\xee\xdd\x88\x60\x6d\xd1\x57\x07\x2e\x98\x0a\x44\x4d\xc1\xd6\x50\xa1\x7a\x0c\x0a\x15\xb5\x5c\x7f\x33\xf0\xf0\xa0\x49\xa6\xe1\xe9\x9b\xdd\x9a\xb2\x2f\x76\x5d\x93\x86\x7d\x01\xab\x8d\x3d\x2f\x3f\x6f\x9e\x50\xec\x84\xf3\x35\x90\x16\x94\x52\x53\x53\x33\x64\x3d\xd3\x9e\xa9\x19\x54\x5a\x25\x22\x86\x8c\x01\x5b\x4a\x4c\xc0\x31\x79\xc8\xc5\x59\xa7\x73\xca\x2a\x5d\xe8\x2f\x75\x4e\xca\x5c\x7f\xb2\xe6\x62\x93\x70\x8c\xeb\x71\x88\x43\x89\xf3\xc5\xe9\xf0\x83\xd5\xc9\x7a\x19\x17\x53\x0a\xe1\xf7\xbf\xfc\xf1\x64\xb7\xb0\x67\xca\xa6\x7b\x5a\xa6\x6b\xca\xa1\x7b\xa2\xa0\x6b\x3e\xa5\x73\xc4\x64\x2a\x2b\x59\x6b\x3d\x35\x08\x5c\xda\xa0\xb7\xac\x9d\xe6\x3d\xd8\x70\x39\x47\x8e\x19\x72\x22\x06\x8a\xc6\x03\xb3\x25\x60\x4a\x51\x42\xd2\x81\x89\xdf\x12\x36\x2e\x67\xa5\xcd\x9b\xab\x0e\xfa\xc6\xb5\xde\x4c\xc3\xaa\x6f\xb0\xa9\x0b\xad\x53\xa4\x62\x11\xad\xf1\x13\xe9\xcb\x5a\xcd\xe8\xf6\x94\x1a\x86\x2b\x25\x51\x1e\xaa\xae\xb1\xc5\x50\x06\xce\xda\x42\x0c\xca\x2b\xe5\x6a\x88\xd6\xde\x0c\x52\xf9\x40\xe2\x7a\x39\x8c\xab\xe1\x38\x9e\x6d\x1e\x4e\x6d\x3e\xfc\xe6\xaa\x5c\xd2\x39\x21\xd7\xc9\x25\x9d\x33\x4f\xaf\x9d\x4b\x5a\x15\x64\x2f\x75\x1a\xbc\xb6\x14\xb4\xe1\x3d\x4d\x32\xab\xb4\xc9\x1c\x1b\x83\x48\xab\xa9\x54\x05\x0e\x06\x41\xeb\x40\x41\x45\xad\x55\xb9\x21\x79\xe6\xa3\x63\x59\x0e\x6b\x19\xe7\x6b\x39\x92\xe5\x38\xc4\x9c\x57\x27\xfb\xb0\xf1\x93\x7f\xfa\x63\xcc\x33\x5d\xd3\x86\xd3\xd8\x60\x67\xed\xc5\xea\x67\x4b\x17\x4a\x29\xda\x33\xa9\x6e\x99\xad\x66\x16\x10\xef\xb1\x55\x3f\x15\xd8\x47\x04\x76\x46\x71\xae\x99\x7c\x56\x37\x03\x1b\xdf\x3e\x59\x0e\x66\x18\xe7\x47\xb2\x19\xe2\xf0\x58\x64\x72\xc7\xeb\x8b\xcf\x7f\xf2\x57\x37\x03\x16\x9d\xa3\x87\x9d\xb0\xe8\x1a\x15\x9d\x84\x05\x19\xd4\x97\x77\xcc\xb6\xb0\x30\xac\xf7\x50\x46\x4c\xd6\x19\x89\x09\x12\x29\x0f\x64\x28\x02\x7b\xcd\xa0\xb9\x64\x66\x2f\x4a\x61\xbd\x31\xf2\xa3\x9c\xc8\x50\xce\xe6\x58\x5f\xff\x8e\x59\xdf\x84\xed\xcd\x11\x20\x93\xca\xc2\x22\x86\x69\x49\x6a\xd8\xd8\x3d\xca\x42\xf9\x14\x6d\x34\x0a\x8c\x04\x07\x24\x52\x80\x4b\xac\xc0\x96\x72\x52\x31\x7a\xa3\x2e\x1c\xcc\xeb\x9c\x13\xbd\xfe\x5c\xe9\x45\x4c\x34\x63\x0e\x7f\xf6\x6a\x85\xf1\xeb\xad\x5a\xbf\xa2\xf6\xf3\xd9\x6c\xe5\xc4\x2e\xd9\xb3\x0b\x93\x9e\x13\x09\xb6\x9a\x00\x59\xb0\x69\xa1\xa6\x5e\x34\x6b\xd0\x39\xa9\x82\x52\xa3\xd3\xf5\x2d\x79\xee\x02\xc9\xbf\x1f\xd7\x45\x96\xf3\xe5\xe1\x9e\xf6\xf3\xbf\x5e\xe2\xf6\xae\x79\xe0\xce\xb1\xde\xc9\x7d\x49\x43\x4a\xe1\xe4\x34\xab\x65\x1f\xf6\x34\x8c\x2d\x1a\xd2\x98\x35\x18\x6b\x22\x10\xd6\xd4\x2c\x5e\x21\x96\xe4\x12\x8a\x4d\xd9\x97\x9b\x61\xf1\x0f\xe5\xf1\x70\x8d\xf1\x91\xd7\x6e\x7e\x22\xa7\xcd\x25\xaa\x6a\x17\xac\x56\x6e\x8f\xf9\x7d\x51\x9a\x1c\x07\x90\x1a\x2a\x50\xae\x1e\xd8\x71\x81\x24\x16\xbd\xb8\x6a\x5c\x71\x37\xc3\xfc\x5f\x17\x29\xc3\xd7\x17\x27\xb5\x9e\x4e\x43\xfe\xa7\xff\xf2\xe6\x21\x8f\x8a\x1d\x4d\x6d\xc5\x63\x60\x6b\xf7\xd8\x3c\x78\xb2\x4e\xa5\x00\xb9\xda\x0a\x44\x2a\xb7\xa2\x46\x01\x5a\xa7\x6d\xf2\xda\xa0\x96\x9b\x61\xf3\xf7\x57\x43\x5d\xad\x9b\x88\x8c\x8b\x3d\x22\xf2\x6f\x7e\xfd\x16\x88\x46\x3b\xeb\x2f\xe9\xf7\x76\xc1\x38\xdc\xb7\x8b\x4c\x21\x54\xcb\x92\x21\x91\x46\xa0\xcc\x04\x9c\x2b\x41\xf0\x11\x9d\xe3\xe0\x39\x5e\x20\x9a\xce\x51\xf9\xeb\x8f\xd6\x5f\x4c\xca\x27\x87\x9b\xb3\xb1\xe3\x57\x68\x56\x77\xb4\x95\xb7\x3b\x5c\x65\xfe\x34\x3f\xcf\x97\x79\x71\x52\xb6\xa7\x5d\x8f\x9f\x65\xd3\xfe\xa9\xf1\xed\x4c\xe9\xd9\xae\x43\xcf\x4c\xf5\xc7\x4f\xae\x2c\x0f\x3e\x3d\xff\x7c\x9d\x13\xc0\x6b\x79\x24\xeb\x8d\x94\xd9\xad\x1a\x17\x1b\x79\x72\xfd\xd6\x78\xe7\x21\x86\x4e\xfb\x75\x8e\x8d\x77\xfd\xda\x29\xfd\x43\x84\x16\x79\x82\x9a\x88\x3c\xba\x7d\xfa\xc7\x8a\xab\xda\x26\x0d\x88\x89\x81\xa4\xf9\x12\x31\x42\x2c\x88\xa6\xa0\xce\xa8\xd4\x5b\x0a\x92\x8b\x45\xae\x34\x77\x0e\xb7\xef\xdd\x1d\x8a\x6c\xe6\x87\xcb\x69\x7a\xfa\xec\xd7\x53\x5d\xf6\x2e\x1f\x75\x1f\x70\xe8\x02\x46\xf7\x09\x95\x2e\x60\x74\xcf\xd6\x77\xe1\xa7\x73\x36\x7d\x72\x4a\xc0\x86\xa0\x98\x27\x86\xbe\xb5\x62\x8b\x7b\x8a\x66\xcf\x21\xb1\xa1\x04\xa4\xd1\x02\xd9\xac\xcf\x0a\xa4\x62\x5c\x72\x21\x0b\x32\xda\x1b\xc2\xc5\xa7\xc3\xb3\x23\x1b\xaf\xbd\x64\xee\xe4\x84\xb7\xc4\x99\x57\x29\xbc\x3b\xc9\xf0\xca\x2c\x67\x94\x6a\x95\xf8\x14\xfc\x88\x8d\xdb\xb3\x65\x14\x0c\xaa\xe2\x48\x41\xae\xc5\x03\xd9\x6a\x80\x33\x6a\x28\x51\xa1\xd2\x5e\xb2\xb2\xe9\x66\xb0\xdc\xdd\xbc\x5a\x6e\x86\xb2\x5a\xfe\xff\x71\x28\xf3\xcd\xf1\x22\xee\x11\xbf\x3f\xfe\xc5\x55\x49\xae\xf3\x48\xd4\x5b\x21\xb9\xce\xc3\x49\x9d\x24\xd7\x75\xb4\x67\x92\xe4\x82\x61\x54\x13\xa5\x95\xd1\x0a\xad\xda\x33\x1e\x91\x98\x03\xb3\x8b\x50\x2b\x07\x20\x31\x8d\xa1\xab\x03\x8b\x41\x9c\xd7\xd1\xd4\xc0\x37\x04\x65\x47\xc7\xab\xf5\x38\xfc\xe9\xf0\xb5\x1f\xb6\x17\xd3\x08\xfb\xbb\x7f\xb8\x2a\xc2\x3a\x4f\xf3\x75\x22\xac\xf3\x28\x66\x27\xc2\x3a\x8f\xd6\xbd\xf6\x34\xda\x7e\x82\x09\x93\x08\x23\xeb\xf7\x14\x92\x29\xe5\xe2\xac\x8f\x10\xb7\x33\xdd\xae\x78\x60\xaf\x2c\xe4\x18\x88\xc9\x63\x55\x12\x6e\x46\x1a\xdd\xce\x73\x9f\x1d\xd1\xfc\x83\xa8\x69\xba\x20\xfb\x7f\x35\xcd\x1b\xaf\x69\x38\x90\x99\x28\xfc\xc9\x2a\xf6\x7a\xcf\xf0\x50\xca\x2e\x45\xaa\x06\xc8\xfa\x6d\x8b\xab\x00\xc7\x10\xa0\x56\x97\x94\xb7\x25\xaa\xf8\xb6\xb2\xfd\x85\x28\xf9\xda\xf1\x3c\x6f\xb6\x47\xa4\xff\x20\x82\xa4\xab\x3c\xfa\xea\x82\xa4\x13\x5d\xff\xdb\x83\xc4\x68\x67\x3c\x4d\x05\x09\x79\x0e\x66\x8f\x24\xce\x5c\xc9\xb2\x2a\x60\x6c\x52\x40\xc6\x6c\x7b\x76\x19\x32\x6b\xc7\xd9\x46\x76\x17\xff\xcc\xc9\xdb\xac\xc8\x5e\x3c\x45\xd5\xab\xef\x3b\x05\x5a\x57\x4f\xe4\x2b\xac\x02\x3b\xe1\xf6\xc7\x57\x05\xbe\xee\xc8\x40\x8b\x76\x72\x87\xc4\x68\x24\xaf\xf6\x88\x2c\xe1\xec\x85\x94\x82\x4c\x31\x02\xa1\x6f\x15\x8d\x30\x44\xf2\x3a\x59\x61\x8d\xf1\x6d\xc9\xf8\x49\x91\x95\xbe\xba\x33\x48\x37\x51\x5c\xbd\x9e\x48\xb9\x3e\xc6\x3b\x83\xab\x0b\xe3\x93\xda\xc7\xa2\x63\x9a\x38\x3a\x47\x96\x8c\xdb\x37\xe3\xd8\x75\x84\xe5\x1c\x78\x3b\x0f\xf0\x5c\x7f\x44\x7f\xfb\xd7\x18\x9f\x83\xf7\x9d\x77\x86\x3b\xf3\x4d\x3e\xd9\x6c\xe6\xab\xe5\x30\x1f\xe5\x68\xf3\xbd\xe5\x9f\x0c\x1f\x2d\xe5\x7b\xcb\xf6\xdf\x3b\xef\x0c\xb7\xf3\x78\xee\xda\xdd\x51\x8e\x06\x18\x56\x8f\x97\xe7\x0f\x61\x5d\xb4\xde\xb3\x73\x2d\x17\xad\xb7\x7b\x61\xd7\x7a\x1f\x3f\xf9\x9f\x00\x00\x00\xff\xff\x3d\x47\x42\xf2\x8e\x59\x00\x00")

func templatesJsonBytes() ([]byte, error) {
	return bindataRead(
		_templatesJson,
		"templates.json",
	)
}

func templatesJson() (*asset, error) {
	bytes, err := templatesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates.json", size: 22926, mode: os.FileMode(420), modTime: time.Unix(1611791102, 0)}
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
	"templates.json": templatesJson,
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
	"templates.json": &bintree{templatesJson, map[string]*bintree{}},
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

