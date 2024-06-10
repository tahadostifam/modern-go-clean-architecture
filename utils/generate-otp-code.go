package utils

import (
	"crypto/rand"
	"io"
	"log"
)
var Table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		log.Fatalln(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = Table[int(b[i])%len(Table)]
	}
	return string(b)
}

