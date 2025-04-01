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

### [21 合并两个有序链表](https://leetcode.com/problems/merge-two-sorted-lists) [简单 链表型]

> You are given the heads of two sorted linked lists `list1` and `list2`.
>
> Merge the two lists into one **sorted** list. The list should be made by splicing together the nodes of the first two lists.
>
> Return *the head of the merged linked list*.

#### 解题思路

- 首刷：归并排序最基本的函数，写不出来可以自裁了（注意处理特殊情况，即**两个链表有一个是空的情况**）

### *[136 只出现一次的数字](https://leetcode.com/problems/single-number) [简单 位运算型型]

> Given a **non-empty** array of integers `nums`, every element appears *twice* except for one. Find that single one.
>
> You must implement a solution with a linear runtime complexity and use only constant extra space.

#### 解题思路

- 首刷：注意到 `nums` 的和一定为 $2*n + k, k 为那个只出现一次的数$，记为 $sum$ 。（别注意了，你是傻逼）
- 看题解后：显然需要让成对的数字相互抵消，而异或运算符合符合我们的要求。

### [64 最小路径和](https://leetcode.com/problems/minimum-path-sum) [中等 动态规划型]

> Given a `m x n` `grid` filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.
>
> **Note:** You can only move either down or right at any point in time.

#### 解题思路

- 首刷：定义一个`m x n ` 的矩阵 `dp`，用来有 $dp[i][j]$ 表示从左上到 $grid[i][j]$ 的最小路径和，则有递推公式 $dp[i][j] = min\{dp[i-1][j], dp[i][j-1]\} + grid[i][j]$，一路规划到右下角即可；同时，可以简化 `dp` 矩阵的大小，即用一个长为 `n` 的数组保存所有在用的子任务信息，最终做到 $O(n^2)$ 的时间复杂度，$O(n)$ 的空间复杂度

### [1 两数之和](https://leetcode.com/problems/two-sum) [简单 数组型]

> Given an array of integers `nums` and an integer `target`, return *indices of the two numbers such that they add up to `target`*.
>
> You may assume that each input would have ***exactly* one solution**, and you may not use the *same* element twice.
>
> You can return the answer in any order.

#### 解题思路

- 首刷：遍历一遍数组，对于遇到的每一个数 `num`，都将 `target-num` 放入哈希表中，同时会先将其 `num` 值与哈希表中前面的 `target-num` 匹配，匹配成功则返回结果

### [160 相交链表](https://leetcode.com/problems/intersection-of-two-linked-lists) [简单 链表型]

> Given the heads of two singly linked-lists `headA` and `headB`, return *the node at which the two lists intersect*. If the two linked lists have no intersection at all, return `null`.
>
> For example, the following two linked lists begin to intersect at node `c1`:
>
> ![img](https://minio.noteikoh.top:443/blog-eikoh/typora/2025/01/18/160_statement.png)
>
> The test cases are generated such that there are no cycles anywhere in the entire linked structure.
>
> **Note** that the linked lists must **retain their original structure** after the function returns.
>
> **Custom Judge:**
>
> The inputs to the **judge** are given as follows (your program is **not** given these inputs):
>
> - `intersectVal` - The value of the node where the intersection occurs. This is `0` if there is no intersected node.
> - `listA` - The first linked list.
> - `listB` - The second linked list.
> - `skipA` - The number of nodes to skip ahead in `listA` (starting from the head) to get to the intersected node.
> - `skipB` - The number of nodes to skip ahead in `listB` (starting from the head) to get to the intersected node.
>
> The judge will then create the linked structure based on these inputs and pass the two heads, `headA` and `headB` to your program. If you correctly return the intersected node, then your solution will be **accepted**.

#### 解题思路：

- 首刷：主要问题在于同步两个链表在相交之前的长度，只需要**让 `NA` 先遍历完 `B` 后再去遍历 `A`， `NB` 先遍历完 `A` 后再去遍历 `B`，这样就可以让相交处前的长度一致**，遍历的同时比较两个结点的地址是否相同即可解决问题。

### [437 路径总和 III](https://leetcode.com/problems/path-sum-iii) [中等 二叉树型]

> Given the `root` of a binary tree and an integer `targetSum`, return *the number of paths where the sum of the values along the path equals* `targetSum`.
>
> The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).

#### 解题思路

- 首刷：前序遍历所有结点，在遍历的同时维护一个哈希表，哈希表里面存**从根节点到当前结点的所有前缀和**；遍历到每个结点时用 `targetSum - node.Val` 去匹配哈希表更新结果，之后将当前结点的前缀和存入哈希表，在遍历完子结点后将前缀和从哈希表中删除。**注意要初始化哈希表（从根节点开始的路径）**

### [226 翻转二叉树](https://leetcode.com/problems/invert-binary-tree) [简单 二叉树型]

> Given the `root` of a binary tree, invert the tree, and return *its root*.

#### 解题思路

- 首刷：进入每个节点，转换左右节点即可，只需要注意边界和特殊值的处理

### [104 二叉树的最大深度](https://leetcode.com/problems/maximum-depth-of-binary-tree) [简单 二叉树型]

> Given the `root` of a binary tree, return *its maximum depth*.
>
> A binary tree's **maximum depth** is the number of nodes along the longest path from the root node down to the farthest leaf node.

#### 解题思路

- 首刷：遍历一遍数，更新深度并返回即可

### [56 合并区间](https://leetcode.com/problems/merge-intervals) [中等 数组型]

> Given an array of `intervals` where `intervals[i] = [starti, endi]`, merge all overlapping intervals, and return *an array of the non-overlapping intervals that cover all the intervals in the input*.

#### 解题思路

- 首刷：有区间的起点在另一区间的终点之前（或重合）且终点在这区间的起点之后（或重合）即可合并，合并后的区间终点用较大值、起点用较小值，如果将区间按照起点排序可以进行简单处理（保证了区间起点在后一区间的终点之前，且如果有能够合并的几个区间，那么它们一定会是连续出现的，避免了遍历查找）

### *[121 买卖股票的最佳时机](https://leetcode.com/problems/best-time-to-buy-and-sell-stock) [简单 数组型]

> You are given an array `prices` where `prices[i]` is the price of a given stock on the `ith` day.
>
> You want to maximize your profit by choosing a **single day** to buy one stock and choosing a **different day in the future** to sell that stock.
>
> Return *the maximum profit you can achieve from this transaction*. If you cannot achieve any profit, return `0`.

#### 解题思路

- 首刷：遍历一遍数组，定义一个与 `prices` 等长的数组 `max`，其中 `max[i]` 为 `i` 天之后的最高价。最后再遍历一遍数组尝试低买高卖即可
- 看解析后：记录最低价则会将时间减半，我是傻逼

#### [206 反转链表](https://leetcode.com/problems/reverse-linked-list) [简单 链表型]

> Given the `head` of a singly linked list, reverse the list, and return *the reversed list*.

#### 解题思路

- 首刷：没啥好多的，反转即可

### [101 对称二叉树](https://leetcode.com/problems/symmetric-tree) [简单 二叉树型]

> Given the `root` of a binary tree, *check whether it is a mirror of itself* (i.e., symmetric around its center).

#### 解题思路

- 首刷：同时递归根节点的左右子节点即可

### [543 二叉树的直径](https://leetcode.com/problems/diameter-of-binary-tree) [简单 二叉树型]

> Given the `root` of a binary tree, return *the length of the **diameter** of the tree*.
>
> The **diameter** of a binary tree is the **length** of the longest path between any two nodes in a tree. This path may or may not pass through the `root`.
>
> The **length** of a path between two nodes is represented by the number of edges between them.

#### 解题思路

- 首刷：直径一定会是某个节点分别从左孩子与右孩子出发到达所有叶结点走过的边的最大值之和，我们在后序遍历的时候可以统计这一值，返回最大的值即可

### [406 根据身高重建队列](https://leetcode.com/problems/queue-reconstruction-by-height) [中等 链表型]

> You are given an array of people, `people`, which are the attributes of some people in a queue (not necessarily in order). Each `people[i] = [hi, ki]` represents the `ith` person of height `hi` with **exactly** `ki` other people in front who have a height greater than or equal to `hi`.
>
> Reconstruct and return *the queue that is represented by the input array* `people`. The returned queue should be formatted as an array `queue`, where `queue[j] = [hj, kj]` is the attributes of the `jth` person in the queue (`queue[0]` is the person at the front of the queue).

#### 解题思路

- 首刷：依据身高排序建立一个链表，链表的内容包括每个 `people` 的身高、期望在队列前的人数、实际在队列前的人数，初始化时前两项照抄最后一项设为0,然后不停遍历链表将符合要求的 `people` 放入结果队列中

### [70 爬楼梯](https://leetcode.com/problems/climbing-stairs) [简单 基础]

> You are climbing a staircase. It takes `n` steps to reach the top.
>
> Each time you can either climb `1` or `2` steps. In how many distinct ways can you climb to the top?

#### 解题思路

- 首刷：对于 $n=2*k+\nu,\nu = 0|1)$，那么从 `i=0` 开始，用 `i` 次爬两步 + `n-2*i` 次爬一步即为一种方案，又有 $C_{n - i}^{i}=(n-i)*(n-i-1)*...*(n-2*i)/i*(i-1)*...*1$ 种排列组合方式，将 `i` 从 `0` 循环到 `k` 得到的排列方式之和即为结果

### [155 最小栈](https://leetcode.com/problems/min-stack) [中等 基础]

> Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
>
> Implement the `MinStack` class:
>
> - `MinStack()` initializes the stack object.
> - `void push(int val)` pushes the element `val` onto the stack.
> - `void pop()` removes the element on the top of the stack.
> - `int top()` gets the top element of the stack.
> - `int getMin()` retrieves the minimum element in the stack.
>
> You must implement a solution with `O(1)` time complexity for each function.

#### 解题思路

- 首刷: 在基本的栈数据结构的基础上，加入一个数组 `min`, `min[n]` 对应栈的前 `n` 个项中的最小元素。同时还需要一个记录栈长度的 `int` 型变量，用于取出最小的元素

### [221 最大正方形](https://leetcode.com/problems/maximal-square) [中等 动态规划型]

> Given an `m x n` binary `matrix` filled with `0`'s and `1`'s, *find the largest square containing only* `1`'s *and return its area*.

#### 解题思路

- 首刷：对于任意位置，如果它对应在 `matrix` 矩阵中的位置是 1，记它的左方\上方\左上方三个位置分别为 `l` \ `t` \ `lt`，如果以这三个位置为右下角的三个 `n-1` 长的全 1 正方形存在，那么就存在 `n` 长的全 1 正方形，对应在代码中即为:`dp[i, j] = min(dp[i-1, j], dp[i, j-1], dp[i, j-1]) + 1`；否则记为 0 即可

### [55 跳跃游戏](https://leetcode.com/problems/jump-game) [中等 数组型]

> You are given an integer array `nums`. You are initially positioned at the array's **first index**, and each element in the array represents your maximum jump length at that position.
>
> Return `true` *if you can reach the last index, or* `false` *otherwise*.

#### 解题思路

- 首刷：倒推法，定义一个标记数组 `flag`，与 `nums` 等长，将其最后一位置为 1，然后不停从后往前遍历，只要其走到的范围内有 1 就可将其置为 1，查看第一位是否能置 1 即可（**能成，但是时间和空间复杂度都不太理想**）
- **优化**：记 `nums` 的长度为 `n` ，那么对于第 `i` 个元素，他想到到达最后一个元素的条件是 `nums[i] >= n - i`，我们从后往前遍历，将每个元素更新为 `nums[i] - n + i`，一旦遇到非负数我们就将 `n` 替换为 `i` ，继续遍历，查看第一个元素的值是否大于 0 即可

### [337 打家劫舍 III](https://leetcode.com/problems/house-robber-iii) [中等 二叉树型]

> The thief has found himself a new place for his thievery again. There is only one entrance to this area, called `root`.
>
> Besides the `root`, each house has one and only one parent house. After a tour, the smart thief realized that all houses in this place form a binary tree. It will automatically contact the police if **two directly-linked houses were broken into on the same night**.
>
> Given the `root` of the binary tree, return *the maximum amount of money the thief can rob **without alerting the police***.

#### 解题思路

- 首刷：从叶结点向上推，对于每个结点，抢了就不能抢子节点，所以给出两个值，`rubNode` 与 `rubNext`，分别是抢这个节点与抢子节点的收入，回到这个节点的父节点，记价值为 `Val`，继续向上给出 `val + L.rubNext + R.rubNext` 与 `max(L.rubNode, L.rubNext) + max(R.rubNode, R.rubNext)`，一直回到根节点计算得出最大值。

### [48 旋转图像](https://leetcode.com/problems/rotate-image) [中等 数组型]

> You are given an `n x n` 2D `matrix` representing an image, rotate the image by **90** degrees (clockwise).
>
> You have to rotate the image [**in-place**](https://en.wikipedia.org/wiki/In-place_algorithm), which means you have to modify the input 2D matrix directly. **DO NOT** allocate another 2D matrix and do the rotation.

#### 解题思路

- 首刷：没有额外空间，意味着只能通过交换达到目的，对于 `i, j` 处的元素，在正方形中与它水平、垂直轴对称的元素分别为 `n - i - 1, j`, `i, n - j - 1`，与它中心点对称的元素为 `n - i - 1, n - j - 1`。它相对于中心的距离为 `i - n/2, j - n/2`，对它来说以中心顺时针旋转的四个元素分别为`i, j`，  `j, n - i - 1`、`n - i - 1, n - j - 1 `、`n - j - 1, i`，在这四个元素按顺序填入 `n - j - 1, i`， `i, j`，  `j, n - i - 1`、`n - i - 1, n - j - 1 `的值即可。为了避免重复的替换，遍历的范围需要限制为整个正方型左上角的四分之一，即 `0~(n/2)+(n%2)`, `0~(n/2)-1`

### *[39 组合总和](https://leetcode.com/problems/combination-sum) [中等 回溯法典型]

> Given an array of **distinct** integers `candidates` and a target integer `target`, return *a list of all **unique combinations** of* `candidates` *where the chosen numbers sum to* `target`*.* You may return the combinations in **any order**.
>
> The **same** number may be chosen from `candidates` an **unlimited number of times**. Two combinations are unique if the frequency of at least one of the chosen numbers is different.
>
> The test cases are generated such that the number of unique combinations that sum up to `target` is less than `150` combinations for the given input.

#### 解题思路

- 首刷：定义一个哈希表 `hash`，`Key` 为 `int` 型，`Val` 为 `int` 数组，遍历 `candidates` 数组，对于每个 `candidate`，遍历它在 $1 \to target$ 范围内的所有整数倍数：记每个倍数为 `multi`，在 `hash` 表中记录 `hash[multi].append(candidate)`... **太复杂了没有可行性，没做出来**
- 看了题解后：完全背包说是，有动态规划和回溯两种解法。
  1. 使用搜索回溯法，用一个 `DFS`  搜下去就可以，只不过需要考虑每个数是可以重复的，因此在 `DFS` 函数的尾部要同时进入 **使用并留在当前元素/不使用并跳过当前元素** 两个分支
  2. 看作一个完全背包，**<u>待补充</u>**

### *[322 零钱兑换](https://leetcode.com/problems/coin-change) [中等 完全背包典型]

> You are given an integer array `coins` representing coins of different denominations and an integer `amount` representing a total amount of money.
>
> Return *the fewest number of coins that you need to make up that amount*. If that amount of money cannot be made up by any combination of the coins, return `-1`.
>
> You may assume that you have an infinite number of each kind of coin.

#### 解题思路

- 首刷：考虑回溯法，找到所有的解法并选择硬币枚数最小的解法，没有则返回 `-1` （**超时了**）
- 优化：**没法剪枝，直接动态规划**，定义 `dp[i]` 为兑换 `i` 元所需的最少硬币数量，如果不能兑换则将值填为 `amount + 1`，这样对于 `i` 可以有递推公式 $dp[i] = min\{dp[i - coin] + 1, amount + 1\}$（`for coin in coins`），如果 `dp[amount] <= amount` 则返回 `dp[amount]` 即可，否则返回 -1

### [148 排序链表](https://leetcode.com/problems/sort-list) [中等 归并排序典型]

> Given the `head` of a linked list, return *the list after sorting it in **ascending order***.

#### 解题思路

- 首刷：归并排序即可，不要手生写不出来

### *[215 数组中的第K个最大元素](https://leetcode.com/problems/kth-largest-element-in-an-array) [中等 数组型]

> Given an integer array `nums` and an integer `k`, return *the* `kth` *largest element in the array*.
>
> Note that it is the `kth` largest element in the sorted order, not the `kth` distinct element.
>
> Can you solve it without sorting?

#### 解题思路

- 首刷：借用快排的思路，以第一个元素为基准，将比它大的放在左边，比它小的放在右边，根据两边的数量进入选择一段继续寻找（**做出来了，但是排名不高，非常狼狈**）

- 优化：优化快排划分的思路，减少不必要的操作：

  1. 进行一个思维转换，第K个最大元素，在排序后的数组中处在 `nums[k-1]` 的位置，我们去找排序后的第 `k-1` 号位置上的元素即可

  2. 通过下面的函数进行划分，划分完毕后，`[left ~ r]` 与 `[r+1, right]` 上的元素在排序后位置锁定（`r` 右边的所有元素一定小于等于 `bench`），**那么我们就可以根据 `r` 的大小选择进入左右两边的哪一边进行查找**

     ```go
     l, r := left-1, right+1
     bench := nums[left]
     for l < r {
         for l++; nums[l] > bench; l++ {
         }
         for r--; nums[r] < bench; r-- {
         }
         if l < r {
             nums[l], nums[r] = nums[r], nums[l]
         }
     }
     ```

  3. 当最后范围缩小到 `left = right` 时，说明我们已经找到了排序后位置锁定在 `nums[k-1]` 上的元素，这个时候返回 `nums[k-1]/nums[left]/nums[right]` 三者之一即可

### [46 全排列](https://leetcode.com/problems/permutations) [中等 深度搜索典型]

> Given an array `nums` of distinct integers, return all the possible permutations. You can return the answer in **any order**.

#### 解题思路

- 首刷：定义一个 `DFS` 函数与一个标记数组 `flag`，然后常规进行深度优先搜索即可

### [538 把二叉搜索树转换为累加树](https://leetcode.com/problems/convert-bst-to-greater-tree) [中等 二叉树型]

> Given the `root` of a Binary Search Tree (BST), convert it to a Greater Tree such that every key of the original BST is changed to the original key plus the sum of all keys greater than the original key in BST.
>
> As a reminder, a *binary search tree* is a tree that satisfies these constraints:
>
> - The left subtree of a node contains only nodes with keys **less than** the node's key.
> - The right subtree of a node contains only nodes with keys **greater than** the node's key.
> - Both the left and right subtrees must also be binary search trees.

#### 解题思路

- 首刷：对于任意一个节点，首先需要加上它右子树上所有节点的和，其次如果它是父节点的左孩子则还需要加上父节点的累加值。因此遍历节点时，先向右孩子传承累积和的同时获取右孩子分支的节点和，并计算出自己的累加值，在将自己的累加值传给左孩子。

### *[399 除法求值](https://leetcode.com/problems/evaluate-division) [中等 并查集典型]

> `[["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]`
>
> You are given an array of variable pairs `equations` and an array of real numbers `values`, where `equations[i] = [Ai, Bi]` and `values[i]` represent the equation `Ai / Bi = values[i]`. Each `Ai` or `Bi` is a string that represents a single variable.
>
> You are also given some `queries`, where `queries[j] = [Cj, Dj]` represents the `jth` query where you must find the answer for `Cj / Dj = ?`.
>
> Return *the answers to all queries*. If a single answer cannot be determined, return `-1.0`.
>
> **Note:** The input is always valid. You may assume that evaluating the queries will not result in division by zero and that there is no contradiction.
>
> **Note:** The variables that do not occur in the list of equations are undefined, so the answer cannot be determined for them.

#### 解题思路

- 首刷：用一个哈希表 `hash` 存储每个字符可以转化的另一个字符，这样对于每个计算请求，将分子和分母统一转换到不能继续转换为止，之后比较分子分母的字符，如果不同说明不能计算，如果相同则计算值（**拼尽全力无法战胜，考虑得太少了**）
- 看题解后：并查集说是，一个并查集就是一个包含了一些节点之间相互关系的有向加权图，有一套特定的实现
