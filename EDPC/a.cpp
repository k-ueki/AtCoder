#include <iostream>
#include <vector>
using namespace std;

int min(int a, int b) {
	if (a > b) return b;
	return a;
}

int abs(int a) {
	if(a < 0) return a * (-1);
	return a;
}

int main() {
	int N;
	cin >> N;
	vector<int> h(N);
	for(int i = 0; i < N; ++i) cin >> h[i];


	vector<int> dp(N);
	dp[0] = 0;
	for(int i = 1; i < N; ++i) {
		if(i == 1){
			dp[i] = abs(h[i] - h[i-1]);
		} else {
			dp[i] = min(dp[i-1] + abs(h[i] - h[i-1]), dp[i-2] + abs(h[i] - h[i-2]));
		}
	}

	cout << dp[N-1] << endl;
}

