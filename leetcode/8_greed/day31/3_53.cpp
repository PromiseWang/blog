#include "Utils.h"

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

int main() {
    Solution solution;
    vector<int> nums = {-1, -2};
    auto result = solution.maxSubArray(nums);
    cout << result;
    return 0;
}