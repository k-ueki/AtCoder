/*
 * CreatedAt: 2024-07-13 20:42:01
 */

#include <bits/stdc++.h>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

struct Hoge {
  int x;
  int y;
};

float dist(Hoge a1, Hoge a2) {
  return pow(a1.x - a2.x, 2) + pow(a1.y - a2.y, 2);
}

int calc(int a, int b, int c) {
  return (b+c)-a;
}

int main() {
  vector<Hoge> a(3);
  rep(i, 3) cin >> a[i].x >> a[i].y;

  vector<float> s(3);
  for(int i = 0; i < 3; ++i) {
    s[i] = dist(a[i], a[(i+1)%3]);
  }

  if(calc(s[0],s[1],s[2])==0) {
    cout << "Yes" << endl;
    return 0;
  }
  if(calc(s[1],s[2],s[0])==0) {
    cout << "Yes" << endl;
    return 0;
  }
  if(calc(s[2],s[0],s[1])==0) {
    cout << "Yes" << endl;
    return 0;
  }

  cout << "No" << endl;
  return 0;
}
