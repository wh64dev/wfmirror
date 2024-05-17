package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	pwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(pwd)
}

func CheckPasswordHash(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
