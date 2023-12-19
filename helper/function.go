package helper

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateUniqueID(username, fullName string) string {
	combinedString := username + fullName
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(combinedString))
	hashedID := fmt.Sprintf("%x", sha256Hash.Sum(nil))
	return hashedID[:20]
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
