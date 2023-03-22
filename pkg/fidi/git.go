package fidi

import (
	_ "embed"

	"github.com/filipVisko/fidi/pkg/git"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

//go:embed snippets/help.md
var helpDoc string

var Cmd = &Z.Cmd{
	Name:        "fidi",
	Summary:     "git workflow for replacing branches with worktrees",
	Description: helpDoc,
	Commands: []*Z.Cmd{
		help.Cmd,
		AddCmd,
		PullCmd,
		RemoveCmd,
		CloneCmd,
	},
}

//go:embed snippets/add.md
var addDoc string

var AddCmd = &Z.Cmd{
	Name:        "add",
	Summary:     "create a new worktree",
	Description: addDoc,
	Commands:    []*Z.Cmd{help.Cmd},
	NumArgs:     1,
	Call: func(method *Z.Cmd, args ...string) error {
		return git.AddWorktree(args[0])
	},
}

//go:embed snippets/pull.md
var pullDoc string

var PullCmd = &Z.Cmd{
	Name:        "pull",
	Summary:     "pull changes from a remote branch into its worktree",
	Description: pullDoc,
	Commands:    []*Z.Cmd{help.Cmd},
	NumArgs:     1,
	Call: func(method *Z.Cmd, args ...string) error {
		return git.PullBranch(args[0])
	},
}

//go:embed snippets/clone.md
var cloneDoc string

var CloneCmd = &Z.Cmd{
	Name:        "clone",
	Summary:     "clones a repo into a newly created directory",
	Description: cloneDoc,
	Commands:    []*Z.Cmd{help.Cmd},
	NumArgs:     1,
	Call: func(method *Z.Cmd, args ...string) error {
		return git.CloneRepo(args[0])
	},
}

//go:embed snippets/remove.md
var removeDoc string

var RemoveCmd = &Z.Cmd{
	Name:        "remove",
	Summary:     "remove a worktree and its branch",
	Description: removeDoc,
	Commands:    []*Z.Cmd{help.Cmd},
	NumArgs:     1,
	Call: func(method *Z.Cmd, args ...string) error {
		return git.RemoveWorktree(args[0], true)
	},
}
