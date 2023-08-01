package types

type List interface {
	Add(index int, val any)
	Append(val any)
	Delete(index int)
}
