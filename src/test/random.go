package test

import (
	"fmt"
	"math/rand"
)

// 生成随机数
func Random() {
	num := rand.Int()
	fmt.Println(num)
}
