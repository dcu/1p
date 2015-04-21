package keychain

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type Item struct {
	Vault     *Vault
	UUID      string
	Type      string
	Name      string
	Url       string
	CreatedAt float64
}

type ItemDetails struct {
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

type ItemEncryptedData struct {
	Fields   []map[string]string `json:"fields,omitempty"`
	URLs     []map[string]string `json:"URLs"`
	Password string              `json:"password,omitempty"`
}

func NewItem(vault *Vault, values []interface{}) *Item {
	item := &Item{Vault: vault}

	item.UUID = values[0].(string)
	item.Type = values[1].(string)
	item.Name = values[2].(string)
	item.Url = values[3].(string)
	item.CreatedAt = values[4].(float64)

	return item
}

func (item *Item) Details() *ItemDetails {
	itemDetails := &ItemDetails{}

	path := item.detailsPath()
	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &itemDetails)
	if err != nil {
		panic(err)
	}

	return itemDetails
}

func (item *Item) EncryptedData() *ItemEncryptedData {
	encryptedData := &ItemEncryptedData{}
	details := item.Details()

	key := item.Vault.FindEncryptionKey()
	jsonData := key.Decrypt(details.Encrypted)
	err := json.Unmarshal(jsonData, &encryptedData)
	if err != nil {
		panic(err)
	}

	return encryptedData
}

func (item *Item) Password() string {
	encryptedData := item.EncryptedData()

	if len(encryptedData.Password) > 0 {
		return encryptedData.Password
	}

	for _, field := range encryptedData.Fields {
		if field["type"] == "P" || field["designation"] == "password" {
			return field["value"]
		}
	}

	return ""
}

func (item *Item) detailsPath() string {
	return filepath.Join(item.Vault.Path, "data/default", item.UUID+".1password")
}
