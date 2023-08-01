package _1

import "errors"

func RemoveSlice[T Number](slice []T, idx int) ([]T, error) {
	if idx < 0 || idx > len(slice) {
		return nil, errors.New("index out of range")
	}
	res := make([]T, 0, len(slice)-1)
	for i := 0; i < idx; i++ {
		res = append(res, slice[i])
	}
	for i := idx + 1; i < len(slice); i++ {
		res = append(res, slice[i])
	}
	return res, nil
}
