package main

import (
	`fmt`
)

type _string struct {
	elements *byte // underlying bytes
	len      int   // number of bytes
}

func main() {
	// 1
	// s1 := "abadalabada\n" // double-quote style, interpreted literals
	// s2 := `abadalabada\n` // back-quote style
	// s3 := `Start: `
	// s3 += s1 + ` ` + s2
	// fmt.Println(s3)


	// 2
	// const World = "world"
	// var hello = "hello"
	
	// var helloWorld = hello + " " + World
	
	// // Compare strings.
	// fmt.Println(hello == "hello")   // true
	// fmt.Println(hello > helloWorld) // false // underlying bytes

	// exercise

	// 3
	// var helloWorld = "hello world!"

	// var hello = helloWorld[:5] // substring
	//// 104 is the ASCII code (and Unicode) of char 'h'.
	// fmt.Println(hello[0])         // 104
	// fmt.Printf("%T \n", hello[0]) // uint8
	//
	//// hello[0] is unaddressable and immutable,
	//// so the following two lines fail to compile.
	///*
	//	hello[0] = 'H'         // error
	//	fmt.Println(&hello[0]) // error
	//*/
	//
	//// The next statement prints: 5 12 true
	// fmt.Println(len(hello), len(helloWorld), utf8.RuneCountInString(helloWorld),
		// strings.HasPrefix(helloWorld, hello))

	// 3.1
	// helloWorld := "Привіт, світе!" // 4
	// fmt.Println(len(helloWorld))

	// 3.2
	// sm := '🙂'
	// fmt.Printf("%T %s", sm, string(sm))

	// 3.3
	// newString := `Hi there, з новим роким!`
	
	// fmt.Println(newString)
	// fmt.Println([]byte(newString))
	// fmt.Println([]rune(newString))

	// exercise
}
