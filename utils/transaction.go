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

func NewTransaction(sender User, rec User, amount int, id int) *Transaction {
	message := fmt.Sprintf("%s => %s (%d)", sender, rec, amount)
	return &Transaction{
		info:      message,
		id:        id,
		publicKey: PublicKeys[sender.Name].key,
		modulus:   PublicKeys[sender.Name].modulus,
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
	return res
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
