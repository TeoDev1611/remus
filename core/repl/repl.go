package repl

import (
	"fmt"

	"github.com/TeoDev1611/remus/core/repl/cmd"
	"github.com/TeoDev1611/remus/lib/icli"
	"github.com/enescakir/emoji"
)

func StartRepl() {
	cli := icli.NewCLI()
	cli.SetWelcomeMessage(fmt.Sprintf("%v Welcome to Remus -%v %v", icli.BLUE, icli.RESET, emoji.Wolf))

	newDBCmd := &icli.BasicCommand{
		Name:        "new",
		Description: "Make a new database",
		Usage:       "new <ARGS>",
		Fn:          cmd.NewCmd,
	}

	deleteDBCmd := &icli.BasicCommand{
		Name:        "delete",
		Description: "Delete a database",
		Usage:       "delete <ARGS>",
		Fn:          cmd.DeleteCmd,
	}

	cli.AddCmd(newDBCmd)
	cli.AddCmd(deleteDBCmd)

	cli.Run()
}
