package day30_234_isPalindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	data := make([]int, 0, 20)
	for t := head; t != nil; t = t.Next {
		data = append(data, t.Val)
	}
	left, right := 0, len(data)-1
	for left < right {
		if data[left] != data[right] {
			return false
		}
		left++
		right--
	}
	return true
}
