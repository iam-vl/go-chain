package main

import (
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchain := NewBlockchain()
	blockchain.Print()

	blockchain.AddTransaction("A", "B", 1.0)
	prevHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(5, prevHash)
	blockchain.Print()

	blockchain.AddTransaction("C", "D", 2.0)
	blockchain.AddTransaction("X", "Y", 3.0)
	prevHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(2, prevHash)
	blockchain.Print()
}
