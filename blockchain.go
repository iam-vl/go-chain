package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (bc *Blockchain) CreateBlock(nonce int, prevHash [32]byte) *Block {
	b := NewBlock(nonce, prevHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) AddTransaction(sender, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

// func (b *Block) MarshalJson() ([]byte, error) {
// 	return json.Marshal(struct {
// 		Timestamp    int64    `json:"timestamp"`
// 		Nonce        int      `json:"nonce"`
// 		PreviousHash string   `json:"previous_hash"`
// 		Transactions []string `json:"transactions"`
// 	}{
// 		Timestamp:    b.timestamp,
// 		Nonce:        b.nonce,
// 		PreviousHash: b.previousHash,
// 		Transactions: b.transactions,
// 	})
// }
