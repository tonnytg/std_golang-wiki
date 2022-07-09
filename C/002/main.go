package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func printRawData(ptr uintptr, size uintptr) {
	fmt.Printf("Printing ptr %016x size %d\n", ptr, size)
	i := ptr
	var offset uintptr
	for i < ptr+size {
		if offset%16 == 0 {
			fmt.Printf("%016x : ", i)
		}

		fmt.Printf("%02x", *(*byte)(unsafe.Pointer(i)))

		i++
		offset++
		if offset%16 == 0 || offset == size {
			fmt.Print("\n")
		} else if offset%8 == 0 {
			fmt.Print("  ")
		} else {
			fmt.Print(" ")
		}
	}
}

func main() {
	fmt.Println("Simple array:")
	myArr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printRawData(uintptr(unsafe.Pointer(&myArr)), unsafe.Sizeof(myArr))

	fmt.Println("\nArray assignment:")
	myArr2 := myArr
	printRawData(uintptr(unsafe.Pointer(&myArr2)), unsafe.Sizeof(myArr2))

	fmt.Println("\nArray pointer:")
	myArrPtr := &myArr
	printRawData(uintptr(unsafe.Pointer(myArrPtr)), unsafe.Sizeof(*myArrPtr))

	fmt.Println("\nSimple slice:")
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printRawData(uintptr(unsafe.Pointer(&mySlice)), unsafe.Sizeof(mySlice))

	fmt.Println("\nSlice headers:")
	mySliceHdr := *(*reflect.SliceHeader)(unsafe.Pointer(&mySlice))
	fmt.Printf("%#v\n", mySliceHdr)
	printRawData(mySliceHdr.Data, uintptr(mySliceHdr.Len*8))

	fmt.Println("\nSlice assignment:")
	mySlice2 := mySlice
	printRawData(uintptr(unsafe.Pointer(&mySlice2)), unsafe.Sizeof(mySlice2))
	mySliceHdr2 := *(*reflect.SliceHeader)(unsafe.Pointer(&mySlice2))
	printRawData(mySliceHdr2.Data, uintptr(mySliceHdr2.Len*8))

	fmt.Println("\nAppend to slice:")
	mySlice2 = append(mySlice2, 11)
	printRawData(uintptr(unsafe.Pointer(&mySlice2)), unsafe.Sizeof(mySlice2))
	mySliceHdr2 = *(*reflect.SliceHeader)(unsafe.Pointer(&mySlice2))
	fmt.Printf("%#v\n", mySliceHdr2)
	printRawData(mySliceHdr2.Data, uintptr(mySliceHdr2.Len*8))
}
