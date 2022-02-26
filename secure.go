package main

import (
	"crypto/sha256"
	"fmt"
)

func createHahs(data ...interface{}) []byte {
	h := sha256.New()
	fmt.Fprint(h, data...)
	return h.Sum(nil)
}

func CreateGoodHash(DL int32, data ...interface{}) ([]byte, string) {
	flag := false
	var h []byte
	var nonce string
	for !flag {
		nonce = GetRandomNonce()
		data[len(data)] = nonce
		h = createHahs(data...)
		flag = isGoodHash(h, DL)
	}
	return h, nonce
}

func isGoodHash(hash []byte, difficultyLevel int32) bool {

}
