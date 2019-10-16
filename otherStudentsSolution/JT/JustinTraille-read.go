package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var file string
	fmt.Println("Enter File Name")
	_, err := fmt.Scan(&file)
	fmt.Println(err)
	if err == nil {
		data, _ := ioutil.ReadFile(file)
		people := strings.Split(string(data), "\n")
		for _, person := range people {
			token := strings.Split(person, " ")
			me := name{token[0], token[1]}
			fmt.Println("First Name " + me.fname)
			fmt.Println("Last Name " + me.lname)
		}
	} else {
		fmt.Println("Invalid file")
	}
}
