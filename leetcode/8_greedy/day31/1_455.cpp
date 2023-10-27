#include "Utils.h"

class Solution {
private:
    vector<vector<int>> result;
    vector<int> path;

    void backTracking() {

    }

public:

};

int main() {
    Solution solution;

    for (auto i: result) {
        for (auto j: i) {
            cout << j << "\t";
        }
        cout << endl;
    }

    return 0;
}