package main

import "fmt"

var a = "df"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var list1 ListNode = ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	var list2 ListNode = ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	lists := mergeTwoLists(&list1, &list2)
	fmt.Printf("2222")
	fmt.Print(lists)

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var res *ListNode = &ListNode{1, nil}
	start := res

	for nil != list1 && nil != list2 {
		if list1.Val < list2.Val {
			start.Next = list1
			list1 = list1.Next
		} else {
			start.Next = list2
			list2 = list2.Next
		}
		start = start.Next
	}
	if list1 == nil {
		start.Next = list2
	}
	if list2 == nil {
		start.Next = list1
	}
	return res.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
