package graph

import (
	"fmt"
	"github.com/qiuhoude/go-web/algorithm/datastructures/queue"
)

/*
a*算法
h(i): 启发函数(heuristic function),是到目标点的曼哈顿距离(Manhattan distance)
	曼哈顿距离: 指两点之间横纵坐标的距离之和 h() = abs(x1-x2) + abs(y1-y2)
g(i): 起点到该点的距离
f(i) = h(i) + g(i)
*/

type point struct {
	f      int    // f(i) = h(i) + g(i)
	g      int    //起点到该点的距离
	x, y   int    //坐标 m[y][x]
	parent *point // 父节点
}

func (p *point) String() string {
	return fmt.Sprintf("(%v,%v,%v)", p.x, p.y, p.g)
}

func newPoint(x, y, g, f int, p *point) *point {
	return &point{
		x:      x,
		y:      y,
		g:      g,
		f:      f,
		parent: p,
	}
}

func comparePoint(v1, v2 interface{}) int {
	p1 := v1.(*point)
	p2 := v2.(*point)
	if p1.f > p2.f {
		return 1
	} else if p1.f < p2.f {
		return -1
	}
	return 0
}

func equalPoint(v1, v2 interface{}) bool {
	p1 := v1.(*point)
	p2 := v2.(*point)
	return p1.x == p2.x && p1.y == p2.y
}

// 曼哈顿距离
func hManhattan(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func xyToId(x, y, width int) int {
	return y*width + x
}

const (
	positive = 10 // 上下左右距离是 10
	opposite = 14 // 对角是距离是 14
)

func checkToXyAvail(x, y int, m [][]uint) bool {
	high := len(m)
	width := len(m[0])
	if y < 0 || y >= high || x < 0 || x >= width { // 超过边界
		return false
	}
	if m[y][x] == 1 { // 障碍物不能走
		return false
	}
	return true
}

func AstarSearch(sx, sy, tx, ty int, m [][]uint) {
	// 都是以左上为原点
	high := len(m)
	width := len(m[0])

	visited := make(map[int]bool)               // 相当于有些教程中的 close表
	que := queue.NewPriorityQueue(comparePoint) //  open表
	visited[xyToId(sx, sy, width)] = true
	que.Enqueue(newPoint(sx, sy, 0, 0, nil)) // 加入起点坐标
	var findPoint *point
out:
	for !que.IsEmpty() {
		curP := que.Dequeue().(*point)
		// 逐个遍历 上下左右 左上右上 左下右下
		for i := -1; i <= 1; i++ { // y轴
			for j := -1; j <= 1; j++ { // x轴
				afY := curP.y + i // 计算后的 x, y
				afX := curP.x + j
				if !checkToXyAvail(afX, afY, m) || (i == 0 && j == 0) {
					continue
				}
				// 有障碍物斜边是否能跨越
				if i == -1 && j == -1 { // 左上
					if !checkToXyAvail(curP.x-1, curP.y, m) || !checkToXyAvail(curP.x, curP.y-1, m) {
						continue
					}
				} else if i == 1 && j == -1 { // 左下
					if !checkToXyAvail(curP.x-1, curP.y, m) || !checkToXyAvail(curP.x, curP.y+1, m) {
						continue
					}
				} else if i == -1 && j == 1 { // 右上
					if !checkToXyAvail(curP.x+1, curP.y, m) || !checkToXyAvail(curP.x, curP.y-1, m) {
						continue
					}
				} else if i == 1 && j == 1 { //右下
					if !checkToXyAvail(curP.x+1, curP.y, m) || !checkToXyAvail(curP.x, curP.y+1, m) {
						continue
					}
				}

				isOpposite := abs(i) == 1 && abs(j) == 1 // 是否为对角
				h := hManhattan(afX, afY, tx, ty)        // 曼哈顿距离
				g := 0
				if isOpposite {
					g = curP.g + opposite // 对角+14
				} else {
					g = curP.g + positive
				}
				f := h + g //f(i) = h(i) + g(i)
				id := xyToId(afX, afY, width)
				np := newPoint(afX, afY, f, g, curP)
				if visited[id] {
					p2 := que.Remove(np, equalPoint)
					if p2 != nil {
						np2 := p2.(*point)
						np2.g = g
						np2.f = f
						// 此处不要更新父节点值
						que.Enqueue(np2)
					}
				} else {
					que.Enqueue(np)
					visited[id] = true // 设置已访问
				}
				if afX == tx && afY == ty {
					findPoint = np
					break out
				}
			}
		}
	}

	if findPoint != nil {
		m2 := make([][]uint, high)
		for i := range m2 {
			m2[i] = make([]uint, width)
			copy(m2[i], m[i])
		}
		cur := findPoint
		var i uint = 2
		for ; cur != nil; cur = cur.parent {
			m2[cur.y][cur.x] = i
			i++
		}
		printMap(m2)
	} else {
		fmt.Println("不可达")
	}
}

func printMap(m [][]uint) {
	//fmt.Printf("%4d", -1)
	//for i := 0; i < len(m[0]); i++ {
	//	fmt.Printf("%4d", i)
	//}
	//fmt.Println()
	for i := range m {
		//fmt.Printf("%4d", i)
		for j := range m[i] {
			if j != 0 {
			}
			fmt.Printf("%4d", m[i][j])
		}
		fmt.Println()
	}
}
