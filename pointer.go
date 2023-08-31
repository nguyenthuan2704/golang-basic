package main

import "fmt"

func main() {
	/*biến i có giá tri là 55*/
	i := 50
	/*biến p trỏ đến ô nhớ của biến i */
	p := &i
	/*khi biến i thay đổi giá trị thì biến p giá trị sẽ thay đổi theo do p đang tham chiếu đến giá trị của i */
	fmt.Println(i)
	fmt.Println(*p)
	/*in ra biến p sẽ ra địa chỉ ô nhớ mà nó tham chiếu đến */
	fmt.Println(p)
	/*in ra địa chỉ ô nhớ của biến i*/
	fmt.Println(&i)
	/*thay đổi giá trị của p sẽ dẫn đến thay đổi giá trị của i*/
	*p = 25
	fmt.Println(i)
	fmt.Println(*p)

}
