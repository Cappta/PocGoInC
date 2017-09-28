package main

import (
	"C"
	"fmt"
)

//export PrintInt
func PrintInt(x int) {
	fmt.Println(fmt.Sprintf("Golang code :D %d", x))
}

func main() {}
