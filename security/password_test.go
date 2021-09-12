package security

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestEncryptPassword(t *testing.T) {
	hashed, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(hashed))
	}
}

func TestVerifyPassword(t *testing.T) {
	err := bcrypt.CompareHashAndPassword([]byte("$2a$10$bmIG45wxDXbugPBUCpipe.HtrmZuGx5WBkdaEkfkQLIB8UjE1DSC6"), []byte("123456"))

	if err != nil {
		fmt.Println("password cannot matched")
	} else {
		fmt.Println("password matched")
	}
}
