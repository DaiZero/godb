package util

import (
	"math/rand"
	"time"
)

func RandName(nameLen int) string {
	var  words = []byte("asdfghjklqwertyuiopzxcvbnm")
	result :=make([]byte,nameLen)
	rand.Seed(time.Now().Unix())
	for i:=range result {
		result[i] =words[rand.Intn(len(words))]
	}
	return string(result)
}