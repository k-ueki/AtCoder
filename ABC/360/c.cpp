/*
 * CreatedAt: 2024-07-14 19:04:55
 */

#include <bits/stdc++.h>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {
  long long n;
  cin >> n;
  vector<long long> a(n), weight(n);
  rep(i,n) cin >> a[i];
  rep(i,n) cin >> weight[i];

  vector<long long> maxs(n+1, 0);
  for(int i = 0; i < n; ++i) {
    maxs[a[i]] = max(maxs[a[i]], weight[i]);
  }
  
  long long ans = 0;
  for(int i = 0; i < n; ++i) {
    if(weight[i] < maxs[a[i]]) ans += weight[i];
  }

  cout << ans << endl;
}
