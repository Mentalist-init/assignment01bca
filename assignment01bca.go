// assignment01bca.go
package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

// Block represents a block in the blockchain.
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
}

// NewBlock creates a new block.
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.CurrentHash = block.CreateHash()
	return block
}

// CreateHash calculates the hash of a block.
func (b *Block) CreateHash() string {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// Blockchain represents a blockchain.
type Blockchain struct {
	Blocks []*Block
}

// DisplayBlocks prints all the blocks in the blockchain.
func (bc *Blockchain) DisplayBlocks() {
	for _, block := range bc.Blocks {
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n\n", block.CurrentHash)
	}
}

// ChangeBlock changes the transaction of a given block.
func (bc *Blockchain) ChangeBlock(index int, newTransaction string) {
	if index >= 0 && index < len(bc.Blocks) {
		block := bc.Blocks[index]
		block.Transaction = newTransaction
		block.CurrentHash = block.CreateHash()
	}
}

// VerifyChain verifies the integrity of the blockchain.
func (bc *Blockchain) VerifyChain() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]
		if currentBlock.PreviousHash != previousBlock.CurrentHash {
			return false
		}
	}
	return true
}

// CalculateHash calculates the hash of a string.
func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", hash)
}
