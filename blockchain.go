package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Timestamp    int64    `json:"timestamp"`
	Nonce        int      `json:"nonce"`
	PreviousHash [32]byte `json:"previous_hash"`
	Transactions []string `json:"transactions"`
}

func NewBlock(nonce int, prevHash [32]byte) *Block {
	b := new(Block)
	b.Timestamp = time.Now().UnixNano()
	b.Nonce = nonce
	b.PreviousHash = prevHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp     %d\n", b.Timestamp)
	fmt.Printf("nonce         %d\n", b.Nonce)
	fmt.Printf("previous_hash %x\n", b.PreviousHash)
	fmt.Printf("transactions  %s\n", b.Transactions)
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	b := &Block{}
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, prevHash [32]byte) *Block {
	b := NewBlock(nonce, prevHash)
	bc.chain = append(bc.chain, b)
	return b
}
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}
func (bc *Blockchain) Print() {
	fmt.Println("===Blockchain start===")
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 10), i, strings.Repeat("=", 10))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 10))
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
