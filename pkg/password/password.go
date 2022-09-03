package password

import "golang.org/x/crypto/bcrypt"

// Hash password to bcrypt hashes.
func HashPassword(s string) (string, error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(s), 8)
	if err != nil {
		return "", err
	}

	return string(passwordBytes), nil
}

// Checking the password hashes with plain password.
func CheckPassword(h string, p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))

	return err == nil
}
