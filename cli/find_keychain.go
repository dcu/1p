package cli

import (
	"fmt"
	"github.com/dcu/1p/keychain"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strconv"
)

func FindVaultsForPassword(password string) []*keychain.Vault {
	paths := FindAllVaultPaths()
	vaults := []*keychain.Vault{}

	for _, path := range paths {
		vault := keychain.NewVault(path)
		key := vault.UnlockEncryptionKey(password)

		if key != nil {
			vaults = append(vaults, vault)
		}
	}

	return vaults
}

func FindDefaultVault() *keychain.Vault {
	path := FindVaultPath()

	return keychain.NewVault(path)
}

func FindAllVaultPaths() []string {
	homeDir := findUserHome()
	pattern := filepath.Join(homeDir, "Dropbox", "1Password", "*.agilekeychain")

	matches, err := filepath.Glob(pattern)
	if err != nil {
		panic(err)
	}

	return matches
}

func FindVaultPath() string {
	candidates := FindAllVaultPaths()

	fmt.Printf("Please select one:\n")
	for index, candidate := range candidates {
		fmt.Printf("%d) %s\n", index+1, candidate)
	}

	answer := AskOption([]string{"1", "2"})

	answerIndex, _ := strconv.Atoi(answer)

	return candidates[answerIndex-1]
}

func findUserHome() string {
	var homeDir string

	user, err := user.Current()
	if err == nil {
		homeDir = user.HomeDir
	}

	if homeDir == "" {
		homeDir = os.Getenv("HOME")
	}

	if homeDir == "" {
		wd, _ := os.Getwd()
		homeRx := regexp.MustCompile(`^/home/[^/]+`)

		matches := homeRx.FindStringSubmatch(wd)
		homeDir = matches[0]
	}

	return homeDir
}
