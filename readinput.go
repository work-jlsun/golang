package main

import (
    "fmt"
    "bufio"
    "os"
)


func main() {
    inputReader  := bufio.NewReader(os.Stdin)

    fmt.Println("please input ... -->")

    input, err  := inputReader.ReadString('\n')
    
    if err == nil {
        fmt.Printf("The input string is %s\n",input)
    }

}
