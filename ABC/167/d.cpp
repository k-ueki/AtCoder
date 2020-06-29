#include <bits/stdc++.h>
#define rep(i,n) for(int i=0;i<(n);i++)
typedef long long ll;

using namespace std;

int main(){
    ios::sync_with_stdio(false);
    int n; ll k;
    cin >> n >> k;
    vector<int> a(n);
    rep(i,n)cin >> a[i];

    vector<int> val(n+1,-1);
    vector<int> s;
    int c = 1, l = 0;
    int v = 1;
    while(val[v] == -1){
        val[v] = s.size();
        s.push_back(v);
        v = a[v-1];
    }
    c = s.size() - val[v];
    l = val[v];

    if (k < l)cout << s[k] << endl;
    else{
        k -= l;
        k %= c;
        cout << s[l+k] << endl;
    }

    return 0;
}
