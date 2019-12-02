package leetcode

import "testing"

//207. 课程表  https://leetcode-cn.com/problems/course-schedule/
// 拓扑排序问题,判断有向图环的问题库 使用kahn算法

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) < 1 {
		return true
	}

	inDegree := make([]int, numCourses) // 入度的数量
	nextCourses := make([][]int, numCourses)

	for i := range prerequisites {
		from := prerequisites[i][1]
		to := prerequisites[i][0]
		inDegree[to]++
		nextCourses[from] = append(nextCourses[from], to)
	}
	var que []int
	for course := range inDegree {
		if inDegree[course] == 0 {
			que = append(que, course)
		}
	}
	if len(que) == 0 {
		return false
	}
	cnt := 0
	for len(que) > 0 {
		course := que[0]
		que = que[1:]
		cnt++
		if nextCourses[course] != nil {
			for _, toCourse := range nextCourses[course] {
				inDegree[toCourse]--
				if inDegree[toCourse] == 0 {
					que = append(que, toCourse)
				}
			}
		}
	}
	return cnt == numCourses
}

func TestCanFinish(t *testing.T) {
	tests := []struct {
		arg1 int
		arg2 [][]int
		want bool
	}{
		{3, [][]int{{1, 0}, {2, 1}}, true},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, true},
	}
	for _, tt := range tests {
		if got := canFinish(tt.arg1, tt.arg2); got != tt.want {
			t.Errorf("canFinish() => got=%v  want=%v", got, tt.want)
		}
	}
}
