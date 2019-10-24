package leetcode

// 234. 回文链表 https://leetcode-cn.com/problems/palindrome-linked-list/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head // 快指针每轮多走1步
	fast := head // 慢指针
	var pre *ListNode = nil
	var temp *ListNode = nil

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		// 反转
		temp = slow.Next
		slow.Next = pre
		pre = slow
		slow = temp
	}
	if fast != nil {
		slow = slow.Next
	}

	for slow != nil {
		if slow.Val != pre.Val {
			return false
		}
		slow = slow.Next
		pre = pre.Next
	}

	return true
}
