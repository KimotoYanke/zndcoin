package main

import "fmt"
import "time"

// Address is an alias type for address
type Address string

// Transaction is a struct for transactions
type Transaction struct {
	Sender    Address
	Recipient Address
	Amount    int
}

// Timestamp is an alias type for timestamps
type Timestamp int64

// Proof is an alias type for proof strings
type Proof string

// PrevHash is an alias type for previous hash strings
type PrevHash string

// Block is a struct for blocks of blockchains
type Block struct {
	Index        int
	Transactions []Transaction
	Timestamp    Timestamp
	Proof        Proof
	PrevHash     PrevHash
}

// Blockchain is a struct for blockchains
type Blockchain struct {
	Chain               []Block
	CurrentTransactions []Transaction
}

func (*bc) GetLastBlock() Block {
	return bc.Chain[len(bc.Chain)-1]
}

// NewTransaction is a function to add a new transaction for the next block.
func (*bc) NewTransaction(sender string, recipient string, amount string) int {
	bc.CurrentTransactions = append(bc.CurrentTransactions, Transaction{Sender: sender, Recipient: recipient, Amount: amount})
	return bc.GetLastBlock().Index + 1
}

func (*bc) NewBlock(proof Proof, prevHash string)

func main() {
	fmt.Println("vim-go")
}
