package leetcode

import (
	"testing"
)

// 硬币问题 leetcode 322
// https://leetcode-cn.com/problems/coin-change/
//coins 硬币, amount 期望的金额, 返回最少需要的硬币数量，如果不可解返回-1
func CoinCharge(coins []int, amount int) int {
	dp := make([]int, amount+1) //dp存储的是 0~amount金额最小值的信息
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if coin <= i && dp[i-coin] != -1 && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
		if dp[i] > amount {
			dp[i] = -1
		}
	}

	return dp[amount]
}

func TestCoinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"[2] => 3", args{[]int{2}, 3}, -1},
		{"[2] => 4", args{[]int{2}, 4}, 2},
		{"[1,2,5] => 11", args{[]int{1, 2, 5}, 11}, 3},
		{"[1,3,5] => 11", args{[]int{1, 3, 5}, 11}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CoinCharge(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("CoinCharge() = %v, want %v", got, tt.want)
			}
		})
	}
}
