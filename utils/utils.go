package utils

import "math/rand"

var Rand *rand.Rand

func GenerateCode() string {
	const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 6)

	for i := range code {
		code[i] = letter[Rand.Intn(len(letter))]
	}

	return string(code)
}
