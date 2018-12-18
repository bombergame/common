package utils

import (
	"testing"

	"github.com/bombergame/common/consts"
)

func TestRandomSequenceGenerator(t *testing.T) {
	g := NewRandomSequenceGenerator()

	const seqLen = 10
	if seq := g.Next(seqLen); seq == consts.EmptyString || len(seq) != seqLen {
		t.Error("wrong sequence")
	}
}
