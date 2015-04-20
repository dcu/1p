package cli

import (
	"fmt"
	"github.com/dcu/1p/keychain"
	"github.com/mgutz/ansi"
	"strconv"
)

type CopyCommand struct {
	Args []string
}

func (command *CopyCommand) Prepare(args []string) bool {
	if len(args) < 3 {
		fmt.Println("The copy command needs a pattern.")
		return false
	}

	command.Args = args[2:]
	return true
}

func (command *CopyCommand) Run(vault *keychain.Vault) {
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

	CopyToClipboard(item.Password())
	fmt.Printf("%s password was copied to clipboard.\n", ansi.Color(item.Name, "green+h:black"))
}

func askItemToUser(items []*keychain.Item) *keychain.Item {
	options := []string{}
	for index, item := range items {
		fmt.Printf(ansi.Color("%d) ", "yellow+h"), index+1)
		fmt.Printf("%s (%s)\n", item.Name, item.Url)

		options = append(options, strconv.Itoa(index+1))
	}

	answer := AskOption(options)
	index, _ := strconv.Atoi(answer)

	return items[index-1]
}
