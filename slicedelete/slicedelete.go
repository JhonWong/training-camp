package slicedelete

import "errors"

func Delete[T any](s []T, index int) ([]T, error) {
	if index < 0 || index >= len(s) {
		err := errors.New("Invalid index")
		return s, err
	}

	newCap := cap(s)
	if len(s)-1 < newCap/2 {
		newCap /= 2
	}

	res := make([]T, len(s)-1, newCap)
	res = append(res, s[:index]...)
	res = append(res, s[index+1:]...)
	return res, nil
}
