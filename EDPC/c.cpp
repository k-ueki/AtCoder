#include <iostream>
#include <vector>
using namespace std;

int maxRange(vector<int> a) {
	int maxNum = 0;
	for(int i = 0;i < (int)a.size(); ++i) {
		maxNum = max(maxNum, a[i]);
	}
	return maxNum;
}

int max(int a, int b) {
	if(a>b) return a;
	return b;
}

int main() {
	int N;
	cin >> N;
	vector<int> a(N), b(N), c(N);
	for(int i = 0; i < N; ++i) cin >> a[i] >> b[i] >> c[i];

	vector<vector<int> > dp(N, vector<int>(3));
	for (int i = 0; i < N; ++i) {
		for (int j = 0; j < 3; ++j) {
			if(i == 0) {
				dp[i][0] = a[i];
				dp[i][1] = b[i];
				dp[i][2] = c[i];
			} else {
				dp[i][0] = a[i] + max(dp[i-1][1], dp[i-1][2]);
				dp[i][1] = b[i] + max(dp[i-1][0], dp[i-1][2]);
				dp[i][2] = c[i] + max(dp[i-1][0], dp[i-1][1]);
			}
		}
	}


	cout << maxRange(dp[N-1]) << endl;
}

