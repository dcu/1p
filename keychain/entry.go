package keychain

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type Entry struct {
	Vault     *Vault
	UUID      string
	Type      string
	Name      string
	Url       string
	CreatedAt float64
}

type EntryDetails struct {
	UUID          string `json:"uuid"`
	UpdatedAt     int64  `json:"updatedAt"`
	SecurityLevel string `json:"securityLevel"`
	ContentsHash  string `json:"contentsHash"`
	Title         string `json:"title"`
	Encrypted     string `json:"encrypted"`
	TxTimestamp   int64  `json:"txTimestamp"`
	CreatedAt     int64  `json:"createdAt"`
	TypeName      string `json:"typeName"`
}

func NewEntry(vault *Vault, values []interface{}) *Entry {
	entry := &Entry{Vault: vault}

	entry.UUID = values[0].(string)
	entry.Type = values[1].(string)
	entry.Name = values[2].(string)
	entry.Url = values[3].(string)
	entry.CreatedAt = values[4].(float64)

	return entry
}

func (entry *Entry) Details() *EntryDetails {
	entryDetails := &EntryDetails{}

	path := entry.detailsPath()
	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &entryDetails)
	if err != nil {
		panic(err)
	}

	return entryDetails
}

func (entry *Entry) detailsPath() string {
	return filepath.Join(entry.Vault.Path, "data/default", entry.UUID+".1password")
}
