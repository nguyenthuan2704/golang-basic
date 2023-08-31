package main

import (
	"fmt"
	"math"
)

func main() {
	/*exported-name : biến được sử dụng bởi 1 package khác gọi vào - tên biến phải viết Hoa - nếu viết thường thì là unexported-name*/
	fmt.Println(math.Pi) /*để gọi số Pi có trong thư viện math , dùng chữ Pi viết hoa*/
}
