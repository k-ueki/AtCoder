/*
 * CreatedAt: 2024-07-13 14:32:05
 */

#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {

  string S;
  cin >> S;

  int upperCount = 0;
  rep(i, S.size()) {
    if(isupper(S[i])) upperCount++;
  }

  if(upperCount > S.size()/2) {
    rep(i, S.size()) S[i] = toupper(S[i]);
  } else {
    rep(i, S.size()) S[i] = tolower(S[i]);
  }
  cout << S << endl;

  return 0;
}
