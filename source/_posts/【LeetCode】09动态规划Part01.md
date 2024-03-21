---
title: 【LeetCode】09动态规划Part01
tags:
  - 算法
  - LeetCode
  - 动态规划
abbrlink: 707fc7db
date: 2024-03-20 15:50:08
---

# 【LeetCode】09动态规划Part01  509.斐波那契数、70. 爬楼梯、746. 使用最小花费爬楼梯



## 动态规划解题步骤

1.   确定dp数组

2.   递推公式

3.   dp数组的初始化

4.   确定遍历顺序

5.   打印dp数组(Debug)



## 509. 斐波那契数

>   题目链接：https://leetcode.cn/problems/fibonacci-number/description/

### 思路

1.   确定dp数组

     `dp[i]`是前两项之和

2.   递推公式

     `dp[i] = dp[i - 1] + dp[i - 2]`

3.   dp数组的初始化

     `dp[0] = 0`
     `dp[1] = 1`
     `dp[i] = 0`

4.   确定遍历顺序

     顺序遍历

5.   打印dp数组(Debug)



### 代码

``` cpp
class Solution {
public:
    int fib(int n) {
        if (n < 2) 
            return n;
        vector<int> dp(n + 1, 0);
        dp[0] = 0;
        dp[1] = 1;
        for (int i = 2; i < n + 1; i++) {
            dp[i] = dp[i - 1] + dp[i - 2];
        }
        return dp[n];
    }
};
```



## 70. 爬楼梯

>   题目链接：https://leetcode.cn/problems/climbing-stairs/description/

### 思路

1.   确定dp数组

     第i个台阶有dp[i]种方法可以上来

2.   递推公式

     `dp[i] = dp[i - 1] + dp[i - 2]`

3.   dp数组的初始化

     `dp[0] = 0`
     
     `dp[1] = 1`
     
     `dp[2] = 2`

     `dp[i] = 0`

4.   确定遍历顺序

     顺序遍历

5.   打印dp数组(Debug)



### 代码

``` cpp
class Solution {
public:
    int climbStairs(int n) {
        if (n <= 2) {
            return n;
        }
        vector<int> dp(n + 1, 0);
        dp[1] = 1;
        dp[2] = 2;
        for (auto i = 3; i < n + 1; i++) {
            dp[i] = dp[i - 1] + dp[i - 2];
        }
        return dp.back();
    }
};
```





## 746. 使用最小花费爬楼梯

>   题目链接：https://leetcode.cn/problems/min-cost-climbing-stairs/description/

### 思路

这道题描述看的不是很懂，理解之后还是有点模糊这个递推公式，AC之后再看看题解。先判断是否是最后一次上台阶（站在最后一个台阶上是，选择一步还是两步）



1.   确定dp数组

     第i个台阶最小花费为dp[i]

2.   递推公式

     如果是最后一次上台阶：则`dp[i] = min(dp[i - 1], cost[i] + dp[i - 2]);`

     否则：`dp[i] = cost[i] + min(dp[i - 1], dp[i - 2]);`

3.   dp数组的初始化

     `dp[0] = cost[0]`

     `dp[1] = cost[1]`

4.   确定遍历顺序

     顺序遍历

5.   打印dp数组(Debug)

### 代码

``` cpp
class Solution {
public:
    int minCostClimbingStairs(vector<int> &cost) {
        vector<int> dp(cost.size(), 0);
        dp[0] = cost[0];
        dp[1] = cost[1];
        if (cost.size() == 2) {
            return min(dp[0], dp[1]);
        }
        for (int i = 2; i < dp.size(); ++i) {
            if (i == dp.size() - 1) {
                dp[i] = min(dp[i - 1], cost[i] + dp[i - 2]);
            }
            else {
                dp[i] = cost[i] + min(dp[i - 1], dp[i - 2]);
            }
        }
        return dp.back();
    }
};
```



### 官方题解

``` cpp
class Solution {
public:
    int minCostClimbingStairs(vector<int>& cost) {
        int n = cost.size();
        vector<int> dp(n + 1);
        dp[0] = dp[1] = 0;
        for (int i = 2; i <= n; i++) {
            dp[i] = min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2]);
        }
        return dp[n];
    }
};

```

主要是对状态转移方程理解的不够透彻

`dp[i] = min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2]);`

