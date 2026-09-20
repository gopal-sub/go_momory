package main

import "fmt"

var global_var = 12



func main(){

	internal_val := 10

	
	fmt.Printf("Global: %v\n", &global_var)
	fmt.Printf("internal: %v", &internal_val)

}