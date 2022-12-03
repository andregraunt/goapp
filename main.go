// package main
//
// import "fmt"
//
// func main() {
// 	name := "vasa"
// 	age := 37
// 	maskoret := 5000.55
// 	fmt.Println("hello go")
// 	fmt.Printf("My name is %s, My age %d and I gonna %f dollars ", name, age, maskoret)
// }
//
// go build -ldflags "-s -w" main.go

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	const col = 30
	// Clear the screen by printing \x0c.
	bar := fmt.Sprintf("\x0c[%%-%vs]", col)
	for i := 0; i < col; i++ {
		fmt.Printf(bar, strings.Repeat("=", i)+">")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Printf(bar+" Done!", strings.Repeat("=", col))
}
