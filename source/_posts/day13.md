---
title: 代码随想录算法训练营第十三天 144.二叉树的前序遍历、145.二叉树的后序遍历、94.二叉树的中序遍历
tags:
  - 算法
  - 代码随想录
  - LeetCode
  - 二叉树
categories: 刷题
abbrlink: 5d0ef500
date: 2023-10-06 19:14:44
---

# 代码随想录算法训练营第十三天 144.二叉树的前序遍历、145.二叉树的后序遍历、94.二叉树的中序遍历

## 二叉树的遍历

>   题目链接：
>
>   -   [144.二叉树的前序遍历](https://leetcode.cn/problems/binary-tree-preorder-traversal/)
>   -   [145.二叉树的后序遍历](https://leetcode.cn/problems/binary-tree-postorder-traversal/)
>   -   [94.二叉树的中序遍历](https://leetcode.cn/problems/binary-tree-inorder-traversal/)
>
>   文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0239.%E6%BB%91%E5%8A%A8%E7%AA%97%E5%8F%A3%E6%9C%80%E5%A4%A7%E5%80%BC.html)
>
>   视频讲解：[每次写递归都要靠直觉？ 这次带你学透二叉树的递归遍历！](https://www.bilibili.com/video/BV1Wh411S7xt)
>
>   状态：AC

### 递归思路

递归还是比较简单的，直接按照遍历规则写递归代码即可

### 代码

``` go
// 递归前序遍历
func preorderTraversal(root *TreeNode) []int {
	var result []int
	getResultPreorder(root, &result)
	return result
}

func getResultPreorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	getResultPreorder(root.Left, result)
	getResultPreorder(root.Right, result)
}
```

``` go
// 递归后序遍历
func postorderTraversal(root *TreeNode) []int {
	var result []int
	getResultPostorder(root, &result)
	return result
}

func getResultPostorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	getResultPostorder(root.Left, result)
	getResultPostorder(root.Right, result)
	*result = append(*result, root.Val)
}
```

``` go
// 递归中序遍历
func inorderTraversal(root *TreeNode) []int {
	var result []int
	getResultInorder(root, &result)
	return result
}

func getResultInorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	getResultInorder(root.Left, result)
	*result = append(*result, root.Val)
	getResultInorder(root.Right, result)
}
```

### 非递归思路

非递归需要借助栈来实现，每次存放树的左右节点（若存在），并判断根节点应该在什么时候存入结果数组。

``` go
// 非递归前序遍历
func preorderTraversal(root *TreeNode) []int {
	stack := arraystack.New()
	result := []int{}
	if root == nil {
		return result
	}
	stack.Push(root)
	for !stack.Empty() {
		value, _ := stack.Pop()
		result = append(result, value.(*TreeNode).Val)
		if value.(*TreeNode).Right != nil {
			stack.Push(value.(*TreeNode).Right)
		}
		if value.(*TreeNode).Left != nil {
			stack.Push(value.(*TreeNode).Left)
		}
	}
	return result
}
```

``` go
// 非递归后序遍历
func postorderTraversal(root *TreeNode) []int {
	stack := arraystack.New()
	result := []int{}
	if root == nil {
		return result
	}
	stack.Push(root)
	for !stack.Empty() {
		value, _ := stack.Pop()
		result = append(result, value.(*TreeNode).Val)
		if value.(*TreeNode).Left != nil {
			stack.Push(value.(*TreeNode).Left)
		}
		if value.(*TreeNode).Right != nil {
			stack.Push(value.(*TreeNode).Right)
		}
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return result
}

```

``` go
// 非递归中序遍历
func InorderTraversal1(root *TreeNode) []int {
	stack := arraystack.New()
	result := []int{}
	if root == nil {
		return result
	}
	p := root
	for !stack.Empty() || p != nil {
		if p != nil {
			stack.Push(p)
			p = p.Left
		} else {
			node, _ := stack.Pop()
			p = node.(*TreeNode)
			result = append(result, p.Val)
			p = p.Right
		}
	}
	return result
}

```

