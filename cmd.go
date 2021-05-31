package cmd

import (
	"github.com/rwxrob/cmdbox"
	"github.com/rwxrob/cmdbox/fmt"
)

func init() {
	x := cmdbox.New("version")
	x.Usage = ``
	x.Summary = `provide version and legal information`

	x.Method = func(args []string) error {
		fmt.Println("would show version")
		return nil
	}

}
