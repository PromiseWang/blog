#include "Utils.h"

//给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
//判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
class Solution {
public:
    bool canJump(vector<int> &nums) {
        int dis = 0;
        if (nums.size() == 1) {
            return true;
        }
        for (int i = 0; i <= dis; ++i) {
            dis = max(nums[i] + i, dis);
            if (dis >= nums.size() - 1) {
                return true;
            }
        }
        return false;
    }
};

int main() {
    Solution solution;
    vector<int> nums = {0, 2, 3};
    auto result = solution.canJump(nums);
    cout << result;
    return 0;
}