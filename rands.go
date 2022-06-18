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

// RandomString returns string of random `length` characters. Default character sets to generate random string are lower and upper alphabets (LowerSet and UpperSet), numbers (NumberSet) and marks (MarkSet). If charSets is provided, RandomString uses only the provided characters to generate string.
func RandomString(length int, charSets ...string) string {
	chars := globalRands.chars
	if len(charSets) > 0 {
		chars = []rune{}
		for _, charSet := range charSets {
			chars = append(chars, []rune(charSet)...)
		}
		chars = Unique(chars)
	}

	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = chars[globalRands.random.Int()%len(chars)]
	}

	return string(runes)
}
