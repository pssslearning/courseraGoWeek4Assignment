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
