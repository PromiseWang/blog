---
title: 代码随想录算法训练营第十五天 104.二叉树的最大深度、559.n叉树的最大深度
abbrlink: e2c426b8
date: 2023-10-07 21:33:49
tags:
  - 算法
  - 代码随想录
  - LeetCode
categories: 刷题
---

# 代码随想录算法训练营第十五天 104.二叉树的最大深度、559.n叉树的最大深度

##  104.二叉树的最大深度

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)
>
> 文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0104.%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E6%9C%80%E5%A4%A7%E6%B7%B1%E5%BA%A6.html)
>
> 视频讲解：[二叉树的高度和深度有啥区别？究竟用什么遍历顺序？很多录友搞不懂 | 104.二叉树的最大深度](https://www.bilibili.com/video/BV1Gd4y1V75u)
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

## 559. N 叉树的最大深度

> 题目链接：[559. N 叉树的最大深度](https://leetcode.cn/problems/maximum-depth-of-n-ary-tree/)
>
> 状态：AC

### 思路

#### 方法一

递归遍历每个子树，叶子节点时候`return 0`，否则下一层深度加一

#### 方法二

每一层遍历时候入队的不是`Left`和`Right`，而是每一个`Children[i]`节点，最后输出`result`的大小

### 代码

``` go
// 递归法
func maxNum(a, b int) int {
     if a > b {
         return a
     }
     return b
 }

func maxDepth(root *Node) int {
    if root == nil {
		return 0
	}
	depth := 1
	for _, v := range root.Children {
		depth = maxNum(depth, maxDepth(v)+1)
	}
	return depth
}
```



``` go
//层先法
func maxDepth(root *Node) int {
    var result [][]int
	if root == nil {
		return 0
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	for !queue.Empty() {
		size := queue.Size()
		var temp []int
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

## 222. 完全二叉树的节点个数

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/count-complete-tree-nodes/)
>
> 文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0222.%E5%AE%8C%E5%85%A8%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E8%8A%82%E7%82%B9%E4%B8%AA%E6%95%B0.html#%E6%80%9D%E8%B7%AF)
>
> 视频讲解：[要理解普通二叉树和完全二叉树的区别！ | LeetCode：222.完全二叉树节点的数量](https://www.bilibili.com/video/BV1eW4y1B7pD)
>
> 状态：AC

### 思路

#### 方法一

使用前序、中序、后序、前序遍历都可，输出节点个数

### 代码

``` go
// 前序遍历
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := arrayqueue.New()
	queue.Enqueue(root)
	count := 0
	for !queue.Empty() {
		size := queue.Size()
		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			count++
			if node.(*TreeNode).Left != nil {
				queue.Enqueue(node.(*TreeNode).Left)
			}
			if node.(*TreeNode).Right != nil {
				queue.Enqueue(node.(*TreeNode).Right)
			}
		}
	}
	return count
}
```

