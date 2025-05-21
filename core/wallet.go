package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	Address    string
}

func NewWallet() *Wallet {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pubBytes, _ := x509.MarshalPKIXPublicKey(&priv.PublicKey)
	addr := hex.EncodeToString(pubBytes)
	return &Wallet{PrivateKey: priv, Address: addr}
}
