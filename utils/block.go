package utils

import (
	"fmt"
	"hash"
	"strings"
	"time"

	"go.uber.org/zap"
)

type Block struct {
	transactions    string
	header          string // prev hash
	work            uint64
	numTransactions int
	next            *Block
	log             *zap.SugaredLogger
}

func NewBlock(header string, log *zap.SugaredLogger) *Block {
	return &Block{"", header, 0, 0, nil, log}
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
	b.log.Info("Computing Work ...")
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
		encoded := encode(string(hashSum)).Text(2)
		padded := padZeros(encoded, hash.Size()*8)
		if padded[:15] == strings.Repeat("0", 15) {
			b.work = nonce // Save the work hash
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
		fmt.Printf("Bad header >> Prev != header: %s != %s\n", prevHash, b.header)
		return false
	}
	encoded := encode(b.GetHash(hash)).Text(2)
	padded := padZeros(encoded, hash.Size()*8)
	return padded[0:15] == strings.Repeat("0", 15)
}

func padZeros(bin string, size int) string {
	return strings.Repeat("0", size-len(bin)) + bin
}

func (b *Block) GetHash(hash hash.Hash) string {
	if hash == nil {
		return ""
	}
	hash.Reset()
	hash.Write([]byte(fmt.Sprintf("%s%s%d", b.header, b.transactions, b.work)))

	return string(hash.Sum(nil))
}

func (b *Block) GetHeader() string { return b.header }
func (b Block) String() string {
	return fmt.Sprintf("Header: %x\nTransactions:\n%s\nWork: %d\n", b.header, b.transactions, b.work)
}
