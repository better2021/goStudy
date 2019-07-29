package randomNum
import(
	"fmt"
	"math/rand"
)

// 生成随机数
func random(){
	num := rand.Int()
	fmt.Println(num)
}

