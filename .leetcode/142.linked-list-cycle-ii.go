/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/* func detectCycle(head *ListNode) *ListNode {
	// 定义快慢节点
	fast, slow := head, head
	// 快指针每次走两步，慢指针每次走一步
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 如果快慢节点相遇，则存在环
		if slow == fast {
			// 寻找 n - k 的位置，即环的长度
			nSubk := head.Next
			slow = slow.Next
			for slow != fast {
				nSubk = nSubk.Next
				slow = slow.Next
			}
			// 寻找 k 的位置，即环的入口点
			for head != nSubk {
				head = head.Next
				nSubk = nSubk.Next
			}
			// 返回环的起始节点
			return head
		}
	}
	return nil
}*/

func detectCycle(head *ListNode) *ListNode {
	// 定义快慢节点
	fast, slow := head, head
	// 快指针每次走两步，慢指针每次走一步
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 如果快慢节点相遇，则存在环
		if slow == fast {
			fast = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}

// @lc code=end

