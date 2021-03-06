package dircopier

//go:generate counterfeiter -o ./fake.go --fake-name Fake ./ DirCopier

import (
	"fmt"
	"github.com/golang-interfaces/iioutil"
	"github.com/golang-interfaces/ios"
	"github.com/golang-utils/filecopier"
	"path/filepath"
)

type DirCopier interface {
	// OS copies an os dir from srcPath to dstPath. Creates or overwrites the destination as needed.
	OS(srcPath string, dstPath string) (err error)
}

func New() DirCopier {
	return dirCopier{
		os:         ios.New(),
		ioutil:     iioutil.New(),
		fileCopier: filecopier.New(),
	}
}

type dirCopier struct {
	os         ios.IOS
	ioutil     iioutil.IIOUtil
	fileCopier filecopier.FileCopier
}

func (dc dirCopier) OS(srcPath string, dstPath string) error {
	// get properties of srcPath
	fi, err := dc.os.Stat(srcPath)
	if nil != err {
		return err
	}

	if !fi.IsDir() {
		return fmt.Errorf("%v is not a dir", srcPath)
	}

	// create dstPath
	if err := dc.os.MkdirAll(
		dstPath,
		fi.Mode(),
	); nil != err {
		return err
	}

	entries, err := dc.ioutil.ReadDir(srcPath)

	for _, entry := range entries {

		sfp := filepath.Join(srcPath, entry.Name())
		dfp := filepath.Join(dstPath, entry.Name())
		if entry.IsDir() {
			// dir copy
			if err := dc.OS(
				sfp,
				dfp,
			); nil != err {
				return err
			}
		} else {
			// file copy
			if err := dc.fileCopier.OS(
				sfp,
				dfp,
			); nil != err {
				return err
			}
		}

	}
	return err
}
