package buildin_types

import "fmt"

func Slice() {
	s1 := []int{1, 2, 3}
	fmt.Printf("slice1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))

	s2 := make([]int, 3)
	fmt.Printf("slice2 %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

}

func SubSlice() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:5]
	s3 := s1[:1]
	s4 := s1[5:]

	s5 := s1[1:5:5]

	fmt.Printf("s2 %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))
	fmt.Printf("s3 %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))
	fmt.Printf("s4 %v, len: %d, cap: %d \n", s4, len(s4), cap(s4))
	fmt.Printf("s5 %v, len: %d, cap: %d \n", s5, len(s5), cap(s5))
}

// ShareSlice 切片共享底层数组，子切片没有扩容,修改切片会影响原数组
func ShareSlice() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6}
	s5 := s1[1:5:5]
	fmt.Printf("s1 %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	s2 := s1[1:5]
	// 修改同一个底层数组
	s2[0] = 100
	fmt.Printf("s1 %v \n", s1)
	fmt.Printf("s2 %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))
	s3 := s1[:1]
	// 没扩容，修改同一个底层数组
	s3 = append(s3, 200)
	fmt.Printf("s1 %v \n", s1)
	fmt.Printf("s3 %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))
	s4 := s1[5:]
	// 发生扩容，s4是新的切片
	s4 = append(s4, 300)
	fmt.Printf("s1 %v \n", s1)
	fmt.Printf("s4 %v, len: %d, cap: %d \n", s4, len(s4), cap(s4))

	// s5对应的还是同一个底层数组的切片
	fmt.Printf("s5 %v, len: %d, cap: %d \n", s5, len(s5), cap(s5))

}
