package funcs

import "fmt"

func Functional1() {
	fmt.Println("functional1")
}

func UseFunctional1() {
	// 函数式编程，函数也是变量
	myFunc := Functional1
	myFunc()
}

func Functional2() {
	// 匿名函数
	fn := func() string {
		return "hello"
	}
	fn()
}
