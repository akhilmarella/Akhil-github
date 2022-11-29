package main

import "fmt"

func main() {
	may := make(chan int)
	go result(may)
	for a := range may {
		fmt.Println("receive:", a)
	}
	write := make(chan string)
	go res(write)
	fmt.Println(<-write)
	read := make(chan int)
	go res2(read)
	for {
		v, valid := <-read
		if valid == false {
			break
		}
		fmt.Println("valid:",v)
	}

}

func result(may chan int) {
	for j := 0; j < 4; j++ {
		may <- j
	}
	close(may)
}
func res(wrote chan<- string) {
	wrote <- "akhil"
}
func res2(read chan int) {
	for i := 0; i < 10; i += 2 {
		read <- i
	}
	close(read)
}
