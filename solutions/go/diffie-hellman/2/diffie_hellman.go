package diffiehellman

import (
    "crypto/rand"
    "math/big"
)

var two = big.NewInt(2)

// Generates a random int64 from [2,p)
// Panics when p <= 2
func PrivateKey(p *big.Int) *big.Int {
    b, err := rand.Int(rand.Reader, new(big.Int).Sub(p, two))
    if err != nil {
        panic("random generator has failed")
    }
	return b.Add(b, two)
}

// Returns A = g**private mod p
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// Returns (private, public) pair
func NewPair(p *big.Int, g int64) (priv *big.Int, pub *big.Int) {
	priv = PrivateKey(p)
    pub = PublicKey(priv, p, g)
    return
}

// Returns s = public2**private1 mod p
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}