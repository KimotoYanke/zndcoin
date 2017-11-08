package main

import "encoding/json"
import "fmt"
import "time"

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
type Proof string

// PrevHash is an alias type for previous hash strings
type PrevHash string

// Block is a struct for blocks of blockchains
type Block struct {
	Index        int           `json:"index"`
	Transactions []Transaction `json:"transactions,omitempty"`
	Timestamp    Timestamp     `json:"timestamp"`
	Proof        Proof         `json:"proof"`
	PrevHash     PrevHash      `json:"prev_hash"`
}

// Blockchain is a struct for blockchains
type Blockchain struct {
	Chain               []Block       `json:"blocks"`
	CurrentTransactions []Transaction `json:"-"`
}

// GetLastBlock is a function to get the last block.
func (bc *Blockchain) GetLastBlock() Block {
	return bc.Chain[len(bc.Chain)-1]
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
func (bc *Blockchain) NewBlock(proof Proof, prevHash PrevHash, timestamp Timestamp) Block {
	block := Block{
		Index:        len(bc.Chain) + 1,
		Timestamp:    timestamp,
		Transactions: bc.CurrentTransactions,
		Proof:        proof,
		PrevHash:     prevHash,
	}
	bc.Chain = append(bc.Chain, block)
	return block
}

// NewBlockchain is a function to make a blockchain.
func NewBlockchain() Blockchain {
	blockchain := Blockchain{}
	return blockchain
}

func main() {
	blockchain := Blockchain{}
	blockchain.NewBlock("hogehoge", "hogehoge", Timestamp(time.Now().UnixNano()))
	b, _ := json.Marshal(blockchain)
	fmt.Printf("%s\n", string(b))
}
