package main

import (
	"C"
	"fmt"
)

//export GoPrintInt
func GoPrintInt(x int) {
	fmt.Println(fmt.Sprintf("Golang code :D %d", x))
}

func main() {}
