---
title: 代码随想录算法训练营第二十六天  39.组合总和、40.组合总和II、131.分割回文串
tags:
  - 算法
  - 代码随想录
  - LeetCode
  - 回溯算法
categories: 刷题
abbrlink: 36f15375
date: 2023-10-19 23:23:02
---

# 代码随想录算法训练营第二十六天  39.组合总和、40.组合总和II、131.分割回文串

## 39.组合总和

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/combination-sum/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0039.%E7%BB%84%E5%90%88%E6%80%BB%E5%92%8C.html)
>
>   视频讲解：[Leetcode:39. 组合总和讲解](https://www.bilibili.com/video/BV1KT4y1M7HJ)
>
>   状态：AC

### 思路

这道题还比较简单，先进行排序，然后开始递归。排序是为了确定当前几个数的和大于了`target`说明后面就不存在结果了。

#### 递归三部曲

1.   确定函数参数：
     -   数组
     -   当前递归位置
     -   总和
     -   目标和
2.   递归终止条件：如果目标和等于总和，那么把所有元素放入`result`
3.   单次递归：总和大于`target`，剪枝，`return`。否则进行下一次递归

### 代码

``` go
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	var path []int
	var backTracking func(candidates []int, index, sum, target int)
	backTracking = func(candidates []int, index, sum, target int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		if sum > target {
			return
		}
		for i := index; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				return
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			backTracking(candidates, i, sum, target)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backTracking(candidates, 0, 0, target)
	return result
}
```

## 40. 组合总和II

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/combination-sum-ii/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0040.%E7%BB%84%E5%90%88%E6%80%BB%E5%92%8CII.html)
>
>   视频讲解：[回溯算法中的去重，树层去重树枝去重，你弄清楚了没？| LeetCode:40.组合总和II](https://www.bilibili.com/video/BV12V4y1V73A)
>
>   状态：TLE

### 思路

<font color="red">从这道题开始换C++写算法了，Go写算法有点难受。最初用Go写算法只是想多练练Go的语法。但是到刷LeetCode还是不太适合。</font>

这道题去重一直想不出来，试了几种方法，要么超时、要么思路错了：

-   将结果放入到一个集合中。由于Go语言没有集合，那么新建一个`maps := map[interface{}]bool{}`。`key`则是每一个结果。但是`key`不能是切片类型。所以不能使用。
-   由于不能使用切片，那么使用一个“曲线救题”的方式，把结果转成字符串（加入结果为[1, 2, 3]，那么就变成“123”）这样就可以存入`key`中。超出了时间限制。该测试用例如下`candidates = [1, 1, ..., 1]（100个）`，`target=30`。去重了但没剪枝。

<font color="green">这题投降</font>

#### 卡哥讲解

对于我来讲想明白了一下两点我的代码就写出来了，假设`candidates = [1, 2(1), 2(2), 2(3), 5]`，括号为了区分是哪个`2`：

-   假如递归到`[1, 2(1), 2(2), 2(3)]`，那么再以后递归时候`2(1)`为第一个元素，那么`[2(1), 2(2), 2(3)]`的情况再刚刚已经出现了，则不需要再递归了。也是剪枝操作。
-   为了防止重复递归，引入一个布尔类型的`used`数组。`used[i]`代表`candidates[i]`在当前递归的树枝上是否使用。那么当`candidates[i] != candidates[i - 1] && used[i] == false`时，才不会重复递归。

#### 递归三部曲

1.   确立递归函数参数
     -   当前递归的和
     -   目标和
     -   当前位置索引
     -   `used`数组，是否使用
     -   `candidates`
2.   递归终止条件：`sum == target`时，将`path`加入到`result`中
3.   单次递归：
     -   先剪枝：当`i > 0 && candidates[i] == candidates[i - 1] && !used[i - 1]`，说明已经递归过了，应该<font color="red">`continue`</font>而不应该`return`。
     -   当前和加上下一个数
     -   下一个数放入到路径`path`中
     -   下一个数标记为`true`，示意已被使用。
     -   下层递归
     -   反向上述操作

### 代码

``` c++
class Solution {
private:
    vector<vector<int>> result;
    vector<int> path;

    void backTracking(int sum, int target, int index, vector<bool> &used, vector<int> &candidates) {
        if (sum == target) {
            result.push_back(path);
            return;
        }
        for (int i = index; i < candidates.size() && sum + candidates[i] <= target; i++) {
            if (i > 0 && candidates[i] == candidates[i - 1] && !used[i - 1]) {
                continue;
            }
            sum += candidates[i];
            path.push_back(candidates[i]);
            used[i] = true;
            backTracking(sum, target, i + 1, used, candidates);
            sum -= candidates[i];
            path.pop_back();
            used[i] = false;
        }
    }

public:
    vector<vector<int>> combinationSum2(vector<int> &candidates, int target) {
        sort(candidates.begin(), candidates.end());
        path.clear();
        result.clear();
        vector<bool> used(candidates.size(), false);
        backTracking(0, target, 0, used, candidates);
        return result;
    }
};
```

## 131. 分割回文串

>   题目链接：[力扣题目链接](https://leetcode.cn/problems/palindrome-partitioning/)
>
>   文章讲解：[代码随想录(https://programmercarl.com)](https://programmercarl.com/0131.%E5%88%86%E5%89%B2%E5%9B%9E%E6%96%87%E4%B8%B2.html)
>
>   视频讲解：[131.分割回文串](https://www.bilibili.com/video/BV1c54y1e7k6)
>
>   状态：AC

### 思路

我这个思路明显是有点绕弯的，不是最好的。头已经大了，已经想不下去最优解了，明天再看。

这道题真的想了一下午。最开始没理解题意，我以为是判断每个子串，如果是回文则输出。写了好久的代码之后发现并不能通过测试用例，重新好好理解下题，然后明白了。是如果一个字符串同时分割出一些子串，如果他们满足则是答案。例如“aaba”，一共有这些种分割方式：

|  分割情况  | 是否满足 |
| :--------: | :------: |
| a\|a\|b\|a |    ✅     |
|  aa\|b\|a  |    ✅     |
|  a\|ab\|a  |    ❌     |
|  a\|a\|ba  |    ❌     |
|  aa\|b\|a  |    ✅     |
|   a\|aba   |    ✅     |
|   aa\|ba   |    ❌     |
|   aab\|a   |    ❌     |

那么只需要知道分割位置即可，变成了从1开始的数组到最后各种组合情况。判断每一种情况是不是回文串

#### 递归三部曲

`vector<vector<string>> result`存放最终结果

`vector<int> path`存放切割位置

`vector<string> pathString`存放切割后的字符串，如果均为回文串，则加入到`result`。

1.   确立函数和参数：
     -   字符串`s`
     -   当前递归位置`index`
2.   递归终止条件：当`index > s.size()`时。将`s`按照`path`存放的数字位置切割，判断切割后每一个串是否均为回文串。如果是加入到`result`，否则清空`pathString`
3.   单次递归：
     -   递归范围[index, s.size()]
     -   `path`添加当前位置
     -   下一层递归
     -   反向操作

### 代码

``` c++
class Solution {
private:
    vector<vector<string>> result;
    vector<int> path;
    vector<string> pathString;
    string temp = "";

    bool judge(string s) {
        if (s.size() == 0) {
            return false;
        }
        if (s.size() == 1) {
            return true;
        }
        for (int i = 0; i < s.size() / 2; i++) {
            if (s[i] != s[s.size() - 1 - i]) {
                return false;
            }
        }
        return true;
    }

    void backTracking(const string &s, int index) {
        if (index > s.size()) {
            int left = 0;
            int flag = true;
            for (int right = 0; right < path.size(); right++) {
                string temp = s.substr(left, path[right] - left);
                pathString.push_back(temp);
                if (!judge(temp)) {
                    flag = false;
                    pathString.clear();
                    return;
                }
                left = path[right];
            }
            if (flag) {
                result.push_back(pathString);
                pathString.clear();
            }
            return;
        }
        for (int i = index; i <= s.size(); ++i) {
            path.push_back(i);
            backTracking(s, i + 1);
            path.pop_back();
        }
    }

public:
    vector<vector<string>> partition(string s) {
        backTracking(s, 1);
        for (auto i: path) {
            cout << i << endl;
        }
        return result;
    }
};
```

