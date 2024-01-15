#include "Utils.h"

class Solution {
public:
    int findMinArrowShots(vector<vector<int>> &points) {
        vector<vector<int>> result;
        result.push_back(points[0]);
        if (points.size() == 1) {
            return 1;
        }
        for (int i = 1; i < points.size(); ++i) {
            for (int j = 0; j < result.size(); ++j) {
                // 无交集、continue
                if (result[j][0] > points[i][1] || result[j][1] < points[i][0]) {
                    continue;
                } else if (result[j][0] <= points[i][0] && points[i][0] <= result[j][1] &&
                           result[j][1] <= points[i][1]) {  // 有交集但不包含
                    result[j][0] = points[i][0];
                } else if () {
                    
                }
                else if (
                        result[j][0] <= points[i][0] && points[i][1] <= result[j][1]) {  // 有交集，包含关系
                    result[j][0] = points[i][0];
                    result[j][1] = points[i][1];
                }
            }
        }
    }
};

int main() {
    Solution solution;
    vector<vector<int>> points = {{10, 16},
                                  {2,  8},
                                  {1,  6},
                                  {7,  12}};
    auto result = solution.findMinArrowShots(points);
    cout << result;
    return 0;
}