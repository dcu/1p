package cli

import (
	"github.com/dcu/1p/keychain"
)

type Command interface {
	Run(vault *keychain.Vault)
	Prepare(args []string) bool
}

func FindCommand(name string) Command {
	switch name {
	case "c", "cp", "copy":
		{
			return &CopyCommand{}
		}
	}

	return &HelpCommand{}
}
