package slicedelete

import "errors"

func Delete[T any](s []T, index int) ([]T, error) {
	if index < 0 || index >= len(s) {
		err := errors.New("index out of range")
		return nil, err
	}

	newCap := cap(s)
	if len(s)-1 < newCap/2 {
		newCap /= 2
	}

	res := make([]T, 0, newCap)
	res = append(res, s[:index]...)
	res = append(res, s[index+1:]...)
	return res, nil
}
