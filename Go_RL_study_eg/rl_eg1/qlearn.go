package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	N_STATES     int     = N   // 状态数
	ACTION_NUM   int     = 4   // 行为数
	EPSILON      float64 = 0.9 // 记忆系数
	ALPHA        float64 = 0.1 // 步长
	GAMMA        float64 = 0.9 // 更新系数
	MAX_EPISODES int     = 13  // 最大训练局数
)

// Q表建立
func build_Qtable() [N_STATES][ACTION_NUM]float64 {
	var ans [N_STATES][ACTION_NUM]float64
	return ans
}

// 选择行为
func choose(state int, qTable [N_STATES][ACTION_NUM]float64) int {
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
	if (state == N_STATES-2 && action == 1) || (state == N_STATES-3 && action == 3) {
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

func qLearn() [N_STATES][ACTION_NUM]float64 {
	q_table := build_Qtable()
	for episode := 0; episode < MAX_EPISODES; episode++ {
		step_counter := 0
		rd := newRoad()
		rd.show()
		for !rd.isWin() {
			S := rd.pos
			A := choose(S, q_table)
			R := reward(S, A)
			rd.act(A)
			q_predict := q_table[S][A]
			var q_target float64
			if !rd.isWin() {
				q_target = R + GAMMA*max(q_table[rd.pos][:])
			} else {
				q_target = R
			}
			q_table[S][A] += ALPHA * (q_target - q_predict)
			step_counter += 1
			rd.show()
		}
		fmt.Println("step =", step_counter)
		fmt.Println(q_table)
		fmt.Println("--------------------", episode, "-------------------")
		time.Sleep(time.Second)
	}
	return q_table
}
