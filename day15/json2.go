package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func main() {
// 	url := "https://api.publicapis.org/entries"

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("err is found:", err)
// 	}
// 	if resp.StatusCode >= 400 {
// 		fmt.Println("err is in statuscode:", err)
// 	}
// 	by, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("err is in read all:", err)
// 	}
// 	fmt.Println(string(by))
// }
