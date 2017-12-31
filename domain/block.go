package domain

import (
	"fmt"
	"time"

	"github.com/danmrichards/snakecoin-go/models"
)

// Manually creates an initial "Genesis" block to start the chain. Note
// arbitrary data and previous block hash.
func CreateGenesisBlock() *models.Block {
	return models.NewBlock(0, time.Now(), []byte("Genesis Block"), []byte("0"))
}

// Creates the next block in the chain based on the previous block.
func NextBlock(b *models.Block) *models.Block {
	nextIndex := b.Index
	nextIndex++

	return models.NewBlock(
		nextIndex,
		time.Now(),
		[]byte(fmt.Sprintf("I'm block %d", nextIndex)),
		b.Hash,
	)
}
