package core

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index        int
	Timestamp    int64
	Transactions []Transaction
	PrevHash     string
	Hash         string
}

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + strconv.FormatInt(block.Timestamp, 10) + block.PrevHash
	for _, tx := range block.Transactions {
		record += tx.Sender + tx.Receiver + strconv.FormatFloat(tx.Amount, 'f', 6, 64) + strconv.FormatInt(tx.Timestamp, 10)
	}
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func NewBlock(transactions []Transaction, prevHash string, index int) Block {
	block := Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     prevHash,
	}
	block.Hash = calculateHash(block)
	return block
}
