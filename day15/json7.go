package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name        string   `json:"name"`
	Designation string   `json:"designation"`
	Details     []Detail `json:"details"`
}
type Detail struct {
	Previouscompany string `json:"previouscompany"`
}

func main() {
	emp := `{
		"Name":"akhil",
    	"Age":25,
		"Designation":"engineer",
		"Details":[
			{
				"Prevoiuscompany":"wipro",
				"salary":25000
			},
			{
				"Previouscompany":"infosys",
				"salary":15000
			}
		]
	}`
	e := []byte(emp)
	var my Employee
	err := json.Unmarshal(e, &my)
	if err != nil {
		fmt.Println("err found in unmarshal:", err)
	}
	fmt.Println(my)
	byt, err := json.Marshal(my)
	if err != nil {
		fmt.Println("err found in marshal", err)
	}
	fmt.Println(string(byt))
}
