#include <iostream>
#include <vector>
using namespace std;

long long INF = 1LL << 60;

int max(int a, int b ) {
	if(a > b) return a;
	return b;
}

int main() {
	int N,W;
	cin >> N >> W;
	vector<long long> weight(N), value(N);
	for(int i = 0; i < N; ++i) cin >> weight[i] >> value[i];


}

