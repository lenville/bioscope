package stack

/*
	mkdir $GOPATH/src/stack
	cp pushpop_test.go $GOPATH/src/stack
	cp stack.go $GOPATH/src/stack
	go test
*/

type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) Push(k int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *Stack) Pop() int {
	s.i--
	return s.data[s.i]
}
