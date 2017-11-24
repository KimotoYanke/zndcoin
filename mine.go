package main

import (
	"fmt"
	"math/big"
	"time"
)

// Mine is a function to mine zndcoin.
func Mine(blockchain *Blockchain, start *int64, ch1 chan ProofHash, ch2 chan Proof) {
	pattern := blockchain.NextBlockPattern()

	f := CreateZndkValidFunc(pattern)
	bitlen := 0
	hash := ProofHash([]byte{0})
	proof := Proof([]byte{0})
	i := big.NewInt(*start)
	one := big.NewInt(1)

	fmt.Printf("pattern: %s\n", pattern)
	for !f(hash.Encode()) {
		proof = Proof(i.Bytes())
		hash = ProofHash(Hash(proof, blockchain.LastBlock().Proof))
		if bitlen != i.BitLen() {
			bitlen = i.BitLen()
		}
		i.Add(i, one)
	}
	blockchain.NewBlock(proof, "", Timestamp(time.Now().UnixNano()))
	// fmt.Printf("%d %s\n", i.Int64(), hash.Encode())
	ch1 <- hash
	ch2 <- proof // + ":" + string(blockchain.SecondLastBlock().Proof)
}
