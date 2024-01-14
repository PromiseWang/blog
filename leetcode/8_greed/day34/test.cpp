//
// Created by Promise on 2023/10/27.
//
#include "Utils.h"

int main() {
    unordered_map<int, int> maps;
    maps[1] += 1;
    for (auto i: maps) {
        cout << i.first << endl << i.second;
    }
    return 0;
}