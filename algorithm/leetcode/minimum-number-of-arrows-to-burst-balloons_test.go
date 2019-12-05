package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

//452. 用最少数量的箭引爆气球 https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/

func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool { // 结束位置升序
		if points[i][1] < points[j][1] {
			return true
		} else {
			if points[i][1] == points[j][1] {
				return points[i][0] < points[j][0]
			}
			return false
		}
	})
	//fmt.Printf("%v\n", points)
	ret := 1
	latestEnd := points[0][1]
	for i := 1; i < n; i++ {
		if points[i][0] > latestEnd {
			ret++
			latestEnd = points[i][1]
		}
	}
	return ret
}

func TestFFindMinArrowShots(t *testing.T) {
	tests := []struct {
		arg  [][]int
		want int
	}{
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}, {13, 15}, {4, 12}}, 3},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v", tt.arg)
		t.Run(name, func(t *testing.T) {
			if got := findMinArrowShots(tt.arg); got != tt.want {
				t.Errorf("findMinArrowShots() => want:%v but got:%v", tt.want, got)
			}
		})
	}

}
