package main

import (
	"io"
	"os"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}
	if len(input) == 0 {
		return
	}

	out := make([]byte, 0, len(input)*2)
	cur := input[0]
	var cnt byte = 1

	for i := 1; i < len(input); i++ {
		x := input[i]
		if x == cur {
			if cnt < 255 {
				cnt++
				continue
			}
			out = append(out, cnt, cur)
			cnt = 1
			continue
		}
		out = append(out, cnt, cur)
		cur = x
		cnt = 1
	}

	out = append(out, cnt, cur)

	if _, err := os.Stdout.Write(out); err != nil {
		os.Exit(1)
	}
}
