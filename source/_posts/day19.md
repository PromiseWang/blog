---
title: 代码随想录算法训练营第十九天  654.最大二叉树、617.合并二叉树、700.二叉搜索树中的搜索、98.验证二叉搜索树
tags:
  - 算法
  - 代码随想录
  - LeetCode
categories: 刷题
abbrlink: 38d5093a
date: 2023-10-12 22:38:04
---

# 代码随想录算法训练营第十九天 654.最大二叉树、617.合并二叉树、700.二叉搜索树中的搜索、98.验证二叉搜索树

## 654.最大二叉树

>   题目链接：[力扣题目地址](https://leetcode.cn/problems/maximum-binary-tree/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0654.%E6%9C%80%E5%A4%A7%E4%BA%8C%E5%8F%89%E6%A0%91.html)
>
>   视频讲解：[又是构造二叉树，又有很多坑！| LeetCode：654.最大二叉树](https://www.bilibili.com/video/BV1MG411G7ox)
>
>   状态：AC

### 思路

1.   找到当前数组最大值`max`以及所在位置`maxIndex`
2.   新建一个节点，将`max`设为节点的值，判断此时的数组长度是否为0
3.   切割左数组和右数组、并向左子树和右子树递归

### 代码

``` go
func constructMaximumBinaryTree(nums []int) *TreeNode {
	max := 0
	maxIndex := 0
	for i, v := range nums {
		if v > max {
			max = v
			maxIndex = i
		}
	}
	root := &TreeNode{
		Val:   max,
		Left:  nil,
		Right: nil,
	}
	if len(nums) == 0 {
		return root
	}
	leftNums := nums[:maxIndex]
	rightNums := nums[maxIndex+1:]
	if len(leftNums) != 0 {
		root.Left = constructMaximumBinaryTree(leftNums)
	}
	if len(rightNums) != 0 {
		root.Right = constructMaximumBinaryTree(rightNums)
	}
	return root
}
```

## 617. 合并二叉树

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/merge-two-binary-trees/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0617.%E5%90%88%E5%B9%B6%E4%BA%8C%E5%8F%89%E6%A0%91.html)
>
>   视频讲解：[一起操作两个二叉树？有点懵！| LeetCode：617.合并二叉树](https://www.bilibili.com/video/BV1m14y1Y7JK)
>
>   状态：AC

### 思路

合并二叉树首先以其中一棵树为基准，将另一棵树“移植”过来。这里以`root1`为准

1.   如果`root1`是空节点，说明应该把`root2`的东西“移植过来”（如果没有也没事）
2.   反之同理，如果`root2`是空节点，则`root1`继续向下遍历
3.   `root1.Val += root2.Val`
4.   `root1.Left`是两棵树向左节点遍历返回的结果。`root1.Right`是两棵树向左节点遍历返回的结果

### 代码

``` go
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}
```

## 700.二叉搜索树中的搜索

>   题目链接：[力扣题目地址](https://leetcode.cn/problems/search-in-a-binary-search-tree/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0700.%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91%E4%B8%AD%E7%9A%84%E6%90%9C%E7%B4%A2.html)
>
>   视频讲解：[不愧是搜索树，这次搜索有方向了！| LeetCode：700.二叉搜索树中的搜索](https://www.bilibili.com/video/BV1wG411g7sF)
>
>   状态：AC

### 思路

这里的搜索并不是返回是否找到，而是要返回子树，总的思路是一样的

1.   `val < root.Val`，向左找；`val > root.Val`，向右找；相等则返回当前结点（包含了子树）

### 代码

``` go
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	var node *TreeNode
	if root.Val == val {
		node = root
		return node
	} else if val < root.Val {
		node = searchBST(root.Left, val)
	} else {
		node = searchBST(root.Right, val)
	}
	return node
}
```

## 98. 验证二叉搜索树

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/validate-binary-search-tree/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://leetcode.cn/problems/validate-binary-search-tree/)
>
>   视频讲解：[你对二叉搜索树了解的还不够！ | LeetCode：98.验证二叉搜索树](https://www.bilibili.com/video/BV18P411n7Q4)
>
>   状态：AC

### 思路

-   <font color="red">错误思路</font>：顺序遍历每个节点，比较是否“左<根<右”，如果均满足返回true，否则是false

    >   错误原因：如果树如下
    >
    >   ​            2
    >
    >   ​         /     \
    >
    >   ​       1        3
    >
    >   ​         \
    >
    >   ​           4
    >
    >   4比2大，但是出现在了2的左边
    >
    >   ``` go
    >   // 错误的代码
    >   func isValidBST(root *TreeNode) bool {
    >   	if root.Left == nil && root.Right != nil {
    >   		if root.Val < root.Right.Val {
    >   			return isValidBST(root.Right)
    >   		} else {
    >   			return false
    >   		}
    >   	} else if root.Left != nil && root.Right == nil {
    >   		if root.Left.Val < root.Val {
    >   			return isValidBST(root.Left)
    >   		} else {
    >   			return false
    >   		}
    >   	} else if root.Left != nil && root.Right != nil {
    >   	  if root.Left.Val < root.Val && root.Val < root.Right.Val {
    >   		return isValidBST(root.Left) && isValidBST(root.Right)
    >   	} else {
    >   		return false
    >   		}
    >       } else {
    >   		return true
    >   	}
    >   }
    >   ```
    >
    >   

-   简单看了下卡哥的文档，知道了用中序遍历，中序遍历得到的结果如果是递增序列则为搜索树

    >   如果树为如下，力扣的测试用例返回的是false，所以不能出现相等的情况
    >
    >   ​            2
    >
    >   ​         /     \
    >
    >   ​       2        2

### 代码

``` go
func isValidBST(root *TreeNode) bool {
	var path []int
	inorder(root, &path)
	for i := 1; i < len(path); i++ {
		if path[i-1] >= path[i] {
			return false
		}
	}
	return true
}
func inorder(root *TreeNode, path *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, path)
	*path = append(*path, root.Val)
	inorder(root.Right, path)
}
```

