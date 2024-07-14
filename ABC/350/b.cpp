/*
 * CreatedAt: 2024-07-14 13:51:34
 */

#include <bits/stdc++.h>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {
  int n, q;
  cin >> n >> q;
  vector<int> t(q);
  rep(i,q) cin >> t[i];

  vector<bool> ss(n+1, true);
  for(int i = 0; i < q; ++i) {
    ss[t[i]] = !ss[t[i]];
  }

  int count = 0;
  for(int i = 1; i <= n; ++i) {
    if(ss[i]) count++;
  }

  cout << count << endl;

}
