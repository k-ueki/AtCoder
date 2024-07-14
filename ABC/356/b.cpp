/*
 * CreatedAt: 2024-07-13 14:39:46
 */

#include <iostream>
#include <vector>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {

  int N, M;
  cin >> N >> M;
  vector<int> A(M);
  rep(i, M) cin >> A[i];
  vector<int> X(M,0);
  rep(i, N) {
    rep(j, M) {
      int value;
      cin >> value;
      X[j] += value;
    }
  }

  
  vector<bool> s(M);
  rep(i, M) {
    if(X[i] < A[i]) {
      cout << "No" << endl;
      return 0;
    }
  }
  
  cout << "Yes" << endl;
  return 0;
}
