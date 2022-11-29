package main

// import "fmt"

// func result(read chan<- int) {
// 	read <- 10
// }
// func res(wrote chan<- string) {
// 	wrote <- "akhil babu"
// }

// func main() {
// 	learn := make(chan int)
// 	go result(learn)
// 	fmt.Println(<-learn)
// 	write := make(chan string)
// 	go res(write)
// 	fmt.Println(<-write)
// 	act := make(chan int)
// 	go res2(act)
// 	for v:=range act{
// 		fmt.Println("received",v)
// 	}
// }
// func res2(activity chan int) {
// 	for i := 0; i <= 10; i++ {
// 		activity <- i
// 	}
// 	close(activity)
// }
