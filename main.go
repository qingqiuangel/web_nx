package main

import (
	"fmt"
	"unsafe"
)

func main() {
	n2 := "g"
	fmt.Printf("n2 的类型 %T n2占中的字节数是 %d", n2, unsafe.Sizeof(n2))
}
