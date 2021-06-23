package cmd

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

func init() {
	x := cmdbox.Add("version")
	x.Usage = `[<subcmd>]`
	x.Summary = `provide version and legal information`
	x.Version = `v1.0.0`
	x.Copyright = `Copyright 2021 Robert S Muhlestein`
	x.License = `Apache-2`

	x.Description = `
		Use this command to print version and legal information. If an
		optional <subcmd> is provided the version information for that
		specific subcommand (if any) will be provided instead (since this
		CmdBox command may have been composed from one or more other
		independent command modules before being statically linked).`

	x.Method = func(args []string) error {
		if len(args) == 0 {
			fmt.Println(x.VersionLine())
			return nil
		}
		subcmd := cmdbox.Get(args[0])
		if subcmd == nil {
			fmt.Println(cmdbox.Unimplemented)
		}
		fmt.Println(subcmd.VersionLine())
		return nil
	}

}
