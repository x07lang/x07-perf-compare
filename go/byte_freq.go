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

	var freq [256]uint32
	for _, b := range input {
		freq[b]++
	}

	out := make([]byte, 0, 256*5)
	var tmp [4]byte
	for i, n := range freq {
		if n == 0 {
			continue
		}
		out = append(out, byte(i))
		binary.LittleEndian.PutUint32(tmp[:], n)
		out = append(out, tmp[:]...)
	}

	if _, err := os.Stdout.Write(out); err != nil {
		os.Exit(1)
	}
}
