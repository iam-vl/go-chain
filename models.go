package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Timestamp    int64          `json:"timestamp"`
	Nonce        int            `json:"nonce"`
	PreviousHash [32]byte       `json:"previous_hash"`
	Transactions []*Transaction `json:"transactions"`
}

type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
}

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewBlock(nonce int, prevHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.Timestamp = time.Now().UnixNano()
	b.Nonce = nonce
	b.PreviousHash = prevHash
	b.Transactions = transactions
	return b
}

func NewBlockchain(blockchainAddress string) *Blockchain {
	bc := new(Blockchain)
	b := &Block{}
	bc.blockchainAddress = blockchainAddress
	bc.CreateBlock(0, b.Hash())
	return bc
}

func NewTransaction(sender, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (b *Block) Print() {
	fmt.Printf("timestamp     %d\n", b.Timestamp)
	fmt.Printf("nonce         %d\n", b.Nonce)
	fmt.Printf("previous_hash %x\n", b.PreviousHash)
	// fmt.Printf("transactions  %s\n", b.Transactions)
	for _, t := range b.Transactions {
		t.Print()
	}
}

func (bc *Blockchain) Print() {
	fmt.Println("===Blockchain start===")
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 10), i, strings.Repeat("=", 10))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 10))
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 15))
	fmt.Printf("  sender_blockchain_address \t%s\n", t.senderBlockchainAddress)
	fmt.Printf("  recipient_blockchain_address \t%s\n", t.recipientBlockchainAddress)
	fmt.Printf("  value \t\t%.1f\n", t.value)
}

func (t *Transaction) MarshalJson() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}
