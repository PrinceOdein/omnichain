package core

import (
	"crypto/ecdsa"
	// "crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	// "encoding/hex"
	"math/big"
)

type Transaction struct {
	From  string
	To    string
	Value int
	SigR  *big.Int
	SigS  *big.Int
}

func NewTransaction(from, to string, value int, priv *ecdsa.PrivateKey) *Transaction {
	tx := &Transaction{From: from, To: to, Value: value}
	h := tx.Hash()
	r, s, _ := ecdsa.Sign(rand.Reader, priv, h[:])
	tx.SigR = r
	tx.SigS = s
	return tx
}

func (tx *Transaction) Hash() [32]byte {
	data := []byte(tx.From + tx.To + string(tx.Value))
	return sha256.Sum256(data)
}

func (tx *Transaction) Verify(pub *ecdsa.PublicKey) bool {
	h := tx.Hash()
	return ecdsa.Verify(pub, h[:], tx.SigR, tx.SigS)
}

func NewGenesisTx() *Transaction {
	return &Transaction{From: "", To: "genesis", Value: 0}
}
