package main

import (
	"C"
	"fmt"
)
import (
	"unsafe"
)

func StringSliceToPtr(slice []string) unsafe.Pointer {
	var ptr unsafe.Pointer
	ptrSize := int(unsafe.Sizeof(ptr))
	arraySize := len(slice) + 1
	ptr = C.malloc(C.size_t(ptrSize * arraySize))
	cSlice := (*[1 << 28]*C.char)(ptr)

	for i := 0; i < len(slice); i++ {
		cSlice[i] = C.CString(slice[i])
	}
	cSlice[len(slice)] = nil

	return ptr
}

func PtrToStringSlice(ptr unsafe.Pointer) (slice []string) {
	cSlice := (*[1 << 28]*C.char)(ptr)

	for i := 0; cSlice[i] != nil; i++ {
		slice = append(slice, C.GoString(cSlice[i]))
	}
	return
}

//export GoTestStringSlice
func GoTestStringSlice() unsafe.Pointer {
	result := []string{"1", "2", "4", "8", "16"}
	ptr := StringSliceToPtr(result)

	fmt.Println("Golang ", result, ": ", ptr)
	return ptr
}

func main() {
	input := []string{"1", "2", "4", "8", "16"}
	fmt.Println("Input ", input)

	ptr := StringSliceToPtr(input)

	fmt.Println(ptr)

	output := PtrToStringSlice(ptr)

	fmt.Println("Output ", output)
}
