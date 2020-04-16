package GettingStarted

import (
	"encoding/json"
	"fmt"
)

func mainas() {
	userProfile := make(map[string]string)
	var name string
	var address string

	fmt.Print("Please Enter Your Name: ")
	fmt.Scan(&name)
	fmt.Println()
	fmt.Print("Please Enter Your Address: ")
	fmt.Scan(&address)

	userProfile["name"] = name
	userProfile["address"] = address

	fmt.Println(userProfile)
	jsonProfile, err := json.Marshal(userProfile)
	if err != nil {
		fmt.Println("There is something wrong with the system")
	} else {
		fmt.Println(string(jsonProfile))
	}
}
