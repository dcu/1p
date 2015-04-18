package keychain

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type EncryptionKeys struct {
	Vault *Vault
	SL1   string           `json:"SL1"`
	SL2   string           `json:"SL2"`
	SL3   string           `json:"SL3"`
	SL4   string           `json:"SL4"`
	SL5   string           `json:"SL5"`
	List  []*EncryptionKey `json:"list"`
}

func NewEncryptionKeys(vault *Vault) *EncryptionKeys {
	encryptedKeys := &EncryptionKeys{Vault: vault}
	encryptedKeys.load()

	return encryptedKeys
}

func (encryptionKeys *EncryptionKeys) UnlockKey(identifier string, password string) *EncryptionKey {
	for _, encryptionKey := range encryptionKeys.List {
		if encryptionKey.Identifier == identifier {
			encryptionKey.Unlock(password)
			return encryptionKey
		}
	}

	return nil
}

func (encryptionKeys *EncryptionKeys) load() {
	path := encryptionKeys.encryptionKeysPath()
	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, encryptionKeys)

	if err != nil {
		panic(err)
	}
}

func (encryptedKeys *EncryptionKeys) encryptionKeysPath() string {
	return filepath.Join(encryptedKeys.Vault.Path, "data/default/encryptionKeys.js")
}
