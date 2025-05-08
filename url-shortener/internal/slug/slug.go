package slug

import (
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const slugLength = 8

func base62Encode(num int) string {
	var encoded string
	for num > 0 {
		encoded = string(charset[num%62]) + encoded
		num /= 62
	}
	return encoded
}

func GenerateSlug(id int) string {
	encoded := base62Encode(id)
	padLength := slugLength - len(encoded)

	if padLength <= 0 {
		return encoded[:slugLength] // or return error if truncation is unacceptable
	}

	// Generate random prefix
	randomPad := make([]byte, padLength)
	for i := range randomPad {
		randomPad[i] = charset[rand.Intn(len(charset))]
	}

	return string(randomPad) + encoded
}
