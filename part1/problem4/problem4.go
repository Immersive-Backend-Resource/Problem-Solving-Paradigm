package main

import "fmt"

func BinarySearch(array []int, x int) int {
	// your code here
}

func main() {
	fmt.Println([]int{1, 1, 3, 5, 5, 6, 7}, 3)                  //2
	fmt.Println([]int{1, 2, 3, 5, 6, 8, 10}, 5)                 //3
	fmt.Println([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)  //6
	fmt.Println([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) //-1
}
