---
title: 代码随想录算法训练营第十八天 513.找树左下角的值、112.路径总和、113.路径总和II、106.从中序与后序遍历序列构造二叉树、105.从前序与中序遍历序列构造二叉树
tags:
  - 算法
  - 代码随想录
  - LeetCode
categories: 刷题
abbrlink: de527742
date: 2023-10-10 19:56:14
---

# 代码随想录算法训练营第十八天  513找树左下角的值、112.路径总和、113.路径总和II、106.从中序与后序遍历序列构造二叉树、105.从前序与中序遍历序列构造二叉树

## 513. 找树左下角的值

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/find-bottom-left-tree-value/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0513.%E6%89%BE%E6%A0%91%E5%B7%A6%E4%B8%8B%E8%A7%92%E7%9A%84%E5%80%BC.html)
>
>   视频讲解：[怎么找二叉树的左下角？ 递归中又带回溯了，怎么办？| LeetCode：513.找二叉树左下角的值](https://www.bilibili.com/video/BV1424y1Z7pn)
>
>   状态：AC

### 思路

这道题对左下角的定义是：最后一行，最左面的元素，只需要层序遍历，输出最后一行的第一个元素即可。

### 代码

``` go
func findBottomLeftValue(root *TreeNode) int {
	queue := arrayqueue.New()
	var result [][]int
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
	fmt.Println(result)
	return result[len(result)-1][0]
}

```

## 112. 路径总和、113. 路径总和II

>   题目链接：
>
>   -   [112.路径总和](https://leetcode.cn/problems/path-sum/)
>   -   [113.路径总和ii](https://leetcode.cn/problems/path-sum-ii/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0112.%E8%B7%AF%E5%BE%84%E6%80%BB%E5%92%8C.html)
>
>   视频讲解：[拿不准的遍历顺序，搞不清的回溯过程，我太难了！ | LeetCode：112. 路径总和](https://www.bilibili.com/video/BV19t4y1L7CR)
>
>   状态：AC

### 思路

1.   找到根节点到叶子节点的路径：这里使用先序遍历（深度搜索）
2.   将该路径所有节点求和，判断是否等于`targetSum`，有则保存至`result`中

### 代码

``` go
// 112题
func hasPathSum(root *TreeNode, targetSum int) bool {
	var result [][]int
	var path []int
	getResult(root, &result, &path)
	for i := 0; i < len(result); i++ {
		sum := 0
		for _, v := range result[i] {
			sum += v
		}
		if sum == targetSum {
			return true
		}
	}
	return false
}

func getResult(root *TreeNode, result *[][]int, path *[]int) {
	if root.Left == nil && root.Right == nil {
		*path = append(*path, root.Val)
		temp := make([]int, len(*path))
		copy(temp, *path)
		*result = append(*result, temp)
		*path = (*path)[:len(*path)-1]
		return
	}
	if root.Left != nil {
		*path = append(*path, root.Val)
		getResult(root.Left, result, path)
		*path = (*path)[:len(*path)-1]
	}
	if root.Right != nil {
		*path = append(*path, root.Val)
		getResult(root.Right, result, path)
		*path = (*path)[:len(*path)-1]
	}
}
```

``` go
// 113题
func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var path []int
	getResult1(root, &result, &path, targetSum)
	return result
}

func getResult1(root *TreeNode, result *[][]int, path *[]int, target int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*path = append(*path, root.Val)
		sum := 0
		for _, v := range *path {
			sum += v
		}
		if sum == target {
			temp := make([]int, len(*path))
			copy(temp, *path)
			*result = append(*result, temp)
		}
		*path = (*path)[:len(*path)-1]
		return
	}
	if root.Left != nil {
		*path = append(*path, root.Val)
		getResult1(root.Left, result, path, target)
		*path = (*path)[:len(*path)-1]
	}
	if root.Right != nil {
		*path = append(*path, root.Val)
		getResult1(root.Right, result, path, target)
		*path = (*path)[:len(*path)-1]
	}
}
```

## 106.从中序与后序遍历序列构造二叉树、105.从前序与中序遍历序列构造二叉树

>   题目链接：
>
>   -   [106.从中序与后序遍历序列构造二叉树](https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)
>   -   [105.从前序与中序遍历序列构造二叉树](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0106.%E4%BB%8E%E4%B8%AD%E5%BA%8F%E4%B8%8E%E5%90%8E%E5%BA%8F%E9%81%8D%E5%8E%86%E5%BA%8F%E5%88%97%E6%9E%84%E9%80%A0%E4%BA%8C%E5%8F%89%E6%A0%91.html)
>
>   视频讲解：[坑很多！来看看你掉过几次坑 | LeetCode：106.从中序与后序遍历序列构造二叉树](https://www.bilibili.com/video/BV1vW4y1i7dn)
>
>   状态：AC

### 思路

这道题自己最开始想的不是很明白，而且自己在考试时做这类题也没有固定的套路，看了下卡哥的讲解

- 后序数组长度为0, 空节点
- 后序数组最后一个元素为父节点元素
- 寻找中序数组位置  作切割点
- 切中序数组
- 切后序数组
- 递归处理左区间 右区间

<div style="position: relative; width: 100%; height: 0;padding-bottom: 100%;"><iframe 
src="../images/day18/力扣106.mp4" scrolling="no" border="0" 
frameborder="no" framespacing="0" allowfullscreen="true" style="position: absolute; width: 100%; 
height: 100%; left: 0; top: 0;"> </iframe></div>

### 代码

``` go
// 中序后序建立二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 1.后序数组长度为0, 空节点
	if len(postorder) == 0 {
		return nil
	}
	// 2.后序数组最后一个元素为父节点元素
	rootValue := postorder[len(postorder)-1]
	root := &TreeNode{
		Val:   rootValue,
		Left:  nil,
		Right: nil,
	}
	if len(postorder) == 1 {
		return root
	}
	// 3.寻找中序数组位置  作切割点
	index := 0
	for ; index < len(inorder); index++ {
		if inorder[index] == rootValue {
			break
		}
	}

	// 4.切中序数组
	leftInorder := inorder[:index]
	rightInorder := inorder[index+1:]
	// 5.切后序数组
	leftPostorder := postorder[:len(leftInorder)]
	rightPostorder := postorder[len(leftPostorder) : len(leftPostorder)+len(rightInorder)]
	// 6.递归处理左区间 右区间
	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)
	return root
}
```

``` go
// 前序中序建立二叉树
func buildTree2(preorder []int, inorder []int) *TreeNode {
	// 1.前序数组为0, 空节点
	if len(preorder) == 0 {
		return nil
	}
	// 2.前序数组第一个元素为节点元素
	rootValue := preorder[0]
	root := &TreeNode{
		Val:   rootValue,
		Left:  nil,
		Right: nil,
	}
	if len(preorder) == 1 {
		return root
	}
	// 3.寻找中序数组位置  作切割点
	index := 0
	for ; index < len(inorder); index++ {
		if inorder[index] == rootValue {
			break
		}
	}

	// 4.切中序数组
	leftInorder := inorder[:index]
	rightInorder := inorder[index+1:]
	// 5.切前序数组
	leftPreorder := preorder[1 : len(leftInorder)+1]
	rightPreorder := preorder[len(leftPreorder)+1:]
	// 6.递归处理左区间 右区间
	root.Left = buildTree2(leftPreorder, leftInorder)
	root.Right = buildTree2(rightPreorder, rightInorder)
	return root
}

```

