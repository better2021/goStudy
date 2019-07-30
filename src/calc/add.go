package calc

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func Params(i int, j string) (int, string) {
	return i, j
}

func init() {
	fmt.Println("init")
}
