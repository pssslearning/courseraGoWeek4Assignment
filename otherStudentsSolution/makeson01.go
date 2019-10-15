package main

import (
        "bufio"
        "encoding/json"
        "fmt"
        "os"
        "strings"
)

func main (){
        data := make(map[string]string)
        reader := bufio.NewReader(os.Stdin)
        fmt.Printf("Name: ")
        name, _ := reader.ReadString('\n')
        fmt.Printf("Address: ")
        address, _ := reader.ReadString('\n')
        data["name"] = strings.Replace( name, "\n", "", 1)
        data["address"] = strings.Replace(address, "\n", "", 1)
        jobject, err := json.Marshal(data)
        if err != nil{
                fmt.Errorf("Error is %s", err)
                os.Exit(1)
        }

        fmt.Printf("%s", jobject)
}