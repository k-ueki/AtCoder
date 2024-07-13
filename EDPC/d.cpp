#include <iostream>
#include <vector>
using namespace std;

long long INF = 1LL << 60;

int max(int a, int b) {
	if(a > b) return a;
	return b;
}

int main() {
	int N, W;
	cin >> N >> W;
	vector<long long> weight(N), value(N);
	for(int i = 0; i < N; ++i) cin >> weight[i] >> value[i];

	vector<vector<long long> > dp(N+1, vector<long long>(W+1, 0));
	for(int i = 0; i < N; ++i) {
		for(int w = 0; w <= W; ++w) {
			if(w - weight[i] >= 0) {
				dp[i+1][w] = max(dp[i][w], dp[i][w - weight[i]] + value[i]);
			}

			dp[i+1][w] = max(dp[i+1][w], dp[i][w]);
		}
	}

	cout << dp[N][W] << endl;

}

