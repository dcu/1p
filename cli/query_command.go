package cli

import (
	"fmt"
	"github.com/dcu/1p/keychain"
	"github.com/mgutz/ansi"
)

type QueryCommand struct {
	Args []string
}

func (command *QueryCommand) Prepare(args []string) bool {
	if len(args) < 3 {
		fmt.Println("The query command needs a pattern.")
		return false
	}

	command.Args = args[2:]
	return true
}

func (command *QueryCommand) Run(vault *keychain.Vault) {
	items := vault.Contents.FindAllItemsByPattern(command.Args[0])
	if len(items) == 0 {
		fmt.Println("Item not found in the keychain.")
		return
	}

	item := items[0]
	if len(items) > 1 {
		item = askItemToUser(items)
	}

	key := vault.FindEncryptionKey()

	if !key.IsUnlocked() {
		fmt.Println("Key is not unlocked.")
		return
	}

	encryptedData := item.EncryptedData()
	details := item.Details()

	fmt.Printf("%s: %s\n", ansi.Color("Type", "cyan+h"), details.TypeName)
	for _, field := range encryptedData.Fields {
		if field["type"] != "P" && field["designation"] != "password" {
			fmt.Printf("%s: %s\n", ansi.Color(field["name"], "cyan+h"), field["value"])
		}
	}
}
