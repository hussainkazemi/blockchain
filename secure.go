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

func CreateGoodHash(DL string, data ...interface{}) ([]byte, string) {
	flag := false
	var h []byte
	var nonce string
	dl := getDifficultyLevel(DL)
	for !flag {
		nonce = GetRandomNonce()
		data[len(data)] = nonce
		h = createHahs(data...)
		flag = isHashOk(h, dl)
	}
	return h, nonce
}

func isHashOk(hash []byte, difficultyLevel []byte) bool {
	if bytes.Equal(hash[28:], difficultyLevel) {
		return true
	} 
	return false
}

func getDifficultyLevel(difficultyLevel string) []byte {
	var ret []byte
	for _, alp := range difficultyLevel {
		ret = append(ret, int(alp))
	}
	return ret
}