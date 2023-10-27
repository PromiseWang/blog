#include "Utils.h"

class Solution {
public:
    int canCompleteCircuit(vector<int> &gas, vector<int> &cost) {
        gas.insert(gas.end(), gas.begin(), gas.end());
        cost.insert(cost.end(), cost.begin(), cost.end() - 1);
        for (int i = 0; i < gas.size() / 2; ++i) {
            int cur = 0;
            bool flag = true;
            for (int j = i; j < i + gas.size() / 2; ++j) {
                if (cur + gas[j] >= cost[j]) {
                    cur += gas[j] - cost[j];
                } else {
                    flag = false;
                    i = j;
                    break;
                }
            }
            if (flag) {
                return i;
            }
        }
        return -1;
    }
};

int main() {
    Solution solution;
    vector<int> gas = {4, 5, 2, 6, 5, 3}, cost = {3, 2, 7, 3, 2, 9};
    auto result = solution.canCompleteCircuit(gas, cost);
    cout << result << endl;
    return 0;
}