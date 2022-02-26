package main

import "fmt"

type BlockChain struct {
	Blocks          []Block
	DifficultyLevel string
}

func (bc BlockChain) String() {
	for _, b := range bc.Blocks {
		fmt.Println(b)
	}
}
