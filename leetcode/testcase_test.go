package leetcode_test

import (
	"leetcode/leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestPrefix(t *testing.T) {
	require.Equal(t, true, leetcode.IsValid("((([])))"))
	require.Equal(t, false, leetcode.IsValid("((([]))"))
	require.Equal(t, true, leetcode.IsValid("((([{}])))"))
	require.Equal(t, true, leetcode.IsValid("[]{}()"))
	require.Equal(t, false, leetcode.IsValid("){"))
}
