package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*block
}

// uppercase : public
// lowercase : private
var b *blockchain
var once sync.Once

func (b *block) calculateHash() {
	b.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(b.Data+b.PrevHash)))
}
func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

func createBlock(data string) *block {
	newBloc := block{data, "", getLastHash()}
	newBloc.calculateHash()
	return &newBloc
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis block")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*block {
	return b.blocks
}
