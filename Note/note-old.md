# 		LeetCode Hot100 刷题笔记

## 	160 [相交链表](https://leetcode.com/problems/intersection-of-two-linked-lists/description/)

给两个单链表的头节点，找出并返回这两个链表相交的起始节点，不存在则返回NULL

![img](http://eikoh-itx:9000/blog-eikoh/typora/2024/09/03/160_statement.png)

如上图的两个链表在C1处相交

解题思路：两个单链表在相交处的节点具有相同的节点地址，也就是有相同的`ListNode *`的值，我们只需要考虑如何在$O(m+n)$时间内找到第一个相同的节点即可完成这道题目

解法一：

```c++
class Solution {
public:
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        ListNode *temp1 = headA;
        ListNode *temp2 = headB;
        while(temp1 != temp2) {
            temp1 = (temp1 == NULL ? headB : temp1->next);
            temp2 = (temp2 == NULL ? headA : temp2->next);
        }
        return temp1;
    }
};
```

时间：$O(m+n)$；空间：$O(1)$

巧妙的点在于两个temp节点并行的访问了ListA + ListB，从而统一了长度：

- 如果两个链表相交前等长，那么在$O(min(m,n))$时间内就能够找到相交的那个节点
- 如果两个链表相交前不等长，那么在两个temp节点访问完ListA/ListB之后转而去访问ListB/ListA时，统一了总的访问长度为$m+n$，从而对齐了末尾的几个节点，那么如果有相交，他们一定会在这个相交的地方相遇。
- 如果两个链表不相交，那么他们会同时在两个链表的末尾取得NULL值，这个时候将NULL值返回。

## 236 [二叉树的最近公共祖先](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/)

给定二叉树的根节点和$p、q$两个节点，返回这两个节点最近的公共祖先，也就是这两节点溯源上去第一次相交的节点

<img src="http://eikoh-itx:9000/blog-eikoh/typora/2024/09/03/image-20240629110943320.png" alt="image-20240629110943320" style="zoom:67%;" />

比如上面，给定6、8两个节点，他们的最近公共祖先为3,而给定5、2两个节点，他们的最近公共祖先为5

解题思路：涉及到二叉树，优先考虑递归能否解。这道题的朴素递归思路应该是从根节点开始，前序深度搜索这颗树，一直搜索到p、q节点或者空后返回。
在某一处节点：

1. 如果向左的深搜返回了一个不为空的节点地址，向右的深搜返回了null，那么说明该节点不是最近公共祖先，向左的深搜返回值是这个最近公共祖先，返回之即可。
2. 如果向右的深搜返回了一个不为空的节点地址，向左的深搜返回了null，那么说明该节点不是最近公共祖先，向右的深搜返回值是这个最近公共祖先，返回之即可。
3. 如果向左、向右的深搜同时返回了不为空的节点地址，则该节点是最近公共祖先，返回之即可。
4. 向左、向右的深搜同时返回了空，说明以该节点为根的子树压根没有p、q节点，返回空。

解法一：

```c++
class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        // 
        if (!root || root == p || root == q) return root;
        TreeNode* left = lowestCommonAncestor(root->left, p, q);
        TreeNode* right = lowestCommonAncestor(root->right, p, q);
        return !left ? right : !right ? left : root;
    }
};
```

## 234 [回文链表](https://leetcode.com/problems/palindrome-linked-list/description/?envType=problem-list-v2&envId=2cktkvj)

给定一个链表，判断其是否为回文链表（head = [1, 2, 2, 1]与head = [1, 2, 3, 2, 1]这种形状就是回文链表）。返回判断结果（是返回true,否返回false）。

解题思路：

1. 快速转置并判断转置后的链表与原链表是否相等（耗费空间大，但时间相对好看）
2. 利用与归并排序相同的拆分办法，通过快慢节点在O(n/2)时间内找到中间节点并拆分为两链表，转置前半部分链表并判断相等（无需与原链表比较，空间耗费少；时间也是线性）

解法一：巧妙的点在与转置了后半部分链表，而且统一了两部分链表的尾节点，从而巧妙的解决了链表长度为奇数时的问题，并且避免了链表中间部分在slow节点之前的问题

```c++
class Solution {
public:
    bool isPalindrome(ListNode* head) {
        // 利用快慢节点的方法找到中间节点
        ListNode *slow = head, *fast = head, *prev, *temp;
        while (fast && fast->next) {
            slow = slow->next;
            fast = fast->next->next;
        }
        // 此时，若链表长度为偶数，则slow节点与slow节点的上一节点之间为中间
        // 若链表长度为奇数，则slow节点为中间节点
        // 将slow节点作为后半部分链表转置的尾节点，让这个单链表变成两条在slow节点汇流的链表
        prev = slow;
        slow = slow->next;
        prev->next = NULL;
        // 转置后半部分链表
        while (slow) {
            temp = slow->next;
            slow->next = prev;
            prev = slow;
            slow = temp;
        }
        // 此时，prev为后半部分转置后的头节点，开始比较两链表是否相同
        fast = head;
        slow = prev;
        while (slow) { // 以后半部分链表为准是因为快慢节点拆分后一定有后半部分链表的长度小于等于前半部分
            if(fast->val != slow->val) return false;
            else {
                fast = fast->next;
                slow = slow->next;
            }
        }
        return true;
    }
};
```

## 739 [每日温度](https://leetcode.com			/problems/daily-temperatures/?envType=problem-list-v2&envId=2cktkvj)

给定一个数组，代表一些天的每日温度，返回一个数组，长度与给定的数组等长，每个数组内的元素是该天到之后最近且更温暖的一天间隔的天数（如果没有更温暖的天数则填0）。如[73,74,75,71,69,72,76,73]需要返回[1,1,4,2,1,1,0,0]。

解题思路：

1. 朴素的想法，每天都遍历这个数组，找到最近的更温暖的天，耗时$O(n^2)$（没试，多半超时）
2. 动态规划的思想，遍历等待的天数，从1天到n-1天，但是在每个元素做有效性检查来剪枝（试了，超时）
3. 用一个栈，栈里是没有找到升温等待时间的天数序号。用O(n)时间遍历数组，每次遍历先用当前元素去确定能否为栈里的那些天数找到升温等待时间（只要当前元素大于栈顶天数对应的日温度即可找到，为i-top()），在解决完所有能解决的之后将该元素的天数也入栈。

解法一：用了栈的先入后出特点，将气温高的天数放在栈的底部，优先解决气温低的天数的升温等待时间，直到栈空（解决完）或者能力不足。

```c++
class Solution {
public:
    vector<int> dailyTemperatures(vector<int>& temperatures) {
        // 获取数组大小，并初始化结果数组以及栈
        int n = temperatures.size();
        vector<int> result(n, 0);
        stack<int> st;
        
        // 遍历数组
        for(int i = 0; i < n; i++){
            // 尝试用该天数的温度为栈里的天数找到升温等待温度
            while(!st.empty() && temperatures[i] > temperatures[st.top()]){
                result[st.top()] = i - st.top();
                st.pop();
            }
            // 该天数入栈
            st.push(i);
        }

        // 栈里残留的就是找不到升温等待天数的，不管它们（有初值0）直接返回结果
        return result;
    }
};
```

## 226 [翻转二叉树](https://leetcode.com/problems/invert-binary-tree/description/?envType=problem-list-v2&envId=2cktkvj)

给定一颗二叉树，将其翻转并返回根节点。（翻转：交换每个节点的左孩子节点与右孩子节点）
<img src="http://eikoh-itx:9000/blog-eikoh/typora/2024/09/03/image-20240713165406040.png" alt="image-20240713165406040" style="zoom:50%;" />

解题思路：考虑递归，用前序遍历的方法操作每个节点即可。

解法一：没啥可说的

```c++
class Solution {
public:
    TreeNode* invertTree(TreeNode* root) {
        // 处理root为空的情况
        if(!root) return root;
        // 先翻转该节点
        TreeNode* tmp = root->left;
        root->left = root->right;
        root->right = tmp;
        // 再翻转两子节点
        invertTree(root->left);
        invertTree(root->right);
        // 返回根节点
        return root;
    }
};
```

## 221 [最大正方形](https://leetcode.com/problems/maximal-square/?envType=problem-list-v2&envId=2cktkvj)

给定m*n的01矩阵，找到其中最大的全由1组成的正方形并返回其面积
![img](https://assets.leetcode.com/uploads/2020/11/26/max1grid.jpg)

如上图，最大的全1正方形由红色和绿色框出，它们的面积为4（2*2）

解题思路：典型的动态规划问题，子问题是从左上角开始，长为i、宽为j的矩形在(i,j)处可组成的最大全1正方形。
$$
dp[i][j] = min\{dp[i-1][j], dp[i][j-1], dp[i-1][j-1]\} + 1 (matrix[i][j] = '1' \and i \ne0\and j\ne0)\\
dp[i][j] = 0 (matrix[i][j] == '0')
$$
这样，对于一个新规划的元素，判断他的上方、左方、左上方的方块各自能够组成的最大全1正方形的边数，取最小值加1即可完成规划。

解法:

```c++
class Solution {
public:
    int maximalSquare(vector<vector<char>>& matrix) {
        // 获取长宽
        int rows = matrix.size();
        int cols = matrix[0].size();
        // 定义dp向量以及最长全1正方形的边
        vector<vector<int>> dp(rows, vector<int>(cols, 0));
        int maxside = 0;

        // 开始动态规划，遍历matrix矩形
        for(int i = 0; i < rows; i++){
            for(int j = 0;j < cols; j++){
                if(matrix[i][j] == '1'){
                    if (i == 0 || j == 0) {
                        dp[i][j] = 1;
                    } else {
                        dp[i][j] = min({dp[i-1][j], dp[i][j-1], dp[i-1][j-1]}) + 1;
                    }
                    // 判断是否更新maxside
                    maxside = (dp[i][j] > maxside ? dp[i][j] : maxside);
                }
            }
        }
        // 返回最大面积
        return maxside * maxside;
    }
};
```

## 215 [数组中第k个最大元素](https://leetcode.com/problems/kth-largest-element-in-an-array/description/?envType=problem-list-v2&envId=2cktkvj)

给定一个乱序数组以及一个数字k,找出这个数组中第k大的元素，返回这个元素的值

解题思路：

1. 借助2分法，每次选择第一个元素为基准值，统计比其小和比其大的元素个数，然后判断条件返回该元素还是进入比其小的那堆亦或是进入大的那堆（试过了，最后一个测试用例超时)
2. 快速选择，以第一个元素为基准值，但是快速处理与基准值相同的元素，这个新的算法尚未消化完毕

解法一：

```c++
	/* 朴素的快排二分思想 */
    int binaryFind(vector<int>& nums, int start, int end, int count, int k) {
        int standard = nums[start];
        int front = start;
        int tail = end;
        int bigger = 0;
        while (front < tail) {
            while (nums[tail] >= standard && front < tail) {
                tail--;
                bigger++;
            }
            switchNum(nums, front, tail);
            while (nums[front] <= standard && front < tail)
                front++;
            switchNum(nums, front, tail);
        }
        if (count+bigger > k - 1) {
            return binaryFind(nums, front+1, end, count, k);
        } else if(count+bigger == k - 1) {
            return nums[front];
        } else {
            return binaryFind(nums, start, front-1, count+bigger+1, k);
        }
    } 
    int findKthLargest(vector<int>& nums, int k) {
        return binaryFind(nums, 0, nums.size()-1, 0, k);
    }
```

解法二：

```c++
	/* 快速选择 */
    int quickSelect(vector<int>& nums, int left, int right, int k) {
        if (left == right)
            return nums[k];
        int front = left - 1, end = right + 1;
        int standard = nums[left];
        while (front < end) {
            do front++; while(nums[front] < standard);
            do end--; while(nums[end] > standard);
            if (front < end) {
                switchNum(nums, front, end);
            }
        }
        if(k <= end) return quickSelect(nums, left, end, k);
        else return quickSelect(nums, end+1, right, k);
    }
    int findKthLargest(vector<int>& nums, int k) {
        int n = nums.size();
        return quickSelect(nums, 0, n-1, n-k);
    }
```

## 208 [实现Trie前缀树](https://leetcode.com/problems/implement-trie-prefix-tree/?envType=problem-list-v2&envId=2cktkvj&)

即实现Trie类，能够完成前缀树的初始化、插入、查询。

前缀树：<img src="http://eikoh-itx:9000/blog-eikoh/typora/2024/09/03/image-20240714172832894.png" alt="image-20240714172832894" style="zoom:67%;" />

解题思路：将字符串里的字母编号，各自对应一条路径，在插入和搜索时对字符串的每个字符进行寻路即可，光做题不太难。

解法一:

```c++
class Trie {
private:
    vector<Trie*> children; // 每个节点的子节点
    bool is_end; // 标识这个节点是否在此终止并构成一个word

    Trie* searchPrefix(string prefix) {
        Trie* node = this;
        for (char ch : prefix) {
            ch -= 'a';
            if (node->children[ch] == nullptr) {
                return nullptr;
            }
            node = node->children[ch];
        }
        return node;
    }

public:
    /* 初始化 */
    Trie() : children(26), is_end(false)  { }
    /* 插入 */
    void insert(string word) {
        Trie* node = this;
        for (char ch : word) {
            ch -= 'a'; // 获取序号
            if (node->children[ch] == nullptr) {
                node->children[ch] = new Trie();
            }
            node = node->children[ch];
        }
        node->is_end = true;
    }
    /* 搜索 */
    bool search(string word) {
        Trie* node = this->searchPrefix(word);
        return node != nullptr && node->is_end;
    }
    /* 查找前缀 */
    bool startsWith(string prefix) {
        return this->searchPrefix(prefix) != nullptr;
    }
};
```

## 207 [课程表](https://leetcode.com/problems/course-schedule/description/?envType=problem-list-v2&envId=2cktkvj&)

需要选n门课程，课程之间存在依赖关系。给定需选课程门树以及课程之间的依赖关系（k*2的矩阵，有k个依赖关系，每个依赖关系中后者依赖前者）。判断是否能够完成n门课程的学习。

解题思路：这是图论的问题，判断图中有无环即可。也即拓扑排序问题。

解法一：每门课程作为一个节点，遍历依赖关系即可获得一张图。用深度搜索的方式搜索这张图，并在搜索的过程中求解拓扑排序。这里求解的方式是置状态位，初始为0表示没有被搜索过，一轮深度搜索中访问过的所有节点置为1,在某一节点的深搜过程中访问到了它自身就说明有环，深度搜索完成后将状态位置为2表示被搜索过而且符合拓扑结构。

```cpp
class Solution {
private:
    vector<vector<int>> edges; // 将依赖关系作为边构造的图
    vector<int> visited; // 节点是否被搜索的状态，0：未被访问，1：这一轮被访问中，2：上一轮被访问
    bool valid = true; // 是否满足拓扑结构，当为false时直接返回
public:
    /* 实现一个图的深度搜索 */
    void dfs(int node) {
        // 标记该节点正在被访问
        visited[node] = 1;
        // 遍历该节点的每条边来深搜
        for (int next: edges[node]) {
            if (visited[next] == 0) {
                // 这条边没被搜过,通过这条边深搜下一个节点
                dfs(next);
                if (!valid) {
                    return ;// 当发现有环时，直接退出
                }
            } else if (visited[next] == 1) {
                // 通过一条边进入这轮搜索访问过的节点，说明发现了环，直接退出
                valid = false;
                return ;
            }
        }
        // 该节点的一轮深搜结束,状态置为2
        visited[node] = 2;
    }
    bool canFinish(int numCourses, vector<vector<int>>& prerequisites) {
        // 先根据依赖关系作图
        edges.resize(numCourses);
        for(const auto& info : prerequisites) {
            // info是一个依赖关系，后者依赖前者
            edges[info[1]].push_back(info[0]);
        }
        // 然后根据图启动dps
        visited.resize(numCourses);
        for (int i = 0;i < numCourses && valid;i++) {
            if(!visited[i]){
                dfs(i);
            }
        }
        return valid;
    }
};
```



## 206 [反转链表](https://leetcode.com/problems/reverse-linked-list/description/?envType=problem-list-v2&envId=2cktkvj&)

给定一个链表，转置后返回

解题思路：转置就行，没啥可说的。

解法一：注意细节，变量定义并赋初值后、开始转置前要将head的next置为`nullptr`。另外要记得更新3个变量的值。

```c++
class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        if(!head || !(head->next)) {
            return head;
        }
        ListNode *prev = head, *node = head->next, *temp;
        prev->next = nullptr;
        while (node) {
            temp = node->next;
            node->next = prev;
            prev = node;
            node = temp;
        }
        return prev;
    }
};
```

## 200 [岛屿数量](https://leetcode.com/problems/number-of-islands/description/?envType=problem-list-v2&envId=2cktkvj&)

给定由‘1’（陆地）和‘0’（水）组成的二维矩阵，计算矩阵网格中的岛屿数量（岛屿被水包围，并且每座岛屿只能由水平方向/或竖直方向上相邻的陆地连接而成）。可以假设网格的四条边全部被水包围

解题思路：深度搜索问题，只要遇到1就启动dps并且把遇到的1都变成0，这样启动dps的数量就是岛屿的数量

解法一：注意边界问题，深度搜索图时列数和行数需要**大于等于**0并**小于**size()。另外四个方向的深度搜索可以通过if进行，8个方向的深度搜索可以通过两个for循环进行。

```cpp
class Solution {
public:
    /* 实现一个深度搜索算法 */
    void dps(vector<vector<char>>& grid, int m, int n) {
        int rows = grid.size();
        int cols = grid[0].size();
        // 将该陆块置为0
        grid[m][n] = '0';

        // 递归搜索上下左右四个方向
        if(m - 1 >= 0 && grid[m-1][n] == '1') dps(grid, m-1, n);
        if(m + 1 < rows && grid[m+1][n] == '1') dps(grid, m+1, n);
        if(n - 1 >= 0 && grid[m][n-1] == '1') dps(grid, m, n-1);
        if(n + 1 < cols && grid[m][n+1] == '1') dps(grid, m, n+1);
        return ;
    }
    int numIslands(vector<vector<char>>& grid) {
        int result = 0;
        // 遍历二维数组，遇到1就启动深搜，深搜一次结果加1
        for(int i = 0;i < grid.size(); i++) {
            for(int j = 0;j < grid[0].size();j++) {
                if(grid[i][j] == '1'){
                    result++;
                    dps(grid, i, j);
                }
            }
        }
        return result;
    }
};
```

## 198 [打家劫舍](https://leetcode.com/problems/house-robber/description/?envType=problem-list-v2&envId=2cktkvj&)

给定一个存放房屋内金额的数组，不能连续取走相邻的两个房屋（数组元素）内的财物。计算并返回在规则限定下能够拾取的最大金额

解题思路：典型的动态规划问题，对于偷n间房间的子问题是$max\{偷第n间+偷前n-2间, 偷前n-1间\}$，即
$$
dp[i] = max\{dp[i-1], dp[i-2]+nums[i]\} (i>=2)\\
dp[i] = nums[0] (i = 1)\\
dp[i] = max\{dp[0], dp[1]\}(i=2)
$$
由此递推关系可以从容的实现解法：

```c++
class Solution {
public:
    int rob(vector<int>& nums) {
        int n = nums.size();
        // 处理特殊情况
        if(n == 1) return nums[0];
        if(n == 2) return max({nums[0], nums[1]});
        // 初始化dp数组
        vector<int> dp(n, 0);
        dp[0] = nums[0];
        dp[1] = max({nums[0], nums[1]});
        // 动态规划求解
        for(int i = 2;i < n;i++)
            dp[i] = max({dp[i-1], dp[i-2] + nums[i]});
        // 返回结果
        return dp[n-1];
    }
};
```

由于在dp求解时固定只会用到两个数组的元素，因此可以有节约空间至O(n)的解法如下：

```cpp
    class Solution {
    public:
        /**
         * 优化空间的动态规划 
         * */
        int rob(vector<int>& nums) {
            int n = nums.size();
            // 处理特殊情况
            if(n == 1) return nums[0];
            // 初始化dp数组
            int dp[2];
            dp[0] = nums[0];
            dp[1] = max({nums[0], nums[1]});
            // 动态规划求解
            for(int i = 2;i < n;i++)
                dp[i%2] = max({dp[(i+1)%2], dp[i%2] + nums[i]});
            // 返回结果
            return dp[(n-1)%2];
        }
    };
```

## 169 [多数元素](https://leetcode.com/problems/majority-element/description/?envType=problem-list-v2&envId=2cktkvj&)

给定一个数组，要求返回其中的多数元素（出现次数大于$\lfloor n/2 \rfloor$的元素）。假设多数元素总是存在，最好在线性时间和O(1)空间内实现。

解题思路：	

1. 将数组排序后取第n/2个元素，由于多数元素共有n/2个，因此第第n/2个元素一定是我们需要求的数。时间复杂度为$O(nlogn)$，手写堆排序的话空间复杂度为$O(1)$
2. 哈希统计法，用一个大小为N的哈希表，在O(n)时间内统计数组内各元素的数量。有空间复杂度和时间复杂度均为O(n)
3. 摩尔投票法：学习中

解法一：用自带的unordered_map求解，即键值为数组内元素，value值是出现次数

```cpp
class Solution {
public:
    int majorityElement(vector<int>& nums) {
        // 建立无序图
        unordered_map<int, int> counts;
        int majority = 0, cnt = 0;
        // 遍历数组
        for (int num : nums) {
            // 数组内元素出现次数加1
            ++counts[num];
            // 更新出现次数最大的元素
            if (counts[num] > cnt) {
                majority = num;
                cnt = counts[num];
            }
        }
        return majority;
    }
};
```

## 238 [除自身以外数组的乘积](https://leetcode.com/problems/product-of-array-except-self/description/?envType=problem-list-v2&envId=2cktkvj&)

给定一个整型数组，返回一个大小相同的数组，结果数组的result[i]是给定数组除去nums[i]外的元素的乘积。需要在线性时间内，而且不允许出现除法

解题思路：

1. 考虑用两个数组分别存储元素i左侧所有数的乘积和右侧所有数的乘积，这将在O(2n)时间完成，最后再用O(n)时间将L[i]*R[i]即可获得我们的结果数组。

解法一：

```cpp
class Solution {
public:
    /* 双数组左右乘积法 */
    vector<int> productExceptSelf(vector<int>& nums) {
        int n = nums.size();
        /* 完成左侧的乘积 */
        vector<int> L(n, 0);
        L[0] = 1;
        for (int i = 1;i < n;i++) {
            L[i] = L[i - 1] * nums[i - 1];
        }
        /* 完成右侧的乘积 */
        vector<int> R(n, 0);
        R[n-1] = 1;
        for (int i = n - 2;i >= 0;i--) {
            R[i] = R[i + 1] * nums[i + 1];
        }
        /* 完成结果数组 */
        vector<int> Result(n, 0);
        for (int i = 0;i < n;i++){
            Result[i] = L[i] * R[i];
        }
        return Result;
    }
};
```

解法二：考虑优化空间，注意到L、R两个数组互不影响，而结果数组在最后才被填入结果。那么我们让L、R两个数组利用结果数组的空间即可将空间复杂度降至O(1)。

```cpp
class Solution {
public:
    /* 优化空间后的左右乘积 */
    vector<int> productExceptSelf(vector<int>& nums) {
        int n = nums.size();
        // 初始化结果数组
        vector<int> Result(n, 0);
        // 计算左边部分的乘积
        Result[0] = 1;
        for (int i = 1;i < n;i++) {
            Result[i] = Result[i-1] * nums[i-1];
        }
        // 计算右边部分的乘积并填入计算结果填入数组
        int right = 1;
        for (int i = n-2;i >= 0;i--) {
            right *= nums[i+1];
            Result[i] *= right;
        }
        return Result;
    }
};
```

## 155 [最小栈](https://leetcode.com/problems/min-stack/description/?envType=problem-list-v2&envId=2cktkvj&)

设计一个支持push、pop、top以及在常数时间内检索到最小元素的最小栈。

解题思路：要在常数时间内检索到最小元素，意味着要保存最小元素，不妨在每次push操作时维护一个向量，如果push进来的数将最小数更新，那么向量就多一个条目存储这个最小数，并且在pop操作时适时删掉这个条目。

解法一：用stack来实现最小栈，即分为普通栈部分和最小数的栈部分，检索最小元素的时间是O(1)但是空间复杂度是O(n)

```cpp
class MinStack {
private:
    stack<int> normal;
    stack<int> min_num;
public:
    MinStack() {
        min_num.push(__INT_MAX__);
    }
    
    void push(int val) {
        normal.push(val);
        min_num.push(min({val, min_num.top()}));
    }
    
    void pop() {
        normal.pop();
        min_num.pop();
    }
    
    int top() {
       return normal.top();
    }
    
    int getMin() {
        return min_num.top();
    }
};
```

解法二：考虑优化空间为O(1)，暂未实现

## 152 [乘积最大子数组](https://leetcode.com/problems/maximum-product-subarray/description/?envType=problem-list-v2&envId=2cktkvj&)

给定一个数组，返回最大的连续非空子数组乘积（这个子数组里所有元素相乘的结果）

解题思路：一眼动态规划，与那道经典的子数组最大和相当类似。不过需要考虑到乘以负数的情况，因此需要更新两个数，分别是最大值和最小值，在遇到乘以负数的情况时交换两者。其余情况类似子数组最大和。

解法一：

```cpp
class Solution {
public:
    int maxProduct(vector<int>& nums) {
        int n = nums.size();
        double max_num = num	s[0], min_num = nums[0];
        double Result = nums[0];
        for (int i = 1;i < n;i++) {
            if (nums[i] < 0) {
                // 遇到负数时，交换最大、最小值
                int tmp = max_num;
                max_num = min_num;
                min_num = tmp;
            }
            // 子问题求解
            max_num = (max_num * nums[i] > nums[i] ? max_num * nums[i] : nums[i]);
            min_num = (min_num * nums[i] < nums[i] ? min_num * nums[i] : nums[i]);
            // 尝试更新最大值
            Result = (max_num > Result ? max_num : Result);
        }
        return Result;
    }
};
```

## 148 [排序链表](https://leetcode.com/problems/sort-list/description/?envType=problem-list-v2&envId=2cktkvj&)

给定一个链表，返回升序排序后的结果链表。

解题思路：从大一写到现在，用归并排序可以实现。

解法一：归并排序与链表的相性相当好，可以在O(nlogn)的时间和O(1)的空间内轻松实现。

```cpp
class Solution {
public:
    /* 实现归并排序的Merge */
    ListNode* mergeList (ListNode* list1, ListNode* list2) {
        // 初始化合并后的链表头
        ListNode* head;
        if (list1->val < list2->val) {
            head = list1;
            list1 = list1->next;
        } else {
            head = list2;
            list2 = list2->next;
        }
        // 启动合并
        ListNode* node = head;
        while (list1 && list2) {
            if (list1->val < list2->val) {
                node->next = list1;
                node = list1;
                list1 = list1->next;
            } else {
                node->next = list2;
                node = list2;
                list2 = list2->next;
            }
        }
        // 处理多出的部分
        if (list1) {
            node->next = list1;
        } else {
            node->next = list2;
        }
        // 返回合并后的链表头
        return head;
    }
    /* 实现归并排序的拆分 */
    ListNode* sortList(ListNode* head) {
        // 处理特殊情况
        if (!head || !head->next) {
            return head;
        }
        // 定义快慢节点
        ListNode* fast = head, *slow = head;
        ListNode* preSlow;
        // 找到链表中点
        while (fast && fast->next) {
            preSlow = slow;
            slow = slow->next;
            fast = fast->next->next;
        }
        // 拆分链表
        preSlow->next = nullptr;
        head = sortList(head);
        slow = sortList(slow);
        return this->mergeList(head, slow);
    }
};
```

## 146 [LRU缓存](https://leetcode.com/problems/lru-cache/description/?envType=problem-list-v2&envId=2cktkvj&)

相当经典的一道题目，需要设计并实现一个满足LRU（最近最少使用）缓存约束的数据结构。即实现一个`LRUCache`类。

解题思路：因为需要保存key值 + value，因此存储缓存的结构应该是哈希表。另外，由于要在O(1)时间内进行`get`和`put`操作，因此需要实现一个双向链表来进行头插和尾删。

解法一：

```cpp
/* 定义LRU缓存的结构 */
struct DLinkNode {
    int key, value;
    DLinkNode* prev;
    DLinkNode* next;
    DLinkNode() : key(0), value(0), prev(nullptr), next(nullptr) {}
    DLinkNode(int _key, int _value) : key(_key), value(_value), prev(nullptr), next(nullptr) {}
};
class LRUCache {
private:
    unordered_map<int, DLinkNode*> cache; // Hash表结构，用来对应键值-节点
    DLinkNode* head; // 双向链表的头节点
    DLinkNode* tail; //双向链表的尾节点
    int size; // 存储目前的缓存占用大小
    int capacity; // 存储缓存的最大容量

    // 将某一节点添加至双向链表的头部
    void addToHead (DLinkNode* node) {
        head->next->prev = node;
        node->next = head->next;
        head->next = node;
        node->prev = head;
    }

    // 将某一节点从双向链表中删除
    void removeNode (DLinkNode* node) {
        node->prev->next = node->next;
        node->next->prev = node->prev;
    }

    // 将某一节点移动至双向链表的头部
    void moveToHead (DLinkNode* node) {
        // 分为两步，先从链表中删除，再添加至头部
        removeNode(node);
        addToHead(node);
    }

    // 缓存大小将要超过容量时，删除尾部的节点
    void deleteTailNode() {
        DLinkNode* node = tail->prev;
        // 链表中删除节点
        removeNode(node);
        // 清空哈希表中的对应项
        cache.erase(node->key);
        // 删除节点内存，防止内存泄漏
        delete(node);
        // 缓存大小减一
        size--;
    }
public:
    // 初始化LRU结构
    LRUCache(int _capacity) : capacity(_capacity), size(0) {
        // 初始化双向链表
        head = new DLinkNode();
        tail = new DLinkNode();
        head->next = tail;
        tail->prev = head;
    }
    
    // 取数的操作
    int get(int key) {
        // 查看是否在缓存中
        if (!cache.count(key)) {
            // 不在，返回-1
            return -1;
        }
        // 在，取出节点将其放置在头部并返回值
        DLinkNode* node = cache[key];
        moveToHead(node);
        return node->value;
    }
    
    void put(int key, int value) {
        // 检查节点是否存在
        if (!cache.count(key)) {
            // 不存在，新建节点
            DLinkNode* node = new DLinkNode(key, value);
            // 添加至哈希表
            cache[key] = node;
            // 插入至链表头部
            addToHead(node);
            // 判断是否超过容量，超的话移除尾部节点
            if (++size > capacity) {
                deleteTailNode();
            }
        } else {
            // 节点已存在，修改值并放到链表头部
            DLinkNode* node = cache[key];
            node->value = value;
            moveToHead(node);
        }
    }
};
```

## 142 [环形链表（加强版）](https://leetcode.com/problems/linked-list-cycle-ii/description/?envType=problem-list-v2&envId=2cktkvj&)

给定一个链表，检测链表中有无环的存在。有则返回顺序遍历链表时环第一次出现的节点，无则返回nullptr。不允许修改原有链表，尽量在O(n)时间和O(1)空间内完成。

解题思路：只要遍历过程中发现一个节点的next指针指向之前访问过的节点即可找到环的起点，问题在于如何保留之前访问过的节点。

1. 使用哈希表可以在O(n)时间和O(n)空间内找到环
2. 使用快慢节点的方式，快节点遇到null说明无环，快慢节点相遇说明有环。

解法一：哈希表的实现

```cpp
class Solution {
public:
    ListNode *detectCycle(ListNode *head) {
        unordered_set<ListNode* > hashMap;
        while (head) {
            if (hashMap.count(head)) {
                // 遇到过的节点，返回之
                return head;
            }
            // 没遇到过的节点，插入哈希表
            hashMap.insert(head);
            head = head->next;
        }
        return nullptr;
    }
};
```

解法二：快慢指针，问题的关键点在于找到第一次出现环的节点，解析如下：

重画链表如下所示，线上有若干个节点。记蓝色慢指针为 slow，红色快指针为 fast。初始时 slow 和 fast 均在头节点处。<img src="http://eikoh-itx:9000/blog-eikoh/typora/2024/09/03/1715514553-RxQrzr-0208_2.png" alt="0208_2.png" style="zoom: 33%;" />
使 slow 和 fast 同时前进，fast 的速度是 slow 的两倍。当 slow 抵达环的入口处时，fast 一定在环上，如下所示。<img src="http://eikoh-itx:9000/blog-eikoh/typora/2024/09/03/1715514558-mCJsmw-0208_3.png" alt="0208_3.png" style="zoom: 33%;" />
其中：

- head 和 A 的距离为 *z*
- 弧 AB (沿箭头方向) 的长度为 *x*
- 同理，弧 BA 的长度为 *y*

可得：

- slow 走过的步数为 *z*
- 设 fast 已经走过了 *k* 个环，*k*≥0，对应的步数为 *z*+*k*(*x*+*y*)+*x*

以上两个步数中，后者为前者的两倍，即 2*z*=*z*+*k*(*x*+*y*)+*x* 化简可得 *z*=*x*+*k*(*x*+*y*)，替换如下所示。

<img src="http://eikoh-itx:9000/blog-eikoh/typora/2024/09/03/1715514562-KmTrNr-0208_4.png" alt="0208_4.png" style="zoom:33%;" />
此时因为 fast 比 slow 快 1 个单位的速度，且*y* 为整数，所以再经过 *y* 个单位的时间即可追上 slow。

即 slow 再走 *y* 步，fast 再走 2*y* 步。设相遇在 C 点，位置如下所示，可得弧 AC 长度为 *y*。<img src="http://eikoh-itx:9000/blog-eikoh/typora/2024/09/03/1715514566-cEsEBC-0208_5.png" alt="0208_5.png" style="zoom:33%;" />
因为此前*x*+*y* 为环长，所以弧 CA 的长度为 *x*。 此时我们另用一橙色指针 ptr (pointer) 指向 head，如下所示。并使 ptr 和 slow 保持 1 个单位的速度前进，在经过 *z*=*x*+*k*(*x*+*y*) 步后，可在 A 处相遇。<img src="http://eikoh-itx:9000/blog-eikoh/typora/2024/09/03/1715514569-ATwmZT-0208_6.png" alt="0208_6.png" style="zoom:33%;" />
再考虑链表无环的情况，fast 在追到 slow 之前就会指向空节点，退出循环即可。

```cpp
class Solution {
public:
    ListNode *detectCycle(ListNode *head) {
        // 处理特殊情况
        if (!head || head->next == head) return head;
        if (!head->next) return nullptr;
        // 定义快慢节点
        ListNode* slow = head, *fast = head;
        // 快慢节点运行直到他们相遇
        while (fast != nullptr && fast->next != nullptr) {
            slow = slow->next;
            fast = fast->next->next;
            if (slow == fast) break;
        }
        if (slow != fast) {
            // 快慢节点不相同就退出循环，说明无环
            return nullptr;
        } else {
            // 快慢节点相同，说明有环，开始寻找环的起点
            while(head != slow) {
                head = head->next;
                slow = slow->next;
            }
            // head 与 slow 相遇的节点就是起始节点，返回之即可
            return head;
        }
    }
};
```

## 141 [环形链表](https://leetcode.com/problems/linked-list-cycle/description/?envType=problem-list-v2&envId=2cktkvj&)

和上一道相同，只不过只检测链表中有无环的存在，有则返回true，无则返回false.

解题思路：快慢链表秒了，没什么可说的

解法一：

```cpp
class Solution {
public:
    bool hasCycle(ListNode *head) {
        // 处理特殊情况
        if (!head || !(head->next)) return false;
        if (head->next == head) return true;
        // 定义快慢节点
        ListNode* slow = head, *fast = head;
        // 循环直到他们相遇
        while (fast != nullptr && fast->next != nullptr) {
            slow = slow->next;
            fast = fast->next->next;
            if (slow == fast) {
                return true;
            }
        }
        return false;
    }
};
```

## 139 [单词拆分](https://leetcode.com/problems/word-break/description/?envType=problem-list-v2&envId=2cktkvj&)

给定一个字符串和一个字符串数组，其中数组对应的是词典，尝试将字符串拆分为多个词典中的单词，拆分成功返回true,失败返回false。

解题思路：

1. 使用之前实现的Trie树对字符串进行分词
2. 动态规划，问题的子问题是$dp[i] = dp[j] \and check(s[j...i-1])$，边界条件有$dp[0] = true$

解法一：用哈希表进行check操作（即检查某一字串是否是单词），有$O(n^2)$时间和$O(n)$空间

```cpp
class Solution {
public:
    bool wordBreak(string s, vector<string>& wordDict) {
        // 初始化哈希表
        unordered_set<string> hashMap;
        for (auto word : wordDict) {
            hashMap.insert(word);
        }
        // 初始化动态规划的数组
        int n = s.size();
        vector<bool> dp(n + 1, false);
        dp[0] = true;
        // 启动动态规划
        for (int i = 1;i <= n;i++) {
            for (int j = 0;j < i;j++) {
                // 循环判断是否能拆分，只要有一种拆分方式就退出循环
                if (dp[j] && hashMap.count(s.substr(j, i-j))) {
                    dp[i] = true;
                    break;
                }
            }
        }
        return dp[n];
    }
};		
```

## 136 [只出现一次的数字](https://leetcode.com/problems/single-number/description/?envType=problem-list-v2&envId=2cktkvj&)

给定非空整型数组，其中只有一个仅出现一次的数字，其余每个元素都出现了两次，找出那个只出现了一次的元素

解题思路：

1. 用哈希表，当有元素出现两次时将其删除，最后留下来的那个元素就是只出现一次的元素。
2. 位运算，因为除要求的数字外都出现了两次，因此将数组的每个数一起做一遍异或运算，最后的结果一定是那个只出现一次的数字

解法一：

```cpp
class Solution {
public:
    int singleNumber(vector<int>& nums) {
        int Result = 0;
        for (int num: nums)
            Result ^= num;
        return Result;
    }
};
```

## 647 [回文子串](https://leetcode.com/problems/palindromic-substrings/description/?envType=problem-list-v2&envId=2cktkvj&)

给你一个字符串 `s` ，请你统计并返回这个字符串中 **回文子串** 的数目。**回文字符串** 是正着读和倒过来读一样的字符串。**子字符串** 是字符串中的由连续字符组成的一个序列。

解题思路：

1. 将字符串转置，然后问题转化为求两个字符串的公共子串数目，这可以用动态规划法解决。
2. 中心拓展，即枚举每一个可能的回文中心，然后用两个指针分别向左右两边拓展，当两个指针指向两个相同的元素时就拓展，不同时就停止拓展。

解法一：中心拓展，需要有序地枚举所有回文中心，即需要考虑奇数长度的回文序列和偶数长度的回文序列。奇数时，中心是一个元素；偶数时，中心是两个元素。可以用巧妙的办法将这两种情况统一在一个循环里完成。

```cpp
class Solution {
public:
    /* 中心拓展法求解 */
    int countSubstrings(string s) {
        int Result = 0, n = s.length();
        // 讨论奇数长度的回文串
        for (int i = 0;i < n;i++) {
            int l = i, r = i;
            while(l >= 0 && r < n && s[l] == s[r]){
                l--;
                r++;
                Result++;
            }
        }
        // 讨论偶数长度的回文串
        for (int i = 0;i < n - 1; i++) {
            int l = i, r = i + 1;
            while(l >= 0 && r < n && s[l] == s[r]) {
                l--;
                r++;
                Result++;
            }
        }
        return Result;
    }
    /* 中心拓展法求解 */
    int countSubstrings(string s) {
        int Result = 0, n = s.length();
        // 合并讨论的版本
        for (int i = 0;i < 2*n - 1;i++) {
            int l = i / 2, r = i/2 + i%2;
            while (l >= 0 && r < n && s[l] == s[r]) {
                l--;
                r++;
                Result++;
            }
        }
        return Result;
    }
};
```

## 128 [最长连续子序列](https://leetcode.com/problems/longest-consecutive-sequence/description/?envType=problem-list-v2&envId=2cktkvj&)

给定一个未排序的整数数组 `nums` ，找出**数字连续**的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 `O(n)` 的算法解决此问题。

解题思路：朴素的想法是对每一个数，不断地查找它的下一个数、上一个数、下一个数的下一个数、上一个数的上一个数，最后可以求得这个数所在的连续子序列长度。

1. 利用哈希表可以用O(n)空间将查找上一个/下一个数的时间降至O(1)，只需要将被查找过的数标记就能够做到O(n)时间内的连续最长序列长度。

解法一：巧妙的点在于查找连续序列时，从最小的数开始查起，用一个`!hashMap.count(num - 1)`就能知道这个数字是不是连续序列的最小数字，从而简化了空间和循环时间

```cpp
class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        int Result = 0, n = nums.size();
        // 建立哈希表
        unordered_set<int> hashMap;
        // 插入哈希表
        for (int num : nums) {
            hashMap.insert(num);
        }
        // 查找连续序列，直接用hashMap里的数（相当于去了重）
        for (int num : hashMap) {
            if (!hashMap.count(num - 1)) {
                // 确定是连续序列的最小数，开始统计连续序列长度
                int length = 1, tmp = num+1;
                while (hashMap.count(tmp)) {
                    length++;
                    tmp++;
                }
                Result = (length > Result ? length : Result);
            }
        }
        return Result;
    }
};
```

## 124 [二叉树中的最大路径和](https://leetcode.com/problems/binary-tree-maximum-path-sum/description/?envType=problem-list-v2&envId=2cktkvj&)

二叉树中的 **路径** 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 **至多出现一次** 。该路径 **至少包含一个** 节点，且不一定经过根节点。
**路径和** 是路径中各节点值的总和。
给你一个二叉树的根节点 `root` ，返回其 **最大路径和** 。

解题思路：看到二叉树首先考虑递归。首先维护每个节点的最大贡献值：叶结点的最大贡献值是它本身的`value`，非叶节点的最大贡献值是`node->val + max{left.maxGain(), right.maxGain()}`。然后根据最大贡献值能够进一步维护最大路径和：对于每一个节点，最大的路径和是`node->val + left.maxGain() > 0 ? left.maxGain() : 0 + right.maxGain() > 0 ? right.maxGain() : 0 `	。在递归的过程中不断维护这两个值，并且更新最大的最大路径和即可获得答案。

解法一：

```cpp
class Solution {
private:
    int Result = 0;
public:
    /* 定义一个递归函数，边递归边计算最大路径和及最大贡献值 */
    int computeNode(TreeNode *node) {
        // 处理特殊情况
        if (!node) return -1;
        // 获取左、右子节点的最大贡献值
        int left = (node->left != nullptr ? computeNode(node->left) : 0);
        int right = (node->right != nullptr ? computeNode(node->right) : 0);
        // 尝试更新最大路径和
        int maxSum = node->val + (left > 0 ? left : 0) + (right > 0 ? right : 0);
        Result = (maxSum > Result ? maxSum : Result);
        // 返回最大贡献值
        return node->val + max({left, right, 0});
    }
    int maxPathSum(TreeNode* root) {
        Result = root->val;
        computeNode(root);
        return Result;
    }
};
```

## 322 [零钱兑换](https://leetcode.com/problems/coin-change/description/?envType=problem-list-v2&envId=2cktkvj&)

 给你一个整数数组 `coins` ，表示不同面额的硬币；以及一个整数 `amount` ，表示总金额。
计算并返回可以凑成总金额所需的 **最少的硬币个数** 。如果没有任何一种硬币组合能组成总金额，返回 `-1` 。
你可以认为每种硬币的数量是无限的。

解题思路：给定特殊的硬币面额可以用贪心算法求解，但是由于硬币的面额随机，所以必须放弃贪心。尝试使用更经典的回溯剪枝或者动态规划。这题动态规划的子问题与单词拆分以及其他的字符串相关的动态规划问题很像，令i、j表示需要凑成的总金额，那么有如下递推关系：
$$
dp[i] = min\{\forall coin \in coins, (dp[i-coin] \ge 0 ? dp[i-coin] + 1 : -1)\} (i-coin >= 0 \and 0\lt i \le n)
$$
问题的边界有$dp[0] = 0$，最后的结果即$dp[amount]$。

解法一：时间复杂度为$O(amount * coins.size())$，空间复杂度为$O(amount)$

```cpp
class Solution {
public:
    int coinChange(vector<int>& coins, int amount) {
        // 初始化dp数组
        vector<int> dp(amount+1, -1);
        dp[0] = 0;
        // 启动dp流程
        for (int i = 1;i <= amount;i++) {
            int min = __INT_MAX__;
            for (int coin : coins) {
                if (i - coin >= 0 && dp[i-coin] != -1) {
                    min = (dp[i-coin] + 1 < min ? dp[i-coin] + 1 : min);
                }
            }
            dp[i] = (min == __INT_MAX__ ? -1 : min);
        }
        // 返回结果
        return dp[amount];
    }
};
```

## 494 [目标和](https://leetcode.com/problems/target-sum/description/?envType=problem-list-v2&envId=2cktkvj&)

给你一个非负整数数组 `nums` 和一个整数 `target` 。
向数组中的每个整数前添加 `'+'` 或 `'-'` ，然后串联起所有整数，可以构造一个 **表达式** ：
例如，`nums = [2, 1]` ，可以在 `2` 之前添加 `'+'` ，在 `1` 之前添加 `'-'` ，然后串联起来得到表达式 `"+2-1"` 。
返回可以通过上述方法构造的、运算结果等于 `target` 的不同 **表达式** 的数目。

解题思路：朴素的思想是深度搜索遍历，但时间复杂度将在`2^n`量级。

1. 因此需要考虑回溯剪枝，即计算剩下的层数能提供的最大值和最小值，当前值超出这个范围则不继续往下搜索。
2. ==其实这道题目是一道0-1背包问题，可以用动态规划法求解，暂时没看==。

解法一：深度搜索配合剪枝硬干

```cpp
class Solution {
private:
    int n, Count, Target;
    vector<int>* maxSum;
public:
    /* 实现基本的深度搜索 */
    void dps(vector<int>& nums, int sum, int i) {
        // 边界条件
        if (i == n - 1) {
            if (sum + nums[i] == Target) Count++;
            if (sum - nums[i] == Target) Count++;
            return ;
        } 
        // 当前数加上正号
        sum += nums[i];
        // 剪枝
        /*
        if (sum + (*maxSum)[i] == Target) Count++;
        if (sum - (*maxSum)[i] == Target) Count++;*/
        if (sum + (*maxSum)[i] >= Target && sum - (*maxSum)[i] <= Target) {
            dps(nums, sum, i+1);
        }
        // 当前数加上负号
        sum -= 2*nums[i];
        // 剪枝
        if (sum + (*maxSum)[i] >= Target && sum - (*maxSum)[i] <= Target) {
            dps(nums, sum, i+1);
        }
    }
    int findTargetSumWays(vector<int>& nums, int target) {
        // 初始化深搜和剪枝用到的
        n = nums.size();
        Count = 0;
        Target = target;
        maxSum = new vector<int>(n, 0);
        int sum = 0;
        for (int i = n-2;i >= 0;i--) {
            sum += nums[i+1];
            (*maxSum)[i] = sum;
        }
        dps(nums, 0, 0);
        return Count;
    }
};
```

## 461 [汉明距离](https://leetcode.com/problems/hamming-distance/description/?envType=problem-list-v2&envId=2cktkvj&)

两个整数之间的 [汉明距离](https://baike.baidu.com/item/汉明距离) 指的是这两个数字对应二进制位不同的位置的数目。
你两个整数 `x` 和 `y`，计算并返回它们之间的汉明距离。

解题思路：异或的位运算能够求得一个仅不同的位上有1的数字，统计这个数字里1的数量即可

解法一：妙的点在与用`num & 1`来计算最后一位是否是1

```cpp
class Solution {
public:
    int hammingDistance(int x, int y) {
        int Count = 0;
        // 异或运算求得不同位
        int num = x ^ y;
        // 统计1的数量
        while(num > 0) {
            Count += num & 1;
            num >>= 1;
            /* 可以优化循环次数
            num &= num - 1;
            Count++;*/
        }
        return Count;
    }
};
```

## 448 [找到所有数组中消失的数字](https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/?envType=problem-list-v2&envId=2cktkvj&)

给你一个含 `n` 个整数的数组 `nums` ，其中 `nums[i]` 在区间 `[1, n]` 内。请你找出所有在 `[1, n]` 范围内但没有出现在 `nums` 中的数字，并以数组的形式返回结果。
在不使用额外空间且时间复杂度为 `O(n)` 的情况下解决这个问题,可以假定返回的数组不算在额外空间内。

解题思路：注意这个数组中可能会重复出现区间[1, n]内的数，因此常规做法是用一个哈希表记录所有出现的区间内的数字，然后遍历`1~n`之间的数字，将没有出现在哈希表中的数字放入结果数组中。

解法一: 哈希表方案，有空间复杂度为O(n)，不符合题目要求

```cpp
class Solution {
public:
    vector<int> findDisappearedNumbers(vector<int>& nums) {
        int n = nums.size();
        // 用集合作为哈希表结构
        unordered_set<int> hashMap;
        // 遍历数组，统计所有出现在[1,n]中的数字
        for (int & num : nums) {
            if (1 <= num && n >= num) {
                // 尝试放入哈希表中
                hashMap.insert(num);
            }
        }
        // 遍历[1,n]，将所有没有出现在哈希表的数字放入结果数组中
        vector<int> resultArr;
        for (int i = 1;i <= n;i++){
            if(!hashMap.count(i)) {
                resultArr.push_back(i);
            }
        }
        return resultArr;
    }
};
```

解法二：符合要求的解法，核心思想是用传入的参数nums作为哈希表。即如果nums[i]中的数字是x，那我们让nums[x-1]里的数字大小+n。这样导致的问题是遍历过程中将会遇到被增加过的数字，此时取模即可解决问题。遍历完一遍后再遍历一遍，只要里面的数字小于等于n就说明从未出现过，放入结果数组中即可。**需要注意的是，在这种取模操作和-1同时进行的情况下，需要先-1在取模，避免越界的问题（取模后是[0, n)）**

```cpp
class Solution {
public:
    vector<int> findDisappearedNumbers(vector<int>& nums) {
        // 遍历第一遍，利用nums数组作为哈希表记录出现过的[1,n]
        int n = nums.size();
        for (int i = 0;i < n;i++) {
            int x = (nums[i] - 1) % n; // 先减一再取模，避免数组越界问题
            nums[x] += n;
        }
        // 遍历第二遍，统计没有出现过的数字并返回
        vector<int> result;
        for (int i = 0;i < n;i++) {
            if (nums[i] <= n) { // 判定边界是小于等于，仔细阅读题目中数组里的元素大小范围
                result.push_back(i+1);
            }
        }
        return result;
    }
};
```

## 438 [找到字符串中所有字母异位词](https://leetcode.com/problems/find-all-anagrams-in-a-string/description/?envType=problem-list-v2&envId=2cktkvj)

给定两个字符串 `s` 和 `p`，找到 `s` 中所有 `p` 的 **异位词** 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
**异位词** 指由相同字母重排列形成的字符串（包括相同的字符串）。

解题思路：滑动窗口，在字符串`s`中构造一个长度与`p`相同的滑动窗口，并在滑动中维护窗口里每种字母的数量。当窗口中每种字母的数量与`p`相同时，即找到了一个异位词。

解法一：普通的滑动窗口，用数组来存储滑动窗口里和`p`中每种字母的数量。

```cpp
class Solution {
public:
    vector<int> findAnagrams(string s, string p) {
        int sLen = s.size(), pLen = p.size();
        // 特殊处理
        if (sLen < pLen) {
            return vector<int>();
        }
        // 记录窗口与p字符串中每种字母数量的数组
        vector<int> sCnt(26, 0);
        vector<int> pCnt(26, 0);
        for (int i = 0;i < pLen;i++) {
            ++sCnt[s[i] - 'a'];
            ++pCnt[p[i] - 'a'];
        }
        // 开始滑动窗口统计异构词
        vector<int> result;
        if (sCnt == pCnt) result.push_back(0);
        for (int i = 0;i < sLen - pLen;i++) {
            // 窗口划走的字母
            --sCnt[s[i] - 'a'];
            // 窗口新增的字母
            ++sCnt[s[i + pLen] - 'a'];
            // 判断是否为异构词，是则添加入结果数组中
            if (sCnt == pCnt) result.emplace_back(i+1);
        }

        return result;
    }
};
```

解法二：优化时间后的滑动窗口，维护一个记录窗口与`p`中不同字母数量的变量，在滑动的过程中维护这个变量

```cpp
class Solution {
public:
    vector<int> findAnagrams(string s, string p) {
        int sLen = s.size(), pLen = p.size();
        // 处理特殊情况
        if (sLen < pLen) return vector<int>();
        // 初始化变量differ以及与之联系的统计数组
        vector<int> cnt(26, 0);
        int differ = 0;
        for (int i = 0;i < pLen;i++) {
            ++cnt[s[i] - 'a']; // 窗口里的字母使统计值加一
            --cnt[p[i] - 'a']; // p字符串里的字母使统计值减一
        }
        for (int j = 0;j < 26;j++) {
            // 每当有一个字母的统计值不为0,differ变量就应该加一（不同的字母多一个）
            if (cnt[j] != 0) differ++;
        }
        // 滑动窗口统计并返回结果数组
        vector<int> result;
        if (differ == 0) result.push_back(0);
        for (int i = 0;i < sLen - pLen;i++) {
            // 将窗口头部的字母滑出，处理统计数组
            int tmp = cnt[s[i] - 'a'];
            --cnt[s[i] - 'a']; // 窗口里少一个字母，少了一个使统计值加一的效果
            differ = (cnt[s[i] - 'a'] ? differ : differ - 1); // 滑出字母后两窗口对应字母同步，differ减1
            differ = (tmp ? differ : differ + 1); // 滑出字母前两窗口对应字母同步，differ增加
            // 将下一个字母滑入，处理统计数组
            tmp = cnt[s[i + pLen] - 'a'];
            ++cnt[s[i + pLen] - 'a'];
            differ = (cnt[s[i + pLen] - 'a'] ? differ : differ - 1); // 滑入字母后两窗口对应字母同步，differ减1
            differ = (tmp ? differ : differ + 1); // 滑出字母前两窗口对应字母同步，differ增加
            // 查看differ的值，为0说明当前窗口是一个异构词
            if (differ == 0) result.emplace_back(i + 1);
        }
        return result;
    }
};
```

## 437 [路径总和——其三](https://leetcode.com/problems/path-sum-iii/description/?envType=problem-list-v2&envId=2cktkvj)

给定一个二叉树的根节点 `root` ，和一个整数 `targetSum` ，求该二叉树里节点值之和等于 `targetSum` 的 **路径** 的数目。
**路径** 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

解题思路：用最朴素的方法，采用前序深度优先遍历，针对每个节点都考虑以它为父节点的所有路径是否符合题意；问题在与如何遍历所有可能的路径，可以用两个函数，一个作为子节点往下遍历路径，一个作为父节点获取所有可能的路径。

解法一：朴素的深度优先搜索，时间复杂度为$O(N^2)$

```cpp
class Solution {
public:
    /* 针对每个节点，将其作为子节点尝试构造路径 */
    int rootSum(TreeNode* root, int targetSum) {
        if (!root) {
            // 处理特殊情况
            return 0;
        }
        int ret = 0;
        // 确认当前以当前节点为路径终点是否符合题意
        if (root->val == targetSum) {
            ++ret; // 注意当前节点作为终点符合题意后，后续路径仍有可能符合题意，不能直接返回1
        }
        // 继续向下探索路径
        ret += rootSum(root->left, targetSum - root->val);
        ret += rootSum(root->right, targetSum - root->val);
        // 返回结果
        return ret;
    }
    /* 针对每个节点，将其作为父节点尝试构造路径 */
    int pathSum(TreeNode* root, int targetSum) {
        if (!root) {
            // 处理特殊情况
            return 0;
        }
        // 获取其子路径可能的解
        int retVal = rootSum(root, targetSum);
        // 将子节点作为父节点尝试构造路径
        retVal += pathSum(root->left, targetSum);
        retVal += pathSum(root->right, targetSum);
        // 返回所有的路径结果
        return retVal;
    }
};
```

解法二：优化思路是减少重复的子问题计算，尝试记忆从根节点到其余每个节点的路径之和，之后即可重复利用这个计算结果，以哪个节点为路径的起点就减去从根节点到起点节点的值即可完成计算。关系式为$Sum_{node节点到当前节点} - Sum_{node节点到路径中的某一节点} = Target$，由于我们在遍历过程中一直能够轻易得到的值是$Sum_{node节点到当前节点} 和 Target$，所以我们需要用数据结构存储全部的$Sum_{node节点到路径中的某一节点}$，称其为前缀值。同时在遍历过程中进行的判定是$Sum_{node节点到当前节点} - Target$是否对应了存储的某一前缀值。需注意的问题是在回溯的过程中需要及时更新数据结构中的内容。

```cpp
class Solution {
private:
    // 存储前缀和的数据结构
    unordered_map<long, int> prefix;
public:
    /* 针对每个节点，获取以其作为路径的终点能够构造的路径数量 */
    int endSum(TreeNode* node, long curVal, int targetSum) {
        if (!node) {
            // 处理特殊情况
            return 0;
        }
        // 更新路径和
        curVal += node->val;
        // 获取以该节点为终点可构造出的路径数量
        int retVal = 0;
        if (prefix.count(curVal - targetSum)) {
            retVal = prefix[curVal - targetSum];
        }
        // 继续以子节点为终点尝试构造路径
        // 更新前缀和
        if (prefix.count(curVal)) prefix[curVal]++;
        else prefix[curVal] = 1;
        // 尝试子节点
        retVal += endSum(node->left, curVal, targetSum);
        retVal += endSum(node->right, curVal, targetSum);
        // 清理前缀和
        prefix[curVal]--;
        return retVal; // 返回最终结果
    }

    int pathSum(TreeNode* root, int targetSum) {
        // 初始化前缀和
        prefix[0] = 1; // 代表从root节点到某一子节点的完整路径
        // 深度优先搜索计算结果
        return endSum(root, 0, targetSum);
    }
};
```

## 42 [接雨水](https://leetcode.com/problems/trapping-rain-water)

给定 `n` 个非负整数表示每个宽度为 `1` 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

![image-20241022184701843](https://minio.noteikoh.top:443/blog-eikoh/typora/2024/10/22/image-20241022184701843.png)

解题思路：按列来求，对于当前在求的列，查看它左边最高的墙(记为$L_{highest}$)和右边最高的墙(记为$R_{highest}$)以及它本身墙的高度(记为$Self$)，需要比较的是$min\{L_{highedst}, R_{highest}\}$与$Self$：如果$min\{L_{highedst}, R_{highest}\} \le Self$，那么当前列显然无法储水（水往低处流走了）；如果$min\{L_{highedst}, R_{highest}\} \gt Self$，那么当前列的储水量就是$min\{L_{highedst}, R_{highest}\} - Self$，多于的水从低处的墙流走了。另外考虑边界条件，最左侧的墙与最右侧的墙都无法储水。最后问题就转化为如何快速地求出左侧、右侧最高的墙，可以利用动态规划的方法解决这一问题：定义两个数组，分别存储$i$处墙的$L_{highest}$及$R_{highest}$，需要占用$O(n)$空间以及$O(n)$时间来初始化这两个数组。最终，我们使用$O(n)$空间和$O(n)$时间解决了这一问题

解法一：按列求水+动态规划求高墙

```go
func trap(height []int) int {
	// 定义并求出每一处的左侧高墙与右侧高墙
	n := len(height)
	l_highest := make([]int, n)
	r_highest := make([]int, n)

	// 求出左侧高墙
	l_highest[0] = height[0]
	for i := 1; i < n-; i++ {
		l_highest[i] = max(l_highest[i-1], height[i])
	}

	// 求出右侧高墙
	r_highest[n-1] = height[n-1]
	for i := n - 2; i > 0; i-- {
		r_highest[i] = max(height[i], r_highest[i+1])
	}

	// 遍历 [1, n-1]，求出总的储水量
	var result int
	for i := 1; i < n-1; i++ {
		tmp := min(l_highest[i], r_highest[i]) - height[i]
		if tmp > 0 {
			result += tmp
		}
	}

	return result
}
```

## 72 [编辑距离](https://leetcode.com/problems/edit-distance)

给你两个单词 `word1` 和 `word2`， *请返回将 `word1` 转换成 `word2` 所使用的最少操作数*  。

你可以对一个单词进行如下三种操作：

- 插入一个字符
- 删除一个字符
- 替换一个字符

![image-20241030234534680](https://minio.noteikoh.top:443/blog-eikoh/typora/2024/10/30/image-20241030234534680.png)

解题思路：一道相当难以看出来的动态规划题目，需要用 `dp[i][j]` 代表 `word1` 到 `i` 位置转换成 `word2` 到 `j` 位置需要最少步数
![Snipaste_2019-05-29_15-28-02.png](https://minio.noteikoh.top:443/blog-eikoh/typora/2024/10/30/76574ab7ff2877d63b80a2d4f8496fab3c441065552edc562f62d5809e75e97e-Snipaste_2019-05-29_15-28-02.png)
从而有如下的递推公式:
$$
dp[i][j] = dp[i-1][j-1](word1[i] == word2[j]) \\
dp[i][j] = min\{dp[i-1][j-1], dp[i-1][j], dp[i][j-1]\} + 1(word1[i] != word2[j])
$$
在第二个递推公式中，$dp[i-1][j]$ 代表删除操作，$dp[i][j-1]$代表插入操作，$dp[i-1][j-1]$代表替换操作

最终的 $dp[m][n]$ 即是最少步数

解法一：

```go
func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	dp := make([][]int, (n1 + 1))
	for i := range dp {
		dp[i] = make([]int, (n2 + 1))
	}

	for i := 0; i <= n1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
			}
		}
	}

	return dp[n1][n2]
}
```



# 数组部分

## 二分查找

### 704 [二分查找](https://leetcode.com/problems/binary-search/description/)

给定一个 `n` 个元素有序的（升序）整型数组 `nums` 和一个目标值 `target` ，写一个函数搜索 `nums` 中的 `target`，如果目标值存在返回下标，否则返回 `-1`。

解题思路：常规的二分查找写法写出即可，要注意的是二分查找的边界条件以及一些细节上的问题。

解法一：朴素的二分查找

```c++
class Solution {
public:
    int search(vector<int>& nums, int target) {
        int n = nums.size();
        int begin = 0, end = n - 1;
        while (begin <= end) { // 边界条件一定是大于
            int middle = (end - begin) / 2 + begin; // 细节，避免相加越界
            if (nums[middle] == target) {
                return middle;
            } else if (nums[middle] < target) {
                begin = middle + 1;
            } else {
                end = middle - 1;
            }
        }
        return -1;
    }
};
```

### 69 [x的平方根](https://leetcode.com/problems/sqrtx/description/)

给你一个非负整数 `x` ，计算并返回 `x` 的 **算术平方根** 。
由于返回类型是整数，结果只保留 **整数部分** ，小数部分将被 **舍去 。**
**注意：**不允许使用任何内置指数函数和算符，例如 `pow(x, 0.5)` 或者 `x ** 0.5` 。

解题思路：首先是可以遍历1～x之间的每个数字，确认平方后与x差距最小的整形并返回。但是注意到这里面的数字天然就已经排序，所以可以用二分查找的方式加速查找。

解法一：二分查找

```c++
class Solution {
public:
    int mySqrt(int x) {
        int start = 1, end = x;
        while (start <= end) {
            int middle = (end - start) / 2 + start;
            int tmp = x / middle;
            if (tmp < middle) {
                end = middle - 1;
            } else if (tmp > middle) {
                start = middle + 1;
            } else {
                return middle;
            }
        }
        return end;
    }
};
```

## 移除元素（前后双指针法的学习）

### 27 [移除元素](https://leetcode.com/problems/remove-element/description/)

给你一个数组 `nums` 和一个值 `val`，你需要 **[原地](https://baike.baidu.com/item/原地算法)** 移除所有数值等于 `val` 的元素。元素的顺序可能发生改变。然后返回 `nums` 中与 `val` 不同的元素的数量。

假设 `nums` 中不等于 `val` 的元素数量为 `k`，要通过此题，您需要执行以下操作：

- 更改 `nums` 数组，使 `nums` 的前 `k` 个元素包含不等于 `val` 的元素。`nums` 的其余元素和 `nums` 的大小并不重要。
- 返回 `k`。

解题思路：谨慎读题，可知需要将等于val的元素全部移到数组的末尾，只需要一边遍历一边交换即可

解法一：手动定义数组的长度，遍历并交换。问题在与会改变元素的顺序

```c++
class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        int length = nums.size();
        // 确认当前数组末尾的元素不是val
        while (length > 0 && nums[length - 1] == val) {
            length--;
        }
        for (int i = 0;i < length;i++) {
            if (nums[i] == val) {
                // 确认是否遍历到数组末尾
                if (i == length - 1) {
                    return length - 1;
                }
                // 移除元素
                nums[i] = nums[length - 1];
                nums[length-- - 1] = val;
                // 确认当前数组末尾的元素不是val
                while (length > 0 && nums[length - 1] == val) {
                    length--;
                }
            }
        }
        return length;
    }
};
```

解法二：通过前后指针法在不改变元素位置的情况下进行删除，关键点在于用前面的指针探寻元素、判断是否进行覆盖。fast与slow的距离就是探寻到的目标元素个数，通过将fast内的非目标元素写入slow来实现将目标元素移到数组末尾的效果（最后fast到数组终点时，fast与slow之间的元素就是为目标元素留出的空位）

```c++
class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        // 快慢指针法
        int slow = 0;
        for (int fast = 0; fast < nums.size(); fast++) {
            if (val != nums[fast]) {
                nums[slow++] = nums[fast];
            }
        }
        return slow;
    }
};
```

### 26 [删除有序数组中的重复项](https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/)

给你一个 **非严格递增排列** 的数组 `nums` ，请你**[ 原地](http://baike.baidu.com/item/原地算法)** 删除重复出现的元素，使每个元素 **只出现一次** ，返回删除后数组的新长度。元素的 **相对顺序** 应该保持 **一致** 。然后返回 `nums` 中唯一元素的个数。

考虑 `nums` 的唯一元素的数量为 `k` ，你需要做以下事情确保你的题解可以被通过：

- 更改数组 `nums` ，使 `nums` 的前 `k` 个元素包含唯一元素，并按照它们最初在 `nums` 中出现的顺序排列。`nums` 的其余元素与 `nums` 的大小不重要。
- 返回 `k` 。

解题思路：双指针法，快的指针探出与前一个元素不同的元素并放到慢指针处

解法一：

```c++
class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int slow = 1;
        for (int fast = 1; fast < nums.size(); fast++) {
            if (nums[fast] != nums[fast - 1]) {
                nums[slow++] = nums[fast];
            }
        }
        return slow;
    }
};
```

## 有序数组的平方（头尾双指针法的学习）

### 977 [有序数组的平方](https://leetcode.com/problems/squares-of-a-sorted-array/description/)

给你一个按 **非递减顺序** 排序的整数数组 `nums`，返回 **每个数字的平方** 组成的新数组，要求也按 **非递减顺序** 排序。

解题思路：题目的难点在于原数组会出现负数，负数的平方会搅乱整个数组的排序。但是根据`y = x^2`的二次函数，我们知道x轴上两端的函数值肯定会高于中间，因此可以使用头尾两个指针，比较绝对值的大小来放入。

解法一：头尾双指针

```c++
class Solution {
public:
    vector<int> sortedSquares(vector<int>& nums) {
        int n = nums.size();
        // 定义头尾指针
        int head = 0, tail = n - 1;
        // 初始化结果数组
        vector<int> Result(n, 0);
        // 比较头尾指针并遍历数组
        while (head <= tail) {
            if (abs(nums[head]) >= abs(nums[tail])) {
                Result[--n] = nums[head] * nums[head++];
            } else {
                Result[--n] = nums[tail] * nums[tail--];
            }
        }
        return Result;
    }
};
```

## 长度最小的子数组（滑动窗口的学习）

### 209 [长度最小的子数组](https://leetcode.com/problems/minimum-size-subarray-sum/description/)

给定一个含有 `n` 个正整数的数组和一个正整数 `target` **。**
找出该数组中满足其总和大于等于 `target` 的长度最小的**子数组**$[nums_l, nums_{l+1}, ..., nums_{r-1}, nums_r]$ ，并返回其长度**。**如果不存在符合条件的子数组，返回 `0` 。

解题思路：关键在于不断搜索连续子数组以及子数组的总和。朴素的想法是用两个for循环暴力求解，但是考虑到这道题目要求子数组的总和大于等于`target`并且是连续的，因此可以用两个指针实现滑动窗口在更低的时间复杂度下求解。

解法一：滑动窗口法

```cpp
class Solution {
public:
    // 滑动窗口法求解
    int minSubArrayLen(int target, vector<int>& nums) {
        // 定义窗口相关信息：子数组总和，窗口左边界、右边界
        int windowSum = 0, start = 0, end = 0;
        int n = nums.size();
        // 定义最终结果：最小的窗口长度
        int result = 0;
        while (end < n) {
            // 计算向右拓展的边界
            windowSum += nums[end];
            // 尝试收缩左边界
            while (windowSum >= target) {
                int length = end - start + 1;
                if (!result) result = length;
                else {
                    result = length < result ? length : result;
                }
                windowSum -= nums[start++];
            }
            // 窗口向右拓展
            end++;
        }
        return result;
    }
};
```

### 904 [水果成蓝](https://leetcode.com/problems/fruit-into-baskets/description/)

你正在探访一家农场，农场从左到右种植了一排果树。这些树用一个整数数组 `fruits` 表示，其中 `fruits[i]` 是第 `i` 棵树上的水果 **种类** 。

你想要尽可能多地收集水果。然而，农场的主人设定了一些严格的规矩，你必须按照要求采摘水果：

- 你只有 **两个** 篮子，并且每个篮子只能装 **单一类型** 的水果。每个篮子能够装的水果总量没有限制。
- 你可以选择任意一棵树开始采摘，你必须从 **每棵** 树（包括开始采摘的树）上 **恰好摘一个水果** 。采摘的水果应当符合篮子中的水果类型。每采摘一次，你将会向右移动到下一棵树，并继续采摘。
- 一旦你走到某棵树前，但水果不符合篮子的水果类型，那么就必须停止采摘。

给你一个整数数组 `fruits` ，返回你可以收集的水果的 **最大** 数目。

解题思路：其实就是找出数组中只由两个数字组成的最长连续子数组，可以用滑动窗口很好地解决

解法一：自实现的滑动窗口法

```c++
class Solution {
public:
    int totalFruit(vector<int>& fruits) {
        int start = 0, end = 0;
        int result = 0; // 最长的长度
        int basket1 = -1, basket2 = -1; // 篮子
        // 滑动窗口求解
        while (end < fruits.size()) {
            // 向右拓展边界
            if (basket1 == -1) {
                basket1 = fruits[end];
            } else if (basket1 != fruits[end] && basket2 == -1) {
                basket2 = fruits[end];
            }
            if (basket1 == fruits[end] || basket2 == fruits[end]) {
                int tmp = end - start + 1;
                result = (tmp > result ? tmp : result);
            } else {
                // 左边界向右收缩
                basket1 = fruits[end];
                start = end - 1;
                basket2 = fruits[start];
                while (start > 0) {
                    if (basket2 == fruits[start - 1]) {
                        start--;
                    } else {
                        break;
                    }
                }
            }
            end++;
        }
        return result;
    }
};
```

## 螺旋矩阵II

### 59 [螺旋矩阵II](https://leetcode.com/problems/spiral-matrix-ii/)

给你一个正整数 `n` ，生成一个包含$1到n^2$所有元素，且元素按顺时针顺序螺旋排列的 `n x n` 正方形矩阵 `matrix` 。

解题思路：考验代码功底的一道题，按照要求通过循环作出图即可

解法一：

```c++
class Solution {
public:
    vector<vector<int>> generateMatrix(int n) {
        // 初始化结果数组
        vector<vector<int>> result(n, vector<int>(n, 0));
        // 开始填充
        int num = 1, target = n*n;
        int t = 0, r = n-1, b = n-1, l = 0;
        while (num <= target) {
            int x = l, y = t;
            result[y][x++] = num++;
            // 从左至右填充
            while (x < r && num <= target) result[y][x++] = num++;
            r--;
            // 从上至下填充
            while (y < b && num <= target) result[y++][x] = num++;
            b--;
            // 底部从右至左填充
            while (x > l && num <= target) result[y][x--] = num++;
            l++;
            // 左边从下至上填充
            while (y > t && num <= target) result[y--][x] = num++;
            t++;
        }
        return result;
    }
};
```

