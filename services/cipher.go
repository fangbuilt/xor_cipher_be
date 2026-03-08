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

	for i := range textBytes {
		resultBytes[i] = textBytes[i] ^ keyBytes[i%len(keyBytes)]
	}

	result := base64.StdEncoding.EncodeToString(resultBytes)

	return result, nil
}

func XORDecipher(cipheredText, key string) (string, error) {
	if len(key) == 0 {
		return "", fmt.Errorf("key cannot be empty")
	}

	decodedText, err := base64.StdEncoding.DecodeString(cipheredText)
	if err != nil {
		return "", fmt.Errorf("invalid base64 input: %w", err)
	}

	keyBytes := []byte(key)

	resultBytes := make([]byte, len(decodedText))

	for i := range decodedText {
		resultBytes[i] = decodedText[i] ^ keyBytes[i%len(keyBytes)]
	}

	return string(resultBytes), nil
}
