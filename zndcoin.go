package main

import "fmt"

type Transaction struct {
	Sender    string
	Recipient string
	Amount    int
}

type Block struct {
	Index        int
	Transactions []Transaction
	Timestamp    time.Time
	Proof        string
	PrevHash     string
}

type Blockchain struct {
	Chain               []Block
	CurrentTransactions []Transaction
}

func (*bc) GetLastBlock() Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (*bc) NewTransaction(sender string, recipient string, amount string) int {
	bc.CurrentTransactions = append(bc.CurrentTransactions, Transaction{Sender: sender, Recipient: recipient, Amount: amount})
	return bc.GetLastBlock().Index + 1
}

func main() {
	fmt.Println("vim-go")
}
