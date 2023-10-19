//
// Created by Promise on 2023/10/19.
//
#include "Utils.h"

using namespace std;
vector<int> nums = {1, 2, 3, 4, 5, 6, 7, 8, 9};
vector<int> path;
vector<vector<int>> result;

void backTracking(vector<int> &nums, int index) {
    result.push_back(path);
    for (int i = index; i < 9; i++) {
        path.push_back(nums[i]);
        backTracking(nums, i+1);
        path.pop_back();
    }
}

int main() {
//    backTracking(nums, 0);
//    cout << result.size();
    string a = "123456789";
    cout << a.substr(8, 1);
    return 0;
}