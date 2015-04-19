package keychain

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"strings"
	"testing"
)

func Test_ParseSaltAndEncryptedDataFromBase64(t *testing.T) {
	for _, expectation := range AES128_Expectations {
		opensslSalt, opensslEncryptedData := ParseSaltAndEncryptedDataFromBase64(expectation["openssl"])

		decodedKey, _ := hex.DecodeString(expectation["key"])
		decodedIV, _ := hex.DecodeString(expectation["iv"])
		decodeEncrypted, _ := base64.StdEncoding.DecodeString(expectation["encrypted"])

		opensslSaltHex := strings.ToUpper(hex.EncodeToString(opensslSalt))
		if len(expectation["salt"]) > 0 && opensslSaltHex != expectation["salt"] {
			t.Errorf("Unexpected salt: %#v != %#v", opensslSaltHex, expectation["salt"])
		}

		derivedKey, iv := PBKDF1([]byte(expectation["password"]), opensslSalt)

		if !bytes.Equal(decodedIV, iv) {
			t.Errorf("Unexpected iv: %#v != %#v", decodedIV, iv)
		}

		if !bytes.Equal(derivedKey, decodedKey) {
			t.Errorf("Unexpected key: %#v != %#v", derivedKey, decodedKey)
		}

		result := AES128_Decrypt(derivedKey, decodedIV, opensslEncryptedData)
		result2 := AES128_Decrypt(derivedKey, decodedIV, decodeEncrypted)
		if !bytes.Equal(result, result2) {
			t.Errorf("Unexpected results: %#v != %#v\n", result, result2)
		}
	}
}
