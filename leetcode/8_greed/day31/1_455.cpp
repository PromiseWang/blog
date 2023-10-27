#include "Utils.h"

//假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
//
//对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j] 。
//如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。
class Solution {
public:
    int findContentChildren(vector<int> &g, vector<int> &s) {
        sort(g.begin(), g.end());
        sort(s.begin(), s.end());
        int index = 0;
        if (s.empty()) {
            return 0;
        }
        int count = 0;
        for (int i = 0; i < g.size() && index < s.size(); ++i) {
            for (int j = index; j < s.size(); ++j) {
                if (s[j] >= g[i]) {
                    index = j + 1;
                    count++;
                    break;
                }
            }
        }
        return count;
    }
};

int main() {
    Solution solution;
    vector<int> g = {1, 1, 1, 1, 1, 1, 1, 1};
    vector<int> s = {3, 4, 5};
    auto result = solution.findContentChildren(g, s);
    cout << result;
    return 0;
}