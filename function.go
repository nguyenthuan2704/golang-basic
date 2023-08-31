package main

import "fmt"

/*có 2 cách viết khai báo tham số và hàm*/
func add(x int, y int) int {
	return x + y
}

func add_plus(x, y int) int {
	return x + y
}

/*function trong golang có thể trả về nhiều kết quả cùng lúc*/
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add_plus(25, 25))

	a, b := swap("go", "rust")
	fmt.Println(a, b)
}
