/*
 * CreatedAt: 2024-07-14 23:39:59
 */

#include <bits/stdc++.h>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {
  int n;
  cin >> n;
  vector<int> b(n-1);
  rep(i,n-1) cin >> b[i];


  int ans = b[0];
  for(int i = 0; i < n-1; ++i){
    if(i == n-2) ans += b[i];
    else ans += min(b[i], b[i+1]);
  }

  cout << ans << endl;
}
