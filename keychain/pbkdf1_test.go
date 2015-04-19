package keychain

import (
	"encoding/hex"
	"strings"
	"testing"
)

func Test_PBKDF1Expectations(t *testing.T) {
	for _, expectation := range AES128_Expectations {
		salt, _ := hex.DecodeString(expectation["salt"])
		derivedKey, iv := PBKDF1([]byte(expectation["password"]), salt)

		derivedKeyHex := strings.ToUpper(hex.EncodeToString(derivedKey))
		ivHex := strings.ToUpper(hex.EncodeToString(iv))

		if expectation["key"] != derivedKeyHex {
			t.Errorf("Failed expectation: %#v. Received: key=%s iv=%s", expectation, derivedKeyHex, ivHex)
		}
	}
}
