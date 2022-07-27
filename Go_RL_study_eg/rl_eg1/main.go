package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	model := qLearn()
	fmt.Println("Q-Table:\n", model)
}
