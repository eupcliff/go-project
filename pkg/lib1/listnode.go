package lib1

import (
	"fmt"
	"log"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func (l *ListNode) Append (val int) {
	for p := l; ; p = p.Next {
		if p.Next == nil {
			p.Next = new(ListNode)
			p.Next.Val = val
			return
		}
	}
}

func (l *ListNode) Dump () {
	i := 0
	str := ""
	for p := l; p != nil ; p = p.Next{
		str += fmt.Sprintf("%d ", p.Val)
//		log.Printf("list[%d]=%d\n", i, p.Val)
		i++
	}
	log.Printf("%s", str)
}

func (l *ListNode) Append2 (val int) *ListNode {
	if l.Val < 0 {
		l.Val = val
		return l
	} else {
		var p = new(ListNode)
		p.Val = val
		p.Next = l
		return p
	}
}

type ListNodeD struct {
    p1 *ListNode
    p2 *ListNode
    next *ListNodeD
    prev *ListNodeD
}

func (l *ListNodeD) MyAppend (prev *ListNodeD, p1 *ListNode, p2 *ListNode) *ListNodeD {
	var next *ListNodeD
	if l.p1 == nil && l.p2 == nil {
		// is head
		l.p1 = p1
		l.p2 = p2
		//l.prev = prev
		next = l
	} else {
		l.next = new(ListNodeD)
		l.next.p1 = p1
		l.next.p2 = p2
		l.next.prev = prev
		next = l.next
	}
	return next
}

func (l *ListNodeD) PrintAll () {
	for p := l; p.next != nil ; p = p.next {
		str := ""
		if p.p1 == nil {
			str += "v1: ,"
		} else {
			str += fmt.Sprintf("v1:%d ,", p.p1.Val)
		}
		if p.p2 == nil {
			str += "v2: "
		} else {
			str += fmt.Sprintf("v2:%d ", p.p2.Val)
		}
		//log.Printf("val=(%s), p1=%p, p2=%p, curr=%p, next=%p, prev=%p", str, p.p1, p.p2, p, p.next, p.prev)
		log.Printf("val=(%s)", str)
	}
}

func (l *ListNodeD) PrintAllReversed () {
	log.Println("print_all_reverse: ")
	for p := l;  ; p = p.prev {
		str := ""
		if p.p1 == nil {
			str += "v1: ,"
		} else {
			str += fmt.Sprintf("v1:%d ,", p.p1.Val)
		}
		if p.p2 == nil {
			str += "v2: "
		} else {
			str += fmt.Sprintf("v2:%d ", p.p2.Val)
		}
		log.Printf("val=(%s), p1=%p, p2=%p, curr=%p, next=%p, prev=%p", str, p.p1, p.p2, p, p.next, p.prev)
		if p.prev == nil {
			break
		}
	}
}
