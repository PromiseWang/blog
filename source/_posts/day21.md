---
title: 代码随想录算法训练营第二十一天  235.二叉搜索树的最近公共祖先、701.二叉搜索树中的插入操作、450.删除二叉搜索树中的节点
tags:
  - 算法
  - 代码随想录
  - LeetCode
categories: 刷题
abbrlink: 1b36a0ee
date: 2023-10-14 22:31:50
---

# 代码随想录算法训练营第二十一天  235.二叉搜索树的最近公共祖先、701.二叉搜索树中的插入操作、450.删除二叉搜索树中的节点

## 235.二叉搜索树的最近公共祖先

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0235.%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E7%9A%84%E6%9C%80%E8%BF%91%E5%85%AC%E5%85%B1%E7%A5%96%E5%85%88.html)
>
>   视频讲解：[二叉搜索树找祖先就有点不一样了！| 235. 二叉搜索树的最近公共祖先](https://www.bilibili.com/video/BV1Zt4y1F7ww?share_source=copy_web)
>
>   状态：AC

### 思路

同上一天的最后一道题：“236.二叉树的最近公共祖先”

[链接跳转](https://promisewang.github.io/post/80a9e5fb.html#236-%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E6%9C%80%E8%BF%91%E5%85%AC%E5%85%B1%E7%A5%96%E5%85%88)

## 701.二叉搜索树中的插入操作

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/insert-into-a-binary-search-tree/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0701.%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E4%B8%AD%E7%9A%84%E6%8F%92%E5%85%A5%E6%93%8D%E4%BD%9C.html)
>
>   视频讲解：[原来这么简单？ | LeetCode：701.二叉搜索树中的插入操作](https://www.bilibili.com/video/BV1Et4y1c78Y?share_source=copy_web)
>
>   状态：AC

## 思路

1.   寻找应该插入的位置，和当前的根节点比较
2.   找到插入

### 代码

``` go
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}

	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}
```

## 450. 删除二叉搜索树中的节点

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/delete-node-in-a-bst/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0450.%E5%88%A0%E9%99%A4%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E4%B8%AD%E7%9A%84%E8%8A%82%E7%82%B9.html)
>
>   视频讲解：[调整二叉树的结构最难！| LeetCode：450.删除二叉搜索树中的节点](https://www.bilibili.com/video/BV1tP41177us?share_source=copy_web)
>
>   状态：AC

### 思路

思路大体上还算好想，其中一种情况困扰了好久，还是看了看题解。在树中查找应该删除的节点

-   1.   当前递归到空节点，说明没找到

-   如果找到

2.   如果是叶子节点，直接删`return nil`

3.   如果左空右不空，右节点直接上位

4.   如果左不空有空，左节点直接上位

5.   如果左右都不空，需要将左子树找到新家

第五点删除过程在下方视频中展示

<div style="position: relative; width: 100%; height: 0;padding-bottom: 100%;"><iframe 
src="../images/day21/二叉排序树删除元素.mp4" scrolling="no" border="0" 
frameborder="no" framespacing="0" allowfullscreen="true" style="position: absolute; width: 100%; 
height: 100%; left: 0; top: 0;"> </iframe></div>

### 代码

``` go
func deleteNode(root *TreeNode, key int) *TreeNode {
	// 1. 没找到删除节点
	if root == nil {
		return root
	}
	if root.Val == key {
		// 2. 该节点是叶子节点: 直接删除
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// 3. 左空右不空: 右节点补位
		if root.Left == nil && root.Right != nil {
			node := root.Right
			return node
		}
		// 4. 左不空右空: 左节点补位
		if root.Left != nil && root.Right == nil {
			node := root.Left
			return node
		}
		// 5. 都不空: 则将删除节点的左子树头结点（左孩子）放到删除节点的右子树的最左面节点的左孩子上，返回删除节点右孩子为新的根节点。
		if root.Left != nil && root.Right != nil {
			node := root.Right
			for node.Left != nil {
				node = node.Left
			}
			node.Left = root.Left
			root = root.Right
			return root
		}
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}
```

