package lib1

import (
	"testing"
)

func TestLib1_hello (t *testing.T) {
//	t.Run("subtest1", func(t *testing.T) {
//		Lib1_hello()
//		t.Log("subtest1 done")
//	})
//	t.Run("subtest2", func(t *testing.T) {
//		c := Lib1_Add(15, 73)
//		t.Log("subtest2 done ", c)
//	})
//	t.Run("subtest3", func(t *testing.T) {
//		nums := []int{1, 15, 22, 3, 8}
//		target := 9
//		result := Lib1_twoSum(nums, target)
//		t.Log("subtest3 done ", result)
//	})
//	nums := []int{1, 15, 22, 3, 8}
//	target := 9
//	result := Lib1_twoSum(nums, target)
//	t.Logf("formualr: %d + %d = %d", nums[result[0]], nums[result[1]], target)
//	t.Log("index: ", result)

	var l1 = ListNode{2, nil}
	l1.Append(4)
	l1.Append(3)
	l1.Append(9)
	l1.Append(1)
	l1.Dump()
	var l2 = ListNode{5, nil}
	l2.Append(7)
	l2.Append(8)
	l2.Dump()
	Lib1_addTwoNumbers(&l1, &l2)
}
