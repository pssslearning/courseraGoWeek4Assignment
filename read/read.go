// Assignment Instructions
// Write a program which reads information from a file and represents
// it in a slice of structs. Assume that there is a text file which
// contains a series of names. Each line of the text file has a first name
// and a last name, in that order, separated by a single space on the line.
//
// Your program will define a name struct which has two fields,
// fname for the first name, and lname for the last name.
// Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file.
// Your program will successively read each line of the text file and
// create a struct which contains the first and last names found in the file.
// Each struct created will be added to a slice, and after all lines have been
// read from the file, your program will have a slice containing one struct
// for each line in the file.
// After reading all lines from the file, your program should iterate through
// your slice of structs and print the first and last names found in each struct.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Required type
type person struct {
	fname string
	lname string
}

func getInputFileNameFromPromt() (inputFilePath string, err error) {

	inputFilePath = ""

	for {
		fmt.Println("\nPlease enter the full path for the input file to read: ")
		fmt.Println("or enter 'END' or press <CTRL + C> to exit program")

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputFilePath = scanner.Text()
			break
		}

		if err := scanner.Err(); err != nil {
			return "", err
		}

		inputFilePath = strings.TrimSpace(inputFilePath)
		if inputFilePath == "" {
			fmt.Println("Nothing entered. Please try again")
			continue
		} else {
			return inputFilePath, nil
		}
	}

}

// Exists reports whether the named file or directory exists and
// have enough permissions to access it
// (Snippet taken from https://github.com/noll/mjau/blob/master/util/util.go#L42)
func fileExists(name string) bool {
	fileInfo, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return !fileInfo.IsDir()
}

func processFile(pathToFile string, personsSlice *[]person) error {
	file, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		aPerson := getPerson(scanner.Text())
		*personsSlice = append(*personsSlice, aPerson)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// DISCLAIMER:
// This will only work properly under the assumption of the Assignment's instructions
// ... Assume that there is a text file which contains a series of names. 
// Each line of the text file has a first name and a last name, in that order, 
// separated by a single space on the line. ...
// It's not well fitted to deal with other line conformation in the file
func getPerson(fileLine string) person {

	fileLine = strings.TrimSpace(fileLine)
	lineComponents := strings.Split(fileLine, " ")

	fname := ""
	lname := ""
	if len(lineComponents) >= 2 {
		fname = lineComponents[0]
		lname = lineComponents[1]
	} else {
		fname = lineComponents[0]
		lname = "(N/A)"
	}

	if len(fname) > 20 {
		fname = substr(fname, 0, 20)
	} else if len(fname) == 0 {
		fname = "(N/A)"
	}

	if len(lname) > 20 {
		lname = substr(lname, 0, 20)
	} else if len(lname) == 0 {
		lname = "(N/A)"
	}

	return person{fname, lname}
}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func main() {

	// The "empty person slice of max. initial capacity 100 "
	personsSlice := make([]person, 0, 100)

	for {

		inputFilePath, err := getInputFileNameFromPromt()

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error condition raised. The program will exit.")
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if inputFilePath == "END" {
			fmt.Println("\nGoodbye !!!")
			os.Exit(0)
		}

		if fileExists(inputFilePath) {
			fmt.Printf("\nProcessing file[%s] ...\n", inputFilePath)
			err := processFile(inputFilePath, &personsSlice)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error condition raised. The program will exit.")
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			} else {
				fmt.Println("\nIterating through the resulting slice of persons:")
				for n, p := range personsSlice {
					fmt.Printf("Entry #[%d] First Name[%s] Last Name[%s]\n", n, p.fname, p.lname)
				}
				os.Exit(0)
			}
		} else {
			fmt.Printf("File [%s] Doesn't exists.\n", inputFilePath)
			continue
		}
	}

}
