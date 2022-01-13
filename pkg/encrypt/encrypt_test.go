package encrypt

import (
	"fmt"
	"testing"
)

func TestGenSecKey(t *testing.T) {
	_, err := GenEncryptKey()
	if err != nil {
		t.Error(err)
	}
}

func TestEncrypt(t *testing.T) {
	password := ""
	keyString, err := GenEncryptKey()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(keyString)
	fmt.Println(password)
	secString := Encrypt(password, keyString)
	fmt.Println(secString)
	decryptStr := Decrypt(secString, keyString)
	fmt.Println(decryptStr)
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
