/*
 * CreatedAt: 2024-07-14 13:42:58
 */

#include <bits/stdc++.h>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {
  int n;
  cin >> n;
  vector<string> a(n);
  vector<string> b(n);
  rep(i,n) {
    cin >> a[i];
  }
  rep(i,n) {
    cin >> b[i];
  }

  for(int i = 0; i < n; ++i) {
    for(int j = 0; j < n; ++j) {
      if(a[i][j] != b[i][j]) {
        cout << i+1 << " " << j+1 << endl;
        return 0;
      }
    }
  }

}
