package main

import "fmt"

func main() {
	var list1 = &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	var list2 = &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	result := mergeTwoLists(list1, list2)

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Dummy/sentinel untuk memudahkan penyambungan
	dummy := &ListNode{}
	curr := dummy

	// Selama keduanya masih ada, pilih nilai lebih kecil
	for list1 != nil && list2 != nil {
		// Pilih node dengan nilai lebih kecil
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		// Pindah ke node berikutnya di hasil
		curr = curr.Next
	}

	// Salah satu list masih tersisa: sambungkan langsung (sudah terurut)
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}

	// Kembalikan hasil, lewati dummy
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
