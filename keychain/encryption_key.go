package keychain

import (
	"crypto/sha1"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
	"log"
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

func (encryptionKey *EncryptionKey) Unlock(password string) bool {
	salt, encryptedKey := ParseSaltAndEncryptedDataFromBase64(encryptionKey.Data)
	log.Printf("Encrypted key: %#v", encryptedKey)

	derivedKey, iv := encryptionKey.pbkdf2([]byte(password), salt)
	log.Printf("Derived key: %#v", derivedKey)

	encryptionKey.decryptedKey = AES128_Decrypt(derivedKey, iv, encryptedKey)

	log.Printf("Decrypted key: %#v", encryptionKey.decryptedKey)

	return encryptionKey.validateDecryptedKey()
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
	log.Printf("Validation: %#v\n", decryptedValidation)

	println(hex.EncodeToString(decryptedValidation))
	println(hex.EncodeToString(encryptionKey.decryptedKey))
	if hex.EncodeToString(decryptedValidation) == hex.EncodeToString(encryptionKey.decryptedKey) {
		return true
	}

	return false
}

func (encryptionKey *EncryptionKey) pbkdf2(password []byte, salt []byte) (key []byte, iv []byte) {
	keyAndIV := pbkdf2.Key(password, salt, encryptionKey.Iterations, KEY_LENGTH, sha1.New)

	return keyAndIV[:KEY_LENGTH/2], keyAndIV[KEY_LENGTH/2:]
}
