package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"time"
)

// Address is an alias type for address
type Address string

// Transaction is a struct for transactions
type Transaction struct {
	Sender    Address `json:"sender"`
	Recipient Address `json:"recipient"`
	Amount    int     `json:"amount"`
}

// Timestamp is an alias type for timestamps
type Timestamp int64

// Proof is an alias type for proof strings
type Proof []byte

// ProofBase64 is an alias type for base64 proof strings
type ProofBase64 []byte

// ProofHash is an alias type for hash of proof strings
type ProofHash []byte

// PrevHash is an alias type for previous hash strings
type PrevHash string

// Block is a struct for blocks of blockchains
type Block struct {
	Index        int           `json:"index"`
	Transactions []Transaction `json:"transactions,omitempty"`
	Timestamp    Timestamp     `json:"timestamp"`
	Proof        Proof         `json:"-"`
	ProofBase64  ProofBase64   `json:"proof"`
	PrevHash     PrevHash      `json:"prev_hash"`
}

// Pattern returns the pattern of block (zun, doko, ki...)
func (bc *Blockchain) Pattern(index int) string {
	for i := 0; i < 5; i++ {
		if zndkPattern[i](Hash(bc.Chain[index].Proof, bc.Chain[index-1].Proof).Encode()) {
			switch i {
			case 0:
				return "zun"
			case 1:
				return "doko"
			case 2:
				return "ki"
			case 3:
				return "yo"
			case 4:
				return "shi"
			}
		}
	}
	return ""
}

// Blockchain is a struct for blockchains
type Blockchain struct {
	Chain               []Block       `json:"blocks"`
	CurrentTransactions []Transaction `json:"-"`
}

// LastBlock is a function to get the last block.
func (bc *Blockchain) LastBlock() *Block {
	if len(bc.Chain) == 0 {
		block := Block{
			Index:        0,
			Timestamp:    Timestamp(time.Now().UnixNano()),
			Transactions: nil,
			Proof:        Proof([]byte{0}),
			PrevHash:     "",
		}
		return &block
	}
	return &bc.Chain[len(bc.Chain)-1]
}

// SecondLastBlock is a function to get the second last block.
func (bc *Blockchain) SecondLastBlock() *Block {
	return &bc.Chain[len(bc.Chain)-2]
}

// NewTransaction is a function to add a new transaction for the next block.
func (bc *Blockchain) NewTransaction(sender Address, recipient Address, amount int) int {
	bc.CurrentTransactions = append(bc.CurrentTransactions,
		Transaction{
			Sender:    sender,
			Recipient: recipient,
			Amount:    amount,
		})
	return bc.LastBlock().Index + 1
}

// NewBlock is a function to add a new block for the blockchain.
func (bc *Blockchain) NewBlock(proof Proof, prevHash PrevHash, timestamp Timestamp) *Block {
	block := Block{
		Index:        len(bc.Chain) + 1,
		Timestamp:    timestamp,
		Transactions: bc.CurrentTransactions,
		Proof:        proof,
		PrevHash:     prevHash,
	}
	bc.CurrentTransactions = []Transaction{}
	bc.Chain = append(bc.Chain, block)
	return &block
}

// NextBlockPattern is a function to get the next block pattern
func (bc *Blockchain) NextBlockPattern() string {
	if len(bc.Chain) < 8 {
		if len(bc.Chain) <= 3 {
			return "zun"
		}
		switch len(bc.Chain) {
		case 4:
			return "doko"
		case 5:
			return "ki"
		case 6:
			return "yo"
		case 7:
			return "shi"
		}
	}
	switch bc.Pattern(len(bc.Chain) - 1) {
	case "doko":
		return "ki"
	case "ki":
		return "yo"
	case "yo":
		return "shi"
	case "shi":
		return "zun"
	case "zun":
		if bc.Pattern(len(bc.Chain)-4) == "zun" {
			return "doko"
		}
		return "zun"
	}
	return "zun"
}

// ToJSON is a function to generate JSON from a block
func (block *Block) ToJSON() []byte {
	block.ProofBase64 = block.Proof.Encode()
	b, _ := json.Marshal(block)
	return b
}

// ParseJSON is a function to generate JSON from a block
func (block *Block) ParseJSON(jsonData []byte) {
	json.Unmarshal(jsonData, block)
	block.Proof = block.ProofBase64.Decode()
}

// NewBlockchain is a function to make a blockchain.
func NewBlockchain() *Blockchain {
	blockchain := Blockchain{}
	return &blockchain
}

// Hash is a function to get the hash of proof
func Hash(proof Proof, lastProof Proof) ProofHash {
	slice := sha256.Sum256(append([]byte(proof), ([]byte(lastProof))...))
	return ProofHash(slice[:])
}

// CreateZndkValidFunc is a function to create function to check a string start with the pattern.
func CreateZndkValidFunc(pattern string) func(string) bool {
	return func(str string) bool {
		if str == "" {
			return false
		}
		bs := []byte(str)
		for i, value := range []byte(pattern) {
			if value != bs[i] && value+'A'-'a' != bs[i] {
				return false
			}
		}
		return true
	}
}

// Encode is a function to encode byte slices
func Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Decode is a function to decode to byte slices
func Decode(str string) []byte {
	d, _ := base64.StdEncoding.DecodeString(str)
	return d
}

// Encode is a function to encode the hash of a proof
func (hash ProofHash) Encode() string {
	return Encode([]byte(string(hash)))
}

// Encode is a function to encode a proof
func (proof Proof) Encode() ProofBase64 {
	return ProofBase64(Encode([]byte(string(proof))))
}

// Decode is a function to decode the base64 proof
func (proofBase64 ProofBase64) Decode() Proof {
	return Decode(string(proofBase64))
}

func main() {
	blockchain := Blockchain{}

	ch1 := make(chan string)
	ch2 := make(chan string)
	start := int64(0)
	for {
		go Mine(&blockchain, &start, ch1, ch2)
		fmt.Printf("hash:\t%s\nproof:\t%s\n len:%d\n", <-ch1, <-ch2, len(blockchain.Chain))
	}
}
