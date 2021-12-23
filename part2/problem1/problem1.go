package main

import "fmt"

func TopDown(n int) int {
	// your code here
}

func main() {
	fmt.Println(TopDown(0))  //0
	fmt.Println(TopDown(1))  //1
	fmt.Println(TopDown(2))  //1
	fmt.Println(TopDown(3))  //2
	fmt.Println(TopDown(5))  //5
	fmt.Println(TopDown(6))  //8
	fmt.Println(TopDown(7))  //13
	fmt.Println(TopDown(9))  //34
	fmt.Println(TopDown(10)) //55
}
