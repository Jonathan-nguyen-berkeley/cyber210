package utils

import (
	"fmt"
	"hash"
	"math/big"
	"time"
)

type Block struct {
	transactions    string
	header          string // prev hash
	work            string
	numTransactions int
	next            *Block
}

func NewBlock(header string) *Block {
	return &Block{"", header, "", 0, nil}
}

func (b *Block) AddTransaction(t *Transaction) bool {
	if !t.verified() {
		return false
	}
	b.transactions += t.String()
	b.numTransactions += 1
	return true
}

func (b *Block) GetTransactionCount() int {
	return b.numTransactions
}

func (b *Block) ComputeWork(hash hash.Hash) time.Duration {
	// first 5 of hash are 0
	if hash == nil {
		return time.Duration(0)
	}
	start := time.Now()
	var nonce uint64
	// hash of header+transactions should start with 00000
	for {

		hash.Reset()
		info := fmt.Sprintf("%s%s%d", b.header, b.transactions, nonce)

		hash.Write([]byte(info))
		hashSum := hash.Sum(nil)
		hashInt := new(big.Int).SetBytes(hashSum)

		if HasLeadingZeros(hashInt) {
			b.work = fmt.Sprintf("%x", nonce) // Save the work hash
			fmt.Printf("Found nonce %x with hash %s\n", nonce, hashInt.String())
			break
		}
		// Increment nonce and try again
		nonce++
	}
	return time.Since(start)
}

func (b *Block) verified(hash hash.Hash, prevHash string) bool {
	if hash == nil {
		fmt.Println("No hash")
		return true
	}
	if prevHash != b.header {
		fmt.Printf("Bad header >> Prev: %s != %s\n", prevHash, b.header)
		return false
	}
	hashInt := new(big.Int).SetBytes([]byte(b.GetHash(hash)))
	return HasLeadingZeros(hashInt)
}

func HasLeadingZeros(hashInt *big.Int) bool {
	// Create a big integer that represents the target value with the required number of leading zeros
	target := new(big.Int).Lsh(big.NewInt(1), uint(8*5))
	return hashInt.Cmp(target) < 0
}

func (b *Block) GetHash(hash hash.Hash) string {
	if hash == nil {
		return ""
	}
	hash.Reset()

	hash.Write([]byte(b.header + b.transactions + b.work))
	return string(hash.Sum(nil))
}

func (b Block) String() string {
	return fmt.Sprintf("Header: %s\nTransactions:\n%s\n%s", b.header, b.transactions, b.work)
}
