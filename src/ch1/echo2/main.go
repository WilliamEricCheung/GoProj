package main

import (
	"fmt"
	"os"
)

/*
以下等价，常用前两个
s := ""
var s string
var s = ""
var s string = ""
 */
func main() {
	s, sep := "", ""
	var tmp = " "
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	tmp = "dfsdf"
	fmt.Println(s)
	fmt.Println(tmp)
}
