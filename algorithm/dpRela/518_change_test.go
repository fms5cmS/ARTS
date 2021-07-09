package dpRela

import (
	"fmt"
	"testing"
)

func change(amount int, coins []int) int {
	// dp 数组，dp[i] 代表了凑成金额 i 的货币组合数为 dp[i]
	// dp 数组初始化，dp[0]=1，非 0 处的值均初始换为 0
	dp := make([]int, amount+1)
	dp[0] = 1
	// 遍历硬币
	for _, coin := range coins {
		fmt.Println(coin)
		for i := coin; i <= amount; i++ {
			// 递推公式
			dp[i] += dp[i-coin]
		}
		fmt.Println(dp)
	}
	return dp[amount]
}

func TestChange(t *testing.T) {
	t.Log(change(5, []int{1, 2, 5}))
}

// 遍历顺序说明（以上面测试用例）：
// 遍历 coins，代表以当前 coin 组合为 amount 的个数
// dp 数组初始化：dp[0] = 1，dp[非零] = 0
// coin = 1，dp[i] 代表了以 1 组合成各个 i 的个数
// coin = 2，dp[i] 代表了以 2、1 组合成各个 i 的个数（dp[i-dp[coin]]） + 以 1 组合成各个 i 的个数（dp[i]）
// coin = 5，dp[i] 代表了以 5、2、1 组合成各个 i 的个数 + 以 2、1 组合成各个 i 的个数