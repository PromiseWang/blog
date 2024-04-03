---
title: 【LeetCode】09动态规划Part02  62. 不同路径、63. 不同路径II
tags:
  - 算法
  - LeetCode
  - 动态规划
abbrlink: e9769661
date: 2024-03-21 14:59:23
---

# 【LeetCode】09动态规划Part02  62. 不同路径、63. 不同路径II



## 62. 不同路径

>   题目链接：https://leetcode.cn/problems/unique-paths/description/

### 思路

到达某个非边缘块（第一行或第一列）只有两种可能，要么从其上方来，要么从其左侧来。那么到达这个块一共就有`dp[i][j] = dp[i - 1][j] + dp[i][j - 1]`个方法。如果是第一行或第一列只可能有一种方法。

#### 1. 确定dp数组

二维dp，`dp[i][j]`代表第i行第j列这个位置有多少种方法可以到达

#### 2. 递推公式

`dp[i][j] = dp[i - 1][j] + dp[i][j - 1]`

#### 3. dp数组的初始化

第一行、第一列均为1

#### 4. 确定遍历顺序

先行后列



### 代码

``` java
class Solution {
public:
    public int uniquePaths(int m, int n) {
        int[][] dp = new int[m][n];
        for (int i = 0; i < m; i++) {
            dp[i][0] = 1;
        }
        for (int i = 0; i < n; i++) {
            dp[0][i] = 1;
        }
        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
            }
        }
        return dp[m - 1][n - 1];
    }
}
```



## 63. 不同路径II

### 思路

思路基本同上，但是多了一层障碍，需要多一层判断，如果位置`(i, j)`有障碍（即`obstacleGrid[i][j] == 1`），那么`dp[i][j]`的值为0

#### 1. 确定dp数组

二维dp，`dp[i][j]`代表第i行第j列这个位置有多少种方法可以到达

#### 2. 递推公式

`dp[i][j] = dp[i - 1][j] + dp[i][j - 1]`

#### 3. dp数组的初始化

第一行、第一列在没有障碍的位置均为1

#### 4. 确定遍历顺序

先行后列



### 代码

``` java
class Solution {
    public int uniquePathsWithObstacles(int[][] obstacleGrid) {
        int n = obstacleGrid.length;
        int m = obstacleGrid[0].length;
        int[][] dp = new int[n][m];
        int flag = 1;
        for (int i = 0; i < n; i++) {
            if (obstacleGrid[i][0] == 1) {
                flag = 0;
            }
            dp[i][0] = flag;
        }
        flag = 1;
        for (int i = 0; i < m; i++) {
            if (obstacleGrid[0][i] == 1) {
                flag = 0;
            }
            dp[0][i] = flag;
        }

        for (int i = 1; i < n; i++) {
            for (int j = 1; j < m; j++) {
                dp[i][j] = obstacleGrid[i][j] == 0 ? dp[i - 1][j] + dp[i][j - 1] : 0;
            }
        }
        return dp[n - 1][m - 1];
    }
};
```

