# 		LeetCode Hot100 刷题笔记

### [283 移动零](https://leetcode.com/problems/move-zeroes) [简单 排序型]

> Given an integer array `nums`, move all `0`'s to the end of it while maintaining the relative order of the non-zero elements.
>
> **Note** that you must do this in-place without making a copy of the array.

#### 解题思路

- 首刷：前后指针法，`front` 指针寻找第一个0，`end` 指针在 `front` 之后寻找第一个数字。直到 `end` 指针遍历到数组末尾
- 刷完后：大可不必如此复杂，**遍历一遍数组将所有非 0 元素移动到数组前部并计数**，之后将数组的后半部分填充为 0 即可

### [169 多数元素](https://leetcode.com/problems/majority-element) [简单 排序型]

> Given an array `nums` of size `n`, return *the majority element*.
>
> The majority element is the element that appears more than `⌊n / 2⌋` times. You may assume that the majority element always exists in the array.

#### 解题思路

- 首刷：**排序后在中间**的元素就是多数元素$O(nlog_n)$ $O(1)$；哈希表计数每个元素，之后选出最多的元素 $O(n)\space O(1)$；
- 看解析后：**摩尔投票法** | **首先**，<u>可以证明最终不会一个数字都不剩</u>。**原因**： 假设两两抵消之后，最终一个数字都不剩。那么就是说一共有偶数个数字，假设有`n`个，那么`n = 2k`，`k`是整数。所以最多会进行k次两两抵消。又因为一定存在众数 (数量超过`⌊n/2⌋ = k`的数字 ，那么这个众数一定会在抵消完毕后留在数组中

### [19 删除链表的倒数第 N 个结点](https://leetcode.com/problems/remove-nth-node-from-end-of-list) [中等 链表型]

> Given the `head` of a linked list, remove the `nth` node from the end of the list and return its head.

#### 解题思路

- 首刷：遍历一遍链表读取长度，之后找到倒数第 N 个结点并删除之

### [141 环形链表](https://leetcode.com/problems/linked-list-cycle) [简单 链表型]

> Given `head`, the head of a linked list, determine if the linked list has a cycle in it.
>
> There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to. **Note that `pos` is not passed as a parameter**.
>
> Return `true` *if there is a cycle in the linked list*. Otherwise, return `false`.

#### 解题思路

- 首刷：快慢节点法，快节点一次前进 2 节点，慢节点一次前进 1 节点，如果两个节点能够相遇那么说明有环。同时可以注意到，假设链表长度为 `n`，并且在 `k` 处有环，而两节点相遇在 `m` 处，那么有下面的等式 $k + m + \mu*(n - k) = 2*(k + m + \nu*(n-k))$ ，即快节点前进的节点数是慢节点的两倍，化简这个式子有 $k+m=(\mu-2\nu)*(n-k)$，其中 $n-k$ 可由慢节点沿环走一圈求出，$m$ 可由 `head` 走到两节点相遇处求出

### [142 环形链表 II](https://leetcode.com/problems/linked-list-cycle-ii) [中等 链表型]

> Given the `head` of a linked list, return *the node where the cycle begins. If there is no cycle, return* `null`.
>
> There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to (**0-indexed**). It is `-1` if there is no cycle. **Note that** `pos` **is not passed as a parameter**.
>
> **Do not modify** the linked list.

#### 解题思路

- 首刷：快慢节点法，接上一题，在已知 $n-k$ 处节点的情况下，让 `head` 节点与 $n-k$ 处节点一起行动，那么在一起行动 $k$ 步后，两节点将会相遇在循环开始的节点，从而解决问题

### [234 回文链表](https://leetcode.com/problems/palindrome-linked-list) [简单 链表型]

> Given the `head` of a singly linked list, return `true` *if it is a* *palindrome* *or* `false` *otherwise*.

#### 解题思路

- 首刷：快慢节点法可以找到链表的中位节点，从而将链表一分为二，然后转置后部分链表后与前半部分进行比较即可