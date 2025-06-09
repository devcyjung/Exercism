package diffiehellman

import (
    "crypto/rand"
    "log/slog"
    "math/big"
    "os"
    "sync"
)

const (
    msgLessThanTwo = "Invalid prime, must be > 2"
    msgFailedKeyGen = "Failed to generate random private key"
)

var (
    two = big.NewInt(2)
	handler = slog.NewTextHandler(os.Stderr, nil)
    once sync.Once
)

func PrivateKey(p *big.Int) *big.Int {
    once.Do(func() {
        slog.SetDefault(slog.New(handler))
    })
    if p.Cmp(two) <= 0 {
        slog.Error(msgLessThanTwo, "p", p)
    }
    b, err := rand.Int(rand.Reader, new(big.Int).Sub(p, two))
    if err != nil {
        slog.Error(msgFailedKeyGen, "p", p)
    }
	return b.Add(b, two)
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	priv := PrivateKey(p)
    pub := PublicKey(priv, p, g)
    return priv, pub
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}