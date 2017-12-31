package main

import (
	"encoding/hex"
	"log"

	"github.com/danmrichards/snakecoin-go/domain"
	"github.com/danmrichards/snakecoin-go/models"
)

var blockChain []*models.Block

const noOfBlocks = 20

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	// Create the blockchain and add the genesis block.
	blockChain = []*models.Block{domain.CreateGenesisBlock()}
	previousBlock := blockChain[0]

	for i := 0; i < noOfBlocks; i++ {
		// Add a new block to the chain.
		nextBlock := domain.NextBlock(previousBlock)
		blockChain = append(blockChain, nextBlock)
		previousBlock = nextBlock

		log.Printf("block %d has been added to the blockchain!\n", nextBlock.Index)
		log.Printf("hash: %s\n'n", hex.EncodeToString(nextBlock.Hash))
	}
}
