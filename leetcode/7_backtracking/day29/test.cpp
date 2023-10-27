//
// Created by Promise on 2023/10/21.
//
#include "Utils.h"

int main() {
    int a = 0;
    for (int i = 1; i < 10; ++i) {
        a += 1 << i;
        cout << a << endl;
    }
    return 0;
}
