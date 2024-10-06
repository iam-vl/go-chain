package main

import (
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchain := NewBlockchain()

	prevHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(5, prevHash)

	prevHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(2, prevHash)
	blockchain.Print()
}
