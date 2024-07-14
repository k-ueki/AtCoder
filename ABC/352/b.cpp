/*
 * CreatedAt: 2024-07-14 13:11:24
 */

#include <bits/stdc++.h>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {
  string s,t;
  cin >> s >> t;

  int si = 0;
  for(int i = 0; i < t.size(); ++i) {
    if(s[si] == t[i]) {
      cout << i+1 << " ";
      si++;
    }
  }

}
