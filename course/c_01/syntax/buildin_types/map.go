package buildin_types

import "fmt"

func Map() {
	m1 := map[string]int{
		"one": 1,
	}
	fmt.Printf("m1: %v, len: %d\n", m1, len(m1))
	m1["two"] = 2
	fmt.Printf("m1: %v, len: %d\n", m1, len(m1))

	m2 := make(map[string]int, 12)
	fmt.Printf("m2: %v, len: %d\n", m2, len(m2))

	// 读取
	v, ok := m1["one"]
	if ok {
		fmt.Printf("m1[\"one\"] = %d\n", v)
	}

	// 读取不存在的key，则值为对应类型的0值
	v, ok = m1["three"]
	fmt.Printf("m1[\"three\"] = %d, ok: %v\n", v, ok)

	delete(m1, "one")
	fmt.Printf("m1: %v, len: %d\n", m1, len(m1))
}
