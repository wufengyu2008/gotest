package main

import "fmt"

func main() {
	/* 这是我的第一个简单的程序 */
	numbers := []int{0,1,2,3,4,5,6,7,8}
	fmt.Println("Hello, World!")
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println("numbers[4:] ==", numbers[8:][0])
}
