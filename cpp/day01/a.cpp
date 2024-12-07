#include <iostream>

using namespace std;

int abs(const int n) {  return n < 0 ? -n : n; }

bool cmp(const int a, const int b) { return a < b; }

int main(int argc, char **argv) {
    freopen(argv[1], "r", stdin);

    int l, r;
    vector<int> lq;
    vector<int> rq;
    while (cin >> l >> r) {
        lq.push_back(l);
        rq.push_back(r);
    }

    sort(lq.begin(), lq.end(), cmp);
    sort(rq.begin(), rq.end(), cmp);

    int sum1 = 0, sum2 = 0;
    for (int i = 0; i < rq.size(); i++) {
        sum1 += abs(rq[i] - lq[i]);
        int count = 0;
        for (int j = 0; j < rq.size(); j++) {
            if (lq[i] == rq[j]) {
                count++;
            }
        }
        sum2 += lq[i] * count;
    }
    cout << "Part 1: " << sum1 << endl;
    cout << "Part 2: " << sum2 << endl;
}