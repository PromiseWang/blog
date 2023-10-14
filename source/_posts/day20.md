---
title: 代码随想录算法训练营第二十天  530.二叉搜索树的最小绝对差、501.二叉搜索树中的众数、236.二叉树的最近公共祖先
tags:
  - 算法
  - 代码随想录
  - LeetCode
categories: 刷题
abbrlink: 80a9e5fb
date: 2023-10-14 17:58:14
---

# 代码随想录算法训练营第二十天  530.二叉搜索树的最小绝对差、501.二叉搜索树中的众数、236.二叉树的最近公共祖先

## 530.二叉搜索树的最小绝对差

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/minimum-absolute-difference-in-bst/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0530.%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E7%9A%84%E6%9C%80%E5%B0%8F%E7%BB%9D%E5%AF%B9%E5%B7%AE.html)
>
>   视频讲解：[二叉搜索树中，需要掌握如何双指针遍历！| LeetCode：530.二叉搜索树的最小绝对差](https://www.bilibili.com/video/BV1DD4y11779)
>
>   状态：AC

### 思路

-   思路一：二叉搜索树中序遍历是递增的，中序遍历后得到结果数组，再遍历结果数组，计算相邻元素的差，找最小值。
-   思路二：使用双指针，看了卡哥的解法豁然开朗。定义当前指针`cur`和上个节点的指针`pre=nil`
    1.   `cur`开始中序遍历时，不断找最左节点的值
    2.   和`pre`指针判断，如果`pre`为空，说明当前是第一个节点，此后将`pre=cur`
    3.   `cur`继续递归，`pre`始终会慢`cur`一层
    4.   比较`cur.Val - pre.Val<minNum`

### 代码

``` go
// 思路一
func getMinimumDifference(root *TreeNode) int {
	var path []int
	inorder(root, &path)
	minNum := math.MaxInt32
	for i := 1; i < len(path); i++ {
		if path[i]-path[i-1] < minNum {
			minNum = path[i] - path[i-1]
		}
	}
	return minNum
}

func inorder(root *TreeNode, path *[]int) {
	if root.Left != nil {
		inorder(root.Left, path)
	}
	*path = append(*path, root.Val)
	if root.Right != nil {
		inorder(root.Right, path)
	}
}
```

``` go
// 思路二
func getMinimumDifference(root *TreeNode) int {
	minNum := 999999999
	var pre *TreeNode
	var inorder2 func(cur *TreeNode)
	inorder2 = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		inorder2(cur.Left)
		if pre != nil && cur.Val-pre.Val < minNum {
			minNum = cur.Val - pre.Val
		}
		pre = cur
		inorder2(cur.Right)
	}
	inorder2(root)
	return minNum
}
```

## 501.二叉搜索树中的众数

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/find-mode-in-binary-search-tree/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0501.%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E4%B8%AD%E7%9A%84%E4%BC%97%E6%95%B0.html)
>
>   视频讲解：[不仅双指针，还有代码技巧可以惊艳到你！ | LeetCode：501.二叉搜索树中的众数](https://www.bilibili.com/video/BV1fD4y117gp)
>
>   状态：AC

### 思路

整体思路和上题基本类似，依旧使用双指针。只不过要多定义些变量，用来存放历史最大值`maxCount`，当前最大值`count`，众数结果`maxNum`

-   如果`cur.Val == pre.Val`，`count++`；
-   否则判断`count`与`maxCount`
    -   `count > maxCount`：清空`maxNum`，放入新的值
    -   `count == maxCount`：追加新的值
    -   `count < maxCount`：不管
    -   更新`count = 1`

### 代码

``` go
func findMode(root *TreeNode) []int {
	maxCount := 0
	maxNum := []int{}
	count := 0
	var pre *TreeNode
	var mid func(cur *TreeNode)
	mid = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		mid(cur.Left)
		if pre == nil {
			count = 1
		} else if pre.Val == cur.Val {
			count++
		} else {
			count = 1
		}
		pre = cur
		if count == maxCount {
			maxNum = append(maxNum, cur.Val)
		} else if count > maxCount {
			maxCount = count
			maxNum = []int{cur.Val}
		}
		mid(cur.Right)
	}
	mid(root)
	return maxNum
}

```

## 236.二叉树的最近公共祖先

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0236.%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E6%9C%80%E8%BF%91%E5%85%AC%E5%85%B1%E7%A5%96%E5%85%88.html)
>
>   视频讲解：[自底向上查找，有点难度！ | LeetCode：236. 二叉树的最近公共祖先](https://www.bilibili.com/video/BV1jd4y1B7E2)
>
>   状态：AC

### 思路

这道题和卡哥的思路不同。我依旧使用了中序遍历。

1.   中序遍历得到结果`path`
2.   查找中序遍历中`p`与`q`的位置，与`root`进行判断
     -   如果`root`在`p`和`q`的左（右）侧，`root`则需要向右（左）节点移动
     -   如果`root`在`p`和`q`之间（或`root`是`p`或`q`），返回即可，具体可以看下方视频

<div style="position: relative; width: 100%; height: 0;padding-bottom: 100%;"><iframe 
src="../images/day20/二叉树公共祖先.mp4" scrolling="no" border="0" 
frameborder="no" framespacing="0" allowfullscreen="true" style="position: absolute; width: 100%; 
height: 100%; left: 0; top: 0;"> </iframe></div>

### 代码

``` go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var path []int
	getInorderPath(root, &path)
	for {
		index := 0
		indexP := 0
		indexQ := 0
		// 找到当前根结点
		for i := 0; i < len(path); i++ {
			if root.Val == path[i] {
				index = i
			}
			if p.Val == path[i] {
				indexP = i
			}
			if q.Val == path[i] {
				indexQ = i
			}
		}
		if indexP == index {
			return p
		}
		if indexQ == index {
			return q
		}
		// p,q 分散在中序遍历的根结点两侧
		if indexP < index && index < indexQ || indexQ < index && index < indexP {
			return root
		}
		// p,q 在根节点左侧
		if indexP < index && indexQ < index {
			root = root.Left
			continue
		}
		// p,q 在根节点右侧
		if index < indexP && index < indexQ {
			root = root.Right
			continue
		}

	}
}
```

