package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Timestamp    int64
	Data         string
	PrevHash     string
	Hash         string
	Nonce        int
	Difficulty   int
}

type Blockchain struct {
	chain  []Block
	difficulty int
}

func calculateHash(b Block) string {
	record := string(b.Timestamp) + b.Data + b.PrevHash + fmt.Sprint(b.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (bc *Blockchain) addBlock(data string) {
	prevBlock := bc.chain[len(bc.chain)-1]
	newBlock := Block{
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash: prevBlock.Hash,
		Nonce:    0,
		Difficulty: bc.difficulty,
	}
	for !isNewBlockValid(newBlock, bc.difficulty) {
		newBlock.Nonce++
		newBlock.Hash = calculateHash(newBlock)
	}
	bc.chain = append(bc.chain, newBlock)
}

func isNewBlockValid(b Block, difficulty int) bool {
	target := fmt.Sprintf("%064d", difficulty)
	hash := calculateHash(b)
	return hash[:difficulty] == target
}

func main() {
	bc := Blockchain{
		chain:      []Block{{Timestamp: 0, Hash: "0"}},
		difficulty: 4,
	}
	bc.addBlock("Transaction 1")
	bc.addBlock("Transaction 2")
	bc.addBlock("Transaction 3")
	for _, block := range bc.chain {
		fmt.Printf("Block %x: %s\n", block.Hash, block.Data)
	}
}