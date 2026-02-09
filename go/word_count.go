package main

import (
	"encoding/binary"
	"io"
	"os"
)

func isSpace(b byte) bool {
	return b == 32 || b == 10 || b == 13 || b == 9
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}

	var cnt uint32
	inWord := false
	for _, ch := range input {
		if isSpace(ch) {
			inWord = false
			continue
		}
		if !inWord {
			cnt++
			inWord = true
		}
	}

	var out [4]byte
	binary.LittleEndian.PutUint32(out[:], cnt)
	if _, err := os.Stdout.Write(out[:]); err != nil {
		os.Exit(1)
	}
}
