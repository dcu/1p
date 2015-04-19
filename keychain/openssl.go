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
		panic("hmmm")
		salt = make([]byte, 0)
		encryptedData = decodedData
	}

	return salt, encryptedData
}
