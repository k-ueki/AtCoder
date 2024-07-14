#include <iostream>
#include <vector>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {
  int N;
  cin >> N;
  vector<long long> l(N), r(N);
  int lsum = 0;
  int rsum = 0;
  rep(i,N) {
    cin >> l[i] >> r[i];
    lsum += l[i];
    rsum += r[i];
  }

  if(lsum <= 0 && rsum >= 0) cout << "Yes" << endl;
  else {
    cout << "No" << endl;
    return 0;
  }


  return 0;
}
