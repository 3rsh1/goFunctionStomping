package main

import (
	"fmt"
	"github.com/JamesHovious/w32"
	"time"
)

func main() {

	w32.GetLastError()
	time.Sleep(20 * time.Second)
	fmt.Println("test1")
	w32.GetLastError()
	fmt.Println("test2")

}
