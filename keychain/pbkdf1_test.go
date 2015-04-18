package keychain

import (
	"encoding/hex"
	"strings"
	"testing"
)

/*
Generate with salt:
openssl enc -aes-128-cbc -p -k "<password>" -S "<hex salt>" -salt -a -p <<<"data"

Generate without salt:
openssl enc -aes-128-cbc -p -k "<password>" -nosalt -a -p <<<"data"

*/
var (
	Expectations = []map[string]string{
		map[string]string{
			"password": "password",
			"salt":     "AAAAAAAAAAAAAAAA",
			"key":      "0471219B5CE9A0F64C27093DEA07B362",
			"iv":       "8070AC35A7FD31ACB828A4FA828F566E",
		},
		map[string]string{
			"password": "password",
			"salt":     "ABABABABABABABAB",
			"key":      "95D1F52CECEBA04EC6F499DBB9A0080F",
			"iv":       "170CE82107B5D230B708BF30FCC0FE52",
		},
		map[string]string{
			"password": "#4Sz{[.4SksZ!",
			"salt":     "A000000000000000",
			"key":      "238973C6764117397E19A9EEA574BFC2",
			"iv":       "B9DC36CBEFAFC11A90D0391ECC4E7C4B",
		},
		map[string]string{
			"password": "#4Sz{[.4SksZ!",
			"salt":     "",
			"key":      "8FE4EFD29E3DB807021C20396FF0DE67",
			"iv":       "B1B5120580211898B1B4526186459B0A",
		},
		map[string]string{
			"password": "",
			"salt":     "",
			"key":      "D41D8CD98F00B204E9800998ECF8427E",
			"iv":       "59ADB24EF3CDBE0297F05B395827453F",
		},
	}
)

func Test_PBKDF1Expectations(t *testing.T) {
	for _, expectation := range Expectations {
		salt, _ := hex.DecodeString(expectation["salt"])
		derivedKey, iv := PBKDF1([]byte(expectation["password"]), salt)

		derivedKeyHex := strings.ToUpper(hex.EncodeToString(derivedKey))
		ivHex := strings.ToUpper(hex.EncodeToString(iv))

		if expectation["key"] != derivedKeyHex {
			t.Errorf("Failed expectation: %#v. Received: key=%s iv=%s", expectation, derivedKeyHex, ivHex)
		}
	}
}
