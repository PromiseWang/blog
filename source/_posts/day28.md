---
title: 代码随想录算法训练营第二十八天  491.递增子序列、46.全排列、47.全排列II
tags:
  - 算法
  - 代码随想录
  - LeetCode
  - 回溯算法
categories: 刷题
abbrlink: 70a8aaa7
date: 2023-10-20 23:59:09
---

# 代码随想录算法训练营第二十八天  491.递增子序列、46.全排列、47.全排列II

## 491.递增子序列

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/non-decreasing-subsequences/)
>
> 文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0491.%E9%80%92%E5%A2%9E%E5%AD%90%E5%BA%8F%E5%88%97.html)
>
> 视频讲解：[回溯算法精讲，树层去重与树枝去重 | LeetCode：491.递增子序列](https://www.bilibili.com/video/BV1EG4y1h78v/)
>
> 状态：AC

### 思路

先查找递增子序列，其次要去重。难点还是在于去重。

查找递增时候，每次需要判断当前的`nums[i]`是否大于`path[-1]`。

去重和之前的还不大一样，因为这里不可以进行排序。应该判断每层的情况

![去重图](https://code-thinking-1253855093.file.myqcloud.com/pics/20201124200229824.png)

### 代码

``` c++
class Solution {
private:
    vector<vector<int>> result;
    vector<int> path;

    void backTracking(vector<int> &nums, int index) {
        if (path.size() > 1) {
            result.push_back(path);
        }
        unordered_set<int> sets;
        for (int i = index; i < nums.size(); ++i) {
            if ((!path.empty() && path.back() > nums[i]) || sets.find(nums[i]) != sets.end()) {
                continue;
            }
            path.push_back(nums[i]);
            sets.insert(nums[i]);
            backTracking(nums, i + 1);
            path.pop_back();
        }
    }

public:
    vector<vector<int>> findSubsequences(vector<int> &nums) {
        backTracking(nums, 0);
        return result;
    }
};
```

## 46. 全排列

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/permutations/)
>
> 文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0046.%E5%85%A8%E6%8E%92%E5%88%97.html)
>
> 视频讲解：[组合与排列的区别，回溯算法求解的时候，有何不同？| LeetCode：46.全排列](https://www.bilibili.com/video/BV19v4y1S79W/)
>
> 状态：AC

### 思路

#### 递归三部曲

保存使用过的数，这里保存的实际上是该数的位置。用的二进制来节省空间。例如，使用了第0个数，括号里指的是该数为二进制，`used=1(2)`。使用了第0、2个数，那么则是`used=101(2)`，使用了第0、1、3、6个数，`used=1001011(2)`。即使用了第i个数，`used += 1 << i`。判断第i个数是否使用过：`1 >> i & used`

1. 确立函数和返回值
    - `nums`
    - `used`，使用二进制保存
2. 递归终止条件：`path`长度等于 `nums`长度
3. 单次递归：如果当前的数没被使用过，加入`path`

### 代码

``` c++
class Solution {
private:
    vector<vector<int>> result;
    vector<int> path;

    void backTracking(vector<int> &nums, int used) {
        if (nums.size() == path.size()) {
            result.push_back(path);
            return;
        }
        for (int i = 0; i < nums.size(); ++i) {
            if (used >> i & 1) {
                continue;
            }
            path.push_back(nums[i]);
            used += 1 << i;
            backTracking(nums, used);
            path.pop_back();
            used -= 1 << i;
        }
    }

public:
    vector<vector<int>> permute(vector<int> &nums) {
        int used = 0;
        backTracking(nums, used);
        return result;
    }
};
```

## 47. 全排列II

> 题目链接：[力扣题目链接](https://leetcode.cn/problems/permutations-ii/)
>
> 文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0047.%E5%85%A8%E6%8E%92%E5%88%97II.html)
>
> 视频讲解：[回溯算法求解全排列，如何去重？| LeetCode：47.全排列 II](https://www.bilibili.com/video/BV1R84y1i7Tm/)
>
> 状态：AC

### 思路

力扣40题[个人博客跳转](https://promisewang.github.io/post/36f15375.html#40-%E7%BB%84%E5%90%88%E6%80%BB%E5%92%8CII) + 上一题。used还是使用数组表示。

### 代码

``` c++
class Solution {
private:
    vector<vector<int>> result;
    vector<int> path;

    void backTracking(vector<int> &nums, vector<bool> used2) {
        if (nums.size() == path.size()) {
            result.push_back(path);
        }
        for (int i = 0; i < nums.size(); ++i) {

            if (i > 0 && nums[i - 1] == nums[i] && !used2[i - 1]) {
                continue;
            }
            if (!used2[i]) {
                used2[i] = true;
                path.push_back(nums[i]);
                backTracking(nums, used2);
                path.pop_back();
                used2[i] = false;
            }
        }
    }

public:
    vector<vector<int>> permuteUnique(vector<int> &nums) {
        sort(nums.begin(), nums.end());
        vector<bool> used2(nums.size(), false);
        backTracking(nums, used2);
        return result;
    }
};
```

