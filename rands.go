package gots

import (
	cryptoRand "crypto/rand"

	"math"
	"math/big"
	"math/rand"
)

const (
	LowerSet  = "abcdefghijklmnopqrstuvwxyz"
	UpperSet  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NumberSet = "0123456789"
	MarkSet   = "!@#$%^&*_-=+;:,.?"

	defaultRandCharSet = LowerSet + UpperSet + NumberSet
)

type rands struct {
	chars  []rune
	random *rand.Rand
}

var globalRands *rands

func init() {
	globalRands = &rands{
		chars: Unique([]rune(defaultRandCharSet)),
	}

	seed, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic("failed to init crypto/rand: " + err.Error())
	}

	// #nosec: using crypto/rand for math/rand.seed
	globalRands.random = rand.New(rand.NewSource(seed.Int64()))
}

func RandomString(length int) string {
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = globalRands.chars[globalRands.random.Int()%len(globalRands.chars)]
	}

	return string(runes)
}
