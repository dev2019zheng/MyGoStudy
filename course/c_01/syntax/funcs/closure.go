package funcs

import "fmt"

// Closure 理解闭包
func Closure(name string) func() string {
	fn := func() string {
		return "hello" + name
	}
	return fn
}

func GetAge() func() int {
	var age = 18
	fmt.Printf(">> %p\n", &age)
	return func() int {
		age++
		fmt.Printf(">> %p\n", &age)
		return age
	}
}

// DeferV1 defer执行是一个栈顺序，后声明先执行
func DeferV1() {
	i := 0

	defer func() {
		fmt.Println("defer1 ", i)
	}()

	defer func(v int) {
		fmt.Println("defer2 ", v)
	}(i)

	defer func(v *int) {
		fmt.Println("defer3 ", *v)
	}(&i)

	i = 1
}

// DeferReturnV1 在Defer中修改返回值不生效
func DeferReturnV1() int {
	a := 1
	defer func() {
		a = 2
	}()
	return a
}

// DeferReturnV2 可以在Defer中修改返回值
func DeferReturnV2() (a int) {
	a = 1
	defer func() {
		a = 2
	}()

	return a
}
