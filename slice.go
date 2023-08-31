package main

import "fmt"

func main() {
	/*khai báo mảng có 6 phần tử nhưng chỉ có 5 giá trị , giá trị không khai báo mặc định là = 0*/
	primes := [6]int{1, 2, 3, 4, 5}
	/*vế đầu trong slice [0] biểu hiện là bỏ qua bao nhiêu phần tử - ở đây không bỏ phần tử nào nên bắt đầu sẽ là 1*/
	/*vế hai trong slice [2] biểu hiện là lấy bao nhiêu phần tử - ở đây là lấy 2 phần tử*/
	var s []int = primes[0:2] // kết quả in ra là [1 2]
	var x []int = primes[1:4] // kết quả in ra là [2 3 4]
	fmt.Println(s)
	fmt.Println(x)
	names := [4]string{"Micheal", "John", "Stephen", "Peter"}
	q := names[0:2]
	w := names[1:3]
	fmt.Println(q, w)

	w[0] = "XXX"
	fmt.Println(q, w)
	fmt.Println(names)
}
