package blockchain

import (
    "fmt"
    "time"
    "crypto/sha256"
)

// Block represents a block in the blockchain
type Block struct {
    Index     int
    Timestamp string
    Data      string
    PrevHash  string
    Hash      string
}

// Blockchain represents the blockchain
type Blockchain struct {
    Chain []Block
}

// NewBlock creates a new block in the blockchain
func (bc *Blockchain) NewBlock(data string) {
    prevBlock := bc.Chain[len(bc.Chain)-1]
    newIndex := prevBlock.Index + 1
    newTimestamp := time.Now().String()
    newHash := calculateHash(newIndex, prevBlock.Hash, newTimestamp, data)

    newBlock := Block{
        Index:     newIndex,
        Timestamp: newTimestamp,
        Data:      data,
        PrevHash:  prevBlock.Hash,
        Hash:      newHash,
    }

    bc.Chain = append(bc.Chain, newBlock)
}

// DisplayAllBlocks displays all blocks in the blockchain
func (bc *Blockchain) DisplayAllBlocks() {
    for _, block := range bc.Chain {
        fmt.Printf("Index: %d\n", block.Index)
        fmt.Printf("Timestamp: %s\n", block.Timestamp)
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("PrevHash: %s\n", block.PrevHash)
        fmt.Printf("Hash: %s\n\n", block.Hash)
    }
}

// ModifyBlock modifies a block in the blockchain
func (bc *Blockchain) ModifyBlock(index int, data string) {
    if index >= 0 && index < len(bc.Chain) {
        bc.Chain[index].Data = data
        bc.Chain[index].Timestamp = time.Now().String()
        bc.Chain[index].Hash = calculateHash(bc.Chain[index].Index, bc.Chain[index].PrevHash, bc.Chain[index].Timestamp, data)
    }
}

// calculateHash calculates the hash of a block
func calculateHash(index int, prevHash string, timestamp string, data string) string {
    hashInput := fmt.Sprintf("%d%s%s%s", index, prevHash, timestamp, data)
    hash := sha256.Sum256([]byte(hashInput))
    return fmt.Sprintf("%x", hash)
}