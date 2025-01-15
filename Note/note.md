# 		LeetCode Hot100 刷题笔记

### [283 移动零](https://leetcode.com/problems/move-zeroes) [简单 排序型]

> Given an integer array `nums`, move all `0`'s to the end of it while maintaining the relative order of the non-zero elements.
>
> **Note** that you must do this in-place without making a copy of the array.

#### 解题思路

- 首刷：前后指针法，`front` 指针寻找第一个0，`end` 指针在 `front` 之后寻找第一个数字。直到 `end` 指针遍历到数组末尾
- 刷完后：大可不必如此复杂，**遍历一遍数组将所有非 0 元素移动到数组前部并计数**，之后将数组的后半部分填充为 0 即可

### *[169 多数元素](https://leetcode.com/problems/majority-element) [简单 排序型]

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

### [94 二叉树的中序遍历](https://leetcode.com/problems/binary-tree-inorder-traversal) [简单 二叉树型]

> Given the `root` of a binary tree, return *the inorder traversal of its nodes' values*.

#### 解题思路

- 首刷：中序遍历的定义为[左子树 -> 根节点 -> 右子树]，按照这个顺序递归遍历即可

### [461 汉明距离](https://leetcode.com/problems/hamming-distance) [简单 位运算型]

> The [Hamming distance](https://en.wikipedia.org/wiki/Hamming_distance) between two integers is the number of positions at which the corresponding bits are different.
>
> Given two integers `x` and `y`, return *the **Hamming distance** between them*.

#### 解题思路

- 首刷：汉明距离的定义为[两个二进制数之间对应位置存在不同的数量]。不停 $mod$ 2并除以2地同时比较 $mod$ 2 的结果即可，最终有一方归零后加上另一方继续除以 2 的次数； 或者**异或后统计 1 的数量**

### [98 验证二叉搜索树](https://leetcode.com/problems/validate-binary-search-tree) [中等 二叉树型]

> Given the `root` of a binary tree, *determine if it is a valid binary search tree (BST)*.
>
> A **valid BST** is defined as follows:
>
> - The left subtree of a node contains only nodes with keys **less than** the node's key.
> - The right subtree of a node contains only nodes with keys **greater than** the node's key.
> - Both the left and right subtrees must also be binary search trees.

#### 解题思路

- 首刷：递归地进行判断即可，判断当前节点，将结果与左节点、右节点的结果相与
- 提交错误后：递归地进行判断，但是考虑在递归的过程中引入边界，**只将当前节点与边界做比较，并在递归的过程中更新边界**

### [448 找到所有数组中消失的数字](https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array) [简单 数组型]

> Given an array `nums` of `n` integers where `nums[i]` is in the range `[1, n]`, return *an array of all the integers in the range* `[1, n]` *that do not appear in* `nums`.
>
> **Follow up:** Could you do it without extra space and in `O(n)` runtime? You may assume the returned list does not count as extra space.

#### 解题思路

- 首刷：新建一个长度为 `n` 的数组 `count`，有 `count[k-1]` 为 `k` 在 `nums` 中出现的次数，初始值为0，遍历数组后筛选出为 0 的项的 `index` 记录在数组前方
- 看题解后：在传入的数组中进行标记操作，若 `nums[i] = k`，则使 `nums[(k - 1) % n] += n`（如果 `k` 一次都没有出现那么 `nums[k-1]` 将会保持值$\le n$），遍历传入数组即可

### [617 合并二叉树](https://leetcode.com/problems/merge-two-binary-trees) [简单 二叉树型]

> You are given two binary trees `root1` and `root2`.
>
> Imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not. You need to merge the two trees into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of the new tree.
>
> Return *the merged tree*.
>
> **Note:** The merging process must start from the root nodes of both trees.

#### 解题思路

- 首刷：定义一个合并函数，传入两树对应位置的节点。若两个节点皆有效，则合并值到新树节点上，并继续合并左子树、右子树后添加到新树节点对应位置；如果存在 `nil` 值，则直接返回另一方的节点

### *[560 和为 K 的子数组](https://leetcode.com/problems/subarray-sum-equals-k) [中等 动态规划型]

> Given an array of integers `nums` and an integer `k`, return *the total number of subarrays whose sum equals to* `k`.
>
> A subarray is a contiguous **non-empty** sequence of elements within an array.

#### 解题思路

- 首刷：定义一个 `n*n` 的矩阵 `dp`，其中 `dp[i, j]` 的含义为 `nums` 从 `i` 到 `j` 的序列和，那么能够递归地有 $dp[i, j] = dp[i, j-1] + nums[j]$，则能在 $O(n^2)$ 时间内解决问题。（或者定义一个大小为 `n` 的数组亦能解决问题，不过代码复杂度有所增加）
- 看到时间有点久的改良：定义 `sum` 数组，有 `sum[i]` 为 `nums` 从 `1` 到 `i` 的和，取代 `dp`，并注意到 `sum[j] - sum[i]` 为 `nums` 从 `i + 1` 到 `j` 的序列和，那么可以在求出 `sum` 数组的同时尝试找到和为 `k` 的子数组（优化了 500ms，提升了 5% 的排名) 
- 看了题解后：继续注意到  `sum[i] + k = sum[j]` 表示从 `i + 1` 到 `j` 的序列和为 `k`，那么我们用哈希表统计 `sum[i] + k` 的值出现的次数，同时去除冗余的 `sum` 数组，每次求出 `sum[i]` 后只需要在哈希表中查找一次即可（优化了遍历的时间），优化后的时间复杂度和空间复杂度都为 $O(n)$，十分巧妙

### [20 有效的括号](https://leetcode.com/problems/valid-parentheses) [简单 栈型]

> Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.
>
> An input string is valid if:
>
> 1. Open brackets must be closed by the same type of brackets.
> 2. Open brackets must be closed in the correct order.
> 3. Every close bracket has a corresponding open bracket of the same type.

#### 解题思路

- 首刷：可以用栈的数据结构解决，发现左括号就加入到栈里面，发现右括号就尝试与栈顶括号匹配，不匹配就返回 `false`。

### [338 比特位计数](https://leetcode.com/problems/counting-bits) [简单 位运算型]

> Given an integer `n`, return *an array* `ans` *of length* `n + 1` *such that for each* `i` (`0 <= i <= n`)*,* `ans[i]` *is the **number of*** `1`***'s** in the binary representation of* `i`.

#### 解题思路

- 首刷：依次处理每个数即可，注意有如下规律：$ans[i] = ans[i-1] + 1, i 为奇数$，$ans[i] = ans[i - 1] + 1 - \{(i-1) 末尾连续的1\}, i为偶数$。

### *[240 搜索二维矩阵 II](https://leetcode.com/problems/search-a-2d-matrix-ii) [中等 数组型]

> Write an efficient algorithm that searches for a value `target` in an `m x n` integer matrix `matrix`. This matrix has the following properties:
>
> - Integers in each row are sorted in ascending from left to right.
> - Integers in each column are sorted in ascending from top to bottom.

#### 解题思路

- 首刷：为了利用矩阵的属性，首先需要一个标记矩阵，记录搜索过的点。对于搜索中的每一个点，直接查看它的右、下、右下三个方向的值：如果有则返回 `true`；如果没有则继续搜索三个值中小于目标值且从未搜索过的位置，没有符合条件的则返回 `false`（始终 de 不完 bug,遂放弃）
- 看解析后：每一行逐行进行二分查找，$O(m * log(n))$解决问题；或者从右上角开始**z型搜索**，对于$matrix[x, y] > target$，抛弃这一列（往下的所有元素都会大于 `target`），对于$matrix[x,y] < target$，左侧的元素肯定不合规，而右侧的元素已经被抛弃过，开始往下尝试搜索

