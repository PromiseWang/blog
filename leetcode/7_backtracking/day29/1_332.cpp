#include "Utils.h"

class Solution {
private:
    vector<string> result;
    vector<int> path;
    string next = "";

    void backTracking(vector<vector<string>> &tickets, int index, vector<bool> used) {
        for (int i = 0; i < tickets.size(); ++i) {
            if (tickets[i][1] == tickets[index][0]) {

            }
        }
    }

public:
    vector<string> findItinerary(vector<vector<string>> &tickets) {
        sort(tickets.begin(), tickets.end());
        vector<bool> used(tickets.size(), false);
        for (int i = 0; i < tickets.size(); ++i) {
            if (tickets[i][0] == "JFK") {
                backTracking(tickets, i, used);
                break;
            }
        }
    }
};

int main() {
    Solution solution;
    vector<vector<string>> tickets = {{"MUC", "LHR"},
                                      {"JFK", "MUC"},
                                      {"SFO", "SJC"},
                                      {"LHR", "SFO"}};
    auto result = solution.findItinerary(tickets);

    for (auto i: result) {
        for (auto j: i) {
            cout << j << "\t";
        }
        cout << endl;
    }
    return 0;
}