package lib

import (
	"crypto/rand"
	"encoding/base64"
)

func RandomString() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
