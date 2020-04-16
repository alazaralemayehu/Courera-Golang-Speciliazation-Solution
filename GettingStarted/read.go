package GettingStarted

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type names struct {
	fname string
	lname string
}

func main() {
	var fileName string

	fmt.Print("Please Enter Your Name: ")
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Something is wrong, Please check your file name")
	}
	var n names
	var namelist = make([]names, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		var str []string = strings.Split(scanner.Text(), " ")
		n.fname = str[0]
		n.lname = str[1]
		namelist = append(namelist, n)
	}
	for _, name := range namelist {
		// fmt.Println(name)
		fmt.Println("First Name :" + name.fname + " Last Name: " + name.lname)
	}
	// fmt.Println(namelist)
}
