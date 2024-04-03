---
title: 【LeetCode】09动态规划Part03  343. 整数拆分、96. 不同的二叉搜索树
tags:
  - 算法
  - 动态规划
  - LeetCode
mathjax: ture
abbrlink: 9e71a6f7
date: 2024-03-24 19:37:38
---

# 【LeetCode】09动态规划Part03  343. 整数拆分、96. 不同的二叉搜索树



## 343. 整数拆分

>   题目链接：https://leetcode.cn/problems/integer-break/description/

### 思路

#### 思路一：贪心算法

根据自己的不断地测试发现，如果一个数n(n>3)，不断地拆出来3，拆到最后会出现3种情况：

1.   `n % 3 == 0`：最大的乘积数为`pow(3, n / 3)`
2.   `n % 3 == 2`：意味着拆到最后一个数为2，那么最大的乘积数是`pow(3, n / 3) * 2`
3.   `n % 3 == 1`：意味着拆到最后一个数为1，那么少拆一个3，使其最后乘以4，最大乘积数为`pow(3, n / 3 - 1) * 4`

确实也AC了这道题目，核心思想就是贪更多的3。至于为什么并不是很理解，也没有证明

但是也尝试过很多策略

-   不断地拆出来2，如果是奇数最后一个3则不拆，这样会获得最大的幂指数。但是得不到最大的乘积
-   不断地拆出来`a = (int)sqrt(n)`，最后一次不足a则加在上一次拆分的结果。同样得不到最大的值

#### 思路二：动态规划

##### 1. 确定dp数组

正整数`i`拆分后得到的最大乘积`dp[i]`。可能是两个数乘积，也可能是多个数的乘积。



##### 2. 递推公式

-   将`i`拆分成`j`和`i - j `，且`i - j`不能继续拆分，此时乘积为`j x (i - j)`
-   将`i`拆分成`j`和`i - j `，且`i - j`可以继续拆分，此时乘积为`j x dp[i - j]`

$$
dp[i] = \max\limits_{1\le j < i}\{\max(j\times (i - j), j \times dp[i - j])\}
$$

>   以n = 6为例
>
>   -   `i = 1`，此时`i`并不能拆分出两个正整数，`dp[1] = 0`
>   -   `i = 2`，此时`i`仅能拆分出$2 = 1 + 1 \Rightarrow 1\times 1=1$，`dp[2] = 1`
>   -   `i = 3`，令`curMax = 0`代表当前拆分出的$\max(j\times (i - j), j \times dp[i - j])$，
>       $$
>       \begin{align*}
>       &j=1\Rightarrow
>       \begin{cases}
>       j\times (i-j)=1\times (3-1) = 2\\\\
>       j\times dp[i-j]=1\times dp[3-1]=1\times 1 = 1
>       \end{cases}\\\\
>       
>       &j=2\Rightarrow
>       \begin{cases}
>       j\times (i-j)=2\times (3-2) = 2\\\\
>       j\times dp[i-j]=2\times dp[3-2]=2\times 0 = 0
>       \end{cases}
>       \end{align*}
>       $$
>       综上，`curMax = 2, dp[3] = 2`，也就是当拆分数字3时，最大乘积为2
>   -   `i = 4`：
>       $$
>       \begin{align*}
>       &j=1\Rightarrow
>       \begin{cases}
>       j\times (i-j)=1\times (4-1) = 3\\\\
>       j\times dp[i-j]=1\times dp[4-1]=1\times 2 = 2
>       \end{cases}\\\\
>       
>       &j=2\Rightarrow
>       \begin{cases}
>       j\times (i-j)=2\times (4-2) = 4\\\\
>       j\times dp[i-j]=2\times dp[4-2]=2\times 1 = 2
>       \end{cases}\\\\
>       
>       &j=3\Rightarrow
>       \begin{cases}
>       j\times (i-j)=3\times (4-3) = 3\\\\
>       j\times dp[i-j]=3\times dp[4-3]=3\times 0 = 0
>       \end{cases}
>       \end{align*}
>       $$
>       综上，`curMax = 4, dp[4] = 4`，也就是当拆分数字4时，最大乘积为4
>   -   `i = 5`：
>       $$
>       \begin{align*}
>       &j=1\Rightarrow
>       \begin{cases}
>       j\times (i-j)=1\times (5-1) = 4\\\\
>       j\times dp[i-j]=1\times dp[5-1]=1\times 4 = 4
>       \end{cases}\\\\
>       
>       &j=2\Rightarrow
>       \begin{cases}
>       j\times (i-j)=2\times (5-2) = 6\\\\
>       j\times dp[i-j]=2\times dp[5-2]=2\times 2 = 4
>       \end{cases}\\\\
>       
>       &j=3\Rightarrow
>       \begin{cases}
>       j\times (i-j)=3\times (5-3) = 6\\\\
>       j\times dp[i-j]=3\times dp[5-3]=3\times 1 = 3
>       \end{cases}\\\\
>       
>       &j=4\Rightarrow
>       \begin{cases}
>       j\times (i-j)=4\times (5-4) = 4\\\\
>       j\times dp[i-j]=4\times dp[5-4]=4\times 0 = 0
>       \end{cases}\\\\
>       \end{align*}
>       $$
>       综上，`curMax = 6, dp[5] = 6`，也就是当拆分数字5时，最大乘积为6
>   -   `i = 6`：
>       $$
>       \begin{align*}
>       &j=1\Rightarrow
>       \begin{cases}
>       j\times (i-j)=1\times (6-1) = 5\\\\
>       j\times dp[i-j]=1\times dp[6-1]=1\times 6 = 6
>       \end{cases}\\\\
>       
>       &j=2\Rightarrow
>       \begin{cases}
>       j\times (i-j)=2\times (6-2) = 8\\\\
>       j\times dp[i-j]=2\times dp[6-2]=2\times 4 = 8
>       \end{cases}\\\\
>       
>       &j=3\Rightarrow
>       \begin{cases}
>       j\times (i-j)=3\times (6-3) = 9\\\\
>       j\times dp[i-j]=3\times dp[6-3]=3\times 2 = 6
>       \end{cases}\\\\
>       
>       &j=4\Rightarrow
>       \begin{cases}
>       j\times (i-j)=4\times (6-4) = 8\\\\
>       j\times dp[i-j]=4\times dp[6-4]=4\times 1 = 4
>       \end{cases}\\\\
>       
>       &j=5\Rightarrow
>       \begin{cases}
>       j\times (i-j)=5\times (6-5) = 5\\\\
>       j\times dp[i-j]=5\times dp[6-5]=5\times 1 = 5
>       \end{cases}\\\\
>       \end{align*}
>       $$
>       综上，`curMax = 9, dp[6] = 9`，也就是当拆分数字6时，最大乘积为9
>
>   返回`dp[n]`

##### 3. dp数组初始化

不需要额外的初始化操作



##### 4. 确定遍历顺序

顺序遍历



##### 5. 打印dp数组（Debug）

```java
// leetcode官解
public int integerBreak(int n) {
    int[] dp = new int[n + 1];
    for (int i = 2; i <= n; i++) {
        int curMax = 0;
        for (int j = 1; j < i; j++) {
            curMax = Math.max(curMax, Math.max(j * (i - j), j * dp[i - j]));
        }
        dp[i] = curMax;
    }
    return dp[n];
}
```



## 96. 不同的二叉搜索树

>   题目链接：https://leetcode.cn/problems/unique-binary-search-trees/description/

自己画了几个树，还是没什么思路，看了题解。

### 动态规划

#### 1. 确定dp数组

`dp[i]`代表：`i`个节点的二叉搜索树有`dp[i]`种



#### 2. 递推公式

假设有3个节点的情况，需要分别考虑1、2、3分别为根节点时，左右子树的情况。

-   当1为根节点时，记做`F[1]`

    左子树为空树，右子树有两个节点，一共`dp[0] * dp[2]`种情况，使用乘法的原因是排列组合

-   当2为根节点时，记做`F[2]`

    左子树有一个节点，右子树有一个节点，一共`dp[1] * dp[1]`种情况

-   当3为根节点时，记做`F[3]`

    左子树有两个节点，右子树为空树，一共`dp[2]*dp[0]`种情况

所以`dp[3] = F[1] + F[2] + F[3]`

推广开来：
$$
\begin{align*}
&dp[i] = \sum\limits_{i=1}^nF[n]\\\\
&F[n]=\prod\limits_{i=0}^{n-1}dp[i]\times dp[n-1-i]
\end{align*}
$$

#### 3. 初始化

`dp[0] = 1`，空树也算一种情况

`dp[1] = 1`，只有一个节点的情况



#### 4. 遍历

顺序遍历



#### 5. 打印dp数组(debug)

``` java
public int numTrees(int n) {
    int[] dp = new int[n + 1];
    dp[0] = 1;
    dp[1] = 1;
    for (int i = 2; i <= n; i++) {
        dp[i] = 0;
        for (int j = 0; j < i; j++) {
            dp[i] += dp[j] * dp[i - 1 - j];
        }
    }
    return dp[n];
}
```

