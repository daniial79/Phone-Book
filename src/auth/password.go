package auth

import "golang.org/x/crypto/bcrypt"

const passwordHashCost = 10

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordHashCost)
	return string(bytes), err
}

func ComparePasswords(hashedPassword, plainTextPassword string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(plainTextPassword),
	)
}
