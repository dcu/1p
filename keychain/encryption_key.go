package keychain

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
)

var (
	ZERO_IV          = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	SALT_PREFIX      = []byte{'S', 'a', 'l', 't', 'e', 'd', '_', '_'}
	MIN_ITERATIONS   = 1000
	KEY_LENGTH       = 32
	SALT_START_INDEX = 8
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
	salt, encryptedKey := parseEncryptedBase64(encryptionKey.Data)
	derivedKey, iv := encryptionKey.pbkdf2(salt, []byte(password))

	encryptionKey.decryptedKey = encryptionKey.aesDecrypt(derivedKey, iv, encryptedKey)
	encryptionKey.validateDecryptedKey()
}

func (encryptionKey *EncryptionKey) Decrypt(b64data string) []byte {
	salt, encryptedKey := parseEncryptedBase64(b64data)
	derivedKey, iv := PBKDF1(encryptionKey.decryptedKey, salt)

	return encryptionKey.aesDecrypt(derivedKey, iv, encryptedKey)
}

func (encryptionKey *EncryptionKey) validateDecryptedKey() {
	decryptedValidation := encryptionKey.Decrypt(encryptionKey.Validation)

	println(hex.EncodeToString(decryptedValidation))
	println(hex.EncodeToString(encryptionKey.decryptedKey))
	if hex.EncodeToString(decryptedValidation) != hex.EncodeToString(encryptionKey.decryptedKey) {
		panic("Key was invalid!")
	}
}

func (encryptionKey *EncryptionKey) aesDecrypt(derivedKey []byte, iv []byte, encryptedKey []byte) []byte {
	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		panic(err)
	}
	aes := cipher.NewCBCEncrypter(block, iv)

	decryptedKey := make([]byte, len(encryptedKey))
	aes.CryptBlocks(decryptedKey, encryptedKey)

	return decryptedKey
}

func (encryptionKey *EncryptionKey) pbkdf2(password []byte, salt []byte) ([]byte, []byte) {
	keyAndIV := pbkdf2.Key(password, salt, encryptionKey.Iterations, KEY_LENGTH, sha1.New)

	return keyAndIV[0 : KEY_LENGTH/2], keyAndIV[KEY_LENGTH/2:]
}

func parseEncryptedBase64(b64data string) (salt []byte, key []byte) {
	decodedData, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		panic(err)
	}

	if bytes.HasPrefix(decodedData, SALT_PREFIX) {
		salt = decodedData[SALT_START_INDEX : KEY_LENGTH/2]
		key = decodedData[KEY_LENGTH/2:]
	} else {
		salt = ZERO_IV
		key = decodedData
	}

	return salt, key
}
