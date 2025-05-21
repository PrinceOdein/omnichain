package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func (b *Block) HashBlock() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	_ = encoder.Encode(b)
	hash := sha256.Sum256(result.Bytes())
	return hash[:]
}

func NewBlock(transactions []*Transaction, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, prevHash, []byte{}, 0}
	block.Hash = block.HashBlock()
	return block
}

func Genesis() *Block {
	return NewBlock([]*Transaction{NewGenesisTx()}, []byte{})
}
