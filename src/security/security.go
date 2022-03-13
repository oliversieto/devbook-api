package security

import "golang.org/x/crypto/bcrypt"

func Hash(phrase string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(phrase), bcrypt.DefaultCost)
}

func ComparePhrases(hashedPhrase, phrase string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPhrase), []byte(phrase))
}
