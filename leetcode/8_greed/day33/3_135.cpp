#include "Utils.h"

//n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
//
//你需要按照以下要求，给这些孩子分发糖果：
//
//每个孩子至少分配到 1 个糖果。
//相邻两个孩子评分更高的孩子会获得更多的糖果。
//请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
class Solution {
public:
    int candy(vector<int> &ratings) {
        vector<int> result(ratings.size(), 1);
        for (int i = 1; i < ratings.size(); ++i) {
            if (ratings[i] > ratings[i - 1]) {
                result[i] = result[i - 1] + 1;
            }
        }
        for (int i = ratings.size() - 1; i >= 1; --i) {
            if (ratings[i - 1] > ratings[i] && result[i - 1] <= result[i]) {
                result[i - 1] = result[i] + 1;
            }
        }
        int sum = 0;
        for (auto i: result) {
            sum += i;
        }
        return sum;
    }
};

int main() {
    Solution solution;
    vector<int> ratings = {1, 0, 2};
    cout << solution.candy(ratings);
    return 0;
}