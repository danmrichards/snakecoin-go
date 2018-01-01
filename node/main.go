package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/danmrichards/snakecoin-go/blockchain/domain"
	"github.com/danmrichards/snakecoin-go/blockchain/models"
)

var (
	nodeTxns   []*models.Transaction
	blockChain []*models.Block
)

const minerAddress = "q3nf394hjg-random-miner-address-34nf3i4nflkn3oi"

// Adds a new transaction to the blockchain.
func transaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	// Extract the transaction data.
	var txn models.Transaction
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("could not read body: %s", err)
		log.Println(msg)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(msg))
		return
	}
	defer r.Body.Close()
	if err = json.Unmarshal(b, &txn); err != nil {
		msg := fmt.Sprintf("could not decode json: %s", err)
		log.Println(msg)

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(msg))
		return
	}

	// Add the transaction to our list.
	nodeTxns = append(nodeTxns, &txn)

	log.Println("new transaction")
	log.Printf("from: %s\n", txn.From)
	log.Printf("to: %s\n", txn.To)
	log.Printf("amount: %d\n\n", txn.Amount)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transaction submission successful"))
}

// Do some mining.
func mine(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	// Get the last proof of work.
	lastBlock := blockChain[len(blockChain)-1]
	lastProof := lastBlock.Data.Proof

	// Find the proof of work for the current block being mined.
	proof := domain.ProofOfWork(lastProof)

	// Once we find a valid proof of work,  we know we can mine a block so we
	// reward the miner by adding a transaction.
	nodeTxns := append(nodeTxns, &models.Transaction{
		From:   "network",
		To:     minerAddress,
		Amount: 1,
	})

	// Now we can gather the data needed to create the new block.
	newBlockData := &models.Data{
		Proof:        proof,
		Transactions: nodeTxns,
	}

	// Create the new mined block.
	minedBlock := models.NewBlock(
		lastBlock.Index+1,
		time.Now(),
		newBlockData,
		lastBlock.Hash,
	)

	blockChain = append(blockChain, minedBlock)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(minedBlock)
}

// Get the current blockchain.
func blocks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blockChain)
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	// Create the blockchain and add the genesis block.
	blockChain = []*models.Block{domain.CreateGenesisBlock()}

	http.HandleFunc("/txn", transaction)
	http.HandleFunc("/mine", mine)
	http.HandleFunc("/blocks", blocks)

	http.ListenAndServe(":8000", nil)
}
