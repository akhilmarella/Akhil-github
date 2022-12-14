package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// type Us struct {
// 	Age   int    `json:"age"`
// 	Count int    `json:"count"`
// 	Name  string `json:"name"`
// }

// func main() {
// 	url := "https://api.agify.io?name=meelad"
// 	res, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("err in http:", err)
// 	}
// 	if res.StatusCode >= 400 {
// 		fmt.Println("err in status code:", err)
// 	}
// 	by, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println("err in body:", err)
// 	}
// 	fmt.Println("no error:", string(by))
// 	u := string(by)
// 	my := Us{Age: 32, Count: 21, Name: "meelad"}
// 	err = json.Unmarshal([]byte(u), &my)
// 	if err != nil {
// 		fmt.Println("err in unmarshal:", err)
// 	}
// 	fmt.Println(my)
// }
