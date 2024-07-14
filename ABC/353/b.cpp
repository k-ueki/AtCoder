/*
 * CreatedAt: 2024-07-14 12:32:09
 */

#include <bits/stdc++.h>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {
  int n, k;
  cin >> n >> k;
  vector<int> a(n);
  rep(i,n) cin >> a[i];


  vector<int> tmp(n);
  // for(int i = 0; i < n-1; ++i) {
  //   if(i == 0) tmp[i] = a[i];
  //   
  //   if(tmp[i]+a[i+1] > k) {
  //     tmp[i+1] = (tmp[i] + a[i+1]) - tmp[i];
  //     count++;
  //   }
  //   else tmp[i+1] = tmp[i] + a[i+1];
  // }
  
  int seat = k, count = 1;
  for(int i = 0; i < n; ++i) {
    if(seat < a[i]) {
      seat = k;
      count++;
    }
    seat -= a[i];
  }
  
  cout << count << endl;
}
