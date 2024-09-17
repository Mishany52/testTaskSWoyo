package utils

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	rand.New(rand.NewSource(time.Now().Unix()))
	s := make([]rune, n)

	for i := range s{
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}