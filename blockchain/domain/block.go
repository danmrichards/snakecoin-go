package domain

import (
	"time"

	"github.com/danmrichards/snakecoin-go/blockchain/models"
)

// Manually creates an initial "Genesis" block to start the chain. Note
// arbitrary data and previous block hash.
func CreateGenesisBlock() *models.Block {
	return models.NewBlock(0, time.Now(), &models.Data{
		Proof:        9,
		Transactions: nil,
	}, []byte("0"))
}

// Simple proof-of-work algorithm for mining.
func ProofOfWork(lastProof int64) int64 {
	inc := lastProof + 1

	// Keep incrementing until the incrementor is divisible by 9 and last proof.
	for inc%9 != 0 || inc%lastProof != 0 {
		inc++
	}

	return inc
}
