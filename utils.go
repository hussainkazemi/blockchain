package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	RandomLen = 5
)

func GetRandomNonce() string {
	alp := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	ln := len(alp)
	var ret string
	for i := 0; i < RandomLen; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		ret = fmt.Sprintf("%s%s", ret, alp[r1.Intn(ln)])
	}
	return ret
}
