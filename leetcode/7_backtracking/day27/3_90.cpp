#include "Utils.h"

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

int main() {
    Solution solution;
    vector<int> nums = {4, 4, 4, 1, 4};
    auto result = solution.subsetsWithDup(nums);

    for (auto i: result) {
        for (auto j: i) {
            cout << j << "\t";
        }
        cout << endl;
    }
    return 0;
}