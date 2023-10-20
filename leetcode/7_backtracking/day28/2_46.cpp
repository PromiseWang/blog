#include "Utils.h"
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
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

int main() {
    Solution solution;
    vector<int> nums = {1, 2, 3};
    auto result = solution.permute(nums);
    for (auto i: result) {
        for (auto j: i) {
            cout << j << "\t";
        }
        cout << endl;
    }

    return 0;
}