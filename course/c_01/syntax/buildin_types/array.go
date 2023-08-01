package buildin_types

import "fmt"

func Array() {
	arr := [3]int{1, 2, 3}
	fmt.Printf("arr: %v, len: %d, cap: %d\n", arr, len(arr), cap(arr))

	arr2 := [3]int{1, 2}
	//arr2[2]  //数组越界直接报错
	fmt.Printf("arr2: %v, len: %d, cap: %d\n", arr2, len(arr2), cap(arr2))

	arr3 := [3]int{}
	//append(arr3, 1) //数组不支持append
	fmt.Printf("arr3: %v, len: %d, cap: %d\n", arr3, len(arr3), cap(arr3))
}
