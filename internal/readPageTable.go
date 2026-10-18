package main


import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	
	f, err := os.Open("/proc/4361/pagemap")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	offset := int64(11256)

	buf := make([]byte, 8)

	_, err = f.ReadAt(buf, offset)
	if err != nil {
		panic(err)
	}

	entry := binary.LittleEndian.Uint64(buf)

	fmt.Printf("pagemap entry: 0x%x\n", entry)
}