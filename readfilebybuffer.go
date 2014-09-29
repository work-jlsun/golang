package main


import (
    "fmt"
    "bufio"
    "os"
)


func main() {

    inputFile, err := os.Open("./readinput.go")

    if err != nil {
        fmt.Printf("open file errror err = %s\n", err.Error())
    }

    inputReader := bufio.NewReader(inputFile)
    buf := make([]byte, 1024)
    for {

        n, err := inputReader.Read(buf)
        if  n == 0 {
            fmt.Printf("reach file end\n");
            break;
        }
       
        if err != nil {
            fmt.Printf("read file error, err = %s\n", err.Error())
        }

        if n > 0 {
            fmt.Printf("%s", string(buf[:n]))
        }

    }
}
