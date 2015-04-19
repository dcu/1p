package keychain

import (
	"bytes"
	"encoding/base64"
)

var (
	SALT_PREFIX        = []byte{'S', 'a', 'l', 't', 'e', 'd', '_', '_'}
	SALT_PREFIX_LENGHT = len(SALT_PREFIX)
	SALT_LENGHT        = 8
)

/*
When encrypting with OpenSSL the result is a base64 string. Whe decoded the format is the following:

	Salted__<salt><encrypted data>

The following method parses that format.
*/
func ParseSaltAndEncryptedDataFromBase64(b64data string) (salt []byte, encryptedData []byte) {
	decodedData, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		panic(err)
	}

	if bytes.HasPrefix(decodedData, SALT_PREFIX) {
		skip := SALT_PREFIX_LENGHT + SALT_LENGHT
		salt = decodedData[SALT_PREFIX_LENGHT:skip]
		encryptedData = decodedData[skip:]
	} else {
		salt = make([]byte, 0)
		encryptedData = decodedData
	}

	return salt, encryptedData
}

func OpensslPadding(src []byte, blockSize int) []byte {
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

func OpensslUnpadding(src []byte, blockSize int) []byte {
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
	return src[:len(src)-int(padding)]
}
