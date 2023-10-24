#include "Utils.h"

FILE *fp = fopen("./result.txt", "a+");
class Solution {
public:
    vector<vector<char>> ans;
private:
    void backTracking(vector<vector<char>> &board, int rowIndex, int colIndex, int *usedRows, int *usedCols,
                      int *usedNine) {
        for (auto i: board) {
            for (auto j: i) {
                fprintf(fp, "%c\t", j);
            }
            fprintf(fp, "\n");
        }
        fprintf(fp, "****************************\n");

        bool flag = true;
        for (int i = 0; i < 9; ++i) {
            if (usedRows[i] != 1022 || usedCols[i] != 1022 || usedNine[i] != 1022) {
                flag = false;
                break;
            }
        }
        if (flag && ans.empty()) {
            ans.assign(board.begin(), board.end());
            return;
        }
        for (int row = rowIndex; row < 9; ++row) {
            for (int col = colIndex; col < 9; ++col) {
                int index = row / 3 * 3 + col / 3;
                if (board[row][col] == '.') {
                    for (int num = 1; num <= 9; ++num) {
                        if (((usedRows[row] >> num & 1) == 0) && ((usedCols[col] >> num & 1) == 0) &&
                            ((usedRows[index] >> num & 1) == 0)) {
                            board[row][col] = '0' + num;
                            usedRows[row] += 1 << num;
                            usedCols[col] += 1 << num;
                            usedNine[index] += 1 << num;
                            if (col < 8) {
                                backTracking(board, row, col + 1, usedRows, usedCols, usedNine);
                            } else if (col == 8 && row < 8) {
                                backTracking(board, row + 1, 0, usedRows, usedCols, usedNine);
                            }
                        }
                    }
//                    if (board[row][col] == '.') {
//                        return;
//                    }
                }
            }
        }
    }

public:
    void solveSudoku(vector<vector<char>> &board) {
        int usedRows[9] = {0}, usedCols[9] = {0}, usedNine[9] = {0};
        for (int row = 0; row < 9; ++row) {
            for (int col = 0; col < 9; ++col) {
                int index = row / 3 * 3 + col / 3;
                if (board[row][col] != '.') {
                    usedRows[row] += 1 << (board[row][col] - '0');
                    usedCols[col] += 1 << (board[row][col] - '0');
                    usedNine[index] += 1 << (board[row][col] - '0');
                }
            }
        }
        backTracking(board, 0, 0, usedRows, usedCols, usedNine);
//        board = ans;
    }
};

int main() {
    Solution solution;
    vector<vector<char>> result = {{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
                                   {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
                                   {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
                                   {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
                                   {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
                                   {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
                                   {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
                                   {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
                                   {'.', '.', '.', '.', '8', '.', '.', '7', '9'}};

    solution.solveSudoku(result);
    fclose(fp);
    for (const auto &i: solution.ans) {
        for (auto j: i) {
            cout << j << '\t';
        }
        cout << endl;
    }
    return 0;
}