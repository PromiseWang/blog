---
title: 代码随想录算法训练营第十四天 102.二叉树的层序遍历、107.二叉树的层次遍历II、199.二叉树的右视图、637.二叉树的层平均值、429.N叉树的层序遍历、515.在每个树行中找最大值、116.填充每个节点的下一个右侧节点指、117.填充每个节点的下一个右侧节点指针II、104.二叉树的最大深度、111.二叉树的最小深度、226.翻转二叉树、101.对称二叉树
tags:
  - 算法
  - 代码随想录
  - LeetCode
categories: 刷题
abbrlink: e1408dd4
date: 2023-10-07 21:33:44
---
# 代码随想录算法训练营第十四天 102.二叉树的层序遍历、107.二叉树的层次遍历II、199.二叉树的右视图、637.二叉树的层平均值、429.N叉树的层序遍历、515.在每个树行中找最大值、116.填充每个节点的下一个右侧节点指、117.填充每个节点的下一个右侧节点指针II、104.二叉树的最大深度、111.二叉树的最小深度

> 文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0102.%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E5%B1%82%E5%BA%8F%E9%81%8D%E5%8E%86.html)
>
> 视频讲解：[讲透二叉树的层序遍历 | 广度优先搜索 | LeetCode：102.二叉树的层序遍历](https://www.bilibili.com/video/BV1GY4y1u7b2)

## 102. 二叉树的层序遍历

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/binary-tree-level-order-traversal/)
>
> 状态：AC

### 思路

层序遍历顾名思义就是要一层一层的输出树的结点，父节点输出后还要将子节点保存一下，到下一层继续使用，这种依次使用，队列的数据结构是最合适的。父节点先入队，将两个子节点依次入队。这一层的父节点入队过后要统计有多少个，为了输出每层分别是谁

### 代码

```go
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	for !queue.Empty() {
		size := queue.Size()
		temp := []int{}
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			temp = append(temp, node.(*TreeNode).Val)
			if node.(*TreeNode).Left != nil {
				queue.Enqueue(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				queue.Enqueue(node.(*TreeNode).Right)
			}
		}
		result = append(result, temp)
	}
	return result
}
```

## 107. 二叉树的层序遍历II

> 题目链接：[107.二叉树的层次遍历II](https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/)
>
> 状态：AC

### 思路

基本同上，这次不过要从下向上输出，翻转最后上一题的`result`即可

### 代码

``` go
func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	for !queue.Empty() {
		size := queue.Size()
		temp := []int{}
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			temp = append(temp, node.(*TreeNode).Val)
			if node.(*TreeNode).Left != nil {
				queue.Enqueue(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				queue.Enqueue(node.(*TreeNode).Right)
			}
		}
		result = append(result, temp)
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return result
}
```

## 199.二叉树的右视图

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/binary-tree-right-side-view/)
>
> 状态：AC

### 思路

将第一次层序遍历的结果使用二维数组逐行写出来，不难发现右视图就是二维数组每一行的最后一个元素

### 代码

``` go
func rightSideView(root *TreeNode) []int {
	var results [][]int
	if root == nil {
		return []int{}
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	for !queue.Empty() {
		size := queue.Size()
		temp := []int{}
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			temp = append(temp, node.(*TreeNode).Val)
			if node.(*TreeNode).Left != nil {
				queue.Enqueue(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				queue.Enqueue(node.(*TreeNode).Right)
			}
		}
		results = append(results, temp)
	}
	result := []int{}
	for _, v := range results {
		result = append(result, v[len(v)-1])
	}
	return result
}

```

## 637.二叉树的层平均值

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/average-of-levels-in-binary-tree/)
>
> 状态：AC

### 思路

层序每一层之后求平均值

### 代码

``` go
func averageOfLevels(root *TreeNode) []float64 {
	var result []float64
	if root == nil {
		return result
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	for !queue.Empty() {
		size := queue.Size()
		sum := 0
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			sum += node.(*TreeNode).Val
			if node.(*TreeNode).Left != nil {
				queue.Enqueue(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				queue.Enqueue(node.(*TreeNode).Right)
			}
		}
		result = append(result, float64(sum)/float64(size))
	}
	return result
}
```

## 429.N叉树的层序遍历

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/n-ary-tree-level-order-traversal/)
>
> 状态：AC

### 思路

层序添加子节点改为循环遍历`Children`数组即可

### 代码

``` go
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	for !queue.Empty() {
		size := queue.Size()
		temp := []int{}
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			temp = append(temp, node.(*Node).Val)
			for j := 0; j < len(node.(*Node).Children); j++ {
				if node.(*Node).Children[j] != nil {
					queue.Enqueue(node.(*Node).Children[j])
				}
			}
		}
		result = append(result, temp)
	}
	return result
}

```

## 515.在每个树行中找最大值

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/find-largest-value-in-each-tree-row/)
>
> 状态：AC

### 思路

每一层找最大，模版依旧

### 代码

``` go
func largestValues(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	for !queue.Empty() {
		size := queue.Size()
		maxNum := math.MinInt32
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			if node.(*TreeNode).Val > maxNum {
				maxNum = node.(*TreeNode).Val
			}
			if node.(*TreeNode).Left != nil {
				queue.Enqueue(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				queue.Enqueue(node.(*TreeNode).Right)
			}
		}
		result = append(result, maxNum)
	}
	return result
}

```

##  116.填充每个节点的下一个右侧节点指针、 117.填充每个节点的下一个右侧节点指针II

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/)
>
> 状态：AC

### 思路

每行遍历出来看做一个链表，两道题代码思路一样

### 代码

``` go

// 与117题代码相同
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := arrayqueue.New()
	root.Next = nil
	queue.Enqueue(root)
	for !queue.Empty() {
		size := queue.Size()
		var p *Node = nil
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()

			if node.(*Node).Left != nil {
				queue.Enqueue(node.(*Node).Left)
			}
			if node.(*Node).Right != nil {
				queue.Enqueue(node.(*Node).Right)
			}
			if p == nil {
				p = node.(*Node)
			} else {
				p.Next = node.(*Node)
				p = p.Next
			}
		}
		p.Next = nil
	}
	return root
}

```

##  104.二叉树的最大深度

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)
>
> 状态：AC

### 思路

#### 方法一

使用递归法，如果递归到叶子节点则返回0，否则使用左右叶子结点继续向下递归，深度加一。取左右结点的最大值

#### 方法二

层序遍历，看`result`二维数组的行数

### 代码

``` go
// 递归
func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxNum(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
```

``` go
// 层序
func maxDepth1(root *TreeNode) int {
	var result [][]int
	if root == nil {
		return 0
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	for !queue.Empty() {
		size := queue.Size()
		temp := []int{}
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			temp = append(temp, node.(*TreeNode).Val)
			if node.(*TreeNode).Left != nil {
				queue.Enqueue(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				queue.Enqueue(node.(*TreeNode).Right)
			}
		}
		result = append(result, temp)
	}
	return len(result)
}
```

## 111.二叉树的最小深度

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/minimum-depth-of-binary-tree/)
>
> 状态：AC

### 思路

#### 方法一

使用递归法，递归终止条件如下：到达了叶子节点`return 1`。到达空节点`return 0`。分别向左和右结点继续递归，计算更小的层数，当前层递归完成后`return 层数加一`

#### 方法二

使用层序遍历。遍历时如果遇到叶子节点说明到达了最低点，直接`return 当前层数`

``` go
//递归
func minNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minN := math.MaxInt32
	if root.Left != nil {
		minN = minNum(minDepth(root.Left), minN)
	}
	if root.Right != nil {
		minN = minNum(minDepth(root.Right), minN)
	}
	return minN + 1
}
```

``` go
//层序
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	count := 0
	for !queue.Empty() {
		count++
		size := queue.Size()
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			if node.(*TreeNode).Left != nil {
				queue.Enqueue(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				queue.Enqueue(node.(*TreeNode).Right)
			}
			if node.(*TreeNode).Left == nil && node.(*TreeNode).Right == nil {
				return count
			}
		}
	}
	return count
}
```

## 226.翻转二叉树

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/invert-binary-tree/)
>
> 状态：AC

### 思路

交换当前结点左右子节点，向下递归

### 代码

``` go
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
```

##  101. 对称二叉树

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/symmetric-tree/)
>
> 文章链接：[代码随想录(programmercarl.com)](https://programmercarl.com/0101.%E5%AF%B9%E7%A7%B0%E4%BA%8C%E5%8F%89%E6%A0%91.html)
>
> 视频链接：[同时操作两个二叉树 | LeetCode：101. 对称二叉树](https://www.bilibili.com/video/BV1ue4y1Y7Mf)
>
> 状态：AC

### 思路

<font color="red">错误思路如下</font>

最开始想的是输出中序遍历的结果，根节点一定是最中间的数，比较两端是否对称即可。但是这种情况没考虑到

``` go
//      1
//     / \
//    2   2
//   /   /
//  2   2
```

所以这个思路是错的

---

<font color="green">正确思路如下</font>

用后序遍历和“伪后序遍历”依次比较子树。后序遍历是左右根，伪后序比那里是右左根，这样可以比较相同的结点。

使用递归需要做如下判断

- 递归函数的构造
    - 比较左右结点，传参是两个节点`func compare(left, right *TreeNode) bool`
- 递归终止条件
    - 两个节点均为空节点，对称`return true`
    - 左空右不空，或，左不空右空，不对称`return false`
    - 左右均不空，但是值不同，不对称`return false`
- 单次递归逻辑
    - 左节点的左 和 右结点的右  向下递归
    - 且 左节点的右 和 右结点的左  向下递归
    - 且二者值均为`true`

### 代码

``` go
func compare(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right == nil {
		return false
	}
	if left == nil && right != nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return compare(left.Left, right.Right) && compare(left.Right, right.Left)

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}
```

