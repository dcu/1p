package keychain

import ()

var (
	DefaultSecurityLevel = "SL5"
)

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

func (vault *Vault) FindEncryptionKey() *EncryptionKey {
	return vault.FindEncryptionKeyBySecurityLevel("SL5")
}

func (vault *Vault) FindEncryptionKeyBySecurityLevel(securityLevel string) *EncryptionKey {
	for _, ekey := range vault.EncryptionKeys.List {
		if ekey.Level == securityLevel {
			return ekey
		}
	}

	return nil
}

func (vault *Vault) UnlockEncryptionKey(password string) *EncryptionKey {
	encryptionKey := vault.FindEncryptionKeyBySecurityLevel(DefaultSecurityLevel)

	if encryptionKey.Unlock(password) {
		return encryptionKey
	}

	return nil
}
