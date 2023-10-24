---
title: 代码随想录算法训练营第二十四天  77.组合
tags:
  - 算法
  - 代码随想录
  - LeetCode
  - 回溯算法
categories: 刷题
abbrlink: 9a97e05d
date: 2023-10-16 11:26:54
---

# 代码随想录算法训练营第二十四天  77.组合

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/combinations/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0077.%E7%BB%84%E5%90%88.html#)
>
>   视频讲解：
>   - [带你学透回溯算法-组合问题（对应力扣题目：77.组合）](https://www.bilibili.com/video/BV1ti4y1L7cv)
>   - [组合问题的剪枝操作](https://www.bilibili.com/video/BV1wi4y157er)
>
>   状态：AC

### 思路

这是一道经典的回溯算法的题目，开始做的时候没有考虑到剪枝的情况。直接请出回溯三步：

1.   确立函数和参数：需要参数`n, k, index`分别表示范围、k个数、当前位置
2.   确定终止条件：路径`path`长度为`k`时，将`path`追加到`append`
3.   单次逻辑：从`index`开始逐渐向`path`追加之后的元素；递归调用；删除刚追加的元素

### 代码

``` go
func combine(n int, k int) [][]int {
	path := []int{}
	result := [][]int{}
	var backTracking func(n, k, index int)
	backTracking = func(n, k, index int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := index; i <= n; i++ {
			path = append(path, i)
			backTracking(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	backTracking(n, k, 1)
	return result
}
```

### 剪枝

来举一个例子，`n = 4, k = 4`，那么第一层for循环的时候，从元素2开始的遍历都没有意义了。 在第二层for循环，从元素3开始的遍历都没有意义了。那么在第一层遍历时候，对`i`的判断应该更改：`i <= n - (k - len(path) + 1)`

``` go
func combine(n int, k int) [][]int {
	path := []int{}
	result := [][]int{}
	var backTracking func(n, k, index int)
	backTracking = func(n, k, index int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := index; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backTracking(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	backTracking(n, k, 1)
	return result
}
```

