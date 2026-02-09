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

	var acc uint32
	for _, b := range input {
		acc += uint32(b)
	}

	var out [4]byte
	binary.LittleEndian.PutUint32(out[:], acc)
	if _, err := os.Stdout.Write(out[:]); err != nil {
		os.Exit(1)
	}
}
