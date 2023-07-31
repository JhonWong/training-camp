package slicedelete

import (
	"errors"
	"reflect"
	"testing"
)

func TestDeleteSlice(t *testing.T) {
	tests := []struct {
		input    []int
		index    int
		expected []int
		cap      int
		err      error
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}, nil},
		{[]int{1, 2, 3, 4, 5}, 0, []int{2, 3, 4, 5}, nil},
		{[]int{1, 2, 3, 4, 5}, -1, nil, errors.New("index out of range")},
		{[]int{1, 2, 3, 4, 5}, 5, nil, errors.New("index out of range")},
	}

	for _, test := range tests {
		output, err := Delete(test.input, test.index)
		if !reflect.DeepEqual(output, test.expected) || !reflect.DeepEqual(err, test.err) {
			t.Errorf("DeleteSlice(%v, %v) = %v, %v, want %v, %v", test.input, test.index, output, err, test.expected, test.err)
		}
	}
}
