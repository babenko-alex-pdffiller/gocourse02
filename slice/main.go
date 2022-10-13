package main

import (
	"fmt"
	"time"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {
	// 0
	// s := [10]int{} // not a pointer
	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	// 0.1
	// s := [10]int{} // pointer on
	// s[1] = 100
	// s[3] = 200
	// fmt.Println(s)
	// fmt.Println(s[0], s[1], s[2], s[3])

	// 0.2
	// s := [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	// c := [...]int{1, 2, 3}
	// fmt.Println(c)

	//1
	// s := make(
	// 	[]int,
	// 	10, // len
	// 	10, // cap
	// )
	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	// 2
	// a := []int{1}
	// b := a[0:1]
	// fmt.Printf("%p\n", &a)
	// fmt.Printf("%p\n", &b)

	// exercise

	// 2.1
	// a := []int{1}
	// b := a[0:1]
	// b[0] = 0
	
	// fmt.Println(a)
	
	// a[0] = 1
	
	// fmt.Println(b)

	// 3
	// a := make([]int, 1, 2)
	// a[0] = 1
	// b := a[0:1]
	
	// a = append(a, 2)
	// a[0] = 10

	// fmt.Println(a, b) // [10 2] [1]

	// 4
	// a := []int{1, 2, 3, 4, 5, 6, 7}
	// b := a[3:5]
	
	// fmt.Println(a, cap(a))
	// fmt.Println(b, cap(b))
	
	// b[0] = 1_000_000
	// fmt.Println(a)

	// exercise

	// 5
	// a := []int{1, 2, 3}
	// b := a[0:3]
	// update_print(a)
	// fmt.Println(`Out func: `, a, b)
	// b[2] = 300
	// time.Sleep(time.Millisecond * 30)

// In func:  [1 200 3]
// Out func:  [1 200 3] [1 200 3]
// In func after 10 milliseconds:  [1 200 300]

	// 6
	//a := []int{1, 2, 3}
	//b := make([]int, 3, 3)
	//copy(b, a)
	//
	//a[2] = 2
	//b[0] = 0
	//fmt.Println(a, b)

	// 7
	//a := []int{1, 2, 3, 4, 5, 6, 7}
	//b := a[0:1:5]
	//fmt.Println(len(b), cap(b))

	// exercise
}

func update_print(s []int) {
	s[1] = 200
	fmt.Println(`In func: `, s)
	go func() {
		time.Sleep(time.Millisecond * 10)
		fmt.Println(`In func after 10 milliseconds: `, s)
	}()
}
