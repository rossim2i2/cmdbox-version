package cmd

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

func init() {
	x := cmdbox.Add("version")
	x.Usage = ``
	x.Summary = `provide version and legal information`

	x.Method = func(args []string) error {
		fmt.Println("would show version")
		return nil
	}

}
