package cli

import (
	"fmt"
	"github.com/dcu/1p/keychain"
	"github.com/mgutz/ansi"
	"strconv"
)

var (
	Commands = map[string]Command{
		"help":  &HelpCommand{},
		"copy":  &CopyCommand{},
		"query": &QueryCommand{},
	}
)

type Command interface {
	Run(vault *keychain.Vault)
	Prepare(args []string) bool
}

func FindCommand(name string) Command {
	switch name {
	case "c", "cp":
		{
			name = "copy"
		}
	case "q":
		{
			name = "query"
		}
	case "h":
		{
			name = "help"
		}
	}

	command := Commands[name]
	if command == nil {
		command = Commands["help"]
	}

	return command
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
