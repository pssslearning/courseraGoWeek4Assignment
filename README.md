# courseraGoWeek4Assignments
Assigment Week #4 
- Assignment #1 makeson.go: https://www.coursera.org/learn/golang-getting-started/peer/0xyF8/module-4-activity-makejson-go
- Assignment #2 read.go: https://www.coursera.org/learn/golang-getting-started/peer/CHgQd/final-course-activity-read-go

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

- on my Debian 10 linux system

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

### My assignment proposal is placed at
```sh
read/read.go
```

### A sample log for its compiling a testing cycle:

```sh
devel1@vbxdeb10mate:~$ cd $GOPATH/src/github.com/pssslearning/courseraGoWeek4Assignment/read/
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoWeek4Assignment/read$ go build read.go
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoWeek4Assignment/read$ ./read

Please enter the full path for the input file to read: 
or enter 'END' or press <CTRL + C> to exit program
/not-existing-dir/not-existing-file.txt
File [/not-existing-dir/not-existing-file.txt] Doesn't exists.

Please enter the full path for the input file to read: 
or enter 'END' or press <CTRL + C> to exit program
../data/SomeNobelPrizes.txt

Processing file[../data/SomeNobelPrizes.txt] ...

Iterating through the resulting slice of persons:
Entry #[0] First Name[cdMichel] Last Name[Mayor]
Entry #[1] First Name[Didier] Last Name[Queloz]
Entry #[2] First Name[Akira] Last Name[Yoshino]
Entry #[3] First Name[John] Last Name[Goodenough]
Entry #[4] First Name[James] Last Name[Peebles]
Entry #[5] First Name[William] Last Name[Kaelin]
Entry #[6] First Name[Peter] Last Name[Ratcliffe]
Entry #[7] First Name[Gregg] Last Name[Semenza]
Entry #[8] First Name[Peter] Last Name[Handke]
Entry #[9] First Name[Abhijit] Last Name[Banerjee]
Entry #[10] First Name[Esther] Last Name[Duflo]
Entry #[11] First Name[Michael] Last Name[Kremer]
Entry #[12] First Name[Arthur] Last Name[Ashkin]
Entry #[13] First Name[Gérard] Last Name[Mourou]
Entry #[14] First Name[Donna] Last Name[Strickland]
Entry #[15] First Name[Frances] Last Name[Arnold]
Entry #[16] First Name[James] Last Name[Allison]
Entry #[17] First Name[Tasuku] Last Name[Honjo]
Entry #[18] First Name[Olga] Last Name[Tokarczuk]
Entry #[19] First Name[Denis] Last Name[Mukwege]
Entry #[20] First Name[Nadia] Last Name[Murad]
Entry #[21] First Name[William] Last Name[Nordhaus]
Entry #[22] First Name[Paul] Last Name[Romer]
Entry #[23] First Name[Rainer] Last Name[Weiss]
Entry #[24] First Name[Barry] Last Name[Barish]
Entry #[25] First Name[Kip] Last Name[Thorne]
Entry #[26] First Name[Jacques] Last Name[Dubochet]
Entry #[27] First Name[Joachim] Last Name[Frank]
Entry #[28] First Name[Richard] Last Name[Henderson]
Entry #[29] First Name[Alice] Last Name[Munro]
Entry #[30] First Name[A-long-first-name-be] Last Name[A-long-surnname-beyo]
```

- On my Windows 10 laptop
  
```sh
Microsoft Windows [Versión 10.0.18362.418]
(c) 2019 Microsoft Corporation. Todos los derechos reservados.

C:\Users\pablo.sancho>cd tmp

C:\Users\pablo.sancho\tmp>go build read.go

C:\Users\pablo.sancho\tmp>read.exe

Please enter the full path for the input file to read:
or enter 'END' or press <CTRL + C> to exit program
C:\Users\pablo.sancho\Documents\Not-Existing-Folder\not-existing-file.txt
File [C:\Users\pablo.sancho\Documents\Not-Existing-Folder\not-existing-file.txt] Doesn't exists.

Please enter the full path for the input file to read:
or enter 'END' or press <CTRL + C> to exit program
C:\Users\pablo.sancho\Documents\SomeNobelPrizes.txt

Processing file[C:\Users\pablo.sancho\Documents\SomeNobelPrizes.txt] ...

Iterating through the resulting slice of persons:
Entry #[0] First Name[cdMichel] Last Name[Mayor]
Entry #[1] First Name[Didier] Last Name[Queloz]
Entry #[2] First Name[Akira] Last Name[Yoshino]
Entry #[3] First Name[John] Last Name[Goodenough]
Entry #[4] First Name[James] Last Name[Peebles]
Entry #[5] First Name[William] Last Name[Kaelin]
Entry #[6] First Name[Peter] Last Name[Ratcliffe]
Entry #[7] First Name[Gregg] Last Name[Semenza]
Entry #[8] First Name[Peter] Last Name[Handke]
Entry #[9] First Name[Abhijit] Last Name[Banerjee]
Entry #[10] First Name[Esther] Last Name[Duflo]
Entry #[11] First Name[Michael] Last Name[Kremer]
Entry #[12] First Name[Arthur] Last Name[Ashkin]
Entry #[13] First Name[Gérard] Last Name[Mourou]
Entry #[14] First Name[Donna] Last Name[Strickland]
Entry #[15] First Name[Frances] Last Name[Arnold]
Entry #[16] First Name[James] Last Name[Allison]
Entry #[17] First Name[Tasuku] Last Name[Honjo]
Entry #[18] First Name[Olga] Last Name[Tokarczuk]
Entry #[19] First Name[Denis] Last Name[Mukwege]
Entry #[20] First Name[Nadia] Last Name[Murad]
Entry #[21] First Name[William] Last Name[Nordhaus]
Entry #[22] First Name[Paul] Last Name[Romer]
Entry #[23] First Name[Rainer] Last Name[Weiss]
Entry #[24] First Name[Barry] Last Name[Barish]
Entry #[25] First Name[Kip] Last Name[Thorne]
Entry #[26] First Name[Jacques] Last Name[Dubochet]
Entry #[27] First Name[Joachim] Last Name[Frank]
Entry #[28] First Name[Richard] Last Name[Henderson]
Entry #[29] First Name[Alice] Last Name[Munro]
Entry #[30] First Name[A-long-first-name-be] Last Name[A-long-surnname-beyo]

```