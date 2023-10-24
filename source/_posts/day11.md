---
title: 代码随想录算法训练营第十一天 20. 有效的括号、1047. 删除字符串中的所有相邻重复项、150. 逆波兰表达式求值。
tags:
  - 算法
  - 代码随想录
  - LeetCode
  - 栈与队列
categories: 刷题
abbrlink: bdbe1e6e
date: 2023-10-03 19:58:17
---

# 代码随想录算法训练营第十一天 20. 有效的括号、1047. 删除字符串中的所有相邻重复项、150. 逆波兰表达式求值。

## 20.有效的括号

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/valid-parentheses/)
>
> 文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0020.%E6%9C%89%E6%95%88%E7%9A%84%E6%8B%AC%E5%8F%B7.html)
>
> 视频讲解：[栈的拿手好戏！| LeetCode：20. 有效的括号](https://www.bilibili.com/video/BV1AF411w78g)
>
> 状态：AC

### 思路

题目保证了输入的字符串只有括号。遇到左括号入栈，遇到右括号出栈，但是要比较出栈的元素和当前的右括号是否匹配，不匹配直接`return false`。最终判断栈是否为空即可。

### 代码

``` go
func isValid(s string) bool {
	stack := arraystack.New()
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			stack.Push(v)
		case ')':
			temp, _ := stack.Pop()
			if temp == '(' {
				continue
			} else {
				return false
			}
		case ']':
			temp, _ := stack.Pop()
			if temp == '[' {
				continue
			} else {
				return false
			}
		case '}':
			temp, _ := stack.Pop()
			if temp == '{' {
				continue
			} else {
				return false
			}
		}
	}
	return stack.Empty()
}

```

## 1047.删除字符串中的所有相邻重复项

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/)
>
> 文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/1047.%E5%88%A0%E9%99%A4%E5%AD%97%E7%AC%A6%E4%B8%B2%E4%B8%AD%E7%9A%84%E6%89%80%E6%9C%89%E7%9B%B8%E9%82%BB%E9%87%8D%E5%A4%8D%E9%A1%B9.html)
>
> 视频讲解：[栈的好戏还要继续！| LeetCode：1047. 删除字符串中的所有相邻重复项](https://www.bilibili.com/video/BV12a411P7mw)
>
> 状态：AC

### 思路

将每个元素与栈顶元素进行比较（前提是栈非空、如果空直接入栈即可），比较相等元素出栈。

### 代码

``` go
func removeDuplicates(s string) string {
	if len(s) == 1 {
		return s
	}
	byteString := []rune(s)
	stack := arraystack.New()
	stack.Push(byteString[0])

	for i := 1; i < len(byteString); i++ {
		temp, _ := stack.Peek()
		if temp == byteString[i] {
			stack.Pop()
		} else {
			stack.Push(byteString[i])
		}
	}
	var newByte []rune
	length := stack.Size()
	for i := 0; i < length; i++ {
		temp, _ := stack.Pop()
		temp1 := temp.(rune)
		newByte = append(newByte, temp1)
	}
	for i := 0; i < len(newByte)/2; i++ {
		newByte[i], newByte[len(newByte)-1-i] = newByte[len(newByte)-1-i], newByte[i]
	}
	return string(newByte)
}
```

此处用的是一个真正意义上的数据结构的栈，存储的是单个字符，写起来还是很长，看了题解之后发现只用字符数组即可，改正后代码

``` go
func removeDuplicates(s string) string {
	var stack []rune
	for _, v := range s {
		if len(stack) > 0 && stack[len(stack)-1] == v {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return string(stack)
}
```

## 150.逆波兰表达式

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/evaluate-reverse-polish-notation/)
>
> 文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0150.%E9%80%86%E6%B3%A2%E5%85%B0%E8%A1%A8%E8%BE%BE%E5%BC%8F%E6%B1%82%E5%80%BC.html)
>
> 视频讲解：栈的最后表演！ | LeetCode：150. 逆波兰表达式求值](https://www.bilibili.com/video/BV1kd4y1o7on)
>
> 状态：AC

### 思路

明白题意之后，发现每次到运算符时候需要将前面两个数进行计算，如：`["a", "b", "+"]`运算结果是`a+b`。不难想到要使用栈。

先将字符串数组转换为数字，可以成功转换则入栈，否则出栈两个元素，并且将两个元素相加之后再重新入栈。

### 代码

``` go

func EvalRPN(tokens []string) int {
	stack := arraystack.New()
	for _, v := range tokens {
		num, err := strconv.Atoi(v)
		if err == nil {  // 如果成功转换数字
			stack.Push(num)
		} else {  // 否则计算后重新入栈
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			if v == "+" {
				stack.Push(a.(int) + b.(int))
			} else if v == "-" {
				stack.Push(a.(int) - b.(int))
			} else if v == "*" {
				stack.Push(a.(int) * b.(int))
			} else if v == "/" {
				stack.Push(a.(int) / b.(int))
			}
		}
	}
	result, _ := stack.Peek()  // 此刻的栈顶为结果
	return result.(int)
}

```

