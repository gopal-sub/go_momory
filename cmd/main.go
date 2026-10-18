package main

import (
	"fmt"
	"os"
	"time"

	// "time"
	"golang.org/x/sys/unix"
)
var global_var = 12



func main(){

	AddFileBuffer()
	time.Sleep(10*time.Minute)

	// internal_val := 10
	// pid := os.Getpid()

	// fmt.Printf("Global: %v\n", &global_var)
	// fmt.Printf("internal: %v\n", &internal_val)
	// fmt.Printf("pid: %v\n", pid)

	// for {
	// 	time.Sleep(10*time.Second)
	// 	fmt.Println("----")
	// }
}

func AddFileBuffer() {
	file, _ := os.OpenFile("./superFile.txt", os.O_RDONLY, 0666)

	// (fd int, offset int64, length int, prot int, flags int) (data []byte, err error)
	data, err := unix.Mmap(
		int(file.Fd()),
		0,
		1024,
		unix.PROT_READ,
		unix.MAP_SHARED,
		
	)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(data))

}