package cmd

import "fmt"

func DeleteCmd(args ...string) error {
	fmt.Println("test called with args:", args)
	return nil
}
