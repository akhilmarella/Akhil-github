package main

import "fmt"

func main() {
	moist := make(chan string)
	go res(moist)
	name := <-moist
	fmt.Println("name is:", name)
	most := make(chan int)
	go res1(most)
	age := <-most
	fmt.Println(age)
	read := make(chan float64)
	go res2(read)
	weight := <-read
	fmt.Println(weight)

}
func res(moist chan string) {
	name := "akhil"
	moist <- name
	fmt.Println("res")
}
func res1(most chan int) {
	age := 25
	most <- age
	fmt.Println("age is")
}
func res2(read chan float64) {
	weight := 80.5
	read <- weight
	fmt.Printf("res2:")

}
