package main

import (
	"fmt"
	"github.com/dcu/1p/cli"
	"os"
)

func main() {
	commandName := "help"
	if len(os.Args) > 1 {
		commandName = os.Args[1]
	}
	command := cli.FindCommand(commandName)

	if !command.Prepare(os.Args[2:]) {
		return
	}

	password := cli.AskPassword("Enter your password: ")
	vaults := cli.FindVaultsForPassword(password)

	if len(vaults) == 0 {
		fmt.Println("Couldn't find a valid keychain for the given password.")
	}

	for _, vault := range vaults {
		fmt.Printf("Using %s\n", vault.Path)
		command.Run(vault)
	}
}
