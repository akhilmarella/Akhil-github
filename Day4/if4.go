package main

import "fmt"

func main() {
	value := 12
	res := check(value)
	fmt.Println(res)
	res1 := result(value)
	fmt.Println(res1)
}
func check(n int) bool {
	if n%2 == 1 {
		return true
	}
	fmt.Println("odd")
	return false
}

func result(age int) bool {
	if age > 10 {
		return true
	}
	fmt.Println("even")
	return false

}
