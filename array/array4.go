package main

import "fmt"

type Rectangle struct {
	Length  int
	Breadth int
	Height  []Square
}
type Square struct {
	Length1  int
	Breadth1 int
}

func main() {
	res, squ := result()
	fmt.Println(res, squ)
	list := []string{"hello"}
	list = append(list, "world")
	list = append(list, "how r u")
	fmt.Println(list)
}

func result() (Rectangle, []Square) {
	r1 := Rectangle{Length: 250, Breadth: 300, Height: []Square{}}
	s1 := Square{Length1: 300, Breadth1: 300}
	s := []Square{s1}
	return r1, s
}
