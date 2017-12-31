package domain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"testing"
	"time"

	"github.com/danmrichards/snakecoin-go/models"
)

func TestCreateGenesisBlock(t *testing.T) {
	expectedBlock := &models.Block{
		Index:        0,
		Timestamp:    time.Now(),
		Data:         []byte("Genesis Block"),
		PreviousHash: []byte("0"),
	}

	h := sha256.New()
	h.Write([]byte(expectedBlock.String()))
	expectedBlock.Hash = h.Sum(nil)

	genesisBlock := CreateGenesisBlock()

	if genesisBlock.String() != expectedBlock.String() {
		t.Errorf("TestCreateGenesisBlock: expected %s, got %s", expectedBlock, genesisBlock)
	}

	if bytes.Compare(expectedBlock.Hash, genesisBlock.Hash) != 0 {
		t.Errorf(
			"TestCreateGenesisBlock: expected %s, got %s",
			hex.EncodeToString(expectedBlock.Hash),
			hex.EncodeToString(genesisBlock.Hash),
		)
	}
}

func TestNextBlock(t *testing.T) {
	previousBlock := CreateGenesisBlock()

	for i := 1; i <= 5; i++ {
		nextBlock := NextBlock(previousBlock)

		if bytes.Compare(nextBlock.PreviousHash, previousBlock.Hash) != 0 {
			t.Errorf(
				"TestNextBlock: block chain failed at index %d, expected %s got %s",
				i,
				hex.EncodeToString(previousBlock.PreviousHash),
				hex.EncodeToString(nextBlock.PreviousHash),
			)
		}

		previousBlock = nextBlock
	}
}
