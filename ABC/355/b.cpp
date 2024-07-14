/*
 * CreatedAt: 2024-07-13 14:49:16
 */

#include <bits/stdc++.h>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

int main() {
  int N,M;
  cin >> N >> M;

  vector<int> A(N), B(M);
  rep(i, N) cin >> A[i];
  rep(i, M) cin >> B[i];

  vector<int> C = A;
  rep(i, M) C.push_back(B[i]);
  sort(C.begin(), C.end());

  set<int> st(A.begin(), A.end());

  for( int i = 0; i < N+M-1; ++i) {
    if(st.contains(C[i]) && st.contains(C[i+1])) {
      cout << "Yes" << endl;
      return 0;
    }
  }

  cout << "No" << endl;
  
  return 0;
}
