---
title: 代码随想录算法训练营第九天 28. 实现 strStr()。459.重复的子字符串
tags:
  - 算法
  - 代码随想录
  - LeetCode
  - 字符串
categories: 刷题
abbrlink: '369230e4'
date: 2023-09-28 15:08:08
---

# 代码随想录算法训练营第九天 28. 实现 strStr()。459.重复的子字符串

## 28.实现strStr()

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/)
>
>   文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0028.%E5%AE%9E%E7%8E%B0strStr.html)
>
>   视频讲解：
>
>   -   [帮你把KMP算法学个通透！B站（理论篇）](https://www.bilibili.com/video/BV1PD4y1o7nd/)
>   -   [帮你把KMP算法学个通透！（求next数组代码篇）](https://www.bilibili.com/video/BV1M5411j7Xx)
>
>   状态：看过视频之后AC

### 思路

具体KMP算法原理看卡哥的视频，讲的很好。

#### KMP中匹配的过程

KMP算法，匹配过程放入到一个小视频当中，每个画面持续3秒。好多视频讲解都是“移动模式串”来讲，自己写代码时候有点蒙，所以自己做了个小动图，不使用“移动”来呈现。

<div style="position: relative; width: 100%; height: 100%;padding-bottom: 100%;"><iframe 
src="../images/day09/匹配.mp4" scrolling="no" border="0" 
frameborder="no" framespacing="0" allowfullscreen="true" style="position: absolute; width: 100%; 
height: 100%; left: 0; top: 0;"> </iframe></div>


#### Next数组构建过程

<div style="position: relative; width: 100%; height: 0;padding-bottom: 100%;"><iframe 
src="../images/day09/next数组构建.mp4" scrolling="no" border="0" 
frameborder="no" framespacing="0" allowfullscreen="true" style="position: absolute; width: 100%; 
height: 100%; left: 0; top: 0;"> </iframe></div>



#### Next数组说明

Next数组中，每个元素表示：

-   截止到目前为止，最长相等前后缀的长度；
-   截止到目前为止，最长前缀的后一位。

### 代码

``` go
func getNext(next []int, s string) {
	j := 0
	next[0] = 0
	for i := 1; i < len(s); i++ {
		for s[i] != s[j] && j > 0 {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
		//fmt.Printf("第%v次循环后的next数组结果为:%v\n", i, next)
	}
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := make([]int, len(needle))
	getNext(next, needle)
	j := 0
	for i := range haystack {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}
```



## 459. 重复的字符串

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/repeated-substring-pattern/)
>
>   文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0459.%E9%87%8D%E5%A4%8D%E7%9A%84%E5%AD%90%E5%AD%97%E7%AC%A6%E4%B8%B2.html)
>
>   视频讲解：[字符串这么玩，可有点难度！ | LeetCode：459.重复的子字符串](https://www.bilibili.com/video/BV1cg41127fw)
>
>   状态：AC

### 思路

如果不使用KMP算法还是比较简单的，有很多东西语言已经帮我们实现好了。看了卡哥的讲解感叹这个思路。

构建一个新字符串`newString`为两个旧串`s`的拼接，但是要掐头去尾一个元素。如果`newString`仍然包含`s`，说明存在子串构成原字符串。

### 代码

``` go
func repeatedSubstringPattern(s string) bool {
	newString := s + s
	newString = newString[1 : len(newString)-2]
	return strings.Contains(newString, s)
}
```

