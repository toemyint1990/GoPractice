package main

import "fmt"

const p = "TOE MYINT"

func main() {
	const q = 432

	const (
		a    = iota
		b    = iota
		iota = iota
	)
	//var s = iota
	//var t = iota
	fmt.Println("p-- ", p)
	fmt.Println("q-- ", q)
	fmt.Println("a-- ", "b-- ", "iota-- ", a, b, iota)
}
