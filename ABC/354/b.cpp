/*
 * CreatedAt: 2024-07-13 15:43:51
 */

#include <bits/stdc++.h>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {

  int N;
  cin >> N;
  vector<string> s(N);
  vector<int> c(N);
  int sum = 0;
  rep(i, N) {
    cin >> s[i] >> c[i];
    sum += c[i];
  }

  sort(s.begin(), s.end());

  cout << s[sum%N] << endl;
}
