package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Day2(t *testing.T) {
	tests := []struct {
		list   []int
		report Report
	}{
		{[]int{7, 6, 4, 2, 1}, Safe},
		{[]int{1, 2, 7, 8, 9}, Unsafe},
		{[]int{9, 7, 6, 2, 1}, Unsafe},
		{[]int{1, 3, 2, 4, 5}, Unsafe},
		{[]int{8, 6, 4, 4, 1}, Unsafe},
		{[]int{1, 3, 6, 7, 9}, Safe},
		{[]int{73, 72, 70, 68, 66}, Safe},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			report := calculateReport(test.list)
			require.Equal(t, report, test.report)
		})
	}
}
