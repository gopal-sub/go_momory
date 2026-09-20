package main

import (
	"fmt"
	"time"
	"os"
)
var global_var = 12



func main(){

	internal_val := 10
	pid := os.Getpid()

	fmt.Printf("Global: %v\n", &global_var)
	fmt.Printf("internal: %v\n", &internal_val)
	fmt.Printf("pid: %v\n", pid)

	for {
		time.Sleep(10*time.Second)
		fmt.Println("----")
	}
}