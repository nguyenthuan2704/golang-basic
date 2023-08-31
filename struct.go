package main

import "fmt"

/*dùng strucs để tạo ra cấu trúc dữ liệu mới dựa trên những cái có sẵn*/

/*
khai báo dùng chữ Hoa ở đầu tên là kiểu khai báo exported: ở ngoài có thể nhìn thấy và dùng nó - trái ngược
với viết thường là unexported: package ngoài nhìn vào sẽ ko thấy
*/
type Vertex struct {
	/*X , Y viết hoa tương tự exported*/
	/*	X int
		Y int*/
	/*cách khai báo khác*/
	X, Y int
}

var (
	v1 = Vertex{4, 6}  // ra kết quả như bình thường {4 6}
	v2 = Vertex{X: 10} // chỉ in ra X {10} và Y = 0
	v3 = Vertex{}      // in ra cả 2 bằng 0 do X và Y đều bằng 0
	z  = &Vertex{7, 8} // in ra z &{7 8}
)

func main() {
	fmt.Println(Vertex{1, 2})
	/*khi sử dụng struct có thể can thiệp vào các biến của struct đó*/
	v := Vertex{1, 2}
	v.X = 4
	v.Y = 5
	fmt.Println(v.X, v.Y)
	/*sử dụng con trỏ pointer cho struct*/
	p := &v
	p.X = 1e9 /*1e9 = 10 mủ 9*/
	fmt.Println(v)

	fmt.Println(v1, v2, v3, z)
}
