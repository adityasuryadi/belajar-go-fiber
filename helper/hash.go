package helpers

import (
	"go-blog/exception"

	"golang.org/x/crypto/bcrypt"
)

func GetHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		exception.PanicIfNeeded(err)
	}
	return string(hash)
}

func ComparePassword(password string, hashPassword string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}
