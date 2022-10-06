package utils

import (
	"math/rand"
	"time"
)

func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	letter := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
