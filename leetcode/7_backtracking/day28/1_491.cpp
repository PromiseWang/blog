#include "Utils.h"

//给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
//
//数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。
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

int main() {
    Solution solution;
    vector<int> nums{4, 6, 7, 7};
    auto result = solution.findSubsequences(nums);
    for (const auto &i: result) {
        for (auto j: i) {
            cout << j << "\t";
        }
        cout << endl;
    }

    return 0;
}