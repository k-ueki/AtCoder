#include <bits/stdc++.h>
#define rep(i,n) for(int i=0;i<(n);i++)
typedef long long ll;

using namespace std;

int main(){
    int n;
    cin >> n;
    vector<int> a(n);
    int s = 0;
    rep(i,n){
        cin >> a[i];
        s ^= a[i];
    }

    rep(i,n)cout << (s^a[i]) << " ";
    cout << endl;
}
