package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

// map的迭代顺序并不确定，从实践来看，该顺序随机，每次运行都会变化
func main() {
	counts := make(map[string] int) //内置函数make创建空map，此外，它还有别的作用
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
