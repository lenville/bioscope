package main

/*
 * Sort two int number in order
 */
func sortTowNumber(a int, b int) (smaller int, bigger int) {
	if a > b {
		return b, a
	} else {
		return a, b
	}
}

func main() {
	var a, b = sortTowNumber(10, 1)
	println(a, b)
}
