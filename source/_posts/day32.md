---
title: 代码随想录算法训练营第三十二天  122. 买卖股票的最佳时机II、55.跳跃游戏、45.跳跃游戏II
tags:
  - 算法
  - 代码随想录
  - LeetCode
  - 贪心算法
abbrlink: 4486d5e0
date: 2023-10-25 03:45:55
---

# 代码随想录算法训练营第三十二天  122. 买卖股票的最佳时机II、55.跳跃游戏、45.跳跃游戏II

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0122.%E4%B9%B0%E5%8D%96%E8%82%A1%E7%A5%A8%E7%9A%84%E6%9C%80%E4%BD%B3%E6%97%B6%E6%9C%BAII.html)
>
>   视频讲解：[贪心算法也能解决股票问题！LeetCode：122.买卖股票最佳时机 II](https://www.bilibili.com/video/BV1ev4y1C7na)
>
>   状态：AC

### 思路

这道题已经给出了一定的提示

>   输入：prices = [7,1,5,3,6,4]
输出：7
解释：
   
   在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4 。
   
   随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3 。  
   
   总利润为 4 + 3 = 7 。

利润是可以叠加计算的。

#### 贪心算法

假设每一天都买股票，第二天都卖股票，计算出第二天的利润，将所有正利润加和即可。

### 代码

``` c++
class Solution {
public:
    int maxProfit(vector<int> &prices) {
        vector<int> profit(prices.size(), 0);
        int result = 0;
        for (int i = 1; i < prices.size(); ++i) {
            if (prices[i] - prices[i - 1] > 0) {
                result += prices[i] - prices[i - 1];
            }
        }
        return result;
    }
};
```

## 55. 跳跃游戏

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/jump-game/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0055.%E8%B7%B3%E8%B7%83%E6%B8%B8%E6%88%8F.html)
>
>   视频讲解：[贪心算法，怎么跳跃不重要，关键在覆盖范围 | LeetCode：55.跳跃游戏](https://www.bilibili.com/video/BV1VG4y1X7kB)
>
>   状态：AC

### 思路

定义距离`dis`，每一格都走最远的路，和`dis`做比较，如果更远更新`dis`，最后比较`dis`能否到`nums.size()-1`

#### 贪心算法

每一格都贪最远的路

### 代码

``` cpp
class Solution {
public:
    bool canJump(vector<int> &nums) {
        int dis = 0;
        if (nums.size() == 1) {
            return true;
        }
        for (int i = 0; i <= dis; ++i) {
            dis = max(nums[i] + i, dis);
            if (dis >= nums.size() - 1) {
                return true;
            }
        }
        return false;
    }
};
```

## 45. 跳跃游戏II

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/jump-game-ii/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0045.%E8%B7%B3%E8%B7%83%E6%B8%B8%E6%88%8FII.html)
>
>   视频讲解：[贪心算法，最少跳几步还得看覆盖范围 | LeetCode： 45.跳跃游戏 II](https://www.bilibili.com/video/BV1Y24y1r7XZ)
>
>   状态：投降

### 思路

这道题最终没有做出来。看了题解。

先尽量远走，如果达不到则向下走一步再向远走。需要两个变量分别记录当前最远距离位置`curDis`和下一次的最远距离位置`nextDis`。

如果`curDis`为当前的位置，则更新成`nextDis`

### 代码

``` cpp
class Solution {
public:
    int jump(vector<int> &nums) {
        vector<int> dp(nums.size(), 0);
        if (nums.size() == 1) {
            return 0;
        }
        int count = 0;
        int curDis = 0;
        int nextDis = 0;
        for (int i = 0; i < nums.size(); ++i) {
            nextDis = max(i + nums[i], nextDis);
            if (i == curDis) {
                curDis = nextDis;
                count++;
                if (nextDis >= nums.size() - 1) {
                    break;
                }
            }
        }
        return count;
    }
};
```

<font color="red">这道题值得反复思考，暂时想通了，可能睡一觉起来又忘了</font>
