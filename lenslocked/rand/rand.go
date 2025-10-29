package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	nRand, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("bytes: %w", err)
	}
	if nRand < n {
		return nil, fmt.Errorf("not enough bytes were returned")
	}

	return b, nil
}

func ToString(n int) (string, error) {
	s, err := Bytes(n)
	if err != nil {
		return "", fmt.Errorf("string: %w", err)
	}

	return base64.URLEncoding.EncodeToString(s), nil
}


