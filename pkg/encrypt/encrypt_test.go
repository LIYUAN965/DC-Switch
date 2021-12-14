package encrypt

import (
	"testing"
)

func TestGenSecKey(t *testing.T) {
	_, err := GenEncryptKey()
	if err != nil {
		t.Error(err)
	}
}

func TestEncrypt(t *testing.T) {
	password := "my-password"
	keyString, err := GenEncryptKey()
	if err != nil {
		t.Error(err)
	}
	secString := Encrypt(password, keyString)
	decryptStr := Decrypt(secString, keyString)
	if decryptStr != password {
		t.Errorf("%v != %v\n", password, decryptStr)
	}
}

func BenchmarkDecrypt(b *testing.B) {
	password := "my-password"
	key, _ := GenEncryptKey()
	pass := Encrypt(password, key)
	for n := 0; n < b.N; n++ {
		Decrypt(pass, key)
	}
}
