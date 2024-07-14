/*
 * CreatedAt: 2024-07-14 18:34:06
 */

#include <bits/stdc++.h>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {
  int n,k;
  cin >> n >> k;
  vector<long long> a(n);
  rep(i,n) cin >> a[i];

  sort(a.begin(), a.end());

  long long ans = INF;
  for(int i = 0; i <= k; ++i) {
    ans = min(ans, a[a.size()-(i+1)] - a[k-i]);
  }

  cout << ans << endl;
}
