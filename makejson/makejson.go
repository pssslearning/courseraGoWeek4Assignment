// --------------------------------------------------------------
// Source code mantained at Github repository for Learning
// https://github.com/pssslearning/courseraGoWeek4Assignment
// --------------------------------------------------------------
//
// Assignment instrictions:
// ------------------------
// Write a program which prompts the user to first enter a name, and then enter an address.
// Your program should create a map and add the name and address to the map using the keys
// “name” and “address”, respectively.
//
// Your program should use Marshal() to create a JSON object from the map,
// and then your program should print the JSON object.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func getStringFromPromt(fieldName string) (scannedString string, err error) {

	var aRuneInInput rune
	scannedString = ""

	fmt.Println("Please enter a string value for field name :", fieldName)

	for {
		_, err := fmt.Scanf("%c", &aRuneInInput)

		if err != nil {
			return "", err
		}

		if aRuneInInput == rune('\n') {
			break
		}

		scannedString += string(aRuneInInput)
	}

	return scannedString, nil
}

func showErrorAndExit(err error) {
	fmt.Println("An error has happenned.")
	fmt.Printf("The error detected was [%s]\n", err)
	fmt.Println("The program will be exited.")
	os.Exit(1)
}

func main() {

	aPerson := make(map[string]string)

	fmt.Println("Enter:")
	promtedString, err := getStringFromPromt("Name")
	if err != nil {
		showErrorAndExit(err)
	} else {
		aPerson["name"] = promtedString
	}

	promtedString, err = getStringFromPromt("Address")
	if err != nil {
		showErrorAndExit(err)
	} else {
		aPerson["address"] = promtedString
	}

	aByteArray, err3 := json.Marshal(aPerson)
	if err3 != nil {
		showErrorAndExit(err3)
	}

	fmt.Println("\nThe PERSON map once marshalled to JSON is:")
	fmt.Println(string(aByteArray))

	fmt.Println("\nGoodbye !!!")
	os.Exit(0)
}
