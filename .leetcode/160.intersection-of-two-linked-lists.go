/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	NA := headB
	NB := headA
	for NA != NB {
		if NA == nil {
			NA = headA
		} else {
			NA = NA.Next
		}
		if NB == nil {
			NB = headB
		} else {
			NB = NB.Next
		}
	}
	return NA
}

// @lc code=end

