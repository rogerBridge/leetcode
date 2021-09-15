package leetcode_test

import (
	"leetcode/leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	require.Equal(t, 6, leetcode.RemoveDuplicates([]int{1, 2, 3, 3, 4, 5, 5, 6}))
}
