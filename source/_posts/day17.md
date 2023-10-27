---
title: 代码随想录算法训练营第十七天 110.平衡二叉树、257. 二叉树的所有路径、404.左叶子之和
tags:
  - 算法
  - 代码随想录
  - LeetCode
  - 二叉树
categories: 刷题
abbrlink: 467977b0
date: 2023-10-09 23:26:52
---

# 代码随想录算法训练营第十七天 110.平衡二叉树、257. 二叉树的所有路径、404.左叶子之和

## 110. 平衡二叉树

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/balanced-binary-tree/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0110.%E5%B9%B3%E8%A1%A1%E4%BA%8C%E5%8F%89%E6%A0%91.html)
>
>   视频讲解：[后序遍历求高度，高度判断是否平衡 | LeetCode：110.平衡二叉树](https://www.bilibili.com/video/BV1Ug411S7my)
>
>   状态：AC

### 思路

由之前的104题：求二叉树最大深度一讲，卡哥提到过深度和高度的区别，结合这道题的视频看懂了。举个例子，我们现在站在地上，看见一幢楼会说：“这楼有多高？”；看见一个坑，会说：“这坑有多深？”。也就是说高度要从下往上看、深度要从上往下看。从下向上自然就需要后序遍历。

1.   每个叶子节点的高度为1，叶子节点的子节点（也就是两个空节点）高度为0；
2.   定义左子树高度为`leftHeight`，左子树递归返回左子树的高度，如果高度为`-1`则执行`return -1`；
3.   定义右子树高度为`rightHeight`，右子树递归返回右子树的高度，如果高度为`-1`则执行`return -1`；
4.   比较左右子树高度差，差的绝对值小于等于1，说明是平衡二叉树，将最高的子树高度加一，回到2、3步继续。否则不为平衡二叉树，`return -1`。

### 代码

``` go
// 平衡二叉树
func isBalanced(root *TreeNode) bool {
	return backTracking(root) != -1
}
// 返回最大值
func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func backTracking(root *TreeNode) int {
    // 空节点
	if root == nil {
		return 0
	}
	// left child
	leftHeight := backTracking(root.Left)
	if leftHeight == -1 {
		return -1
	}
	// right child
	rightHeight := backTracking(root.Right)
	if rightHeight == -1 {
		return -1
	}
	// 不是平衡二叉树
	if rightHeight-leftHeight > 1 || rightHeight-leftHeight < -1 {
		return -1
	} else {
		return maxNum(rightHeight, leftHeight) + 1  //高度加一继续返回
	}
}
```

## 257. 二叉树的所有路径

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/binary-tree-paths/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0257.%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E6%89%80%E6%9C%89%E8%B7%AF%E5%BE%84.html)
>
>   视频讲解：[递归中带着回溯，你感受到了没？| LeetCode：257. 二叉树的所有路径](https://www.bilibili.com/video/BV1ZG411G7Dh)
>
>   状态：AC

### 思路

回溯算法的比较典型的应用，要输出所有的路径，应该使用先序遍历。走到叶子结点后将路径添加至最后的结果。

卡哥回溯三部曲：

1.   确立函数参数和返回值：由于是Go语言，全局变量在LeetCode中会出现一些问题，具体可以查看这篇我的这篇博客[Go语言刷LeetCode使用全局变量的问题](https://promisewang.github.io/post/bc862a56.html#%E9%81%87%E5%88%B0%E6%9C%80%E5%A4%B4%E5%A4%A7%E7%9A%84%E9%97%AE%E9%A2%98%EF%BC%81%EF%BC%81%EF%BC%81%EF%BC%81%EF%BC%81%EF%BC%81)。函数变量包括：
     -   根节点`root`
     -   这道题的结果`result`
     -   一条路径`path`
2.   确定终止条件：当前节点为子节点时，保存`path`到`result`，`return`
3.   单次递归：左节点不为空，向左递归；右节点不为空；向右递归

### 代码

``` go
func binaryTreePaths(root *TreeNode) []string {
	result := []string{}
	path := []int{}
	getResult(root, &result, &path)
	return result
}

func getResult(root *TreeNode, result *[]string, path *[]int) {
    // 终止条件，并且将切片类型的path按要求写成字符串，添加到result中。
	if root.Left == nil && root.Right == nil {
		*path = append(*path, root.Val) //添加新的节点
		var temp string
		if len(*path) == 1 {
			temp = "1"
		} else {
			temp = strconv.Itoa((*path)[0])
			for i := 1; i < len(*path); i++ {
				temp += "->" + strconv.Itoa((*path)[i])
			}
		}
		*path = (*path)[:len(*path)-1]
		*result = append(*result, temp)
		return
	}
    // 左节点不为空，向左递归
	if root.Left != nil {
		*path = append(*path, root.Val)
		getResult(root.Left, result, path)
		*path = (*path)[:len(*path)-1]
	}
    
    // 右节点不为空，向右递归
	if root.Right != nil {
		*path = append(*path, root.Val)
		getResult(root.Right, result, path)
		*path = (*path)[:len(*path)-1]
	}
}

```

## 404.左叶子之和

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/sum-of-left-leaves/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0404.%E5%B7%A6%E5%8F%B6%E5%AD%90%E4%B9%8B%E5%92%8C.html)
>
>   视频讲解：[二叉树的题目中，总有一些规则让你找不到北 | LeetCode：404.左叶子之和](https://www.bilibili.com/video/BV1GY4y1K7z8)
>
>   状态：AC

### 思路

要判断的是所有左叶子的和，选择先序遍历。递归三部曲请出来：

1.   确立函数参数和返回值：此处没有用全局变量，需要使用一个指针存放和。如何判断是左节点，我这里引入一个变量`flag`为`bool`类型。如果向左节点递归，则为`true`，否则为`false`。这样当到达叶子节点时候判断`flag`值就能知道是不是左叶子
     -   根节点`root`
     -   和指针`sum`
     -   是否为左节点`flag`
2.   确定终止条件：如果是叶子节点，并且`flag == true`，`*sum += root.Val`，返回。
3.   单次递归：如果左节点不为空，向左递归，`flag = true`；如果右结点不为空，向右递归，`flag = false`

### 代码

``` go
func sumOfLeftLeaves(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 0
	}
	sum := 0
	getSum(root, &sum, false)
	return sum
}

// flag为true, 代表左节点; 否则是右结点
func getSum(root *TreeNode, sum *int, flag bool) { 
	if root.Left == nil && root.Right == nil && flag {
		*sum += root.Val
		return
	}
	if root.Left != nil {
		getSum(root.Left, sum, true)
	}
	if root.Right != nil {
		getSum(root.Right, sum, false)
	}
}
```

