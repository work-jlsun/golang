package main

import (
    "fmt"
    "bufio"
    "os"
    "io"
)


func main() {

    inputfile, err := os.Open("./readinput.go")
    
    if err != nil {
        fmt.Printf("open file error = %s\n", err.Error())
        return
    }

    inputReader := bufio.NewReader(inputfile)

    for {
        inputstring , err := inputReader.ReadString('\n')

        if (err == nil) {
            fmt.Printf("%s\n", inputstring)
        } else {
            if err == io.EOF {
                break;
            }
            fmt.Printf("input error, err = %s", err.Error())
            break;
        }

    }
}
