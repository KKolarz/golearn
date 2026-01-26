package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Sum of a fixed value array", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("Got %v, want: %v", got, want)
		}

	})

}

func TestSumAll(t *testing.T) {
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5, 6}

	got := SumAll(numbers1, numbers2)
	want := []int{6, 15}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	t.Run("Testing a sum of tails", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{2, 4, 7}

		got := SumAllTails(numbers1, numbers2)
		want := []int{5, 11}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Testing a pass of an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 5})
		want := []int{7}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
