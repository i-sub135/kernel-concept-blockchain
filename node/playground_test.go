package node

import (
	"fmt"
	"testing"
)

func TestPlayground(t *testing.T) {
	var v int = 0
	for !checkNum(v) {
		v = v + 1
		fmt.Println("loop numb => ", v)
	}
	fmt.Println("final numb => ", v)
}
func checkNum(n int) bool {
	if n == 10 {
		return true
	} else {
		return false
	}

}
