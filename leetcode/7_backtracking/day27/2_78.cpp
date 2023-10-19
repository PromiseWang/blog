//
// Created by Promise on 2023/10/20.
//
#include "Utils.h"

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

int main() {
    Solution solution;
    vector<int> nums = {1, 2, 2};
    auto result = solution.subsets(nums);
    for (auto i: result) {
        for (auto j: i) {
            cout << j << "\t";
        }
        cout << endl;
    }
    return 0;
}