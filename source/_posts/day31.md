---
title: 代码随想录算法训练营第三十一天  455.分发饼干、376.摆动序列、53.最大子数组和
tags:
  - 算法
  - 代码随想录
  - LeetCode
  - 贪心算法
categories: 刷题
abbrlink: 5aa98f1b
date: 2023-10-25 03:12:41
---

# 代码随想录算法训练营第三十一天  455.分发饼干、376.摆动序列、53.最大子数组和

## 455. 分发饼干

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/assign-cookies/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0455.%E5%88%86%E5%8F%91%E9%A5%BC%E5%B9%B2.html)
>
>   视频讲解：[贪心算法，你想先喂哪个小孩？| LeetCode：455.分发饼干](https://www.bilibili.com/video/BV1MM411b7cq)
>
>   状态：AC

### 思路

先将两个数组进行排序。先遍历胃口`g`，再遍历`s`，如果`g[i] <= s[j]`时，那么可以将`s[j]`喂给`g[i]`，记录此时`s`的位置放到`index`，下次`s`从`index`开始循环。

#### 贪心思路

充分利用饼干喂饱当前的孩子。

### 代码

``` c++
class Solution {
public:
    int findContentChildren(vector<int> &g, vector<int> &s) {
        sort(g.begin(), g.end());
        sort(s.begin(), s.end());
        int index = 0;
        if (s.empty()) {
            return 0;
        }
        int count = 0;
        for (int i = 0; i < g.size() && index < s.size(); ++i) {
            for (int j = index; j < s.size(); ++j) {
                if (s[j] >= g[i]) {
                    index = j + 1;
                    count++;
                    break;
                }
            }
        }
        return count;
    }
};
```

看了题解之后发现从后向前遍历写起来更简单一些，而且只需要一层循环即可。

## 376.摆动序列

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/wiggle-subsequence/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0376.%E6%91%86%E5%8A%A8%E5%BA%8F%E5%88%97.html)
>
>   视频讲解：[贪心算法，寻找摆动有细节！| LeetCode：376.摆动序列](https://www.bilibili.com/video/BV17M411b7NS)
>
>   状态：无

### 思路

<font color="red">错误：这道题一直想的只是当前位置和前后数字进行比较，如果高于或低于旁边的两个数字判定为摆动。但是并没有思考到上一次的振幅改变不在旁边的情况，所以不能只和身边的两个数进行比较</font>

定义两个差：当前差`cur`和上次的差`pre`。上次差意味着出现了连续递增或递减，比较两次差是否摆动

#### 贪心思路

递增或递减时贪最远的情况

### 代码

``` c++
class Solution {
public:
    int wiggleMaxLength(vector<int> &nums) {
        if (nums.size() < 2) {
            return nums.size();
        }
        int cur = 0;
        int pre = 0;
        int count = 1;
        for (int i = 0; i < nums.size() - 1; ++i) {
            cur = nums[i + 1] - nums[i];
            if ((pre <= 0 && cur > 0) || (pre >= 0 && cur < 0)) {
                count++;
                pre = cur;
            }
        }
        return count;
    }
};
```

## 53.最大子数组和

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/maximum-subarray/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0053.%E6%9C%80%E5%A4%A7%E5%AD%90%E5%BA%8F%E5%92%8C.html)
>
>   视频讲解：[贪心算法的巧妙需要慢慢体会！LeetCode：53. 最大子序和](https://www.bilibili.com/video/BV1aY4y1Z7ya)
>
>   状态：AC

### 思路

贪心算法没太想懂，使用动态规划还是很简单的。

#### 动态规划：

定义`dp`数组，`dp[i]`代表前`i`个元素最大和，如果前n项和小于0，重新开始，否则累加。`dp[i] = max(dp[i-1] + nums[i], nums[i])`

#### 代码

``` c++
class Solution {
public:
    int maxSubArray(vector<int> &nums) {
        if (nums.size() < 2) {
            return nums[0];
        }
        vector<int> dp(nums.size(), 0);
        dp[0] = nums[0];
        int maxSum = dp[0];
        for (int i = 1; i < nums.size(); ++i) {
            dp[i] = max(dp[i - 1] + nums[i], nums[i]);
            if (dp[i] > maxSum) {
                maxSum = dp[i];
            }
        }
        return maxSum;
    }
};
```

#### 贪心算法

看了卡哥的讲解：如果连续和小于0，那么重新开始（因为负数相加会越来越小）。否则继续累加。

``` c++
class Solution {
public:
    int maxSubArray(vector<int> &nums) {
        if (nums.size() < 2) {
            return nums[0];
        }
        int sum = 0;
        int max = INT32_MIN;
        for (int i = 0; i < nums.size(); ++i) {
            sum += nums[i];
            if (sum > max) {
                max = sum;
            }
            if (sum < 0) {
                sum = 0;
            }
        }
        return max;
    }
};
```

