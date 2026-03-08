package services

import (
	"encoding/base64"
	"fmt"
)

type CipherFunc func(text, key string) (string, error)

func XORCipher(text, key string) (string, error) {
	if len(key) == 0 {
		return "", fmt.Errorf("key cannot be empty")
	}

	textBytes := []byte(text)
	keyBytes := []byte(key)

	resultBytes := make([]byte, len(textBytes))

	for i := 0; i < len(textBytes); i++ {
		resultBytes[i] = textBytes[i] ^ keyBytes[i%len(keyBytes)]
	}

	return base64.StdEncoding.EncodeToString(resultBytes), nil
}
