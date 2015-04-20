package cli

import (
	"fmt"
	"github.com/dcu/1p/keychain"
	"strconv"
)

type CopyCommand struct {
	Args []string
}

func (command *CopyCommand) Prepare(args []string) bool {
	if len(args) == 0 {
		fmt.Println("The copy command needs a pattern.")
		return false
	}

	command.Args = args
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

	fmt.Println(item.Name)
}

func askItemToUser(items []*keychain.Item) *keychain.Item {
	options := []string{}
	for index, item := range items {
		fmt.Printf("%d) %s (%s)\n", index+1, item.Name, item.Url)

		options = append(options, strconv.Itoa(index+1))
	}

	answer := AskOption(options)
	index, _ := strconv.Atoi(answer)

	return items[index-1]
}
