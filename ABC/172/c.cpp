#include <bits/stdc++.h>
#define rep(i,n) for(int i=0;i<(n);i++)
typedef long long ll;

using namespace std;

int main(){
    ios::sync_with_stdio(false);
    int n,m,k;
    cin >> n >> m >>k;
    stack<int> a;
    stack<int> b;
    int tmp;
    rep(i,n){
        cin >> tmp;
        a.push(tmp);
    }
    rep(i,m){
        cin >> tmp;
        b.push(tmp);
    }

    int sum = 0;
    while (sum <= k ||(!a.empty() && !b.empty())){

    }

    return 0;
}