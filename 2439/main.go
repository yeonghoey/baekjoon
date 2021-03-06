package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	writer := bufio.NewWriter(os.Stdout)
	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Fprint(writer, " ")
		}
		for j := 0; j < i; j++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprint(writer, "\n")
	}
	writer.Flush()
}
