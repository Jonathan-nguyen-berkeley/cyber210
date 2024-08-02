package utils

import "math/big"

// Convert string to int
func encode(message string) *big.Int {
	messageBytes := []byte(message)
	// Convert the byte slice to a big integer.
	messageBigInt := new(big.Int).SetBytes(messageBytes)
	return messageBigInt
}

// Convert integer to string
func decode(message *big.Int) string {
	// Convert the big integer to a byte slice.
	messageBytes := message.Bytes()
	// Return the string representation of the byte slice.
	return string(messageBytes)
}

func AddTransactionHelper(block *Block, sender User, receiver User, amount int) {
	transaction := NewTransaction(sender, receiver, amount, block.GetTransactionCount()+1)
	transaction.Sign(PrivateKeys[sender.Name].key)
	block.AddTransaction(transaction)
}
