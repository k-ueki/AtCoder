#include <bits/stdc++.h>
#define rep(i,n) for(int i=0;i<(n);i++)
typedef long long ll;

using namespace std;

const int INF = 1001001001;
int a[12][12];

int main(){
    ios::sync_with_stdio(false);
    int n,m,x;
    cin >> n >> m >> x;
    vector<int> c(n);
    rep(i,n){
        cin >> c[i];
        rep(j,m)cin >> a[i][j];
    }

    int ans = INF;
    rep(i,1<<n){
        int cost = 0;
        vector<int> exp(m);
        rep(j,n){
            if(i>>j&1){
                cost += c[j];
                rep(k,m)exp[k]+=a[j][k];
            }
        }
        bool ok = true;
        rep(j,m)if (exp[j] < x) ok =false;
        if(ok)ans = min(ans,cost);
    }

    if(ans == INF)cout << -1 << endl;
    else cout << ans << endl;
    return 0;
}