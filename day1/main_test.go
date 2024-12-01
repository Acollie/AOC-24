package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Day1(t *testing.T) {
	list1 := []int{3, 4, 2, 1, 3, 3}
	list2 := []int{4, 3, 5, 3, 9, 3}

	diff := calculateDiff(list1, list2)
	require.Equal(t, diff, 11)
}
