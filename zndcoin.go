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

// ProofHash is an alias type for hash of proof strings
type ProofHash string

// PrevHash is an alias type for previous hash strings
type PrevHash string

// Block is a struct for blocks of blockchains
type Block struct {
	Index        int           `json:"index"`
	Transactions []Transaction `json:"transactions,omitempty"`
	Timestamp    Timestamp     `json:"timestamp"`
	Proof        Proof         `json:"-"`
	ProofBase64  string        `json:"proof"`
	PrevHash     PrevHash      `json:"prev_hash"`
}

// Blockchain is a struct for blockchains
type Blockchain struct {
	Chain               []Block       `json:"blocks"`
	CurrentTransactions []Transaction `json:"-"`
}

// GetLastBlock is a function to get the last block.
func (bc *Blockchain) GetLastBlock() *Block {
	return &bc.Chain[len(bc.Chain)-1]
}

// GetSecondLastBlock is a function to get the second last block.
func (bc *Blockchain) GetSecondLastBlock() *Block {
	return &bc.Chain[len(bc.Chain)-2]
}

// NextPattern is a function to get the next block's pattern.
// uncompleted
func (bc *Blockchain) NextPattern() string {
	var hashes [8]ProofHash
	for i := -7; i <= 0; i++ {
		hashes[i+7] = Hash(bc.Chain[len(bc.Chain)+i].Proof, bc.Chain[len(bc.Chain)+i-1].Proof)
	}
	ValidProofHash(hashes)
	return ""
}

// NewTransaction is a function to add a new transaction for the next block.
func (bc *Blockchain) NewTransaction(sender Address, recipient Address, amount int) int {
	bc.CurrentTransactions = append(bc.CurrentTransactions,
		Transaction{
			Sender:    sender,
			Recipient: recipient,
			Amount:    amount,
		})
	return bc.GetLastBlock().Index + 1
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

// ToJSON is a function to generate JSON from a block
func (bc *Block) ToJSON() []byte {
	bc.ProofBase64 = base64.StdEncoding.EncodeToString(bc.Proof)
	b, _ := json.Marshal(bc)
	return b
}

// NewBlockchain is a function to make a blockchain.
func NewBlockchain() *Blockchain {
	blockchain := Blockchain{}
	return &blockchain
}

// Hash is a function to get the hash of proof
func Hash(proof Proof, lastProof Proof) ProofHash {
	slice := sha256.Sum256(append([]byte(proof), ([]byte(lastProof))...))
	return ProofHash(string(base64.StdEncoding.EncodeToString(slice[:])))
}

// CreateZndkValidFunc is a function to create function to check a string start with the pattern.
func CreateZndkValidFunc(pattern string) func([]byte) bool {
	return func(str []byte) bool {
		for i, value := range []byte(pattern) {
			if value != str[i] && value+'A'-'a' != str[i] {
				return false
			}
		}
		return true
	}
}

// ValidProofHash is a function to check the proof is valid
func ValidProofHash(proofHashes [8]ProofHash) (bool, int) {
	zndkPattern := [5]func([]byte) bool{
		CreateZndkValidFunc("zun"),
		CreateZndkValidFunc("doko"),
		CreateZndkValidFunc("ki"),
		CreateZndkValidFunc("yo"),
		CreateZndkValidFunc("shi"),
	}

	zndk := [8]func([]byte) bool{
		zndkPattern[0],
		zndkPattern[0],
		zndkPattern[0],
		zndkPattern[0],
		zndkPattern[1],
		zndkPattern[2],
		zndkPattern[3],
		zndkPattern[4],
	}

	i := 0
	for i < 8 {
		if zndk[i]([]byte(proofHashes[7])) {
			break
		} else if i == 0 { // if index is 0 but not matched, go to 4
			i = 4
			continue
		}
		i++
	}
	if i == 8 {
		return false, -1
	}

	// Search Shi and Zun
	j := -1

	if zndk[0]([]byte(proofHashes[0])) && zndk[7]([]byte(proofHashes[7])) {
		j = 0
	} else {
		for i = 0; i < 7; i++ {
			if zndk[7]([]byte(proofHashes[i])) &&
				zndk[0]([]byte(proofHashes[i+1])) {
				j = i + 1
			}
		}
	}

	if j < 0 {
		return false, -1
	}

	for k := 0; k < 8; k++ {
		if !zndk[k]([]byte(proofHashes[(j+k)%8])) {
			return false, -1
		}
	}

	return true, i
}

// Mine is a function to mine zndcoin.
func Mine(blockchain Blockchain, pattern string) {
	f := CreateZndkValidFunc(pattern)
	var i big.Int
	bitlen := 0
	hash := []byte{0}
	one := big.NewInt(1)
	for !f([]byte(base64.StdEncoding.EncodeToString(hash))) {
		hash = []byte(Hash(Proof(i.Bytes()), blockchain.GetLastBlock().Proof))
		if bitlen != i.BitLen() {
			bitlen = i.BitLen()
			fmt.Printf("bitlen: %d\n", bitlen)
		}
		i.Add(&i, one)
	}
	fmt.Printf("%d %s", i.Int64(), base64.StdEncoding.EncodeToString(hash))
}

func main() {
	blockchain := Blockchain{}
	blockchain.NewBlock(Proof([]byte{}), "", Timestamp(time.Now().UnixNano()))
	Mine(blockchain, "doko")
}
