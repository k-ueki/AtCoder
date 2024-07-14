/*
 * CreatedAt: 2024-07-14 22:49:17
 */

#include <bits/stdc++.h>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {

  int n,m ;
  cin >> n >> m;
  vector<int> h(n);
  rep(i,n) cin >> h[i];

  int count = 0;
  rep(i,n) {
    count += h[i];
    if(count > m) {
      cout << i << endl;
      return 0;
    }
  }

  cout << n << endl;

}
