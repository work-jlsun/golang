package main 

import (
    "flag"
    "fmt"
    "os"
)

func cat(f *os.File) {
    const NBUF  = 512
    var buf  [NBUF]byte

    for {
        
        switch bytes, err := f.Read(buf[:]); true {
        case bytes < 0: //read error
            fmt.Fprintf(os.Stderr, "cat : reading err = %s\n", err.Error())
            os.Exit(1)
        case bytes == 0: // EOF
            return
        case bytes > 0:
            if wn, err := os.Stdout.Write(buf[:bytes]); wn != bytes {
                fmt.Fprintf(os.Stderr, "cat : writing err = %s\n",err.Error() )
                return
            }
        }
    }
}


func main() {
    flag.Parse()
    
    if  flag.NArg() == 0 {
        fmt.Fprintf(os.Stdout, "choose you files\n")
        return
    }

    for i := 0; i < flag.NArg(); i++ {
        f, err := os.Open(flag.Arg(i))
        
        if f == nil {
            fmt.Fprintf(os.Stderr, "cat : can't open %s , err = %s\n",
            flag.Arg(i), err.Error())
            os.Exit(1)
        }
        cat(f)
        f.Close()
    }
}
