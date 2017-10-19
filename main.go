package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	mr "math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
)

var (
	keyLen int
	dict   Dictionary
)

const (
	// DefaultKeyLen is the default key length
	DefaultKeyLen = 32
	// Charset represents the available characters for the string method
	Charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789"
)

type (
	// Dictionary contains the web2 Dictionary contents
	Dictionary struct {
		Total int
		Words []string
	}
)

func init() {
	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dict.Total++
		dict.Words = append(dict.Words, strings.ToLower(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

var size int
var number int
var method string

func main() {
	flag.IntVar(&size, "size", 32, "rank -size 32")
	flag.IntVar(&number, "n", 1, "rank -n 5")
	flag.StringVar(&method, "method", "base64", "rank -method words|hex|base64|string")
	flag.Parse()

	for n := 0; n < number; n++ {
		if number == 1 {
			fmt.Print(randomKey(size, method))
			return
		}
		fmt.Println(randomKey(size, method))
	}
}

// Generated is the struct to send back to the browser
type Generated struct {
	Password string `json:"password"`
}

func genHandler(c echo.Context) error {
	lengthParam := c.Param("len")
	method := c.Param("method")

	keyLength, err := strconv.Atoi(lengthParam)
	if err != nil {
		c.Logger().Warn(err)
	}

	if method == "" {
		method = "string"
	}

	var pass string
	if keyLength == 0 {
		pass = randomKey(DefaultKeyLen, method)
	} else {
		pass = randomKey(keyLength, method)
	}

	g := Generated{Password: pass}

	return c.JSON(http.StatusOK, g)
}

// RandHexOfSize returns a random hexadecimal string
func randomKey(size int, method string) string {
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
	case "words":
		return strings.ToLower(randWords(size))
	default:
		str = base64.URLEncoding.EncodeToString(b)
	}

	// truncate to the provided size
	return str[:size]
}
func randWords(size int) (pass string) {
	mr.Seed(time.Now().UnixNano())

	var words []string
	var sep = " "

	for i := 0; i < size; i++ {
		w := dict.Words[mr.Intn(dict.Total-1)]
		words = append(words, w)
	}

	return strings.Join(words, sep)
}

func randBytes() ([]byte, error) {
	byt := 64

	if keyLen > 64 {
		byt = DefaultKeyLen
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
		b[i] = Charset[mr.Intn(len(Charset))]
	}

	return string(b)
}
