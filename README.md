# courseraGoWeek4Assignment
Assigment Week #4 Course https://www.coursera.org/learn/golang-getting-started/peer/0xyF8/module-4-activity-makejson-go

# pssslearning courseraGoWeek4Assignment
Assigment Week 3 - Course  hhttps://www.coursera.org/learn/golang-getting-started/peer/0xyF8/module-4-activity-makejson-go

## Coursera course: Programming with Google Go (University of California, Irvine) 

[Coursera - Getting Started with Go](https://www.coursera.org/learn/golang-getting-started/home/welcome)

## Week 4 - Module 4 Activity 1: makejson.go

### Goals

The goal of this program is to practice working with RFCs in Go, using JSON as an example.

### Assignment directions

Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

### My assignment proposal is placed at
```sh
makejson/makejson.go
```

### A sample log for its compiling a testing cycle:

```sh
devel1@vbxdeb10mate:~$ cd $GOPATH/src/github.com/pssslearning/courseraGoWeek4Assignment/makejson/
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoWeek4Assignment/makejson$ go build makejson.go 
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoWeek4Assignment/makejson$ ./makejson 
Enter:
Please enter a string value for field name : Name
Pablo M.F. Sancho
Please enter a string value for field name : Address
Fake Street,35 1st.floor 12345-VALENCIA (SPAIN)  

The PERSON map once marshalled to JSON is:
{"address":"Fake Street,35 1st.floor 12345-VALENCIA (SPAIN)","name":"Pablo M.F. Sancho"}

Goodbye !!!
```

## Week 4 - Module 4 Activity 2 (Final Course Activity): read.go

### Goals

The goal of this assignment is to practice working with the ioutil and os packages in Go.

### Assignment directions

Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.