package keychain

import (
	"crypto/sha1"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
)

var (
	MIN_ITERATIONS = 1000
	KEY_LENGTH     = 32
)

type EncryptionKey struct {
	Data       string `json:"data"`
	Validation string `json:"validation"`
	Level      string `json:"level"`
	Identifier string `json:"identifier"`
	Iterations int    `json:"iterations"`

	decryptedKey []byte
}

func (encryptionKey *EncryptionKey) Unlock(password string) {
	salt, encryptedKey := ParseSaltAndEncryptedDataFromBase64(encryptionKey.Data)
	derivedKey, iv := encryptionKey.pbkdf2(salt, []byte(password))

	encryptionKey.decryptedKey = AES128_Decrypt(derivedKey, iv, encryptedKey)
	encryptionKey.validateDecryptedKey()
}

func (encryptionKey *EncryptionKey) Decrypt(b64data string) []byte {
	salt, encryptedKey := ParseSaltAndEncryptedDataFromBase64(b64data)
	derivedKey, iv := PBKDF1(encryptionKey.decryptedKey, salt)

	return AES128_Decrypt(derivedKey, iv, encryptedKey)
}

func (encryptionKey *EncryptionKey) validateDecryptedKey() {
	decryptedValidation := encryptionKey.Decrypt(encryptionKey.Validation)

	println(hex.EncodeToString(decryptedValidation))
	println(hex.EncodeToString(encryptionKey.decryptedKey))
	if hex.EncodeToString(decryptedValidation) != hex.EncodeToString(encryptionKey.decryptedKey) {
		panic("Key was invalid!")
	}
}

func (encryptionKey *EncryptionKey) pbkdf2(password []byte, salt []byte) ([]byte, []byte) {
	keyAndIV := pbkdf2.Key(password, salt, encryptionKey.Iterations, KEY_LENGTH, sha1.New)

	return keyAndIV[0 : KEY_LENGTH/2], keyAndIV[KEY_LENGTH/2:]
}
