package main

/*Khai báo interface*/
type Speaker interface {
	Speak() string
}

/*Bất kỳ kiểu dữ liệu nào miễn là có method Speak() string (không tham số và trả về string) thì đều được gọi là Speaker*/
/*Interface trong Golang được implement một cách ngầm định (implicitly)
type Foo struct{}

func (Foo) Speak() string {
	return "Hello, I am Foo"
}

func main() {
	var someSpeaker Speaker = Foo{}
	fmt.Println(someSpeaker.Speak())
}*/
/*==============================================================*/
/*cách tạo kiểu any trong go : sử dụng khai báo một interface rỗng
var i interface{}
describe(i)

i = 42
describe(i)

i = "hello"
describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)*/
/*==============================================================*/
/*ngoài ra còn có thể dùng làm value cho kiểu dữ liệu map của Golang
product := make(map[string]interface{}, 0)

product["name"] = "Iphone 13 Pro Max"
product["price"] = 31000000
product["quantity"] = 40

fmt.Println(product)*/
/*==============================================================*/
/*Ép kiểu dữ liệu Interface
var i interface{} = "hello"

s := i.(string)
fmt.Println(s)

s, ok := i.(string)
fmt.Println(s, ok)

f, ok := i.(float64)
fmt.Println(f, ok)

f = i.(float64) // panic
fmt.Println(f)*/
/*==============================================================*/
/*Nhầm lẫn giữa tham trị và tham chiếu (value type và pointer type)
Trong ví dụ sau, struct Foo implement Speaker với một receiver là value type.
Nhưng struct Bar lại implement Speaker với một receiver là pointer type.
Kết quả là biến someSpeaker (có kiểu Speaker Interface) nhận được cho cả value và pointer cho Foo nhưng lại chỉ dùng với pointer Bar mà thôi.
type Speaker interface {
	Speak() string
}

type Foo struct{}

func (Foo) Speak() string {
	return "Hello, I am Foo"
}

type Bar struct{}

func (*Bar) Speak() string {
	return "Hello, I am Bar"
}

func main() {
	var someSpeaker Speaker = Foo{}

	fmt.Println(someSpeaker.Speak())

	someSpeaker = &Foo{} // will be OK

	fmt.Println(someSpeaker.Speak())

	someSpeaker = Bar{} // PANIC HERE

	fmt.Println(someSpeaker.Speak())
}*/
/*==============================================================*/
/*Interface nên được khai báo ở tại nơi sử dụng nó để tránh lỗi cycle import (các pakage import lẫn nhau)
package b

type Speaker interface {
	Speak() string
}

func someFunc(spk Speaker) {
	spk.Speak()
}*/
