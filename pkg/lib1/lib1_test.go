package lib1

import (
	"testing"
)

func TestLib1_hello (t *testing.T) {
	t.Run("LeetCodeTwoSumTest", func(t *testing.T) {
		result := LeetCodeTwoSumTest()
		t.Log("LeetCodeTwoSumTest done ", result)
	})

	t.Run("LeetCodeAddTwoSumTest", func(t *testing.T) {
		l := LeetCodeAddTwoNumbersTest()
		l.Dump()
		t.Log("LeetCodeAddTwoSumTest done")
	})
}
