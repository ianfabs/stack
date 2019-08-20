package main

//IntArrayStack is a stack implementation of type int
type IntArrayStack struct {
	data []int
}

func (stack *IntArrayStack) clear() {
	stack.data = []int{}
}

func (stack *IntArrayStack) push(i int) {
	stack.data = append([]int{i}, stack.data...)
}

func (stack *IntArrayStack) pop() int {
	p := stack.data[0]
	stack.data = stack.data[1:len(stack.data)]
	return p
}

func (stack *IntArrayStack) size() int {
	return len(stack.data)
}

func (stack *IntArrayStack) find(el int) bool {
	for _, n := range stack.data {
		if el == n {
			return true
		}
	}
	return false
}

func (stack *IntArrayStack) indexOf(el int) int {
	for i, n := range stack.data {
		if el == n {
			return i
		}
	}
	return -1
}
