---
title: 代码随想录算法训练营第六天| 242.有效的字母异位词、349. 两个数组的交集、202. 快乐数、1. 两数之和 。
tags:
  - 算法
  - 代码随想录
  - LeetCode
categories: 刷题
abbrlink: 20198d61
date: 2023-09-25 10:09:40
---

# 代码随想录算法训练营第六天| 242.有效的字母异位词、349. 两个数组的交集、202. 快乐数、1. 两数之和 。

## 242.有效的字母异位词

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/valid-anagram/)
>
>   文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0242.%E6%9C%89%E6%95%88%E7%9A%84%E5%AD%97%E6%AF%8D%E5%BC%82%E4%BD%8D%E8%AF%8D.html)
>
>   视频讲解：[学透哈希表，数组使用有技巧！Leetcode：242.有效的字母异位词](https://www.bilibili.com/video/BV1YG411p7BA)
>
>   状态：AC

### 思路

-   方法一：使用`map`即可。键是字母的`ASCII`码、值为频率。分别用两个串构建两个`map`，再比较两个`map`是否相同。

-   方法二：看了卡哥给的代码发现优化空间还是很大的，使用一个`map`即可，存放`s`的情况。异位词满足两个条件：
    -   两个串等长。
    -   两个串中字母出现的频率相同。这一点可以用`s`串的`map`的`value`自减。

### 代码

``` go
//麻烦的方法
func isAnagram(s string, t string) bool {
	words1 := map[byte]int{}
	words2 := map[byte]int{}
	for i := 0; i < len(s); i++ {
		_, ok := words1[s[i]]
		if ok {
			words1[s[i]] += 1
		} else {
			words1[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		_, ok := words2[t[i]]
		if ok {
			words2[t[i]] += 1
		} else {
			words2[t[i]] = 1
		}
	}

	return reflect.DeepEqual(words1, words2)
}
```

``` go
// 简单的代码
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}
```



## 349. 两个数组的交集

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/intersection-of-two-arrays/)
>
>   文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0349.%E4%B8%A4%E4%B8%AA%E6%95%B0%E7%BB%84%E7%9A%84%E4%BA%A4%E9%9B%86.html)
>
>   视频讲解：[学透哈希表，set使用有技巧！Leetcode：349. 两个数组的交集](https://www.bilibili.com/video/BV1ba411S7wu)
>
>   状态：AC

### 思路

从求交集可以看出是要使用集合的结构，但是在Go语言中<font color="#ff0000">并没有</font>集合，所以我使用了`map`（并不是用map模拟set）。map中，键为`nums1`的每个元素，值只有三种状态：第一次出现是`0`，多次出现是`1`，在`nums2`也出现过是`2`，最后遍历map，找到值为2的key。

利用const和iota模拟枚举类型。

### 代码

``` go
const (
	NEW = iota
	EXIST1
	EXIST2
)

func intersection(nums1 []int, nums2 []int) []int {
	set := map[int]int{}  //类集合操作
	for _, v := range nums1 {
		if _, ok := set[v]; !ok {
			set[v] = NEW  // 第一个出现
		} else {  // 多次出现
			set[v] = EXIST1
		}
	}
	for _, v := range nums2 {
		if _, ok := set[v]; ok {  // nums2也出现
			set[v] = EXIST2
		}
	}
	var result []int
	for k, v := range set {
		if v == EXIST2 {
			result = append(result, k)
		}
	}
	return result
}
```

## 202. 快乐数

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/happy-number/)
>
>   文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0202.%E5%BF%AB%E4%B9%90%E6%95%B0.html#%E6%80%9D%E8%B7%AF)
>
>   状态：AC

### 思路

1.   构造一个集合，用于存放每一次拆数求平方和的结果
2.   拆数
3.   判断结果是否是1，如果是`return true`
4.   如果不是1，判断结果是否在集合中，如果存在则说明出现了循环，`return false`



### 代码

``` go
func isHappy(n int) bool {
	set := map[int]bool{}  // 模拟集合
	for {
		sum := 0
        //拆数求平方和
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
        //得到正确答案
		if sum == 1 {
			return true
		}
        
        //结果不在集合中则放入集合
		if _, ok := set[sum]; !ok {
			set[sum] = true
		} else {// 否则退出循环 return false
			break
		}
		n = sum // 新一轮的数
	}
	return false
}
```

## 1.两数之和

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/two-sum/)
>
>   文章讲解：[代码随想录(programmercarl.com)](https://programmercarl.com/0001.%E4%B8%A4%E6%95%B0%E4%B9%8B%E5%92%8C.html)
>
>   视频讲解：[梦开始的地方，Leetcode：1.两数之和](https://www.bilibili.com/video/BV1aT41177mK)
>
>   状态：AC

### 思路

1.   遍历`nums`将元素存为`map`的`key`，将元素的索引存放为`map`的`value`；
2.   再遍历`nums`，查看`target - v`是否在`key`中，并返回两个值：`nums`元素的索引和`map[target-v]`

遍历两次还是思路上走弯路了，一次即可。

### 代码

```go
//遍历两次
func twoSum(nums []int, target int) []int {
    mapNums := map[int]int{}
    for i, v := range nums {
       mapNums[v] = i
    }
    var result []int
    for i, v := range nums {
       if _, ok := mapNums[target-v]; ok && mapNums[target-v] != i {
          result = append(result, i)
          result = append(result, mapNums[target-v])
          break
       }
    }
    return result
}
```

``` go
func twoSum(nums []int, target int) []int {
	mapNums := map[int]int{}
	for i, v := range nums {
		if value, ok := mapNums[target-v]; ok {
			return []int{i, value}
		} else {
			mapNums[v] = i
		}
	}
	return []int{}
}
```

## 小结

-   这次的题目不算很难，做起来就很快，算上写博客的时间不到两个小时

-   对Go语言的map有了新的理解：

    ``` go
    test = map[int]int{}
    if value, ok := test[key]; ok {
    	// key存在的情况
    } else {
    	// key不存在的情况
    }
    ```

    ``` go
    // 使用反射来判断两个map是否相等
    import reflect
    reflect.DeepEqual(map1, map2)
    ```

    ``` go
    test = map[int]int{}
    test[1]--  // 即使不存在1，默认新增键值对{1:0}，然后再自减
    ```

    

-   第一次使用Go模拟枚举

    ``` go
    const (
    	A iota
        B
        C
    )
    ```

    
