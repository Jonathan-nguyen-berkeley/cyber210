package utils

import (
	"math/big"
)

type keyPair struct {
	key     *big.Int
	modulus *big.Int
}

type User struct {
	Name string
}

var PublicKeys = map[string]keyPair{}
var PrivateKeys = map[string]keyPair{}

func (u User) GeneratePrivateKey(publicKey *big.Int) *big.Int {

	p := new(big.Int).Exp(big.NewInt(2), big.NewInt(127), nil)
	p.Sub(p, big.NewInt(1))

	q := new(big.Int).Exp(big.NewInt(2), big.NewInt(521), nil)
	q.Sub(q, big.NewInt(1))

	psub := new(big.Int).Sub(p, big.NewInt(1))
	qsub := new(big.Int).Sub(q, big.NewInt(1))
	phi := new(big.Int).Mul(psub, qsub)
	n := new(big.Int).Mul(p, q)
	priv := new(big.Int)
	priv.ModInverse(publicKey, phi)
	PublicKeys[u.Name] = keyPair{publicKey, n}
	PrivateKeys[u.Name] = keyPair{priv, n}
	return priv
}
