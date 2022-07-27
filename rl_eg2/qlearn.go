package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	N_STATES     int     = N * N        // 状态数
	ACTION_NUM   int     = M            // 行为数
	EPSILON      float64 = 0.8          // 记忆系数
	ALPHA        float64 = 0.1          // 步长
	GAMMA        float64 = 0.9          // 更新系数
	MAX_EPISODES int     = 20           // 最大训练局数
	QTABLE_PATH  string  = "QTable.txt" // Q表保存路径
)

// Q表
type Qtable [N_STATES][ACTION_NUM]float64

// 打印Q表
func (this Qtable) show() {
	for i := 0; i < N_STATES; i++ {
		for j := 0; j < ACTION_NUM; j++ {
			fmt.Printf("%.2f ", this[i][j])
		}
		fmt.Println()
	}
}

// 保存Q表
func (this Qtable) save() {
	filePath := QTABLE_PATH
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0000)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < N_STATES; i++ {
		for j := 0; j < ACTION_NUM; j++ {
			writer.WriteString(fmt.Sprintf("%.2f ", this[i][j]))
		}
		writer.WriteByte('\n')
	}
	writer.Flush()
}

// Q表建立
func build_Qtable() Qtable {
	var ans Qtable
	return ans
}

// 选择行为
func choose(state int, qTable Qtable) int {
	var res int
	var s_action [ACTION_NUM]float64
	copy(s_action[:], qTable[state][:])
	if rand.Float64() > EPSILON || s_action == [ACTION_NUM]float64{0, 0, 0, 0} {
		res = rand.Intn(ACTION_NUM)
	} else {
		for i := range s_action {
			if s_action[res] < s_action[i] {
				res = i
			}
		}
	}
	return res
}

// 状态奖励函数
func reward(state int, action int) float64 {
	if (state == N_STATES-2 && action < 2) || (state == N_STATES-1-N && action >= 2 && action < 4) {
		return 1
	}
	return 0
}

func max(arr []float64) float64 {
	res := 0.0
	for i := range arr {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}

func qLearn() Qtable {
	episode := 0
	q_table := build_Qtable()
	env := newMaze()
	env.show()
	ui := newUI(env)
	go func() {
		for true {
			ui.flush()
		}
	}()
	go func() {
		for episode = 0; episode < MAX_EPISODES; episode++ {
			env.arr[N][N] = 0
			env.pos[0], env.pos[1] = 1, 1
			step_counter := 0
			for !env.isWin() {
				S := env.getState()
				A := choose(S, q_table)
				R := reward(S, A)
				env.act(A)
				q_predict := q_table[S][A]
				var q_target float64
				if !env.isWin() {
					q_target = R + GAMMA*max(q_table[env.getState()][:])
				} else {
					q_target = R
				}
				q_table[S][A] += ALPHA * (q_target - q_predict)
				step_counter += 1
				env.show()
			}
			fmt.Println("step =", step_counter)
			q_table.show()
			fmt.Println("--------------------", episode, "-------------------")
			time.Sleep(time.Second)
		}
	}()
	ui.myWindow.ShowAndRun()
	if episode == MAX_EPISODES {
		q_table.save()
	}
	return q_table
}
