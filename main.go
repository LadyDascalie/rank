package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
)

var keyLen int
var passNum int
var method string

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
	default:
		str = base64.URLEncoding.EncodeToString(b)
	}

	// truncate to the provided size
	return str[:size]
}

func randBytes() ([]byte, error) {
	b := make([]byte, 64) // always generate 64 bytes
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
