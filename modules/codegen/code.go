package codegen

import (
	"math/rand"
	"strconv"
	"time"
)

var rng *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomNumber() string {
	min := 100000
	max := 999999
	code := rng.Intn(max-min+1) + min
	c := strconv.Itoa(code)
	return c
}

func GenerateRandomAlphanumeric(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rng.Intn(len(charset))]
	}
	return string(b)
}
