package main

import (
	"fmt"
	"github.com/dcu/1p/cli"
	//"github.com/dcu/1p/keychain"
)

func main() {
	password := cli.AskPassword("Enter your password: ")

	vaults := cli.FindVaultsForPassword(password)

	for _, vault := range vaults {
		fmt.Printf("Using %s\n", vault.Path)
	}
}
