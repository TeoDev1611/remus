package cmd

import (
	"fmt"

	"github.com/TeoDev1611/remus/core/db/new"
)

func NewCmd(args ...string) error {
	var data new.UIData
	if len(args) == 0 {
		data = new.GetInfo("", "", true)
	} else {
		data = new.GetInfo(args[0], args[1], false)
	}
	fmt.Println(data.Name)
	fmt.Println(data.Password)
	fmt.Println(data.Date)
	fmt.Println(data.UID)
	return nil
}
