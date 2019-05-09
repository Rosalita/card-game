package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcRatio(t *testing.T) {
	tests := []struct {
		from  int
		to    int
		scale float64
	}{
		{800, 600, 1.3333333333333333},
	}

	for _, test := range tests {
		scale := calcRatio(test.from, test.to)
		assert.Equal(t, test.scale, scale)
	}
}
