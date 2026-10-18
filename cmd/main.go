package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	pid := os.Getpid()
	fmt.Println(pid)
	arr := make([]byte, 5<<30) // 1 GiB
	time.Sleep(10*time.Second)

	for i := 0; i < len(arr); i += 4096 {
		arr[i] = 10
		fmt.Printf("Touched byte %d\n", i)
	}

	time.Sleep(10*time.Minute)

}