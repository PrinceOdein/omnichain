package api

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/PrinceOdein/omnichain/core"
	"net/http"
)

var blockchain = core.NewBlockchain()

func StartServer() {
	http.HandleFunc("/chain", handleChain)
	http.HandleFunc("/mine", handleMine)
	fmt.Println("API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleChain(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blockchain.Blocks)
}

func handleMine(w http.ResponseWriter, r *http.Request) {
	wlt := core.NewWallet()
	tx := core.NewTransaction("genesis", wlt.Address, 50, wlt.PrivateKey)
	blockchain.AddBlock([]*core.Transaction{tx})
	w.Write([]byte("Block mined!\n"))
}
