package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// name - struct to store first and last name
type name struct {
	fname string // first name
	lname string // last name
}

// Given filename, reads line by line, getting first name and last name and storing in slice of structs.
// When done reading iterate thru slice and print out the names.
func main() {
	// Get filename
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter filename -> ")
	filename, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Problem reading filename. %s\n", err.Error())
		os.Exit(1)
	}
	filename = strings.Replace(filename, "\n", "", -1)
	if filename == "" {
		fmt.Printf("Filename needs to be set (e.g. names.txt).\n")
		os.Exit(1)
	}

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Problem opening file %s. %s\n", filename, err.Error())
		os.Exit(1)
	}
	defer file.Close()

	names := make([]name, 0, 0) // slice of name structs

	// Scan each line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		// save names to struct and append to slice
		pos := strings.Index(line, " ")
		var n name
		if pos < 0 {
			n.fname = line
		} else {
			n.fname = line[:pos]
			n.lname = line[(pos + 1):]
		}
		names = append(names, n)
	}
	err = scanner.Err()
	if err != nil {
		fmt.Printf("Problem scanning file. %s\n", err.Error())
		os.Exit(1)
	}

	// iterate thru slice and print out names
	fmt.Printf("%-20s %-20s\n", "First Name", "Last Name")
	fmt.Printf("%-20s %-20s\n", "----------", "---------")
	for _, n := range names {
		fmt.Printf("%-20s %-20s\n", n.fname, n.lname)
	}
}
