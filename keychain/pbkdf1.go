package keychain

import (
	"crypto/md5"
)

func PBKDF1(key []byte, salt []byte) ([]byte, []byte) {
	derivedKey := md5.Sum(append(key, salt...))
	iv := md5.Sum(append(append(derivedKey[:], key...), salt...))

	return derivedKey[:], iv[:]
}
