package main

import (
	"encoding/json"
	"fmt"
)

type row struct {
	Source  string `json:"source"`
	Message string `json:"message"`
}

// Encode just marshalls the output to json
func Encode(m *string) {
	r := &row{
		Source:  "container",
		Message: *m,
	}
	o, err := json.Marshal(r)
	if err != nil {
		fmt.Println("{'message':'error encoding'}")
	}
	fmt.Println(string(o))
}
