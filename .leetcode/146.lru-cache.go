/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
type doublyLink struct {
	key  int
	val  int
	prev *doublyLink
	next *doublyLink
}

type LRUCache struct {
	hash map[int]*doublyLink
	size int
	cap  int
	head *doublyLink
	end  *doublyLink
}

func Constructor(capacity int) LRUCache {
	link := &doublyLink{
		key:  -1,
		val:  -1,
		prev: nil,
		next: nil,
	}
	return LRUCache{
		hash: make(map[int]*doublyLink),
		size: 0,
		cap:  capacity,
		head: link,
		end:  link,
	}
}

func (this *LRUCache) moveToHead(node *doublyLink) {
	if this.head.next == node {
		// 已经在头部，跳过
		return
	}
	if this.head != this.end {
		prev, next := node.prev, node.next
		if prev != nil {
			prev.next = next
		}
		if next != nil {
			next.prev = prev
		}
		if this.end == node {
			this.end = prev
		}
		this.head.next.prev = node
	} else {
		this.end = node
	}
	node.prev = this.head
	node.next = this.head.next
	this.head.next = node
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.hash[key]
	if exist {
		this.moveToHead(node)
		return node.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	node, exist := this.hash[key]
	if exist {
		node.val = value
		this.moveToHead(node)
	} else {
		if this.size == this.cap {
			// 装满了，要杀掉最后链表的最后一个
			delete(this.hash, this.end.key)
			this.end = this.end.prev
			this.end.next = nil
		} else {
			this.size++
		}
		// 放入 cache 中
		new := &doublyLink{
			key:  key,
			val:  value,
			prev: nil,
			next: nil,
		}
		this.hash[key] = new
		this.moveToHead(new)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

