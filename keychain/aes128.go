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

	return opensslUnpadding(decryptedData, aes.BlockSize())
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

	data = opensslPadding(data, aes.BlockSize())

	encryptedData := make([]byte, len(data))
	aes.CryptBlocks(encryptedData, data)

	return encryptedData
}

func AES128_EncryptToBase64(key []byte, iv []byte, data []byte) string {
	encryptedData := AES128_Encrypt(key, iv, data)

	return base64.StdEncoding.EncodeToString(encryptedData)
}

func opensslPadding(src []byte, blockSize int) []byte {
	padding := blockSize - (len(src) % blockSize)
	if padding == 0 {
		padding = blockSize
	}
	src = append(src, byte(0xa))
	for i := 1; i < padding; i++ {
		src = append(src, byte(padding-1))
	}
	return src
}

func opensslUnpadding(src []byte, blockSize int) []byte {
	if len(src) == 0 {
		return nil
	}

	padding := src[len(src)-1]
	if int(padding) > len(src) || int(padding) > blockSize {
		return nil
	} else if padding == 0 {
		return nil
	}

	for i := len(src) - 1; i > len(src)-int(padding)-1; i-- {
		if src[i] != padding {
			return nil
		}
	}
	return src[:len(src)-int(padding)-1]
}
