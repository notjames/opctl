package core

import (
	"github.com/opctl/opctl/util/cliexiter"
	"path/filepath"
)

func (this _core) OpCreate(
	path string,
	description string,
	name string,
) {
	if err := this.opCreator.Create(
		filepath.Join(path, name),
		name,
		description,
	); nil != err {
		this.cliExiter.Exit(cliexiter.ExitReq{Message: err.Error(), Code: 1})
		return // support fake exiter
	}
}
