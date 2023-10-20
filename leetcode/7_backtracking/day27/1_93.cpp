//
// Created by Promise on 2023/10/20.
//有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
//例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
//给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。
// 你 不能 重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。
#include "Utils.h"

class Solution {
private:
    vector<string> result;
    vector<int> path;

    void backTracking(string s, int index) {
        if (path.size() == 3) {
            string s1 = s.substr(0, path[0]);
            if (s1.size() > 3) {
                return;
            }
            string s2 = s.substr(path[0], path[1] - path[0]);
            if (s2.size() > 3) {
                return;
            }
            string s3 = s.substr(path[1], path[2] - path[1]);
            if (s3.size() > 3) {
                return;
            }
            string s4 = s.substr(path[2]);
            if (s4.size() > 3) {
                return;
            }
            int n1 = getNum(s1);
            int n2 = getNum(s2);
            int n3 = getNum(s3);
            int n4 = getNum(s4);
            if (0 <= n1 && n1 <= 255 &&
                0 <= n2 && n2 <= 255 &&
                0 <= n3 && n3 <= 255 &&
                0 <= n4 && n4 <= 255) {
                result.push_back(s1 + "." + s2 + "." + s3 + "." + s4);
            }
            return;
        }
        for (int i = index; i <= s.size(); ++i) {
            path.push_back(i);
            backTracking(s, i + 1);
            path.pop_back();
        }
    }

    int getNum(string s) {
        int num = 0;
        if (s == "0") {
            return 0;
        }
        if (s[0] - '0' == 0 && s.size() > 1 || s.size() == 0) {
            return -1;
        }
        for (char i: s) {
            num *= 10;
            num += (i - '0');
        }

        return num;
    }

public:
    vector<string> restoreIpAddresses1(string s) {
        if (s.size() < 4) {
            return result;
        }
        int n1, n2, n3, n4;
        for (int i = 1; i <= 3; ++i) {
            n1 = getNum(s.substr(0, i));
            if (0 <= n1 && n1 <= 255) {
                for (int j = i + 1; j <= 6 || j < s.size() - 1; ++j) {
                    if (j - i > 3) {
                        break;
                    }
                    n2 = getNum(s.substr(i, j - i));
                    if (0 <= n2 && n2 <= 255) {
                        for (int k = j + 1; k <= 9 || k < s.size(); ++k) {
                            if (k - j > 3) {
                                break;
                            }
                            if (s.size() - k > 3) {
                                continue;
                            }
                            n3 = getNum(s.substr(j, k - j));
                            if (0 <= n3 && n3 <= 255) {
                                n4 = getNum(s.substr(k));
                                if (0 <= n4 && n4 <= 255) {
                                    result.push_back(to_string(n1) + "." + to_string(n2) + "." + to_string(n3) + "." +
                                                     to_string(n4));
                                }
                            }
                        }
                    }
                }
            }
        }
        return result;
    }

    vector<string> restoreIpAddresses(string s) {
        if (s.size() < 4) {
            return result;
        }
        backTracking(s, 1);
        return result;
    }
};

int main() {
    Solution solution;
//    string s = "25525511135";
    string s = "101023";
    auto result = solution.restoreIpAddresses(s);
    for (auto i: result) {
        cout << i << endl;
    }
    return 0;
}