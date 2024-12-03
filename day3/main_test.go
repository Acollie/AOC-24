package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_main(t *testing.T) {
	result := fetchMuls("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	require.Len(t, result, 4)
	total := mutiplyMuls(result)
	require.Equal(t, total, 161)
}
