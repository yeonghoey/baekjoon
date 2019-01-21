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
	N := nextInt()
	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		A[i] = nextInt()
	}

	D := make([]int, N+1)
	for i := 1; i <= N; i++ {
		for j := 0; j < i; j++ {
			if A[i] > A[j] {
				D[i] = max(D[i], D[j]+A[i])
			}
		}
	}

	fmt.Println(max(D[0], D[1:]...))
}

func max(a0 int, as ...int) int {
	x := a0
	for _, a1 := range as {
		if a1 > x {
			x = a1
		}
	}
	return x
}
