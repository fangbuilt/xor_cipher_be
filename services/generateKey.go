package services

import (
	"math/rand"
	"time"
)

type GenerateKeyFunc func() string

func GenerateKey() string {
	rand.Seed(time.Now().UnixNano())

	charset := "abcdefghijklmnopqrstuvwxyz"

	shuff := []rune(charset)
	rand.Shuffle(len(shuff), func(i, j int) {
		shuff[i], shuff[j] = shuff[j], shuff[i]
	})

	return string(shuff)
}
