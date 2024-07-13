#include <iostream>
#include <vector>
using namespace std;

long long INF = 1LL << 60;

int min(long long a, long long b) {
	if (a > b) return b;
	return a;
}

int main() {
	int N, K;
	cin >> N >> K;
	vector<int> h(N);
	for(int i = 0; i < N; ++i) cin >> h[i];

	vector<long long> dp(N, INF);
	dp[0] = 0;
	for(int i = 0; i < N; ++i) {
		for(int j = 1; j <= K && i + j < N; ++j) {
			dp[i+j] = min(dp[i+j], dp[i] + abs(h[i+j] - h[i]));
		}
	}

	cout << dp[N-1] << endl;
}

