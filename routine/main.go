package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("concurrency with GoRoutines")
	go compute(5)
	go compute(5)
	//time.Sleep(100 * time.Millisecond)
	fmt.Scanln() //持续扫描，等到新行或者EOF
}

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}

}
