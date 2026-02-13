package services

import (
	"encoding/base64"
	"fmt"
)

type CipherFunc func(text, key string) (string, error)

func XOREncrypt(text, key string) (string, error) {
	if len(key) == 0 {
		return "", fmt.Errorf("key cannot be empty")
	}

	result := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ key[i%len(key)]
	}
	return base64.StdEncoding.EncodeToString(result), nil
}

func XORDecrypt(encoded, key string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	result := make([]byte, len(data))
	for i := range data {
		result[i] = data[i] ^ key[i%len(key)]
	}
	return string(result), nil
}
