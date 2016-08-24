package main

import (
	"fmt"
	"runtime"
)

func main() {
	var a = runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(a, "개입니다")
	fmt.Println(runtime.GOMAXPROCS(0))
	s := "Hello, world"
	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}
	fmt.Scanln()
}
