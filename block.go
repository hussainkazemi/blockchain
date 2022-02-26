package main

import (
	"fmt"
	"time"
)

type Block struct {
	Index     int32
	TimeStamp time.Time
	Data      string
	Perv_Hash []byte
	Hash      []byte
	Nonce     int32
}

func (b Block) String() string {
	ret := fmt.Sprintf(string(b.Hash))
	return ret
}
