package types

type node struct {
	val  any
	next *node
}

type LinkedList struct {
	head *node
	tail *node

	Len int
}

func (l *LinkedList) Add(index int, val any) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Append(val any) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Delete(index int) {
	//TODO implement me
	panic("implement me")
}

//
//// Add 值复制
//func (l LinkedList) Add(index int, val any) {
//
//}
//
//// AddV1 值引用地址，修改原对象
//func (l *LinkedList) AddV1(index int, val any) {
//
//}
