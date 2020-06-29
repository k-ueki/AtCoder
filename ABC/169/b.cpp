#include <bits/stdc++.h>
#define rep(i,n) for(int i=0;i<n;++i)
typedef long long ll;
using namespace std;

const ll Max = 1e18;

int main(){
    ios::sync_with_stdio(false);
    int n; cin >> n;
    vector<int> a(n);
    rep(i,n)cin >> a[i];

    ll ans = 1;
    rep(i,n)ans *= a[i];

    if(ans >= Max)cout << -1 << endl;
    else cout << ans << endl;

    return 0;
}