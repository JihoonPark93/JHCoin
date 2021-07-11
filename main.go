package main

import (
	"fmt"

	"github.com/JihoonPark93/JHCoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("SecondBlock")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")

	for _, block := range chain.AllBlocks() {
		fmt.Println("")
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
	}
}
