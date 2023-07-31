package main

import (
	"fmt"

	"github.com/jhonwong/training-camp/slicedelete"
)

func main() {
	s := make([]int, 2, 4)
	ns, err := slicedelete.Delete[int](s, 1)
	if err != nil {
		fmt.Println("Erase error!")
		return
	}

	fmt.Println("After delete:", ns, "\tCap:", cap(ns))
}
