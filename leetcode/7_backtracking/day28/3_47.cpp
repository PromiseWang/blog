#include "Utils.h"

// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
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

int main() {
    Solution solution;
    vector<int> nums = {1, 1, 2};
    auto result = solution.permuteUnique(nums);
    for (auto i: result) {
        for (auto j: i) {
            cout << j << "\t";
        }
        cout << endl;
    }

    return 0;
}