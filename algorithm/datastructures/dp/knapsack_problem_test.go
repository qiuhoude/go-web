package dp

import "testing"

//背包问题9讲 https://www.kancloud.cn/kancloud/pack/70125

/*
0-1 背包问题 求总总量最大
现在我们有 n 个物品，每个物品的重量不等，并且不可分割。我们现在期望选择几件物
品，装载到背包中。在不超过背包所能装载重量的前提下，如何让背包中物品的总重量最大

我们可以把物品依次排列,整个问题就分解为了 n 个阶段，每个阶段
对应一个物品怎么选择。先对第一个物品进行处理，选择装进去或者不装进去，然后再递归地处
理剩下的物品

使用回溯方式
i 表示考察到哪个物品了
curW 当前背包的重量
limitW 重量限制
weight 物品列表
*/
func bag01Problem(i, cw, limitW int, weight []int, ret *int) {
	al1Cnt++
	if cw == limitW || // 表示装满了
		i == len(weight) { // 表示已经考察完所有的物品
		if cw > *ret { //找到了
			*ret = cw
		}
		return
	}

	bag01Problem(i+1, cw, limitW, weight, ret) // 选择不装第 i 个物品
	if cw+weight[i] <= limitW {                // 已经超过可以背包承受的重量的时候，就不要再装了
		bag01Problem(i+1, cw+weight[i], limitW, weight, ret) //选择 第 i 个物品
	}
}

/*
使用回溯方式记忆化搜索的方式,将f(i,cw) 重复的递归调用去除掉
f(i,cw) 建立 i,cw 的状态数组
*/
func bag01ProblemRS(limitW int, weight []int, ret *int) {
	// 初始化state
	state := make([][]bool, len(weight))
	for i := range state {
		state[i] = make([]bool, limitW+1)
		for j := range state[i] {
			state[i][j] = false
		}
	}
	bag01RememberSearch(0, 0, limitW, weight, ret, state)
}

func bag01RememberSearch(i, cw, limitW int, weight []int, ret *int, state [][]bool) {
	al2Cnt++
	if cw == limitW || // 表示装满了
		i == len(weight) { // 表示已经考察完所有的物品
		if cw > *ret { //找到了
			*ret = cw
		}
		return
	}
	if state[i][cw] { //之前已经访问过
		return
	}
	state[i][cw] = true                                      // 已经访问过
	bag01RememberSearch(i+1, cw, limitW, weight, ret, state) // 不选择i
	if cw+weight[i] <= limitW {
		bag01RememberSearch(i+1, cw+weight[i], limitW, weight, ret, state)
	}
}

/*
动态规划的方式

把整个求解过程分为 n 个阶段，每个阶段会决策一个物品是否放到背包中。每个物品决策
（放入或者不放入背包）完之后，背包中的物品的重量会有多种情况，也就是说，会达到多种不
同的状态，对应到递归树中，就是有很多不同的节点

把每一层重复的状态（节点）合并，只记录不同的状态，然后基于上一层的状态集合，来推
导下一层的状态集合。我们可以通过合并每一层重复的状态，这样就保证每一层不同状态的个数
都不会超过 w 个（w 表示背包的承载重量）

用一个二维数组 states[n][w+1]，来记录每层可以达到的不同状态

第 0 个（下标从 0 开始编号）物品的重量是 2，要么装入背包，要么不装入背包，决策完之
后，会对应背包的两种状态，背包中物品的总重量是 0 或者 2。我们用 states[0][0]=true 和
states[0][2]=true 来表示这两种状态

第 1 个物品的重量也是 2，基于之前的背包状态，在这个物品决策完之后，不同的状态有 3
个，背包中物品总重量分别是 0(0+0)，2(0+2 or 2+0)，4(2+2)。我们用 states[1][0]=true，
states[1][2]=true，states[1][4]=true 来表示这三种状态

以此类推，直到考察完所有的物品后，整个 states 状态数组就都计算好了。我把整个计算的过
程画了出来，你可以看看。图中 0 表示 false，1 表示 true。我们只需要在最后一层，找一个值
为 true 的最接近 w（这里是 9）的值，就是背包中物品总重量的最大值。

然后通过当前阶段的状态集合，来推导下一个阶段的状态集合,动态地往前推进

*/
func knapsack01ProblemDP1(w int, weight []int) int {
	n := len(weight)
	// 初始化状态
	states := make([][]bool, n)
	for i := range states {
		states[i] = make([]bool, w+1)
		for j := range states[i] {
			states[i][j] = false
		}
	}
	// 第0个weight 的两种状态
	states[0][0] = true         //不放0个物品, 第一行的数据要特殊处理，可以利用哨兵优化
	states[0][weight[0]] = true // 放第0个物品

	// 动态规划状态转移
	for i := 1; i < n; i++ {
		for j := 0; j <= w; j++ { // 不把第 i 个物品放入背包
			if states[i-1][j] { // 前一个物品已经放入
				states[i][j] = states[i-1][j]
			}
		}
		for j := 0; j <= w-weight[i]; j++ { //  第 i 个物品放入背包
			if states[i-1][j] {
				states[i][j+weight[i]] = true
			}
		}
	}
	// 结果
	for i := w; i >= 0; i-- {
		if states[n-1][i] {
			return i
		}
	}
	return 0
}

/*
使用 1维数组来做
*/
func knapsack01ProblemDP2(w int, weight []int) int {
	n := len(weight)
	// 初始化状态
	states := make([]bool, w+1)
	for i := range states {
		states[i] = false
	}
	states[0] = true
	states[weight[0]] = true
	// 动态规划状态转移,
	for i := 1; i < n; i++ {
		//j 需要从大到小来处理。如果我们按照 j 从小到大处理的话，会出现 for 循环重复计算的问题
		for j := w - weight[i]; j >= 0; j-- { //  第 i 个物品放入背包
			if states[j] {
				states[j+weight[i]] = true
			}
		}
	}
	// 结果
	for i := w; i >= 0; i-- {
		if states[i] {
			return i
		}
	}
	return 0
}

// 带价值的背包问,不能使用记忆化搜索的方式解答

// 使用回溯的方式
func knapsack02(i, cw, cv, limitW int, weight, value []int, ret *int) {
	al1Cnt++
	if cw == limitW || // 表示装满了
		i == len(weight) { // 表示已经考察完所有的物品
		if cv > *ret { //找到了
			*ret = cv
		}
		return
	}

	knapsack02(i+1, cw, cv, limitW, weight, value, ret) // 选择不装第 i 个物品
	if cw+weight[i] <= limitW {                         // 已经超过可以背包承受的重量的时候，就不要再装了
		knapsack02(i+1, cw+weight[i], cv+value[i], limitW, weight, value, ret) //选择 第 i 个物品
	}
}

// 使用动态规划的方式
func knapsack02Dp(w int, weight, value []int) int {
	n := len(weight)
	// f(i,cw,cv) cv 就是存储的值
	// 初始化states
	states := make([][]int, n)
	for i := range states {
		states[i] = make([]int, w+1)
		for j := range states[i] {
			states[i][j] = -1
		}
	}
	states[0][0] = 0                // 第物品放不放进背包
	states[0][weight[0]] = value[0] // 第一放
	for i := 1; i < n; i++ {
		for j := 0; j <= w; j++ { // 不把第 i 个物品放入背包
			if states[i-1][j] >= 0 { // 前一个物品已经放入
				states[i][j] = states[i-1][j]
			}
		}
		for j := 0; j <= w-weight[i]; j++ { //  第 i 个物品放入背包
			if states[i-1][j] >= 0 {
				// 上一val放入背包的值 + 本次背包的值
				v := states[i-1][j] + value[i]
				if v > states[i][j+weight[i]] { // 相同总量下选取 value值大的
					states[i][j+weight[i]] = v
				}
			}
		}
	}
	// 在最后找出最大值
	maxV := 0
	for _, v := range states[n-1] {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

var al1Cnt int = 0
var al2Cnt int = 0

func TestBagProblem(t *testing.T) {
	weight := []int{2, 2, 4, 6, 3}
	limit := 9

	ret1 := 0
	bag01Problem(0, 0, limit, weight, &ret1)
	t.Log("maxW ret1:", ret1, "al1Cnt:", al1Cnt)

	ret2 := 0
	bag01ProblemRS(limit, weight, &ret2)
	t.Log("maxW ret2:", ret2, "al2Cnt:", al2Cnt)

	ret3 := knapsack01ProblemDP1(limit, weight)
	t.Log("maxW ret3:", ret3)

	ret4 := knapsack01ProblemDP2(limit, weight)
	t.Log("maxW ret4:", ret4)

	// 带价值的背包问题
	value := []int{3, 4, 8, 9, 6}
	ret5 := 0
	knapsack02(0, 0, 0, limit, weight, value, &ret5)
	t.Log("maxV ret5:", ret5)

	ret6 := knapsack02Dp(limit, weight, value)
	t.Log("maxV ret6:", ret6)

}
