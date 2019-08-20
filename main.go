package main

import "fmt"

func main() {
	stack := IntArrayStack{data: []int{}}
	stack.push(10)
	stack.push(11)
	stack.push(1)
	stack.push(0)
	stack.push(5)

	fmt.Println(stack.size())
	fmt.Println(stack.pop())
	//printSlice(stack.data)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
