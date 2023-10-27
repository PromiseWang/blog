#include "Utils.h"

//给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
//
//每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:
//
//0 <= j <= nums[i]
//i + j < n
//返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
class Solution {
public:
    int jump(vector<int> &nums) {
        vector<int> dp(nums.size(), 0);
        if (nums.size() == 1) {
            return 0;
        }
        int count = 0;
        int curDis = 0;
        int nextDis = 0;
        for (int i = 0; i < nums.size(); ++i) {
            nextDis = max(i + nums[i], nextDis);
            if (i == curDis) {
                curDis = nextDis;
                count++;
                if (nextDis >= nums.size() - 1) {
                    break;
                }
            }
        }
        return count;
    }
};

int main() {
    Solution solution;
    vector<int> nums = {1, 2, 1, 1, 1};
    auto result = solution.jump(nums);
    cout << result;
    return 0;
}