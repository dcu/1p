package keychain

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Contents struct {
	Vault   *Vault
	Entries []*Item
}

func NewContents(vault *Vault) *Contents {
	contents := &Contents{Vault: vault}
	contents.load()

	return contents
}

func (contents *Contents) FindItemById(uuid string) *Item {
	for _, item := range contents.Entries {
		if item.UUID == uuid {
			return item
		}
	}

	return nil
}

func (contents *Contents) FindAllItemsByPattern(pattern string) []*Item {
	items := []*Item{}
	for _, item := range contents.Entries {
		patternLC := strings.ToLower(pattern)
		itemNameLC := strings.ToLower(item.Name)

		if strings.Contains(strings.ToLower(itemNameLC), patternLC) {
			items = append(items, item)
		}

		if len(items) >= 10 {
			break
		}
	}

	return items
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
		item := NewItem(contents.Vault, values)
		contents.Entries = append(contents.Entries, item)
	}
}

func (contents *Contents) contentsPath() string {
	return filepath.Join(contents.Vault.Path, "data/default/contents.js")
}
