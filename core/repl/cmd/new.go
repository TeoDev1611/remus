package cmd

import (
	"github.com/TeoDev1611/remus/core/db/new"
)

func NewCmd(args ...string) error {
	var data new.UIData
	if len(args) == 0 {
		data = new.GetInfo("", "", true)
	} else {
		data = new.GetInfo(args[0], args[1], false)
	}
	new.StartDB(data)
	return nil
}
