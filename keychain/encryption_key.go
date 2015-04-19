package keychain

import (
	"crypto/sha1"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
)

var (
	MIN_ITERATIONS = 1000
	KEY_SIZE       = 16
)

type EncryptionKey struct {
	Data       string `json:"data"`
	Validation string `json:"validation"`
	Level      string `json:"level"`
	Identifier string `json:"identifier"`
	Iterations int    `json:"iterations"`

	decryptedKey []byte
}

func (encryptionKey *EncryptionKey) Unlock(password string) bool {
	salt, encryptedKey := ParseSaltAndEncryptedDataFromBase64(encryptionKey.Data)
	derivedKey, iv := encryptionKey.pbkdf2([]byte(password), salt)

	encryptionKey.decryptedKey = AES128_Decrypt(derivedKey, iv, encryptedKey)

	if encryptionKey.validateDecryptedKey() {
		return true
	}

	encryptionKey.Lock()
	return false
}

func (encryptionKey *EncryptionKey) Lock() {
	encryptionKey.decryptedKey = nil
}

func (encryptionKey *EncryptionKey) Decrypt(b64data string) []byte {
	salt, encryptedKey := ParseSaltAndEncryptedDataFromBase64(b64data)
	derivedKey, iv := PBKDF1(encryptionKey.decryptedKey, salt)

	return AES128_Decrypt(derivedKey, iv, encryptedKey)
}

func (encryptionKey *EncryptionKey) validateDecryptedKey() bool {
	if len(encryptionKey.decryptedKey) == 0 {
		return false
	}

	decryptedValidation := encryptionKey.Decrypt(encryptionKey.Validation)

	if hex.EncodeToString(decryptedValidation) == hex.EncodeToString(encryptionKey.decryptedKey) {
		return true
	}

	return false
}

func (encryptionKey *EncryptionKey) isUnlocked() bool {
	if len(encryptionKey.decryptedKey) > 0 {
		return true
	}

	return false
}

func (encryptionKey *EncryptionKey) pbkdf2(password []byte, salt []byte) (key []byte, iv []byte) {
	keyAndIV := pbkdf2.Key(password, salt, encryptionKey.Iterations, KEY_SIZE*2, sha1.New)

	return keyAndIV[:KEY_SIZE], keyAndIV[KEY_SIZE:]
}
