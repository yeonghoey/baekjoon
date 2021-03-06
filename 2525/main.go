package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

func nextInt() int {
	scanner.Scan()
	sign, x := 1, 0
	for _, b := range scanner.Bytes() {
		if b == '-' {
			sign = -1
			continue
		}
		x *= 10
		x += int(b - '0')
	}
	return sign * x
}

func main() {
	A := nextInt()
	B := nextInt()
	C := nextInt()
	x := (A*60 + B + C) % (24 * 60)
	h := x / 60
	m := x % 60
	fmt.Printf("%d %d\n", h, m)
}
