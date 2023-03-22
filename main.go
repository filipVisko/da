package main

import (
	"github.com/filipVisko/da/pkg/bon"
	"github.com/filipVisko/da/pkg/fidi"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/rwxrob/uniq"
)

func main() {
	var Cmd = &Z.Cmd{
		Name:    "da",
		Summary: "filipVisko's personal command tree monolith",
		Site:    "https://filip.dev",
		Commands: []*Z.Cmd{
			help.Cmd,
			fidi.Cmd,
			uniq.Cmd,
			bon.BonCmd,
		},
	}
	Cmd.Run()
}
