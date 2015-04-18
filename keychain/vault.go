package keychain

import ()

type Vault struct {
	Path           string
	EncryptionKeys *EncryptionKeys
	Contents       *Contents
}

func NewVault(path string) *Vault {
	vault := &Vault{
		Path: path,
	}
	vault.load()

	return vault
}

func (vault *Vault) load() {
	vault.EncryptionKeys = NewEncryptionKeys(vault)
	vault.Contents = NewContents(vault)
}

func (vault *Vault) FindById(uuid string) *Entry {
	for _, entry := range vault.Contents.Entries {
		if entry.UUID == uuid {
			return entry
		}
	}

	return nil
}
