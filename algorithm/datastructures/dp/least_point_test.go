package dp

import (
	"math"
	"sort"
	"testing"
)

/*

分治思想
分解：将原问题分解成一系列子问题；
解决：递归地求解各个子问题，若子问题足够小，则直接独立求解；
合并：将子问题的结果合并成原问题

最近点问题
借助 归并排序的思想
*/

// 1. 求 数轴上最近两点的距离 (有序的数)
func Test_leastPointOnAxis(t *testing.T) {
	axis := []float64{1, 5, 8, 11, 13, 15}
	t.Log(leastPointOnAxis(axis))
}

func leastPointOnAxis(axis []float64) float64 {
	n := len(axis)
	if n < 2 {
		return 0
	}
	return calcLeastPoint(axis, 0, n-1)
}

func calcLeastPoint(axis []float64, p, r int) float64 {
	if p+1 >= r { // 只有2个点时
		return math.Abs(axis[r] - axis[p])
	}
	if p+2 >= r { // 三个点的时候
		return math.Min(math.Abs(axis[p+1]-axis[p]), math.Abs(axis[r]-axis[p+1]))
	}
	mid := p + ((r - p) >> 1)
	min1 := calcLeastPoint(axis, p, mid)   //左边最小
	min2 := calcLeastPoint(axis, mid+1, r) // 有右边最小
	// 左右缝隙的距离
	return math.Min(min1, math.Min(min2, math.Abs(axis[mid+1]-axis[mid])))
}

//2. n个点在公共空间中，求出所有点对的欧几里得距离最小的点对。
//思路: 一样使用分治法, 按照 x ,y 两个轴进行切分
// https://blog.csdn.net/sinat_35678407/article/details/82874216

type Point struct {
	x, y float64
}

// 两点之间的距离
func distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func leastPointDistance(points []Point) float64 {
	// 找出x y分别的最大最小值
	n := len(points)
	if n < 2 {
		return 0
	}
	if n == 2 {
		return distance(points[0], points[1])
	}
	if n == 3 {
		return math.Min(distance(points[0], points[1]),
			math.Min(distance(points[0], points[2]), distance(points[1], points[2])))
	}
	// 进行区域划分
	mid := n>>1 - 1
	d1 := leastPointDistance(points[:mid+1])   // 左区域
	d2 := leastPointDistance(points[n-mid-1:]) // 右区域
	d := math.Min(d1, d2)                      // 左右区域缝隙的点进行比较
	return merge(points, d, mid)
}

// 合并比较
func merge(points []Point, minDic float64, mid int) float64 {
	var left, right []Point
	// 左右进行分区 对x轴进距离筛选,距离要小于minDic
	for i := range points {
		if points[i].x <= points[mid].x && math.Abs(points[i].x-points[mid].x) < minDic { // 左边区域
			left = append(left, points[i])
		} else if points[i].x > points[mid].x && math.Abs(points[i].x-points[mid].x) < minDic { //右边区域
			right = append(right, points[i])
		}
	}
	// 对右边y轴进行排序
	sort.Slice(right, func(i, j int) bool {
		return right[i].y > right[j].y
	})

	// 用左边区域的每个点跟右边区域符合条件的点进行比较
	for i := range left {
		index := 0
		for ; index < len(right) && left[i].y < right[index].y; index++ {
			// 左边的点y轴的高度 要高于 右边区域的才进行比较
		}
		for j := range right { // 右边区域
			if j > 6 || index+j >= len(right) { // 遍历右边坐标y轴方向上 距离上界最近的的6个点
				break
			}
			calcDis := distance(left[i], right[j+index])
			if calcDis < minDic {
				minDic = calcDis
			}
		}
	}
	return minDic
}

func Test_leastPointDistance(t *testing.T) {
	var points []Point
	points = append(points, Point{0, 15})
	points = append(points, Point{0, 23})
	points = append(points, Point{1, 0.7})
	points = append(points, Point{0, 1.8})

	sort.Slice(points, func(i, j int) bool {
		return points[i].x > points[j].x
	})
	leastDis := leastPointDistance(points)
	t.Log("最短距离 ", leastDis)
}
