package main

import (
	"crypto/rand"
	mr "math/rand"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"time"
)

var keyLen int
var passNum int
var method string

const CHARSET = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"0123456789"

func main() {
	flag.IntVar(&keyLen, "l", 32, "randkey -l 32")
	flag.IntVar(&passNum, "n", 1, "randkey -n 5")
	flag.StringVar(&method, "m", "base64", "randkey -m hex")
	flag.Parse()

	for i := 1; i <= passNum; i++ {
		fmt.Println(randomKey(keyLen))
	}
}

// RandHexOfSize returns a random hexadecimal string
func randomKey(size int) string {
	b, err := randBytes()
	if err != nil {
		panic("Cannot generate random bytes")
	}

	var str string

	switch method {
	case "hex":
		str = hex.EncodeToString(b)
	case "base64":
		str = base64.URLEncoding.EncodeToString(b)
	case "string":
		str = randString(size)
	default:
		str = base64.URLEncoding.EncodeToString(b)
	}

	// truncate to the provided size
	return str[:size]
}

func randBytes() ([]byte, error) {
	byt := 64

	if keyLen > 64 {
		byt = keyLen
	}

	b := make([]byte, byt) // always generate 64 bytes
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func randString(n int) string {
	mr.Seed(time.Now().UnixNano())

	b := make([]byte, n)

	for i := range b {
		b[i] = CHARSET[mr.Intn(len(CHARSET))]
	}

	return string(b)
}
