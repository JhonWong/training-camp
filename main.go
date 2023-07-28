package main

import (
	"fmt"

	"github.com/jhonwong/training-camp/slicedelete"
)

func main() {
	s := []int{1, 2, 3}
	ns, err := slicedelete.Delete[int](s, 1)
	if err != nil {
		fmt.Println("Erase error!")
		return
	}

	fmt.Println("After delete:", ns)
}
