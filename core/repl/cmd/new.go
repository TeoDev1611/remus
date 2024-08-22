package cmd

import "fmt"

func NewCmd(args ...string) error {
	fmt.Println("test called with args:", args)
	return nil
}
