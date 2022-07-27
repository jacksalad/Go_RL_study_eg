package main

import (
	"fmt"
	"time"
)

const N int = 10 // 路线长度

// 环境结构体
type road struct {
	pos int
	arr [N]int
}

// 构造函数
func newRoad() *road {
	pos := 0
	var arr [N]int
	arr[pos] = 1
	arr[3] = 2
	arr[5] = 2
	return &road{pos, arr}
}

// 打印输出
func (this *road) show() {
	for i := 0; i < N; i++ {
		fmt.Print(this.arr[i], " ")
	}
	fmt.Println()
	time.Sleep(time.Second / 5)
}

// 行为执行
func (this *road) act(x int) {
	switch x {
	case 0: // left run
		if this.pos-1 >= 0 && this.arr[this.pos-1] != 2 {
			this.arr[this.pos] = 0
			this.arr[this.pos-1] = 1
			this.pos--
		}
	case 1: // right run
		if this.pos+1 < N && this.arr[this.pos+1] != 2 {
			this.arr[this.pos] = 0
			this.arr[this.pos+1] = 1
			this.pos++
		}
	case 2: // left jump
		if this.pos-2 >= 0 && this.arr[this.pos-2] != 2 {
			this.arr[this.pos] = 0
			this.arr[this.pos-2] = 1
			this.pos -= 2
		}
	case 3: // right jump
		if this.pos+2 < N && this.arr[this.pos+2] != 2 {
			this.arr[this.pos] = 0
			this.arr[this.pos+2] = 1
			this.pos += 2
		}
	}
}

// 判断获胜
func (this *road) isWin() bool {
	if this.pos == N-1 {
		return true
	}
	return false
}
