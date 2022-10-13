package main

import (
	"fmt"
    "math/rand"
    "time"
)

func main() {
	// Declaration means that variable is only declared and memory is allocated, but no value is set.
	// However, definition means the variables has been initialized

	// termins:
	// init
	// operand
	// operator
	// condition
	// expression

	// namespace
	// var i = 100
	// {
	// 	i := 1
	// 	fmt.Println(i)
	// }
	// fmt.Println(i)

	// if 1
	// rand.Seed(time.Now().UnixNano())
	// if ; isItCorrect(rand.Int(), rand.Int()) {
	// 	fmt.Printf(`Correct!`)
	// }
	
	// // if 2
	// if a, b := rand.Int(), rand.Int(); isItCorrect(a, b) {
	// 	fmt.Printf(`Correct!`)
	// } else if a%b == 0 {
	// 	fmt.Printf(`Not correct!`)
	// } else {
	// 	fmt.Printf(`Nothing!`)
	// }
	//
	//// for loop 1
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	
	// for loop 2
	// var i = 0
	// for ; i < 10; {
	// 	fmt.Println(i)
	// 	i++
	// }
	// fmt.Printf("\nand: %d\n\n",i)
	// for i < 20 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//
	//// for loop 3
	// for i := 0; ; i++ { // <=> for i := 0; true; i++ {
	// 	if i >= 10 {
	// 		// "break" statement will be explained below.
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// exercise

	// for loop 4
	// for ; true; {
	// }
	// for true {
	// }
	// for ; ; {
	// }
	// for /*true*/ {
	// }

	// for loop 5
	// for i := 0; ; i++ {
	// 	if i >= 10 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// for loop 6
	//for i := 0; i < 3; i++ {
	//	fmt.Print(i)
	//	i := i
	//	i = 10
	//	_ = i
	//}

	// for loop 7
	// for i := 0; i < 10; i++ {
	// 	if i % 2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Print(i)
	// }

	// for loop 7
	// var str = "world"
	// for i, b := range []byte(str) {
	// 	fmt.Println(i, ":", b)
	// }

	// for loop 8
	// s := "éक्षिaπ囧"
	// for i, b := range []byte(s) {
	// 	fmt.Printf("The byte at index %v: 0x%x \n", i, b)
	// }

	// for loop 8.1
	// s := "éक्षिaπ囧"
	// for i, b := range s {
	// 	fmt.Printf("The byte at index %v: 0x%x \n", i, b)
	// }

	// exercise

	// switch 1
	// rand.Seed(time.Now().UnixNano())
	// switch n := rand.Intn(100); n % 9 {
	// case 0:
	// 	fmt.Println(n, "is a multiple of 9.")
	// case 1, 2, 3:
	// 	fmt.Println(n, "mod 9 is 1, 2 or 3.")
	// 	break // is nonsense
	// case 4, 5, 6, 7:
	// 	fmt.Println(n, "mod 9 is 4, 5 or 6.")
	// default:
	// 	fmt.Println(n, "mod 9 is 7 or 8.")
	// }

	// switch 2
	// rand.Seed(time.Now().UnixNano())
	// switch n := rand.Intn(100) % 5; n {
	// case 0, 1, 2, 3, 4:
	// 	fmt.Println("n =", n)
	// 	fallthrough
	// case 5, 6, 7, 8:
	// 	n := 99               // for current branch code block
	// 	fmt.Println("n =", n) // 99
	// 	fallthrough
	// 	//_ = n
	// default:
	// 	fmt.Println("n =", n)
	// }

	// switch 3
	// switch n := 5; n {
	// }
	
	// switch 5 {
	// }
	
	// switch _ = 5; {
	// }
	
	// switch /*true*/ {
	// }

	// switch 4
	// switch /*cond*/ {
	// case true:
	// 	fmt.Println("hello")
	// default:
	// 	fmt.Println("bye")
	// }

	// switch 5 // i love it
	// rand.Seed(time.Now().UnixNano())
	// switch n := rand.Intn(3); n {
	// default:
	// 	//fmt.Println("n == 2")
	// 	fallthrough
	// case 0:
	// 	fmt.Println("n == 0")
	// case 1:
	// 	fmt.Println("n == 1")
	// }

	// exercise

	// i := 0
	// Next: // only in namespace
	// fmt.Println(i)
	// i++
	// if i < 5 {
	// 	goto Next
	// }

	// 	rand.Seed(time.Now().UnixNano())
	// Begin:
	// 	switch n := rand.Intn(100); n % 9 {
	// 	case 0:
	// 		fmt.Println(n, "is a multiple of 9.")
	// 	case 1, 2, 3:
	// 		fmt.Println(n, "mod 9 is 1, 2 or 3.")
	// 		break Begin
	// 	case 4, 5, 6:
	// 	default:
	// 		fmt.Println(n, "mod 9 is 7 or 8.")
	// 	}

	// exercise
}

func isItCorrect(a, b int) bool {
	return a > b
}
