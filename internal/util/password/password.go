package password

import "golang.org/x/crypto/bcrypt"

func HashPassword(password, salt string) (string, error) {
	password += salt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(password string, hash string, salt string) bool {
	password += salt
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
