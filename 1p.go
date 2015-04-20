package main

import (
	"fmt"
	"github.com/dcu/1p/cli"
	"github.com/mgutz/ansi"
	"os"
)

func main() {
	commandName := "help"
	if len(os.Args) > 1 {
		commandName = os.Args[1]
	}
	command := cli.FindCommand(commandName)

	if !command.Prepare(os.Args) {
		return
	}

	password := cli.AskPassword("Enter your password: ")
	vaults := cli.FindVaultsForPassword(password)

	if len(vaults) == 0 {
		fmt.Println("Couldn't find a valid keychain for the given password.")
	}

	for _, vault := range vaults {
		fmt.Printf("Using %s\n", ansi.Color(vault.Path, "green+h:black"))
		command.Run(vault)
	}
}
