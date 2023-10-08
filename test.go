package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

// Block represents a single block in the blockchain.
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

// Blockchain is a slice of blocks.
type Blockchain struct {
	Blocks []*Block
}

// CalculateHash calculates the hash of a string.
func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", hash)
}

// NewBlock creates a new block and adds it to the blockchain.
func (bc *Blockchain) NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.Hash = bc.CreateHash(block)
	bc.Blocks = append(bc.Blocks, block)
	return block
}

// CreateHash calculates the hash of a block.
func (bc *Blockchain) CreateHash(b *Block) string {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// DisplayBlocks prints all blocks in the blockchain.
func (bc *Blockchain) DisplayBlocks() {
	for i, block := range bc.Blocks {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("Transaction: %s\n", block.Transaction)

		fmt.Printf("Prev Hash: %s\n", block.PreviousHash)
		fmt.Printf("Curr Hash: %s\n\n", block.Hash)

		fmt.Printf("Nonce: %d\n", block.Nonce)

	}
}

// ChangeBlock changes the transaction of a specific block.
func (bc *Blockchain) ChangeBlock(blockIndex int, newTransaction string) {
	if blockIndex >= 0 && blockIndex < len(bc.Blocks) {
		bc.Blocks[blockIndex].Transaction = newTransaction
		bc.Blocks[blockIndex].Hash = bc.CreateHash(bc.Blocks[blockIndex])
	}
}

// VerifyChain verifies the integrity of the blockchain.
func (bc *Blockchain) VerifyChain() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		if bc.Blocks[i].PreviousHash != bc.Blocks[i-1].Hash {
			return false
		}
	}
	return true
}
