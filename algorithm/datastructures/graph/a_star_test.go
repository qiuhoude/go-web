package graph

import (
	"testing"
)

func buildMap() [][]uint {
	// 8 * 6
	m := make([][]uint, 6)
	for i := range m {
		m[i] = make([]uint, 8)
	}
	m[0][4] = 1
	m[1][4] = 1
	//m[2][4] = 1
	m[3][4] = 1
	m[4][4] = 1
	m[5][4] = 1
	return m
}

func TestGenMap(t *testing.T) {
	m := buildMap()
	printMap(m)
}

/*
0,0,0,0,0,0,0,0
0,0,0,0,1,0,0,0
0,0,0,0,1,0,0,0
0,0,s,0,1,0,t,0
0,0,0,0,1,0,0,0
0,0,0,0,1,0,0,0
*/

func TestAstarSearch(t *testing.T) {
	m := buildMap()
	AstarSearch(2, 3, 6, 3, m)
}
