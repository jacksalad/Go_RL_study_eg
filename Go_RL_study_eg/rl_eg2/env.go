package main

import (
	"fmt"
	"time"
)

const (
	N int = 10 // 迷宫长度
	M int = 6  // 行为数
)

// 环境结构体
type Maze struct {
	pos [2]int
	arr [N + 2][N + 2]int
}

// 构造函数
func newMaze() *Maze {
	var pos [2]int = [2]int{1, 1}
	var arr [N + 2][N + 2]int
	arr[1][1] = 1
	for i := 0; i < N+2; i++ {
		arr[i][0] = 3
		arr[i][N+1] = 3
		arr[0][i] = 3
		arr[N+1][i] = 3
	}
	arr[3][4] = 2
	arr[3][5] = 2
	arr[4][4] = 2
	arr[7][4] = 2
	arr[6][4] = 2
	arr[5][5] = 2
	arr[1][8] = 2
	arr[2][8] = 2
	arr[3][8] = 2
	arr[3][7] = 2
	arr[3][6] = 2
	arr[6][9] = 2
	arr[4][9] = 2
	arr[7][8] = 2
	arr[7][9] = 2
	arr[1][8] = 2
	arr[10][6] = 2
	arr[9][6] = 2
	arr[10][1] = 2
	arr[10][2] = 2
	arr[5][1] = 2
	arr[5][2] = 2
	return &Maze{pos, arr}
}

// 打印输出
func (this *Maze) show() {
	for i := 0; i < N+2; i++ {
		for j := 0; j < N+2; j++ {
			fmt.Print(this.arr[i][j], " ")
		}
		println()
	}
	fmt.Println()
	time.Sleep(time.Second / 5)
}

// 行为执行
func (this *Maze) act(x int) {
	switch {
	case x < 2: // Right
		if this.pos[1]+1 <= N && this.arr[this.pos[0]][this.pos[1]+1] != 2 {
			this.arr[this.pos[0]][this.pos[1]] = 0
			this.arr[this.pos[0]][this.pos[1]+1] = 1
			this.pos[1]++
		}
	case x < 4: // Down
		if this.pos[0]+1 <= N && this.arr[this.pos[0]+1][this.pos[1]] != 2 {
			this.arr[this.pos[0]][this.pos[1]] = 0
			this.arr[this.pos[0]+1][this.pos[1]] = 1
			this.pos[0]++
		}
	case x == 4: // Left
		if this.pos[1]-1 > 0 && this.arr[this.pos[0]][this.pos[1]-1] != 2 {
			this.arr[this.pos[0]][this.pos[1]] = 0
			this.arr[this.pos[0]][this.pos[1]-1] = 1
			this.pos[1]--
		}
	case x == 5: // Up
		if this.pos[0]-1 > 0 && this.arr[this.pos[0]-1][this.pos[1]] != 2 {
			this.arr[this.pos[0]][this.pos[1]] = 0
			this.arr[this.pos[0]-1][this.pos[1]] = 1
			this.pos[0]--
		}
	}
}

// 判断获胜
func (this *Maze) isWin() bool {
	if this.pos == [2]int{N, N} {
		return true
	}
	return false
}

// 获取状态数
func (this Maze) getState() int {
	return (this.pos[0]-1)*N + this.pos[1] - 1
}
