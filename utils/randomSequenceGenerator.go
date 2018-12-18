package utils

import (
	"math/rand"
)

//RandomSequenceGenerator provides methods to generate random sequence
type RandomSequenceGenerator struct {
	numRunes int
	runes    []rune
}

//NewRandomSequenceGenerator creates generator instance
func NewRandomSequenceGenerator() *RandomSequenceGenerator {
	runes := []rune(`abcdefghijklmnopqrstuvwxyz1234567890@#$^&*()_-=+`)
	return &RandomSequenceGenerator{
		runes:    runes,
		numRunes: len(runes),
	}
}

//Next returns new sequence
func (g *RandomSequenceGenerator) Next(seqLen int) string {
	key := make([]rune, seqLen)

	for i := 0; i < seqLen; i++ {
		key[i] = g.runes[rand.Intn(g.numRunes)]
	}

	return string(key)
}
