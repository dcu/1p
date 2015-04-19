package keychain

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AES128_Decrypt(key []byte, iv []byte, encryptedData []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	aes := cipher.NewCBCDecrypter(block, iv)

	decryptedData := make([]byte, len(encryptedData))
	aes.CryptBlocks(decryptedData, encryptedData)

	return OpensslUnpadding(decryptedData, aes.BlockSize())
}

func AES128_DecryptFromBase64(key []byte, iv []byte, encryptedDataB64 string) []byte {
	encryptedData, _ := base64.StdEncoding.DecodeString(encryptedDataB64)

	return AES128_Decrypt(key, iv, encryptedData)
}

func AES128_Encrypt(key []byte, iv []byte, data []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	aes := cipher.NewCBCEncrypter(block, iv)

	data = OpensslPadding(data, aes.BlockSize())

	encryptedData := make([]byte, len(data))
	aes.CryptBlocks(encryptedData, data)

	return encryptedData
}

func AES128_EncryptToBase64(key []byte, iv []byte, data []byte) string {
	encryptedData := AES128_Encrypt(key, iv, data)

	return base64.StdEncoding.EncodeToString(encryptedData)
}
