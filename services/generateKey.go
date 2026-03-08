package services

import (
	"math/rand"
)

type GenerateKeyFunc func() string

func GenerateKey() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	shuffle := []rune(charset)

	rand.Shuffle(len(shuffle), func(i, j int) {
		shuffle[i], shuffle[j] = shuffle[j], shuffle[i]
	})

	return string(shuffle)
}
