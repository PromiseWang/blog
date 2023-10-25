---
title: 代码随想录算法训练营第三十三天  1005.K次取反后最大化的数组和、134.加油站、135.分发糖果
tags:
  - 算法
  - 代码随想录
  - LeetCode
  - 贪心算法
abbrlink: af9e53d4
date: 2023-10-25 17:14:53
---

# 代码随想录算法训练营第三十三天  1005.K次取反后最大化的数组和、134.加油站、135.分发糖果

## 1005.K次取反后最大化的数组和

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/1005.K%E6%AC%A1%E5%8F%96%E5%8F%8D%E5%90%8E%E6%9C%80%E5%A4%A7%E5%8C%96%E7%9A%84%E6%95%B0%E7%BB%84%E5%92%8C.html)
>
>   视频讲解：[贪心算法，这不就是常识？还能叫贪心？LeetCode：1005.K次取反后最大化的数组和](https://www.bilibili.com/video/BV138411G7LY)
>
>   状态：AC

### 思路

1.   对数组排序，统计负数个数`count`，遍历时再记录绝对值最小的数。
2.   如果`count >= k`，那么前k个数全都取反
3.   如果`count < k`，说明将负数全都取反后，还有额外的操作次数，判断该次数`k - count`是奇数还是偶数，如果是偶数不用操作，如果是奇数，翻转绝对值最小的数

#### 贪心算法

贪尽可能多的正数

### 代码

``` cpp
class Solution {
public:
    int largestSumAfterKNegations(vector<int> &nums, int k) {
        sort(nums.begin(), nums.end());
        int count = 0;  // 负数个数
        int minIndex = 0;
        for (int i = 0; i < nums.size(); ++i) {
            if (nums[i] < 0) {
                count++;
            }
            if (abs(nums[i]) < abs(nums[minIndex])) {
                minIndex = i;
            }
        }
        // count >= k 前k个负数全都翻
        if (count >= k) {
            for (int i = 0; i < k; ++i) {
                nums[i] = -nums[i];
            }
        } else {
            // count < k
            //   1. 所有负数都翻转
            //   2.1 if (k - count) % 2 == 0 输出
            //   2.2 else 翻转绝对值最小的数
            for (int i = 0; i < count; ++i) {
                nums[i] = -nums[i];
            }
            if ((k - count) % 2 == 1) {
                nums[minIndex] = -nums[minIndex];
            }
        }
        int sum = 0;
        for (auto i: nums) {
            sum += i;
        }
        return sum;
    }
};
```

## 134. 加油站

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/gas-station/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0134.%E5%8A%A0%E6%B2%B9%E7%AB%99.html)
>
>   视频讲解：[贪心算法，得这么加油才能跑完全程！LeetCode ：134.加油站](https://www.bilibili.com/video/BV1jA411r7WX)
>
>   状态：AC

### 思路

1.   将环路展开成直线的路，即将`gas`数组向后拼接一个自己，`cost`向后拼接一个自己，但是去掉一位。
2.   外层遍历`gas`，范围从0到`gas.size()/2`，索引为`i`。内层遍历`cost`，范围从`i`到`i+gas.sise()/2`
3.   定义当前的油量`cur`，如果`cur + gas[j] >= cost[j]`，那么下一次的剩余油量则为`cur += gas[j] - cost[j]`。否则从`i`开始的路不能走一圈，继续向下遍历

#### 贪心算法

从一点出发贪最远的路

### 代码

``` cpp
class Solution {
public:
    int canCompleteCircuit(vector<int> &gas, vector<int> &cost) {
        gas.insert(gas.end(), gas.begin(), gas.end());
        cost.insert(cost.end(), cost.begin(), cost.end() - 1);
        for (int i = 0; i < gas.size() / 2; ++i) {
            int cur = 0;
            bool flag = true;
            for (int j = i; j < i + gas.size() / 2; ++j) {
                if (cur + gas[j] >= cost[j]) {
                    cur += gas[j] - cost[j];
                } else {
                    flag = false;
                    i = j;
                    break;
                }
            }
            if (flag) {
                return i;
            }
        }
        return -1;
    }
};
```

## 135. 分发糖果

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/candy/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0135.%E5%88%86%E5%8F%91%E7%B3%96%E6%9E%9C.html)
>
>   视频讲解：[贪心算法，两者兼顾很容易顾此失彼！LeetCode：135.分发糖果](https://www.bilibili.com/video/BV1ev4y1r7wN)
>
>   状态：AC

### 思路

这道题没有想明白，看了题解后懂了。前后两次遍历评分数组。每个人糖果初始化为1。

前向遍历时如果后一项大于前一项，那么后一项的糖果值为前一项加1。

反向遍历时，如果前一项大于后一项时，并且前一项糖果也不超过后一项时，那么前一项糖果等于后一项加1。

### 代码

``` cpp
class Solution {
public:
    int candy(vector<int> &ratings) {
        vector<int> result(ratings.size(), 1);
        for (int i = 1; i < ratings.size(); ++i) {
            if (ratings[i] > ratings[i - 1]) {
                result[i] = result[i - 1] + 1;
            }
        }
        for (int i = ratings.size() - 1; i >= 1; --i) {
            if (ratings[i - 1] > ratings[i] && result[i - 1] <= result[i]) {
                result[i - 1] = result[i] + 1;
            }
        }
        int sum = 0;
        for (auto i: result) {
            sum += i;
        }
        return sum;
    }
};
```

