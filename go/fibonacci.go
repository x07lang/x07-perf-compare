package main

import (
	"encoding/binary"
	"io"
	"os"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}
	if len(input) < 4 {
		os.Exit(1)
	}

	n := binary.LittleEndian.Uint32(input[:4])
	var result uint32
	if n < 2 {
		result = n
	} else {
		var a uint32 = 0
		var b uint32 = 1
		for i := uint32(2); i <= n; i++ {
			tmp := a + b
			a = b
			b = tmp
		}
		result = b
	}

	var out [4]byte
	binary.LittleEndian.PutUint32(out[:], result)
	if _, err := os.Stdout.Write(out[:]); err != nil {
		os.Exit(1)
	}
}
