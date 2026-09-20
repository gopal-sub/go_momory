package main

import (
	"fmt"
	"time"
)
var global_var = 12



func main(){

	internal_val := 10


	fmt.Printf("Global: %v\n", &global_var)
	fmt.Printf("internal: %v", &internal_val)

	for {
		time.Sleep(3*time.Second)
		fmt.Println("----")
	}
}