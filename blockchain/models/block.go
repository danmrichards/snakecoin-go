package models

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Index        int64     `json:"index"`
	Timestamp    time.Time `json:"timestamp"`
	Data         *Data     `json:"data"`
	PreviousHash []byte    `json:"-"`
	Hash         []byte    `json:"hash"`
}

// Returns a string representation of the block.
func (b *Block) String() string {
	return fmt.Sprintf(
		"%d%s%s%s",
		b.Index,
		b.Timestamp.Format(time.RFC3339),
		b.Data.String(),
		hex.EncodeToString(b.PreviousHash),
	)
}

// Custom JSON Marshal to return the hash as a string.
func (b *Block) MarshalJSON() ([]byte, error) {
	type Alias Block
	return json.Marshal(&struct {
		Hash string `json:"hash"`
		*Alias
	}{
		Hash:  hex.EncodeToString(b.Hash),
		Alias: (*Alias)(b),
	})
}

// Creates a new SHA256 bash of the block.
func (b *Block) generateHash() []byte {
	h := sha256.New()
	h.Write([]byte(b.String()))

	return h.Sum(nil)
}

// Creates a new Block.
func NewBlock(index int64, timestamp time.Time, data *Data, previousHash []byte) *Block {
	b := &Block{
		Index:        index,
		Timestamp:    timestamp,
		Data:         data,
		PreviousHash: previousHash,
	}

	b.Hash = b.generateHash()
	return b
}
