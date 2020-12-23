package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CalculateFuel(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("CalculateFuel(%d)", test.input), func(t *testing.T) {
			got := CalculateFuel(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_SumOfFuels(t *testing.T) {
	input := []int{12, 14, 1969, 100756}
	want := 34241
	got := SumOfFuels(input)
	assert.Equal(t, want, got)
}
