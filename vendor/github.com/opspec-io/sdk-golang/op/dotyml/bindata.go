// Code generated by go-bindata.
// sources:
// github.com/opspec-io/spec/spec/op.yml.schema.json
// DO NOT EDIT!

package dotyml

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

var _githubComOpspecIoSpecSpecOpYmlSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x7b\x73\xdb\x36\x12\xff\xdf\x9f\x02\xa3\x7a\x5a\xe9\x62\x51\x76\x1e\xbe\xd6\x9d\x4e\xc6\x4d\x9c\xab\x6f\x9a\xd8\x93\xd7\xcd\xd5\x72\x1a\x98\x84\x24\xd4\x24\xc1\x02\xa0\x6c\xb7\x97\xef\x7e\x03\x80\x6f\x02\x7c\x99\x8c\x95\x87\x67\x3a\x8d\x08\x2c\x88\xfd\xed\x62\xb1\xbb\x78\xf0\xef\x2d\x00\x46\xdb\xcc\x5e\x21\x0f\x8e\x0e\xc0\x68\xc5\x79\x70\x30\x9b\xfd\xc1\x88\x3f\x55\x4f\x2d\x42\x97\x33\x87\xc2\x05\x9f\xee\x3e\x9c\xa9\x67\xdf\x8c\x76\x24\x1d\x76\x62\x1a\x76\x30\x9b\x91\x80\x05\xc8\xb6\x30\x99\xed\x5a\x7b\xd6\xfe\x8c\x04\xd6\x8d\xe7\x5a\x51\x33\xa2\x49\x45\xc6\x31\x77\x91\x20\x54\x15\xd4\x43\x07\x31\x9b\xe2\x80\x63\xe2\x8b\xa2\xa7\x68\x81\x7d\xc4\x00\x04\xc1\xe5\x52\xd5\x08\x28\x09\x10\xe5\x18\xb1\xd1\x01\x10\xfd\x06\x60\xe4\x43\x0f\x25\xbf\xca\xad\xbc\x80\x1e\x02\x64\x01\xf8\x0a\x25\xed\xc8\x7a\xfc\x26\x90\x3d\x60\x9c\x62\x7f\x39\x92\x8f\x3f\xa8\xd2\x42\x1b\xa6\xa6\x9f\xa6\x3f\xb5\x6f\xd8\xa6\x68\x21\xea\x7d\x33\x73\x04\x2b\x58\x54\x64\x33\x0f\xd2\x4b\x87\x5c\xf9\xf9\x37\x62\x3f\x08\x39\xcb\xbe\x4c\x4f\x1d\x40\x0a\xbd\x3c\x29\x09\x79\x67\x5a\x1a\xfa\xf5\x74\xcc\x2e\xc0\xb3\x46\x94\x55\x43\xf3\x56\xd5\x88\x61\x21\x41\x1d\x2a\x0c\x79\x6f\x11\x2d\x30\x26\x95\xa9\xd9\x5b\xce\x54\xe5\xf3\x71\x49\x13\x27\x20\x64\xc8\x01\x17\x37\x4d\x05\x94\xeb\xca\x56\xd4\x9d\x11\x45\x7f\x86\x98\x22\xa1\xed\x67\x19\xc5\xdb\x02\xe0\x5c\x96\x43\xc7\x91\xf4\xd0\x3d\xcd\x2a\xe9\x02\xba\x0c\x45\xda\x9d\xbc\x22\x55\x5e\x48\x29\xbc\x39\x95\x72\xc9\xb0\x99\x8c\x8e\x4c\xf1\x8e\x01\x83\x43\x51\x05\x48\xd1\x22\x8e\xa8\xc0\x02\xfa\x39\xc0\x63\x45\x27\x17\x7f\x20\x9b\xa7\xcf\x35\x83\x29\xed\x53\xee\x91\xb9\x72\xc5\x68\x49\x8a\x9b\x0c\x83\xf8\xef\xc3\x4e\xb1\xe9\x05\x0c\x5d\xae\x6b\x36\x66\x4b\x75\xb7\xb2\x15\xcc\x5e\x21\x9b\x22\x6d\x33\x05\x38\x8f\x95\xc2\xca\x46\x01\x66\x80\x29\xc2\x1d\xd3\xdb\x2f\x08\x71\x11\xac\xe1\xc2\x26\x3e\xe3\x14\x62\x9f\x97\xd1\x33\x02\x24\xbb\xf0\x24\x43\x99\x7f\xc5\x96\xe1\x75\x95\x8a\xb8\x55\x24\x4f\x48\x6b\xf5\x57\x56\x2a\x0d\x02\x90\x2a\x4c\xf4\xfb\x3c\x37\x84\x4b\x4c\x18\xb5\x3c\x5b\xa9\xbb\xea\x26\x4c\x1c\x73\xe4\x15\xc1\x2e\xca\xfa\xdf\xaf\x4e\x5e\x80\x57\x72\x6a\x02\x67\x05\x52\x70\x89\x6e\xae\x08\x75\x52\x83\xc2\x09\x71\x99\x85\x11\x5f\xc8\xe9\x70\xc5\x3d\x37\x9a\x13\xaf\x28\x5e\xae\xf8\x34\x33\x61\x4e\xd7\xd0\xc5\x0e\x14\xed\x4d\x77\xf7\xbe\x61\xc8\x96\xff\xdc\xb7\xf6\x76\x27\x39\x5d\x32\xc8\x5e\xf0\xad\x15\x7d\x46\xd4\x23\xdc\x96\x45\x3c\x20\x63\x3f\x14\xf8\x82\xfe\xcd\xc9\x22\xa7\x26\xe2\xaf\xa1\xea\x1b\xd9\x2f\x69\xbb\xa6\x49\x1d\x2c\xdd\xdf\x96\x1f\x69\xf9\x5f\xe7\x5a\xb1\x78\xf0\xba\xb5\xf2\xc5\x34\x3d\x0a\x67\x37\x11\xce\xa3\xb2\xd6\xc5\xe3\x0a\xfb\x1c\x2d\x11\xcd\x17\x7a\xd8\xc7\x5e\x28\x26\xa4\x5d\x3d\x83\xd8\x6f\xcf\x60\x44\x33\x14\x83\x7b\x7d\x32\x18\xfa\xf8\xcf\x10\xb5\xe6\x31\x43\x36\x94\xf5\x78\x60\x60\xb3\x34\x0b\xb5\xb3\xef\x39\x93\x1d\xb5\x65\x76\x4a\x72\x15\x4c\x6e\xc9\xcf\xaa\x52\xaf\x8e\x49\xcc\xe5\xc6\xbb\x26\xe5\x40\x46\xd4\x04\x6b\xe8\x86\xa8\xad\x2b\xb1\x41\xf3\x7c\xa1\x83\xf9\x99\xde\xc1\xd4\xac\x32\x49\xa1\x49\x5d\x9e\x62\x8a\x6c\x4e\x68\xbf\x9e\xac\x83\xe9\xa7\xad\x2c\x3f\x0a\x04\xe0\x05\x23\x6e\xc8\x11\x08\x20\x5f\x81\x05\x25\x1e\xa0\x84\x70\x81\x4f\x70\xb9\x04\x84\x02\x8a\x5c\xc8\xf1\x3a\xaa\x21\x6c\x1e\x0d\x28\xe2\xc8\x51\xb5\x85\x43\xeb\x60\x0a\xb0\x0f\xae\x56\xd8\x5e\x45\x21\x99\x70\x6f\x45\xfc\x67\x54\xc8\x6c\x64\x6c\x62\xac\xbd\x6b\xed\x24\xa2\xee\xec\x5e\x6f\xd0\x98\x10\x2a\xa6\x1d\x0f\x0b\xec\x22\xf3\x80\x48\x4b\x4d\x23\xe2\x19\x76\x51\xaf\x83\x41\xbc\xf2\xeb\x68\xd8\xb4\xd1\x20\xa4\xf2\x59\x0c\x04\xa9\x5e\xda\x91\x10\x84\xae\xfb\x84\x22\x27\x1f\xfc\x19\xb4\xb7\x80\x92\xa0\x43\x3e\xc7\xd0\x65\x2a\x8d\xe3\x84\x42\x0a\x00\x86\x7c\x25\x9e\xdb\xd2\x4b\x02\x57\x98\x2b\x39\x32\x12\x52\x1b\x45\xa3\x05\x7b\x70\x89\x84\x46\xe4\x92\x3e\xa6\xf1\x11\x32\x44\x0b\xb9\x44\x5d\x8f\xd0\x75\x40\x11\x93\x69\x27\x9b\x20\x6a\xe3\x0b\x17\x01\x4e\x80\x52\x0f\xcd\x1c\x6f\x18\x29\x69\x3b\xfa\xe8\x2e\x80\x8c\x09\x17\xf2\x2e\xbb\x53\xd2\x0f\xbd\xe8\x13\xe4\x74\xdd\x8f\x55\xa2\xbd\x23\xea\x87\xde\x05\xa2\x75\xc9\x83\x72\xad\xee\xd9\x03\xd7\x95\x31\x6b\xf3\x9c\x81\x20\x18\x28\xa4\xb9\x7f\xdf\xe0\xeb\xab\x6c\x4b\xae\x48\x1f\xf2\x1a\x24\x5d\x06\x2c\x6b\x45\xb4\xba\x18\x07\xf3\xcd\x81\x11\x04\x43\x01\x63\x0a\x82\xee\x00\x18\xe4\x87\x5e\x1b\x5c\x44\xfd\xa1\x60\x31\xc5\xf8\xcd\x61\x89\x29\x14\x10\xf5\xdc\x2f\x08\xf5\x60\x71\xae\x1b\x11\x1f\x35\xc9\xfc\x24\x03\x58\x17\xa2\xeb\x80\x7c\xa9\x6c\x0f\x93\x76\x5e\x75\x11\x5c\x20\x69\xe7\x4d\x2d\x14\xa6\xee\x52\x79\x24\xbe\xb3\xc2\x73\x90\x76\xaa\x50\x72\x6e\x9c\x7e\x8d\xb9\xa0\x28\xd3\xd0\x26\x15\x24\x48\x86\xd2\x12\x83\x92\x14\x45\x5e\xc8\xf7\xb4\x66\x42\x91\x0c\xc4\xc4\xc3\x2e\x4c\x84\x2e\xc7\x81\x8b\xda\xd9\xb1\x94\x6a\xa8\xc4\x55\x07\x56\x7c\x52\x1a\x73\x55\x3c\xf8\x84\x0f\xa5\x4c\x8f\x1a\x25\xb3\x2b\xec\x6a\x96\xad\xd8\x6e\x34\x66\x4c\x12\x0c\xc5\x9a\x49\xc7\x3e\xd6\x24\xd3\xca\x35\xd7\xb8\x4d\xe6\xd0\x33\x5b\x6e\x72\xbb\x5f\x28\xf3\xda\x67\xf8\x19\x69\xf4\xc6\x07\xa0\xe6\x49\xb0\x8f\x70\x2f\x9a\xb7\xee\x74\x61\xb1\x52\x05\x37\x2b\xa8\xcc\x0b\x21\x1f\x56\x2a\xc5\xab\x0b\x0f\xca\xb5\x7a\x58\x5c\x3c\x35\x69\x6d\xc3\x15\xc6\x94\xbe\x8d\xf9\x5a\x41\xdf\xa1\xe8\x8a\x35\x30\x60\xfb\xd6\x23\x6b\xbf\x60\xc1\x9a\xfa\x65\x2d\xf4\xaf\x97\xe5\xbc\x5a\x27\xea\xcb\x08\xcb\xca\x8a\xfa\x35\x2c\xeb\x0a\x8c\x83\x02\xe4\x3b\xc8\xb7\x5b\x8e\xd0\x2c\xdd\x50\x0e\x5f\x71\xa1\xbc\xe1\xb0\xfc\xb8\x0b\xe5\x55\x39\xcf\xf6\x2b\xe2\x9f\x65\x94\x5c\x0c\x1a\x47\x7e\xe8\xba\xe5\xe9\x3c\x9a\x5c\x72\x8f\xcf\x6b\x15\xd8\x83\xd7\xdd\xe6\x98\x1c\xe1\x50\x2a\x6c\x1a\xe9\x5d\x77\x13\x74\x64\x35\x4b\x38\x14\xab\xa6\x30\xa0\x13\xab\x9f\x5a\xe0\x56\x61\x79\xbf\xbc\xc0\xad\xc3\x34\x64\x0c\x6e\x2a\xd1\x09\x06\x57\xea\xa2\x67\x68\x70\x84\x41\x43\xa7\xb7\xcb\xec\x64\x00\x0c\x72\x8e\x68\x47\x7b\x50\x22\x1e\x0a\xbe\x7f\x6e\x2a\x7c\x99\x18\xaa\x31\x6a\x31\xcd\x50\x60\x15\x2d\x4d\xf7\xac\x74\xd9\x17\xe9\x25\x5d\xa2\xa4\x66\x4e\x97\x64\xcb\x4d\xe9\x92\x13\x59\xa7\xd7\x74\x49\x54\xef\x93\x49\x97\xe8\xfc\x9c\xdb\xa7\x4b\x54\xab\x77\x9b\x2e\xa9\x34\xfc\x9b\x95\x2e\xc9\x0b\x21\x9f\x2e\x61\xc4\xbe\x44\x15\x7a\x9e\x2d\xaf\xd5\xda\x82\xb4\x5e\x49\xda\x4a\xfd\x37\xe9\xb9\x7a\xed\x86\xe8\x79\x7b\x05\x55\xdd\xff\x2c\x36\x70\x44\x92\xd0\x2b\x8f\xb4\xbd\x75\xb9\xb6\x72\xad\xaf\x4b\xf1\xd1\x63\xc3\x49\xa3\x12\x60\xb5\x73\xfc\x97\x91\xf3\xe9\x00\xcc\x67\x99\x64\xa8\x76\x7a\x6a\x97\xe2\x8b\x66\x3a\x0c\x10\x65\x48\xee\x95\xcb\x61\xa1\xa8\x07\x41\xa3\xe8\x2c\xb7\xdd\x1d\xe0\x40\x8e\xa6\x1c\x7b\x9a\x7d\xd3\x95\x99\xbc\x98\x0c\x28\xde\xfa\xe5\xc9\x7a\x50\x5c\xb8\xd5\x09\xad\xc5\x76\x83\x94\xcb\x42\xd9\x79\xd5\x7c\x55\x81\x9a\xb0\xe5\x74\x2a\x37\xc0\x4d\xc5\x08\xab\x03\xef\x10\x28\x92\x68\xcf\x1c\x45\x0b\x44\x91\x6f\x23\x00\x19\x90\x03\x53\x9d\xa0\x3c\x5b\x62\xbe\x0a\x2f\x2c\x9b\x78\x33\x45\x30\x73\xb0\x60\xf7\x22\x14\x2d\xcd\x12\xba\x14\xef\x1a\x0a\x4e\x11\x8a\x0b\xf6\xac\xbd\x07\x69\x13\xfd\x02\x5c\x04\xa4\x1f\x9c\x91\x07\xb1\x26\xe9\x57\x69\x77\x04\xc9\x50\x5a\x79\xbf\x57\xd0\x14\x77\xfd\x20\xb5\x22\x8c\x17\xf6\x08\x36\x00\x2b\xa6\x1a\x0a\xaf\x07\xbd\xe2\x95\xf0\xd8\x0f\x64\x38\x58\x3f\x6c\x07\x97\xa0\x18\x0a\xaa\x87\xbd\x42\x25\x79\xeb\x0d\xa6\xfd\xd6\x30\xed\x0f\x05\xd3\xa3\xbe\x61\xda\xef\x09\xa6\x90\xe2\x76\x28\x85\x14\x0f\x05\xd2\x7e\xaf\x20\x09\xce\xfa\xc1\x88\x21\x6f\xdd\x60\x27\xe2\x21\x60\xc8\x83\x3e\xc7\x36\x88\x6e\x42\x28\x4e\x93\xaa\x21\x81\x91\xc2\xee\x60\x36\x4b\x1f\xcd\x7a\xe5\x3e\xea\x73\x35\x00\x5b\xba\x92\xc2\x82\xd3\xaf\xc8\x5f\xf2\x55\xcb\xc5\x26\x45\x34\x90\x1f\x6d\xca\x53\xd7\x2c\xbe\xec\xe9\x39\xc4\x7e\x07\x0e\x63\xa2\x81\x38\x34\xa5\x92\xeb\x96\x97\x76\xf2\x0c\xc4\xd9\xb9\xcf\x63\xd9\xa9\x22\xf8\xfb\xf2\x96\x9d\x3a\x44\xc2\xd1\x42\x48\x87\xb5\x93\x81\xc0\xf9\xde\x80\x8d\xc6\xd8\xa5\x81\xec\x88\xa2\x25\xba\xee\xe5\x44\xb3\x7a\x4f\x45\xea\x33\x53\xde\x3a\xf5\xa9\xce\xd9\x74\x4a\x7d\x2a\xf6\x37\x23\xf5\xd9\x20\xc5\x3f\xcc\x01\xb8\xe8\xa0\xd2\x9d\xa6\xf8\x2b\x07\xd9\x86\x65\x69\x73\x42\x28\x1c\xb4\x2b\x6a\x78\x01\xf1\xd3\x2e\xeb\x53\x95\x4b\xb2\xa3\xb3\xe9\xef\x16\x9c\xfe\x75\x38\xfd\x6d\x77\xfa\xc3\xf9\xbd\x8e\x67\x41\x2a\x2e\xc0\x39\x4d\xef\xac\x32\x88\xbc\x61\x6b\xb9\x5b\x0b\x7a\x68\x2f\x39\xd2\xde\x43\x5b\xe9\x69\xe0\x1e\x1a\xcb\xee\xef\xee\xa1\xb9\xec\xfa\x67\x0f\xcd\x65\x97\x99\xfa\x68\x2e\x63\xba\x9b\xb8\xbc\xdd\x27\x91\xe2\x6a\xbc\x6e\x22\x29\xd6\x31\x4d\x1a\xe9\x40\xb4\x75\xb5\xcb\x77\xe7\xf4\x35\x43\x54\x6f\x7e\x4c\xca\x7a\x9d\xff\xf7\x9b\xe4\x4a\x2b\x6d\x7a\x8c\x6f\x2b\x76\x24\xd1\xa6\x31\xa2\xea\xb6\xe3\xe3\x26\xe8\x97\x8d\x47\xd6\xfd\x0a\x3e\x34\x11\x66\xd9\x85\x8d\x0a\xf4\xbb\x43\x2b\x63\xef\xd1\x15\xc5\x1c\x9d\xf8\x6e\xf1\x82\xb9\x5a\x1c\x12\xc2\x9e\xb7\xca\xef\xed\x56\x66\x01\xeb\x97\x6d\x75\x3e\xf8\xdf\xf5\x51\x8e\xf9\x7a\xb7\x76\xed\x34\x3b\x5d\xd5\xa0\xa1\x66\xbb\xfd\x1a\x34\x54\x15\xbf\x69\xfd\x16\x66\x2f\xf5\x7e\xb9\xbd\x34\x9a\xd0\xb3\x57\x1c\x72\x6c\x03\x1b\xba\x2e\x58\x52\x18\xac\x52\xbd\x40\xbe\x75\x85\x2f\x71\x80\x1c\xac\xae\x2e\x15\xbf\x66\x4f\xa0\xeb\xfe\x2e\x6b\x4e\x34\xfe\x4f\xd9\x53\x69\xc0\xa9\x4d\x7c\x0e\xb1\x8f\xa8\x68\xbb\x33\xee\xc1\x6d\xa8\x85\xcf\xe7\xba\xc8\xbd\x4d\x1b\x0c\x51\x0c\x8b\x2d\x68\x25\x95\x67\x58\x27\xb3\x7c\x8d\xce\xeb\xff\x49\x33\x6d\xc2\x23\xdb\x2b\x6e\x7d\xd3\x69\xce\x13\xe2\x79\xd0\x77\x00\x0d\x7d\x70\x71\x03\x20\x48\xde\xf5\x23\x20\x6b\x44\x29\x76\x10\x03\xd0\xbf\x01\x0c\x71\x00\xb9\x8c\x52\xd4\xb2\x98\x8b\xd6\x48\xb3\xdc\x63\x8e\xf5\x81\x39\xde\xd7\x75\xad\xf5\x65\x0e\x95\x62\xd5\x5d\xe9\x90\x17\x6e\xf4\xab\x18\x0a\x62\xaa\x8d\x9d\x2a\xf6\x39\xea\x98\x89\xaf\x74\xc2\x88\x01\xec\x4b\x14\x53\xa9\x96\x88\xeb\xb6\x7e\x46\xd5\xde\x8d\xcf\x54\xb8\x71\x7e\x30\x79\x2c\x82\x8f\xf9\x7c\x96\x89\x3f\xb6\xb5\x54\xc6\x40\x24\xfe\xd3\x91\xe8\x58\x1a\x5f\x61\xd7\x05\x17\x08\x5c\x90\xd0\x77\xa4\x64\xa0\x97\xdc\x33\x03\x82\xcb\x65\x79\x26\x29\xc1\x27\xcf\x09\x68\x2b\x7d\xd0\xd3\x36\xed\x9d\x49\x7b\x1c\x4c\x95\xea\x80\x6f\x67\x84\x02\x66\x93\x40\xae\xf0\xca\xfe\x23\x0e\xc2\x80\xf8\x00\x5d\xe3\xb2\x48\x93\x37\xb5\x55\xb0\x88\x1f\xcd\xd3\xf3\xd2\xb3\x62\xad\x12\x0a\xcd\xa2\x6c\x0d\xe9\x08\xf9\xeb\xb7\xb0\x17\x65\x3e\xf2\xd7\x98\x12\xdf\x43\x3e\x07\x6b\x48\x31\xbc\x70\x7b\x55\xeb\xb3\x77\x3f\xdd\x81\xf6\x62\x3f\xa3\x0d\x57\x33\xa5\xcd\x3e\xf4\x34\x6b\xef\x25\xe0\x3e\xba\x1a\xd7\x18\xc1\xa8\xb1\x4f\x54\x53\x17\xd8\xd5\xea\x46\x5b\x3d\x7d\x86\xfb\xd5\xcb\xaf\xe6\xd6\xd8\x3b\x93\x9e\xca\x1b\xbd\xbe\x4c\x7b\x2b\xbd\xa4\x2e\x5a\x5c\xe1\xd5\x01\x95\x0b\x2d\xae\xf4\x24\x45\xa5\x5b\x62\xe2\x9d\x4b\x9c\x24\x77\x80\x69\x81\xee\x00\xb2\x46\x63\xb4\xb7\x9b\x35\x78\x4d\x4a\xd6\x41\x40\xda\xd4\x70\x0e\xac\x62\x4c\xde\x9f\x8c\x35\x57\xa4\x01\x8d\x1c\xe4\x67\x17\x72\x46\x08\xd8\xd0\x17\x83\x39\xd9\xe0\x25\x97\xca\xe5\x95\x7c\x84\xaf\x54\x2e\x4c\xd5\x64\xb7\xcb\xb2\x04\x84\xea\x93\xff\xc5\x2c\x9c\xa8\x17\xd9\x96\xe4\x66\xc0\xb4\xbb\x9c\xc8\x07\x2b\xc2\x2a\xd6\x24\x8c\x0a\xdd\xcc\xbe\x9e\x49\x33\x3a\x9e\xaa\xff\x4f\x1e\x8f\xb9\x1d\xfc\x2f\x74\x82\xc9\xe3\x86\xea\xfe\x0b\x61\x1c\x08\x86\xc7\x6c\x22\x7a\x7c\x81\xa5\xa1\xd4\x2b\x7c\xcd\x26\x83\x6c\xc7\x65\x88\x5d\xe8\x5c\x17\x4d\xed\xac\x66\x2a\x4f\xdc\x69\x4a\x6c\x8a\xfd\x81\x79\xe1\x22\xa9\x54\x8a\x1d\x63\xed\x88\x36\xfe\x43\xc7\x11\xd6\x02\x78\x30\x08\x90\x9c\xa2\x60\x5c\xa4\xdb\x76\x09\xea\x74\x79\x68\x54\xb9\x73\x44\x8b\x31\x75\x9f\xa0\xbe\xb3\xcc\x4e\x81\x19\x4b\xee\x20\x4a\x41\x40\xd1\x02\x5f\xe7\xa1\x54\x3e\xdf\x80\x50\x3a\x28\xa0\xc8\x86\x5c\x1a\x53\x4e\x43\xd4\x2b\xd8\x27\x61\x93\x43\x53\x1f\x1b\x6c\x12\xf2\xcf\x0e\xec\x2b\x42\x2f\x9f\x96\xae\x7b\xd6\x41\xf1\x1f\x42\x2f\x05\x9f\x4e\xe6\xca\x69\xbe\x02\xe3\x7c\xee\x27\xb3\x9b\x4b\xba\x10\x6d\x33\xff\x5b\x86\xae\x9a\xe7\xee\xc8\x7b\xca\x3c\x3b\xef\x63\xad\x59\xbf\x8c\x9c\x86\x06\x5b\x85\x77\xb5\x38\x39\x29\x13\x97\x27\xe5\xaf\x0f\x15\x17\xfa\x85\xcf\x03\x5d\xfc\x17\x62\xe0\xc5\xe1\xf3\x23\x35\xd9\x9e\xbc\x79\x7d\xfa\xe6\xf5\xef\xf2\x01\xf6\xa3\x9d\x86\xe0\x3b\xf1\xfb\x20\x5b\xf8\x9d\x05\x8e\x17\xf9\xda\x0c\x08\xa7\x7e\x07\x60\x0e\x9e\xbf\x79\xf5\x5a\xde\x02\xc8\x58\xe8\x21\x47\xb5\xff\x13\xd8\x1e\x67\x08\x26\x1b\xb2\xca\x6d\x8e\x48\x6a\x37\x28\x36\xd4\xb4\x5e\x16\x3a\x95\x58\x8f\x4b\xdf\xa3\xaa\x90\xea\xf1\x8b\x44\x38\x52\xb6\x6f\x0f\x7f\x7d\x93\x93\x6a\x5a\xe1\x40\x15\x2a\xa9\x46\xf5\x2a\xe4\xa9\x6a\x08\x81\xa6\x4d\x7c\xfa\xf2\x6c\x17\x79\x0c\x24\x61\xfd\x69\xe7\xe0\x76\x79\x7b\x12\xdc\x12\xcc\xae\x11\x62\xe9\xfb\x69\x0d\x00\xcf\x69\x7a\xa3\x60\xaf\xfc\xa9\xb5\xc6\xaf\x89\xed\xe4\x5d\x05\x95\x9a\xb7\x74\x8f\xaa\x33\x2b\x67\x39\xb2\x26\x11\x46\xba\xc3\x2f\xa4\x78\x9a\x44\x83\x9f\x58\x1c\xdc\x21\x51\x57\xe7\xfb\x7c\x61\x2a\x7e\x59\xdc\xfa\x98\x14\xd5\x40\xd9\x04\x2c\x50\xad\xdf\xa0\xbb\x8e\x83\x86\x7a\x0e\x5a\xe9\x3a\x30\xe7\x23\x6b\x8d\x01\xe8\x64\x10\x80\x21\x61\xa8\xed\x45\xf5\x48\x03\x86\xd1\x06\x74\x23\x0e\xb4\x1a\x75\xba\x5e\xb6\xb5\x03\x42\xd3\x7a\xb4\x03\x0d\x67\x63\xc3\x0d\x0d\xc1\x2d\x1c\xee\xdc\x5a\xbf\x6e\xf6\xce\x55\xe8\x3c\x87\xc7\xad\x14\x67\xf2\xde\xb6\xb4\xdb\x8d\x6f\x71\xd1\x83\x98\x74\xb0\x3b\x94\x99\x2d\x0f\xda\x9d\x27\x69\x71\x67\x18\x55\x1b\x1b\x0b\x62\xd4\xbd\xee\x10\x4a\x0e\x8e\x52\x9f\xd5\x1c\x2a\xa4\x95\xa2\x2f\xae\xa0\x35\x74\x43\xc8\x11\x8b\x4c\x6d\x1e\x8c\xba\xbd\x38\x7d\x41\xa8\xf7\xb7\xeb\xf7\x73\x19\xe7\x0d\xc5\x9c\x83\x12\x3b\x9f\xe3\x2f\xb3\x4a\xd4\xf8\xab\x1c\x9a\x19\xa3\x61\xaf\x9a\x61\x9e\xe9\x93\xee\x30\x46\xd5\xfe\xae\xa3\x8a\x0f\x86\xe4\x37\x0d\x45\x5b\xe8\xfa\x50\x15\x10\x6f\xc7\x4b\xd5\xb5\x7a\xaf\xac\xf9\x03\x79\x3d\x09\x37\xee\xd1\x46\x49\x57\xd3\xa9\xe1\xc4\xab\x36\x20\xf6\x22\xdd\xe8\xba\xe7\xb6\xc2\xad\xb8\x2c\xff\x76\xb2\x8d\x6e\xeb\xde\x28\xd1\x96\xfa\x34\x9c\x64\xd5\x4c\xd7\x8f\x89\x57\x6d\xc9\xac\x4f\x40\xc9\x1a\x3b\xc8\x01\x90\xa5\x25\x3b\x20\x9a\x47\x6f\xe4\x06\x14\x06\xbe\x55\x60\x33\x10\x6f\x02\x50\x1f\xf2\x22\xae\x0c\xa1\xda\xce\x15\xba\xbb\xf9\xea\xb3\xec\xa6\x0c\x7b\xc5\xfe\x86\x16\x41\x53\xcb\x95\xe6\x26\x71\xfa\xd1\x9f\x21\x5e\x43\x17\xf9\x5c\xc0\x3e\x1f\x9d\xbe\x3c\x39\x3d\x7a\xf9\xfa\xbf\xf3\xd1\x01\x98\x8f\xb6\xc7\xf1\xef\xc9\x7c\x54\x1d\xc0\xeb\x77\x4d\x14\x3d\x71\xf3\x29\xe3\x2e\xc7\xa4\x7a\x9b\x6e\xa3\x5b\xe3\x36\x69\xd8\x6a\x3a\x35\xdc\xb8\x4d\xb9\x30\x0f\xd8\xf7\xdb\xe3\xe3\xa7\x47\x2f\x5e\x1f\x3f\x3b\x3e\x7a\x39\x79\xbf\x03\x72\x0f\xac\x57\x6f\x7e\xb6\x12\x65\x29\x96\xce\x4e\x0f\x5f\xff\x22\x9e\x12\x2a\x0a\xa2\x9f\x60\x7c\xb5\x42\x14\x25\x5f\xe3\x7b\x3f\x7b\x0f\x28\x12\x7d\x46\x3e\x67\x72\x17\x62\xb4\xea\x80\xfd\x25\x78\x4f\x02\xeb\xc6\x73\xdf\x4f\x2c\xf0\x02\x31\x2e\x1e\x85\x81\x80\x8a\xf8\xd1\x3e\x5b\x79\x24\x30\x0c\x02\x42\x39\x72\x76\x00\xb6\x90\x25\xdf\xb6\x3d\x7e\x79\xf4\x6c\x32\x79\x5f\xca\xf6\x0a\xb6\xde\xcd\xe7\xdb\xf3\xf9\xd8\xba\x37\x9f\x4f\xb6\xcb\xbe\x7b\x36\x81\x9e\x3f\x1d\xda\xcb\xfc\x55\x10\x69\xd3\xf9\xab\x98\xd6\xef\x6d\xfe\x8a\x56\xf4\x7a\x1a\x08\x5a\x5d\x43\x7d\xe0\x96\xef\x56\x1d\x6c\x15\xe7\x29\xb4\x63\xc4\x80\x66\xe5\x31\xc1\xdb\x37\x54\x72\x8b\x3a\xb6\x53\x9a\x84\x3b\xb6\xd3\xd6\x88\x30\xe4\xbd\xcd\x6d\xbf\xcf\x5d\x93\xf1\x36\xeb\xa6\x15\x8f\x66\xac\xef\x5b\xbb\xd6\x6e\xe9\x92\x0c\xdd\x55\x18\x2c\x40\xf6\x4c\xd5\xb7\x56\xdc\x73\x35\x6b\x3b\xc5\x41\x95\x1d\xed\xe3\x68\x6f\xca\x7c\x6e\x69\xfe\x39\x7e\x7c\x30\x9e\xcf\xe5\xfe\x95\xc3\xe9\x6f\x70\xfa\xd7\xf4\xfc\xde\xf8\xf1\xc1\x7c\x6e\xe5\x1e\x4d\xfe\x31\x99\x3c\x96\xcf\xef\x65\x9e\xcf\xe7\xd3\xf9\xdc\x3a\xbf\x37\x79\xbc\x9d\xb7\x17\xc9\x79\x6a\x1d\x34\x49\xa1\x09\x9c\xe7\x51\x05\x80\x7d\x70\xb6\xde\xb5\xee\x7f\x0f\x9e\x10\xcf\x23\xbe\x28\x00\xec\xc6\xe7\xf0\x3a\x05\x2a\x40\xb6\x65\xcb\x62\xd1\xb0\x44\x4c\x90\xcc\x26\x00\xfb\xb6\x1b\x3a\x62\x74\xff\xeb\xd9\x73\xc0\xa1\x18\xdd\xe8\x9a\x23\x5f\xca\xb6\xda\xf0\x6d\x89\xff\x3e\xfc\x3f\x00\x00\xff\xff\x18\xc5\x2d\x07\xf7\x87\x00\x00")

func githubComOpspecIoSpecSpecOpYmlSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_githubComOpspecIoSpecSpecOpYmlSchemaJson,
		"github.com/opspec-io/spec/spec/op.yml.schema.json",
	)
}

func githubComOpspecIoSpecSpecOpYmlSchemaJson() (*asset, error) {
	bytes, err := githubComOpspecIoSpecSpecOpYmlSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "github.com/opspec-io/spec/spec/op.yml.schema.json", size: 34807, mode: os.FileMode(420), modTime: time.Unix(1521604342, 0)}
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
	"github.com/opspec-io/spec/spec/op.yml.schema.json": githubComOpspecIoSpecSpecOpYmlSchemaJson,
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
	"github.com": &bintree{nil, map[string]*bintree{
		"opspec-io": &bintree{nil, map[string]*bintree{
			"spec": &bintree{nil, map[string]*bintree{
				"spec": &bintree{nil, map[string]*bintree{
					"op.yml.schema.json": &bintree{githubComOpspecIoSpecSpecOpYmlSchemaJson, map[string]*bintree{}},
				}},
			}},
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
