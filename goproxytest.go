package main

import (
    "net/http"
    "log"
    "io"
    "fmt"
    "runtime"
    "io/ioutil"
    "net"
//    "time"
)


var  tr *http.Transport = &http.Transport{
            Dial: func(network, addr string) (net.Conn, error) { 
                //log.Println("dial!") 
                return net.Dial(network, addr) 
            },
        DisableKeepAlives: false,
        MaxIdleConnsPerHost: 5000,
 }

var  httpClient *http.Client = &http.Client{Transport: tr}



func ProxyServer(w http.ResponseWriter, req *http.Request) {
    //log.Print("Path =", req.URL.Path)
    //log.Print("query =", req.URL.RawQuery)
    //req, err := http.NewRequest("GET", "http://10.160.134.175:28086" + 
    //req.URL.Path + "?" + req.URL.RawQuery, nil)
    //fmt.Println(req.URL.Path)
    //newPath := "http://noscoder2.server.163.org:8081" + req.URL.Path + "?" + req.URL.RawQuery

    // create request 
    newPath := "http://172.17.2.201:8686/ab-test" 
    newReq, err := http.NewRequest("GET", newPath, nil)
    newReq.Header.Set("Connection", "Keep-Alive")

    if err != nil {
        fmt.Printf("new Request error\n")
        return
    }

    /*
    tr := &http.Transport{
            Dial: func(network, addr string) (net.Conn, error) { 
                log.Println("dial!") 
                return net.Dial(network, addr) 
            },
        DisableKeepAlives: false,
        MaxIdleConnsPerHost: 10,
    }

    client := &http.Client{Transport: tr}
   */

    //do request
    resp, err := httpClient.Do(newReq)
    if (err !=  nil) {
        io.WriteString(w, err.Error() + ", error\n")
        return
    }else {
        //defer resp.Body.Close()
       // fmt.Printf("client Do ok")
    }

    //return headers
    for headerKey, headerVal := range resp.Header {
        for _, value := range headerVal {
           w.Header().Add(headerKey, value)
        }
    }
    w.Header().Set("Connection", "keep-alive")

    statusCode :=  resp.StatusCode
    w.WriteHeader(statusCode)
    // get the content Length
    /*
    contentLength, ok := resp.Header["Content-Length"]
    if ok {
        fmt.Print("i hava add the header\n")
        //add the header to return, and return header
        w.Header().Add("Content-Length", contentLength[0])
        w.WriteHeader(200)
    }
    */
    //return body
    bytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("read all from nody error")
        return
    }
    resp.Body.Close()
    writeNum, err := w.Write(bytes)
    if err != nil || writeNum == 0 {
        fmt.Printf("write  error, err = %s\n", err.Error())
        return
    }

   /* 
    buf := make([]byte, 1024) 
    
    for {
        n, err := resp.Body.Read(buf)
        if  n == 0 {
            fmt.Printf("reach  end\n");
            break;
        }

        if err != nil {
            fmt.Printf("read  error, err = %s\n", err.Error())
            break
        }
        fmt.Printf("have read %s bytes\n", n)
        // sleep
        //time.Sleep(1 * time.Second)
        writeNum, err := w.Write(buf[:n])
        if err != nil || writeNum == 0 {
            fmt.Printf("write  error, err = %s\n", err.Error())
        }
    }
    */
    
}

func main() {
    runtime.GOMAXPROCS(8)
    fmt.Printf("will listen on 8066")
    http.HandleFunc("/", ProxyServer) 
    err :=  http.ListenAndServe(":8066", nil)
    if err != nil {
        log.Fatal("ListenAndServer error: ", err)
    }
}
