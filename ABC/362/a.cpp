/*
 * CreatedAt: 2024-07-13 20:41:53
 */

#include <bits/stdc++.h>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

int main() {
  int R,G,B;
  cin >> R >> G >> B;
  string c; 
  cin >> c;

  if(c=="Red") cout << min(G,B) << endl;
  if(c=="Blue") cout << min(R,G) << endl;
  if(c=="Green") cout << min(R,B) << endl;

  return 0;
}
