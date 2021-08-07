package blockchain

import (
	"sync"

	"github.com/JihoonPark93/JHCoin/db"
	"github.com/JihoonPark93/JHCoin/utils"
)

const (
	defaultDifficulty  int     = 2
	difficultyInterval int     = 5
	blockInterval      int     = 2
	intervalRange      float32 = 0.1
)

type blockchain struct {
	NewestHash        string `json:"newestHash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveCheckpoint(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func (b *blockchain) recalculateDifficulty() int {
	allBlock := b.Blocks()
	newestBlock := allBlock[0]
	lastRecalculatedBlock := allBlock[difficultyInterval-1]
	actualTime := float32((newestBlock.Timestamp - lastRecalculatedBlock.Timestamp)) / 60
	expectedTime := float32(difficultyInterval * blockInterval)
	if actualTime < expectedTime*(1+intervalRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime > expectedTime*(1-intervalRange) {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return b.recalculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{
				Height: 0,
			}
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis block")
			} else {
				b.restore(checkpoint)
			}
		})
	}
	return b
}
