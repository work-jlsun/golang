package main

import (
    "net/http"
    "log"
    "fmt"
    "runtime"
)


var  buf  []byte = make([]byte, 512)


func handler512(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Connection", "keep-alive")
//  a := []byte("aaaaa...512...aaa")
    w.Header().Set("Content-Length", fmt.Sprintf("%d", len(buf)))
    w.Header().Set("Content-Type", "application/octet-stream")
    w.Header().Set("Server", "ngx_openresty/1.4.3.6")
    w.Write(buf)
}

func main() {
    runtime.GOMAXPROCS(8)
    
    for i := 0; i != 512; i++ {
        buf[i] = 'a';
    }

    http.HandleFunc("/512b", handler512)

    log.Fatal(http.ListenAndServe(":8081", nil))
}
