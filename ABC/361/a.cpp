/*
 * CreatedAt: 2024-07-14 22:41:40
 */

#include <iostream>
#include <vector>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {
  int n,k,x;
  cin >> n >> k >> x;
  vector<int> a(n);
  rep(i,n) cin >> a[i];
  
  vector<int> ans;
  for(int i = 0; i < n; ++i) {
    ans.push_back(a[i]);
    if(i==k-1) ans.push_back(x);
  }

  rep(i,n+1) cout << ans[i] << " ";
  return 0;
}
