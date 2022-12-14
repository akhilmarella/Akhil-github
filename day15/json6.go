package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Dimension struct {
// 	//	Height int `json:"height"`
// 	Width int `json:"width"`
// }
// type Bird struct {
// 	Species string `json:"species"`
// 	//	Description string `json:"description`
// 	Dimensions []Dimension `json:"dimension`
// }

// func main() {
// 	Bjson := `{
//               "Species":"pigeon",
// 			  "Description":"likes to perch on rocks",
// 			  "Dimensions":[
// 				{
//                  "Height":24,
// 				 "Width":10
// 				},
// 				{
// 					"Height":20,
// 					"Width":6
// 				}
// 			  ]	
// 	        }`
// 	bytdata := []byte(Bjson)
// 	var a Bird
// 	err := json.Unmarshal(bytdata, &a)
// 	if err != nil {
// 		fmt.Println("err is there in unmarshal:", err)
// 	}
// 	fmt.Println(a)
// 	byt, err := json.Marshal(a)
// 	if err != nil {
// 		fmt.Println("err in Marshal:", err)
// 	}
// 	fmt.Println(string(byt))
// }
