#include <iostream>
#include <vector>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {
  int N, A;
  cin >> N >> A;
  vector<int> T(N);
  rep(i, N) cin >> T[i];

  vector<int> ans(N);
  for(int i = 0; i < N; ++i) {
    if(i == 0) ans[i] = T[i] + A;
    else {
      ans[i] = max(ans[i-1] + A, T[i] + A);
    }
  }

  rep(i, N) cout << ans[i] << endl;
  return 0;
}
