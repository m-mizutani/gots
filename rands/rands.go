package rands

import (
	"github.com/m-mizutani/gots/slice"
)

const (
	LowerSet  = "abcdefghijklmnopqrstuvwxyz"
	UpperSet  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NumberSet = "0123456789"
	MarkSet   = "!@#$%^&*_-=+;:,.?"

	defaultRandCharSet = LowerSet + UpperSet + NumberSet
)

var defaultChars []rune

func init() {
	defaultChars = []rune(defaultRandCharSet)
}

// New returns string of random `length` characters. Default character sets to generate random string are lower and upper alphabets (LowerSet and UpperSet), numbers (NumberSet) and marks (MarkSet). If charSets is provided, New uses only the provided characters to generate string.
func New(length int, charSets ...string) string {
	chars := defaultChars
	if len(charSets) > 0 {
		chars = []rune{}
		for _, charSet := range charSets {
			chars = append(chars, []rune(charSet)...)
		}
		chars = slice.Unique(chars)
	}

	random, err := RandomUint64Array(length)
	if err != nil {
		panic("failed to generate random string: " + err.Error())
	}

	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = chars[random[i]%uint64(len(chars))]
	}

	return string(runes)
}

func Alphabet(length int) string {
	return New(length, LowerSet, UpperSet)
}

func LowerCase(length int) string {
	return New(length, LowerSet)
}

func UpperCase(length int) string {
	return New(length, UpperSet)
}

func AlphaNum(length int) string {
	return New(length, LowerSet, UpperSet, NumberSet)
}

func Number(length int) string {
	return New(length, NumberSet)
}
