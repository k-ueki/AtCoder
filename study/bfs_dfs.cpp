/*
 * CreatedAt: 2024-07-14 23:30:44
 */

#include <bits/stdc++.h>
using namespace std;

#define rep(i,n) for(int i = 0; i < n; ++i)

long long INF = 1LL << 60;

void search(const Graph &G, int s) {
  int N = (int)G.size();

  vector<bool> seen(N, false);
  queue<int> todo;

  seen[s] = true;
  todo.push(s);
  
  while(!todo.empty()){
    int v = todo.front();
    todo.pop();

    for(int x : G[v]) {
      if (seen[x]) continue;

      seen[x] = true;
      todo.push(x);
    }
  }
  
}

int main() {

}
