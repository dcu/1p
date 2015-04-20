package keychain

import (
	"encoding/base64"
	"encoding/hex"
	"strings"
	"testing"
)

func Test_AES128_Encrypt(t *testing.T) {
	for _, expectation := range AES128_Expectations {
		iv, err := hex.DecodeString(expectation["iv"])
		if err != nil {
			panic(err)
		}
		key, _ := hex.DecodeString(expectation["key"])
		if err != nil {
			panic(err)
		}

		encryptedData := AES128_EncryptToBase64(key, iv, []byte(expectation["plain"]))

		if encryptedData != expectation["encrypted"] {
			t.Errorf("Error encrypting data -> %s Expected: %s Obtained: %s", expectation["plain"], expectation["encrypted"], encryptedData)
		}

		_, opensslEncryptedData := ParseSaltAndEncryptedDataFromBase64(expectation["openssl"])
		opensslEncryptedDataB64 := base64.StdEncoding.EncodeToString(opensslEncryptedData)

		if opensslEncryptedDataB64 != encryptedData {
			decodedEncrypted, _ := base64.StdEncoding.DecodeString(encryptedData)
			t.Errorf("Error comparing with openssl encrypting data ->  Expected:\n%#v\nObtained:\n%#v", opensslEncryptedData, decodedEncrypted)
		}
	}
}

func Test_AES128_Decrypt(t *testing.T) {
	for _, expectation := range AES128_Expectations {
		iv, _ := hex.DecodeString(expectation["iv"])
		key, _ := hex.DecodeString(expectation["key"])

		decryptedData := AES128_DecryptFromBase64(key, iv, expectation["encrypted"])

		if strings.TrimSuffix(string(decryptedData), "\n") != expectation["plain"] {
			t.Errorf("Error decrypting data -> `%s` Expected: `%#v` Obtained: `%#v`", expectation["encrypted"], expectation["plain"], string(decryptedData))
		}
	}
}
