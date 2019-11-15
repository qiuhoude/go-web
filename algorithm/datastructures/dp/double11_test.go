package dp

import (
	"fmt"
	"testing"
)

/*

双11 满减问题

购物车中 n 个商品 , 满200减50 ,挑选其中的商品 最小达到满减的条件

w 满减的券的额度
price 购物车中商品的价格
*/
func double11(w int, price []int) {
	n := len(price)

	// 定义状态 f(商品id,商品价格)
	states := make([][]bool, n)
	wc := 3 * w
	for i := range states {
		states[i] = make([]bool, wc+1) // 满减是要超过,所以价格要大于w
		for j := range states[i] {
			states[i][j] = false
		}
	}

	// 首个状态
	states[0][0] = true        // 首个商品不买
	states[0][price[0]] = true // 首个商品买
	// 将问题分解成 n个步骤来解决
	for i := 1; i < n; i++ { // 遍历每个商品
		for j := 0; j <= wc; j++ { // 第i个商品不买的情况
			if states[i-1][j] { // 前一个商品已购买的情况
				states[i][j+0] = states[i-1][j] // 不购买 j不用加上价格
			}
		}

		for j := 0; j <= wc-price[i]; j++ { // 第i个商品需要购买
			if states[i-1][j] {
				states[i][j+price[i]] = true
			}
		}
	}

	min := -1
	for j := w; j < wc+1; j++ {
		if states[n-1][j] { // 最小满足条件的最小价格
			min = j // 找到最低价格
			break
		}
	}
	if min == -1 {
		fmt.Println("无解:", min)
		return
	}
	fmt.Println("最低价格是:", min)
	tmp := min
	// 逆推有哪些商品
	sum := 0
	for i := n - 1; i >= 0; i-- {
		if tmp-price[i] >= 0 && states[i][tmp-price[i]] {
			// 最小商品价格减去一个商品的价格,去状态表查看中是否有购买,如果有购买就是的
			fmt.Printf("需要购买第 %v个商品 价格是 %v\n", i, price[i])
			tmp -= price[i]
			sum += price[i]
		}
	}
	fmt.Println("sum:", sum)
}

func TestDouble11(t *testing.T) {
	sl := []int{3, 5, 7, 1, 2, 8, 9}
	double11(31, sl)
}
