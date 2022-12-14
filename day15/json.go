package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // struct tags
// type Details struct {
// 	Name           string `json:"name"`
// 	Age            int `json:"age"`
// 	Qualifications []Qualification
// }

// type Qualification struct {
// 	Degree     string  `json:"degree"`
// 	Percentage float64 `json:"percentage"`
// }

// func main() {

// 	jsonData := `{
// 		"name" : "akhil",
// 		"age" : 20,
// 		"qualifications" : [
// 			{
// 				"degree": "inter",
// 				"percentage": 44.1
// 			},
// 			{
// 				"degree" : "btech",
// 				"percentage" : 99

// 			}
// 		]
// 	}`

// 	byteData := []byte(jsonData)

// 	var my Details

// 	err := json.Unmarshal(byteData, &my)
// 	if err != nil {
// 		fmt.Println("err in unmarshal:", err)
// 	}

// 	fmt.Println(my)

// 	ret, err := json.Marshal(my)
// 	if err != nil {
// 		fmt.Println("err in marshall", err)
// 	}

// 	fmt.Println(string(ret))

// }
