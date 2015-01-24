/*
 * Terrific string reverse
 */

package main
import "fmt"

func main() {
	s := "foobar"
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i] //这里用到了平行赋值的方法
	}
	fmt.Printf("Previous  string: %s\n", s)
	fmt.Printf("Following string: %s\n", string(a))
}
