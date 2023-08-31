package main

import "fmt"

/*Golang có các kiểu dữ liệu : bool , string , int , uint(số nguyên không nhận dấu trừ , tức là không phải số âm) , byte , rune , float , complex */

/*khai báo biến nhưng không gán giá trị thì mặc định là 0 với kiểu int và false với kiểu bool*/
var c, python, java bool
var i int

/*khai báo biến cách 2 : dùng var thì phải xác định kiểu của biến và sử dụng dấu = để gán giá trị*/
var x, j int = 1, 2

func main() {
	var p, rust, javascript = true, false, true
	fmt.Println(i, c, python, java)
	fmt.Println(x, j, p, rust, javascript)
	/*cách 3: không sử dụng var thì dùng dấu := để gán giá trị , chương trình tự động hiểu kiểu cuủa biến*/
	k := 3
	fmt.Println(k)
}
