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

	data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		data = []byte(text)
	}

	keyBytes := []byte(key)
	result := make([]byte, len(data))

	for i := range data {
		result[i] = data[i] ^ keyBytes[i%len(keyBytes)]
	}

	return base64.StdEncoding.EncodeToString(result), nil
}
