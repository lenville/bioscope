/*
 * Caculate a slice float64's average value
 * 编写计算一个类型是 float64 的 slice 的平均值的代码。在稍候的练习 Q5 中 将会改写为函数。
 */

package main
import "fmt"

func main() {
	xs := []float64{1.0, 2.0, 3.0, 4.0}
	sum := 0.0
	avg := 0.00
	fmt.Print("Slice: [ ")
	switch len(xs) {
	case 0:
		avg = 0
	default:
		for _, v := range xs {
			fmt.Printf("%f ", v)
			sum += v
		}
		avg = sum / float64(len(xs))
		fmt.Print("]\n")
	}

	fmt.Printf("Average: %f\n", avg)
}
