package domain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"testing"
	"time"

	"github.com/danmrichards/snakecoin-go/blockchain/models"
)

func TestCreateGenesisBlock(t *testing.T) {
	expectedBlock := &models.Block{
		Index:     0,
		Timestamp: time.Now(),
		Data: &models.Data{
			Proof:        9,
			Transactions: nil,
		},
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
