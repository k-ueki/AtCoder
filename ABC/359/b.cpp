#include <iostream>
#include <vector>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {
  int N;
  cin >> N;
  vector<int> A(2*N);
  rep(i, 2*N) cin >> A[i];

  int ans = 0;
  for(int c = 1; c <= N; ++c) {
    int pre=-1;
    for(int i = 0; i < A.size(); ++i) {
      if(A[i] == c) {
        if(pre == -1) pre = i;
        else {
          if(i-pre == 2) ans++;
        }
      }
    }
  }
  cout<< ans << endl;
  return 0;
}
