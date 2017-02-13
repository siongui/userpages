package mylib

import (
	"math/rand"
	"time"
)

func RandomString2(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := ""
	for i := 0; i < strlen; i++ {
		index := rand.Intn(len(chars))
		result += chars[index : index+1]
	}
	return result
}
