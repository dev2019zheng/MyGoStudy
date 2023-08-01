package _1

import (
	"errors"
)

type Number interface {
	int | uint | int8 | int16 | int32 | int64 | float32 | float64
}

func Max[T Number](vals ...T) (T, error) {
	if len(vals) == 0 {
		var t T
		return t, errors.New("empty vals")
	}

	max := vals[0]
	for _, val := range vals[1:] {
		if val > max {
			max = val
		}
	}
	return max, nil
}

func AddSlice[T any](slice []T, idx int, val any) []T {
	if idx < 0 || idx > len(slice) {
		panic("index out of range")
	}
	res := make([]T, 0, len(slice)+1)
	for i := 0; i < idx; i++ {
		res = append(res, slice[i])
	}
	slice[idx] = val
	for i := idx; i < len(slice); i++ {
		res = append(res, slice[i])
	}
	return res
}
