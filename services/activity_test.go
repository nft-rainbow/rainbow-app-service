package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcNextTokenId(t *testing.T) {
	next := calcNextTokenId(10, [][2]uint{{11, 13}, {15, 16}})
	assert.Equal(t, uint(14), next)

	next = calcNextTokenId(10, [][2]uint{{16, 20}, {16, 19}, {11, 15}})
	assert.Equal(t, uint(21), next)
}
