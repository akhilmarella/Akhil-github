package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// func main() {
// 	url := "https://api.ipify.org?format=json"
// 	res, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("error in url:", err)
// 	}
// 	if res.StatusCode >= 400 {
// 		fmt.Println("error in status code:", err)
// 	}
// 	by, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println("error in body:", err)
// 	}
// 	fmt.Println(string(by))
// }
