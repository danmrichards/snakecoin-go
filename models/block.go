package models

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Index        int64
	Timestamp    time.Time
	Data         []byte
	PreviousHash []byte
	Hash         []byte
}

// Returns a string representation of the block.
func (b *Block) String() string {
	return fmt.Sprintf(
		"%d%s%s%s",
		b.Index,
		b.Timestamp.Format(time.RFC3339),
		string(b.Data),
		string(b.PreviousHash),
	)
}

// Creates a new SHA256 bash of the block.
func (b *Block) generateHash() []byte {
	h := sha256.New()
	h.Write([]byte(b.String()))

	return h.Sum(nil)
}

// Creates a new Block.
func NewBlock(index int64, timestamp time.Time, data []byte, previousHash []byte) *Block {
	b := &Block{
		Index:        index,
		Timestamp:    timestamp,
		Data:         data,
		PreviousHash: previousHash,
	}

	b.Hash = b.generateHash()
	return b
}
