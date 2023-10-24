---
title: 代码随想录算法训练营第二十七天  93.复原IP地址、78.子集、90.子集II
tags:
  - 算法
  - 代码随想录
  - LeetCode
  - 回溯算法
categories: 刷题
abbrlink: ff1e328
date: 2023-10-20 02:43:48
---

# 代码随想录算法训练营第二十七天  93.复原IP地址、78.子集、90.子集II

## 93.复原IP地址

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/restore-ip-addresses/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0093.%E5%A4%8D%E5%8E%9FIP%E5%9C%B0%E5%9D%80.html)
>
>   视频讲解：[回溯算法如何分割字符串并判断是合法IP？| LeetCode：93.复原IP地址](https://www.bilibili.com/video/BV1XP4y1U73i/)
>
>   状态：AC

### 思路

因为IP地址有一定的局限性（在0\~255范围内），这道题我采用了两种方法，一种直接暴力算法，三重循环。另一种使用回溯。题目中给定字符串`s`为潜在的IP地址，其长度在[1, 20]。那么他的合法长度在[4, 12]之间，可以先做一些判定。无论是**暴力算法**还是**回溯算法**，查找的均为字符串的分割处。

#### 1. 暴力算法

<font color="red">数学的范围都为<b>闭区间</b>，字符串的范围均为<b>左闭右开</b>！！！</font>

<font color="red">数学的范围都为<b>闭区间</b>，字符串的范围均为<b>左闭右开</b>！！！</font>

<font color="red">数学的范围都为<b>闭区间</b>，字符串的范围均为<b>左闭右开</b>！！！</font>

-   第一层范围[1, 3]，使用指针`i`
    -   分割出字符串`s1 = s[0: i]`，判断是否有前导0，判断转换成数字`n1`后是否在[0, 255]之间
-   第二层范围[`i`+1, 6]，使用指针`j`
    -   分割出字符串`s2 = s[i: j]`，判断是否有前导0，判断长度是否大于3，判断转换成数字`n2`后是否在[0, 255]之间
-   第三层范围[`j`+1, 9]，使用指针`k`
    -   分割出字符串`s3 = s[j: k]`，判断是否有前导0，判断长度是否大于3，判断转换成数字`n3`后是否在[0, 255]之间
    -   分割出字符串`s4 = s[k:]`，判断是否有前导0，判断长度是否大于3，判断转换成数字`n4`后是否在[0, 255]之间

均合法，则拼接IP地址`n1 + "." + n2 + "." + n3 + "." + n4`

#### 代码

``` c++

class Solution {
private:
    vector<string> result;
    string path;

    void backTracking(string s, int index, int count) {
        if (count > 3) {
            return;
        }
    }

    int getNum(string s) {
        int num = 0;
        if (s == "0") {
            return 0;
        }
        if (s[0] - '0' == 0 && s.size() > 1 || s.size() == 0) {
            return -1;
        }
        for (char i : s) {
            num *= 10;
            num += (i - '0');
        }

        return num;
    }

public:
    vector<string> restoreIpAddresses(string s) {
        if (s.size() < 4) {
            return result;
        }
        int n1, n2, n3, n4;
        for (int i = 1; i <= 3; ++i) {
            n1 = getNum(s.substr(0, i));
            if (0 <= n1 && n1 <= 255) {
                for (int j = i + 1; j <= 6 || j < s.size() - 1; ++j) {
                    if (j - i > 3) {
                        break;
                    }
                    n2 = getNum(s.substr(i, j - i));
                    if (0 <= n2 && n2 <= 255) {
                        for (int k = j + 1; k <= 9 || k < s.size(); ++k) {
                            if (k - j > 3) {
                                break;
                            }
                            if (s.size() - k > 3) {
                                continue;
                            }
                            n3 = getNum(s.substr(j, k - j));
                            if (0 <= n3 && n3 <= 255) {
                                n4 = getNum(s.substr(k));
                                if (0 <= n4 && n4 <= 255) {
                                    result.push_back(to_string(n1) + "." + to_string(n2) + "." + to_string(n3) + "." +
                                                     to_string(n4));
                                }
                            }
                        }
                    }
                }
            }
        }
        return result;
    }
};
```

**这代码真是又臭又长啊。**



#### 2. 回溯算法

##### 递归三部曲

`vector<string> result`存放最终结果

`vector<int> path`存放待切割的位置

1.   确立函数和返回值：
     -   字符串`s`
     -   当前位置`index`
2.   递归终止条件：
     -   当`path`长度为3，意味着可以切割四块，开始切割操作，切割后判断合法性，均合法加入`result`中
     -   当`path`长度大于3，`return`
3.   单次递归：
     -   `path`添加切割位置
     -   下一层递归，切割位置加一
     -   反向操作

#### 代码

``` c++

class Solution {
private:
    vector<string> result;
    vector<int> path;

    void backTracking(string s, int index) {
        if (path.size() == 3) {
            string s1 = s.substr(0, path[0]);
            if (s1.size() > 3) {
                return;
            }
            string s2 = s.substr(path[0], path[1] - path[0]);
            if (s2.size() > 3) {
                return;
            }
            string s3 = s.substr(path[1], path[2] - path[1]);
            if (s3.size() > 3) {
                return;
            }
            string s4 = s.substr(path[2]);
            if (s4.size() > 3) {
                return;
            }
            int n1 = getNum(s1);
            int n2 = getNum(s2);
            int n3 = getNum(s3);
            int n4 = getNum(s4);
            if (0 <= n1 && n1 <= 255 &&
                0 <= n2 && n2 <= 255 &&
                0 <= n3 && n3 <= 255 &&
                0 <= n4 && n4 <= 255) {
                result.push_back(s1 + "." + s2 + "." + s3 + "." + s4);
            }
            return;
        }
        for (int i = index; i <= s.size(); ++i) {
            path.push_back(i);
            backTracking(s, i + 1);
            path.pop_back();
        }
    }

    int getNum(string s) {
        int num = 0;
        if (s == "0") {
            return 0;
        }
        if (s[0] - '0' == 0 && s.size() > 1 || s.size() == 0) {
            return -1;
        }
        for (char i: s) {
            num *= 10;
            num += (i - '0');
        }

        return num;
    }

public:

    vector<string> restoreIpAddresses(string s) {
        if (s.size() < 4) {
            return result;
        }
        backTracking(s, 1);
        return result;
    }
};
```

结构清晰了很多，可读性很好。

## 78. 子集

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/subsets/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0078.%E5%AD%90%E9%9B%86.html)
>
>   视频讲解：[回溯算法解决子集问题，树上节点都是目标集和！ | LeetCode：78.子集](https://www.bilibili.com/video/BV1U84y1q7Ci)
>
>   状态：AC

### 思路

组合问题，这道题的每个元素不重复，做起来很简单，求出所有情况。

#### 递归三部曲

`vector<vector<int>> result`存放最终结果

`vector<int> path`存放单次集合

1.   确定函数和参数：
     -   数组`nums`
     -   当前位置`index`
2.   确定终止条件：第一层的递归结束（或无）
3.   单次递归：
     -   单次集合加入到最终结果中
     -   单次集合加入新的元素
     -   下一层递归
     -   反向操作

### 代码

``` c++
class Solution {
private:
    vector<vector<int>> result;
    vector<int> path;

    void backTracking(vector<int> &nums, int index) {
        result.push_back(path);
        for (int i = index; i < nums.size(); ++i) {
            path.push_back(nums[i]);
            backTracking(nums, i + 1);
            path.pop_back();
        }
    }

public:
    vector<vector<int>> subsets(vector<int> &nums) {
        backTracking(nums, 0);
        return result;
    }
};
```



## 90. 子集II

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/subsets-ii/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0090.%E5%AD%90%E9%9B%86II.html)
>
>   视频讲解：[回溯算法解决子集问题，如何去重？| LeetCode：90.子集II](https://www.bilibili.com/video/BV1vm4y1F71J/)
>
>   状态：AC

### 思路

基本同上一题，但是增加了可重复元素，需要去重。去重在我的博客[40. 组合总和II](https://promisewang.github.io/post/36f15375.html#40-%E7%BB%84%E5%90%88%E6%80%BB%E5%92%8CII)

这道题做之前要先排序

`vector<vector<int>> result`：存放最终结果

`vector<int> path`：存放单次集合

#### 递归三部曲

1.   确定函数和参数
     -   数组`nums`
     -   当前索引`index`
     -   判断元素是否使用过的`used`
2.   递归终止条件：无
3.   单次递归：
     -   单次集合加入最终结果
     -   当`i > 0 && nums[i] == nums[i - 1] && used[i - 1] == false`说明之前已经遍历过，`continue`
     -   `path`加入下个元素
     -   `used[i] == true`下个元素标记被使用过
     -   下一层递归
     -   反向操作



### 代码

``` c++
class Solution {
private:
    vector<vector<int>> result;
    vector<int> path;

    void backTracking(vector<int> &nums, int index, vector<bool> &used) {
        result.push_back(path);
        for (int i = index; i < nums.size(); ++i) {
            if (i > 0 && nums[i] == nums[i - 1] && !used[i - 1]) {
                continue;
            }
            path.push_back(nums[i]);
            used[i] = true;
            backTracking(nums, i + 1, used);
            path.pop_back();
            used[i] = false;
        }
    }

public:
    vector<vector<int>> subsetsWithDup(vector<int> &nums) {
        sort(nums.begin(), nums.end());
        vector<bool> used(nums.size(), false);
        backTracking(nums, 0, used);
        return result;
    }
};
```

