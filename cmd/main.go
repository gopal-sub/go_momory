package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	pid := os.Getpid()
	fmt.Println(pid)
	arr := make([]byte, 1<<30) // 1 GiB

	for i := 0; i < len(arr); i += 4096 {
		arr[i] = 10
		fmt.Printf("Touched byte %d\n", i)
		time.Sleep(time.Second)
	}
}