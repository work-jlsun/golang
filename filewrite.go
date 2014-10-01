package main

import (
    "os"
    "fmt"
)


func main() {

    os.Stdin.WriteString("hello world")

    openFile, err := os.OpenFile("test", os.O_CREATE|os.O_WRONLY,0666)
    
    if err != nil {
        fmt.Printf("%s\n", err.Error())
    }
    defer openFile.Close()

    openFile.WriteString("hello world")
}
