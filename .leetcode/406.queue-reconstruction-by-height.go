/*
 * @lc app=leetcode id=406 lang=golang
 *
 * [406] Queue Reconstruction by Height
 */

// @lc code=start

type Queue struct {
	Height  int
	Target  int
	Current int
	Next    *Queue
}

func findMin(people [][]int, index int) (rst []int) {
	tmp := index
	target := index
	rst = people[index]
	index++
	for ; index < len(people); index++ {
		if people[index][0] < rst[0] {
			rst = people[index]
			target = index
		}
	}
	people[target] = people[tmp]
	return
}

func buildQueue(people [][]int) (head *Queue) {
	head = &Queue{}
	node := head
	for i := 0; i < len(people); i++ {
		min := findMin(people, i)
		node.Next = &Queue{Height: min[0], Target: min[1]}
		node = node.Next
	}
	return head
}

func putPeople(head *Queue, people *Queue, queue [][]int, index int) {
	// 更新状态
	for head.Next != people {
		node := head.Next
		node.Current++
		head = node
	}
	// 放入队列
	queue[index] = []int{people.Height, people.Target}
	head.Next = people.Next
	// 继续更新后面的状态
	head = people.Next
	for head != nil && head.Height == people.Height {
		head.Current++
		head = head.Next
	}
}

func findPeople(head *Queue, queue [][]int, index int) {
	people := head.Next
	for people.Current != people.Target {
		people = people.Next
	}
	putPeople(head, people, queue, index)
}

func reconstructQueue(people [][]int) [][]int {
	queue := buildQueue(people)
	for i := 0; i < len(people); i++ {
		findPeople(queue, people, i)
	}
	return people
}

// @lc code=end

