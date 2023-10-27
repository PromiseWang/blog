#include "Utils.h"


int main() {
    vector<int> nums = {1, 2, 3, 4, 5};
    nums.insert(nums.end(), nums.begin(), nums.end() - 1);
    for (auto i: nums) {
        cout << i << endl;
    }
    return 0;
}