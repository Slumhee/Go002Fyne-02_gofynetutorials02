package main

import (
 "fmt"
 "unsafe"
)

func main() {

 arr := []int{10, 20, 30, 40, 50}

 ptr := unsafe.Pointer(&arr[0])

 fmt.Println(ptr)
 fmt.Printf("%d\n", ptr)
 fmt.Println(uintptr(ptr))

 ptr = unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(arr[0]))
 fmt.Printf("%d\n", ptr)


 *(*int)(ptr) = 25

 fmt.Println(arr) 
}