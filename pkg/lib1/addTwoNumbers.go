package lib1

func LeetCodeAddTwoNumbersTest() *ListNode {
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
	return LeetCodeAddTwoNumbers(&l1, &l2)
}
func LeetCodeAddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 = new(ListNodeD)
	p1 := l1
	p2 := l2
	var tail_l1 *ListNodeD
	var tail_l2 *ListNodeD
	p3 := l3
	for {
		p3 = p3.MyAppend(p3, p1, p2)
		if p1 != nil {
			tail_l1 = p3
			p1 = p1.Next
		}
		if p2 != nil {
			tail_l2 = p3
			p2 = p2.Next
		}
		if p1 == nil && p2 == nil {
			p3.MyAppend(p3, p1, p2)
			break
		}
	}
	r1 := tail_l1
	r2 := tail_l2
	var z = new(ListNode)
	z.Val = -1
	second := 0
	first := 0
	for {
		if r1 == nil && r2 == nil {
			break
		}
		sum := second
		if r1 != nil {
			sum += r1.p1.Val
			r1 = r1.prev
		}
		if r2 != nil {
			sum += r2.p2.Val
			r2 = r2.prev
		}
		first = (sum%10)
		second = (sum/10)
		z = z.Append2(first)
	}
	z.Dump()
	return nil
}

func LeetCodeTwoSumTest() []int {
		nums := []int{1, 15, 22, 3, 8}
		target := 9
		return LeetCodeTwoSum(nums, target)
}

func LeetCodeTwoSum (nums []int, target int) []int {
	for i, v0 := range nums[0:] {
		for j, v1 := range nums[i+1:] {
			if v0 + v1 == target {
				return []int{i, i+j+1}
			}
		}
	}
	return []int{0, 0}
}