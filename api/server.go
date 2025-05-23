package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/PrinceOdein/omnichain/core"
)

var blockchain = core.NewBlockchain()

func StartServer() {
	http.HandleFunc("/chain", handleChain)
	http.HandleFunc("/mine", handleMine)
	http.HandleFunc("/transactions/new", handleNewTransaction)

	log.Println("API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleChain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blockchain.Blocks)
}

func handleMine(w http.ResponseWriter, r *http.Request) {
	block := blockchain.MineBlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(block)
}

func handleNewTransaction(w http.ResponseWriter, r *http.Request) {
	var tx core.Transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "Invalid transaction", http.StatusBadRequest)
		return
	}
	tx.Timestamp = time.Now().Unix()
	blockchain.AddTransaction(tx)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tx)
}
