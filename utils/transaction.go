package utils

import (
	"crypto/sha256"
	"fmt"
	"math/big"

	"github.com/fatih/color"
)

type Transaction struct {
	info      string
	signature *big.Int // should be m^e mod p
	id        int
	publicKey *big.Int
	modulus   *big.Int
}

func NewTransaction(sender string, rec string, amount int, id int) *Transaction {
	message := fmt.Sprintf("%s => %s (%d)", sender, rec, amount)
	return &Transaction{
		info:      message,
		id:        id,
		publicKey: PublicKeys[sender].publicKey,
		modulus:   PublicKeys[sender].modulus,
	}
}
func (t *Transaction) verified() bool {

	decrypted := new(big.Int)
	decrypted.Exp(t.signature, t.publicKey, t.modulus)

	decoded := decode(decrypted)
	expected := hashMessage(fmt.Sprintf("%s%d", t.info, t.id))

	return decoded == expected
}

func (t *Transaction) String() string {
	res := color.New(color.FgGreen).Sprint("============\n")
	res += fmt.Sprintf("Id: %d\nInfo: %s\nSignature: %d\n", t.id, t.info, t.signature)
	res += color.New(color.FgGreen).Sprint("============")
	return res
}

func encode(message string) *big.Int {
	messageBytes := []byte(message)
	// Convert the byte slice to a big integer.
	messageBigInt := new(big.Int).SetBytes(messageBytes)
	return messageBigInt
}

func decode(message *big.Int) string {
	// Convert the big integer to a byte slice.
	messageBytes := message.Bytes()
	// Return the string representation of the byte slice.
	return string(messageBytes)
}

func hashMessage(message string) string {
	hash := sha256.New()
	hash.Write([]byte(message))
	return string(hash.Sum(nil))
}

func (t *Transaction) Sign(privateKey *big.Int) {
	hashedMessage := hashMessage(fmt.Sprintf("%s%d", t.info, t.id))
	encodedMessage := encode(hashedMessage)
	signature := new(big.Int)
	signature.Exp(encodedMessage, privateKey, t.modulus)
	t.signature = signature
}
