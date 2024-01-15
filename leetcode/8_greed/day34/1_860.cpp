#include "Utils.h"

//在柠檬水摊上，每一杯柠檬水的售价为 5 美元。顾客排队购买你的产品，（按账单 bills 支付的顺序）一次购买一杯。
//
//每位顾客只买一杯柠檬水，然后向你付 5 美元、10 美元或 20 美元。你必须给每个顾客正确找零，也就是说净交易是每位顾客向你支付 5 美元。
//
//注意，一开始你手头没有任何零钱。
//
//给你一个整数数组 bills ，其中 bills[i] 是第 i 位顾客付的账。如果你能给每位顾客正确找零，返回 true ，否则返回 false 。
class Solution {
public:
    bool lemonadeChange(vector<int> &bills) {
        int a[3] = {0};
        for (int bill: bills) {
            switch (bill) {
                case 5:
                    a[0] += 1;
                    break;
                case 10:
                    if (a[0] < 1) {
                        return false;
                    } else {
                        a[0] -= 1;
                        a[1] += 1;
                    }
                    break;
                case 20:
                    if (a[0] >= 1 && a[1] >= 1) {
                        a[0] -= 1;
                        a[1] -= 1;
                    } else if (a[0] >= 3) {
                        a[0] -= 3;
                    } else {
                        return false;
                    }
                    a[2] += 1;
                    break;
            }
        }
        return true;
    }
};

int main() {
    Solution solution;
    vector<int> bills = {10, 10};
    cout << solution.lemonadeChange(bills);

    return 0;
}