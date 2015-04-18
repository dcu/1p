package keychain

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type Contents struct {
	Vault   *Vault
	Entries []*Entry
}

func NewContents(vault *Vault) *Contents {
	contents := &Contents{Vault: vault}
	contents.load()

	return contents
}

func (contents *Contents) load() {
	path := contents.contentsPath()
	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	jsonContent := make([][]interface{}, 0)
	err = json.Unmarshal(data, &jsonContent)
	if err != nil {
		panic(err)
	}

	for _, values := range jsonContent {
		entry := NewEntry(contents.Vault, values)
		contents.Entries = append(contents.Entries, entry)
	}
}

func (contents *Contents) contentsPath() string {
	return filepath.Join(contents.Vault.Path, "data/default/contents.js")
}
