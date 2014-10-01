package main

import (
    "os"
    "bufio"
    "fmt"
)

func main () {
    outputFile, err := os.OpenFile("output.data",
    os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0666)

    if err != nil {
        fmt.Printf("open file err = %s\n",err.Error())
    }
    defer outputFile.Close()

    outputWriter := bufio.NewWriter(outputFile)
    
    for i := 0; i < 10; i++ {
        outputWriter.WriteString("hello world\n")
    }
    outputWriter.Flush()

}
