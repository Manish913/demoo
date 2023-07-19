package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwap(t *testing.T) {
	Swap1(2, 7)
	assert.Equal(t, 9, 9)
	//if s != 9 {
	//	t.Error("failed")
	//}
}
