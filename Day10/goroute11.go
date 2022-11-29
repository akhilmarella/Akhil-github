package main

// import "fmt"

// func res(number int, square chan int) {
// 	sum := 0
// 	for number != 0 {
// 		digit := number % 10
// 		sum += digit * digit
// 		number /= 10
// 	}
// 	square <- sum
// }
// func res1(number int, cube chan int) {
// 	sum := 0
// 	for number != 0 {
// 		digit := number % 10
// 		sum += digit * digit * digit
// 		number /= 10
// 	}
// 	cube <- sum

// }
// func main() {
// 	number := 1080
// 	sq := make(chan int)
// 	cu := make(chan int)
// 	go res(number, sq)
// 	go res1(number, cu)
// 	square, cube := <-sq,<- cu
// 	fmt.Println("value",square+cube)
// }
