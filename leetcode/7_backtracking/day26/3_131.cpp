//
// Created by Promise on 2023/10/19.
//
#include "Utils.h"

class Solution {
private:
    vector<vector<string>> result;
    vector<int> path;
    vector<string> pathString;
    string temp = "";

    bool judge(string s) {
        if (s.size() == 0) {
            return false;
        }
        if (s.size() == 1) {
            return true;
        }
        for (int i = 0; i < s.size() / 2; i++) {
            if (s[i] != s[s.size() - 1 - i]) {
                return false;
            }
        }
        return true;
    }

    void backTracking(const string &s, int index, string temp) {
        if (index > s.size()) {
            int left = 0;
            int flag = true;
            for (int right = 0; right < path.size(); right++) {
                string temp = s.substr(left, path[right] - left);
                pathString.push_back(temp);
                if (!judge(temp)) {
                    flag = false;
                    pathString.clear();
                    return;
                }
                left = path[right];
            }
            if (flag) {
                result.push_back(pathString);
                pathString.clear();
            }
            return;
        }
        for (int i = index; i <= s.size(); ++i) {
            path.push_back(i);
            backTracking(s, i + 1, temp);
            path.pop_back();
        }
    }

public:
    vector<vector<string>> partition(string s) {
        backTracking(s, 1, "");
        for (auto i: path) {
            cout << i << endl;
        }
        return result;
    }
};

int main() {
    Solution s;
    s.partition("aab");
    return 0;
}
