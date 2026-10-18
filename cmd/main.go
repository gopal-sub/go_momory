package main

import (
	"fmt"
	// "os"
	"time"

	// "time"
	// "golang.org/x/sys/unix"
)
var global_var = 12



func main(){
	arra_hold := make([]byte, 10737418240)

	for i := 0; i < len(arra_hold); i += 4096{
		held :=arra_hold[i]
		fmt.Println(held)


		if i% 2560 == 0{
			fmt.Printf(`touched byte %d`, i)
			time.Sleep(1* time.Second)
		}
	}
}