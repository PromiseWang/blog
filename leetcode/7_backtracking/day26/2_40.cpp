//
// Created by Promise on 2023/10/19.
//
#include "Utils.h"


class Solution {
private:
    vector <vector<int>> result;
    vector<int> path;

    void backTracking(int sum, int target, int index, vector<bool> &used, vector<int> &candidates) {
        if (sum == target) {
            result.push_back(path);
            return;
        }
        for (int i = index; i < candidates.size() && sum + candidates[i] <= target; i++) {
            if (i > 0 && candidates[i] == candidates[i - 1] && !used[i - 1]) {
                continue;
            }
            sum += candidates[i];
            path.push_back(candidates[i]);
            used[i] = true;
            backTracking(sum, target, i + 1, used, candidates);
            sum -= candidates[i];
            path.pop_back();
            used[i] = false;
        }
    }

public:
    vector <vector<int>> combinationSum2(vector<int> &candidates, int target) {
        sort(candidates.begin(), candidates.end());
        path.clear();
        result.clear();
        vector<bool> used(candidates.size(), false);
        backTracking(0, target, 0, used, candidates);
        return result;
    }
};


int main() {
    Solution s;
    vector<int> candidates = {10, 1, 2, 7, 6, 1, 5};
    int target = 8;
    vector <vector<int>> result = s.combinationSum2(candidates, target);
    for (const auto &i: result) {
        for (auto j: i) {
            cout << j << "\t";
        }
        cout << endl;
    }
    return 0;
}