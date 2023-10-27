#include "Utils.h"

//给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：
//
//选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
//重复这个过程恰好 k 次。可以多次选择同一个下标 i 。
//
//以这种方式修改数组后，返回数组 可能的最大和 。
class Solution {
public:
    int largestSumAfterKNegations(vector<int> &nums, int k) {
        sort(nums.begin(), nums.end());
        int count = 0;  // 负数个数
        int minIndex = 0;
        for (int i = 0; i < nums.size(); ++i) {
            if (nums[i] < 0) {
                count++;
            }
            if (abs(nums[i]) < abs(nums[minIndex])) {
                minIndex = i;
            }
        }
        // count >= k 前k个负数全都翻
        if (count >= k) {
            for (int i = 0; i < k; ++i) {
                nums[i] = -nums[i];
            }
        } else {
            // count < k
            //   1. 所有负数都翻转
            //   2.1 if (k - count) % 2 == 0 输出
            //   2.2 else 翻转绝对值最小的数
            for (int i = 0; i < count; ++i) {
                nums[i] = -nums[i];
            }
            if ((k - count) % 2 == 1) {
                nums[minIndex] = -nums[minIndex];
            }
        }
        int sum = 0;
        for (auto i: nums) {
            sum += i;
        }
        return sum;
    }
};

int main() {
    Solution solution;
    vector<int> nums = {-4, -2, -3};
    int k = 4;
    auto result = solution.largestSumAfterKNegations(nums, k);
    cout << result;
    return 0;
}