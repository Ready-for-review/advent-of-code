package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
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

func Test_ToNumberList(t *testing.T) {
	t.Run("Test that list of numbers is transformed as expected", func(t *testing.T) {
		want := []int{1, 2, 3}
		got, err := ToNumberList(strings.NewReader("1\n2\n3"))
		assert.Nil(t, err)
		assert.EqualValues(t, want, got)

	})
}

func Test_CalculateFuelFromInputFile(t *testing.T) {
	input := "testinput.txt"
	want := 34241
	got, err := CalculateFuelFromInputFile(input)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
