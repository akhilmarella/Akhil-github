package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type User struct {
// 	Email   string `json:"email"`
// 	Pass    string `json:"pass"`
// 	Details []Detail
// }
// type Detail struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// func main() {
// 	jData := `{
// 		"Email": "qwerty@12.com",
// 		"Pass": "asdf@1233",
// 		"Details":[
// 			{
// 			"Name": "akhil",
// 			"Age" : 25
// 			},
// 			{
// 				"Name": "babu",
// 				"Age": 20
// 			}
// 		]		
//     }`
// 	bytedata := []byte(jData)

// 	var my User
// 	err := json.Unmarshal(bytedata, &my)
// 	if err != nil {
// 		fmt.Println("error found in unmarshal:", err)
// 	}
// 	fmt.Println("un marsha:",my)
// 	res, err := json.Marshal(my)
// 	if err != nil {
// 		fmt.Println("error found in marshal:", err)
// 	}
// 	fmt.Println("marshal:",string(res))
// }



