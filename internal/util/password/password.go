package password

import "golang.org/x/crypto/argon2"

func HashPassword(password string, salt []byte) string {
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	return string(hash)
}

func VerifyPassword(password string, hash string, salt []byte) bool {
	return hash == HashPassword(password, salt)
}
