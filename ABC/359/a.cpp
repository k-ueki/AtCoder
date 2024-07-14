/*
 * CreatedAt: 2024-07-14 22:45:21
 */

#include <bits/stdc++.h>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {

  int n;
  cin >> n;
  vector<string> s(n);
  rep(i,n) cin >> s[i];
  
  int count = 0;
  rep(i,n) {
    if(s[i]=="Takahashi") count++;
  }

  cout << count << endl;

}
