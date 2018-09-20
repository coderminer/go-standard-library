package main

import (
	"fmt"
)

func main() {
	n := 0
	if true {
		n := 1
		n += 1
	}
	fmt.Println(n)

	var dst, src []int
	src = []int{1, 2, 3}
	count := copy(dst, src)
	fmt.Println("src:", src)
	fmt.Println("dst:", dst)
	fmt.Println("copied count:", count)
}
