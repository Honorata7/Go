//Program to calculate roots of a square trinomial from command line arguments

package main

import (
	"flag"
	"fmt"
)

type variables struct {
	a int
	b int
	c int
}

func main() {
	//do it completely different
	vars := variables{}
	flag.IntVar(&vars.a, "a", 0, "a")
	flag.IntVar(&vars.b, "b", 0, "b")
	flag.IntVar(&vars.c, "c", 0, "c")
	flag.Parse()
	fmt.Println("a:", vars.a, "b:", vars.b, "c:", vars.c)
}
