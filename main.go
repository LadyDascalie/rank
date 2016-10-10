package main

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var keyLen int
var passNum int
var method string

func main() {
	flag.IntVar(&keyLen, "l", 32, "randkey -l 32")
	flag.IntVar(&passNum, "n", 1, "randkey -n 5")
	flag.StringVar(&method, "m", "hex", "randkey -m base64")
	flag.Parse()

	for i := 1; i <= passNum; i++ {
		fmt.Println(RandHexOfSize(keyLen))
	}
}

func RandHexOfSize(size int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	b, err := randBytes(size)
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
		str = hex.EncodeToString(b)
	}

	// truncate to the provided size
	return str[:size]
}

func randBytes(size int) ([]byte, error) {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, size)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}
