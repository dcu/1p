package cli

import (
	"fmt"
	"github.com/dcu/1p/keychain"
)

type HelpCommand struct {
}

func (command *HelpCommand) Prepare(args []string) bool {
	usage := `Usage: 1p <command> <args>

Copy command
    Usage: 1p copy <pattern>
    Aliases: c, cp

    Copies the first match to the clipboard.
`

	fmt.Println(usage)

	return false
}

func (command *HelpCommand) Run(vault *keychain.Vault) {
}
