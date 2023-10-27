---
title: 代码随想录算法训练营第二十二天  669.修建二叉搜索树、108.将有序数组转化为二叉搜索树、538.把二叉搜索树转化为累加树
tags:
  - 算法
  - 代码随想录
  - LeetCode
  - 二叉树
categories: 刷题
abbrlink: b6e45431
date: 2023-10-15 20:53:48
---

# 代码随想录算法训练营第二十二天  669.修建二叉搜索树、108.将有序数组转化为二叉搜索树、538.把二叉搜索树转化为累加树

## 669.修建二叉搜索树

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/trim-a-binary-search-tree/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0669.%E4%BF%AE%E5%89%AA%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91.html)
>
>   视频讲解：[你修剪的方式不对，我来给你纠正一下！| LeetCode：669. 修剪二叉搜索树](https://www.bilibili.com/video/BV17P41177ud?share_source=copy_web)
>
>   状态：AC

### 思路

与昨天最后一道题类似，昨天最后一道题要求删除一个节点，这道题则为批量删除节点，多加一个判断条件即可

[链接跳转](https://promisewang.github.io/post/1b36a0ee.html#450-%E5%88%A0%E9%99%A4%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E4%B8%AD%E7%9A%84%E8%8A%82%E7%82%B9)

先使用前序遍历得到树的所有节点，再以此判断改节点是否应该删除

### 代码

``` go
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	path := []int{}
	pre(root, &path)
	for _, v := range path {
		if !(low <= v && v <= high) {
			root = deleteNode(root, v)
		}
	}
	return root
}

func pre(root *TreeNode, path *[]int) {
	if root == nil {
		return
	}
	*path = append(*path, root.Val)
	pre(root.Left, path)
	pre(root.Right, path)
}
```

## 108.将有序数组转化为二叉搜索树

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0108.%E5%B0%86%E6%9C%89%E5%BA%8F%E6%95%B0%E7%BB%84%E8%BD%AC%E6%8D%A2%E4%B8%BA%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91.html)
>
>   视频讲解：[构造平衡二叉搜索树！| LeetCode：108.将有序数组转换为二叉搜索树](https://www.bilibili.com/video/BV1uR4y1X7qL?share_source=copy_web)
>
>   状态：AC

### 思路

根据二叉搜索树的定义，比根节点小的值在根节点左侧，比根节点大的值在根节点右侧。而且还要建立平衡搜索树，那么根节点应该为数组最中间的那个值。切割数组，然后递归

### 代码

``` go
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := backTracking(nums, 0, len(nums)-1)
	return root
}

func backTracking(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  backTracking(nums, left, mid-1),
		Right: backTracking(nums, mid+1, right),
	}
	return root
}
```

## 538.把二叉搜索树转化为累加树

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/convert-bst-to-greater-tree/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0538.%E6%8A%8A%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E8%BD%AC%E6%8D%A2%E4%B8%BA%E7%B4%AF%E5%8A%A0%E6%A0%91.html)
>
>   视频讲解：[普大喜奔！二叉树章节已全部更完啦！| LeetCode：538.把二叉搜索树转换为累加树](https://www.bilibili.com/video/BV1d44y1f7wP?share_source=copy_web)
>
>   状态：AC

### 思路

看到树的构建过程，实际上是对树进行“右根左”的遍历。使用双指针，一个指向当前节点，另一个指向过去的节点。当前结点的新的值要加上过去节点的值。总体思路与530题思路一样。[跳转链接](https://promisewang.github.io/post/80a9e5fb.html#530-%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E7%9A%84%E6%9C%80%E5%B0%8F%E7%BB%9D%E5%AF%B9%E5%B7%AE)

### 代码

``` go
func convertBST(root *TreeNode) *TreeNode {
	var pre *TreeNode
	var function func(cur *TreeNode)
	function = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		function(cur.Right)
		if pre != nil {
			cur.Val += pre.Val
		}
		pre = cur
		function(cur.Left)
	}
	function(root)
	return root
}
```



