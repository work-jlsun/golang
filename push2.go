package main

import (
    "fmt"
    "strconv"
    "net"
)

func GenTestBody() string {

    // cache key
    pushURL := "/docdocdoc/mytest12q1243242323"

    // cache value
    body := "{\"Type\":JPEG,\"Width\":1280,\"Height\":850,\"Size\":114328}"
    bodyLen := len(body)
    response := "HTTP/1.1 200 OK\r\n" + "Content-type: application/json;charset=UTF-8\r\n" + "Content-length: " + strconv.Itoa(bodyLen) + "\r\n\r\n" + body


    // push request body
    responseLen := len(response)
    pushBody := "PUSH " + pushURL  + " HTTP/1.1\r\n" + "Host: nos.netease.com\r\n" + "Content-Length: " + strconv.Itoa(responseLen) +"\r\n\r\n" + response
    return pushBody
}

func push() {
    //resolve addr
    tcpAddr, err := net.ResolveTCPAddr("tcp4", "172.17.2.201:80")
    if err != nil {
        fmt.Println("ResolveTCPAddr err = ", err.Error())
        return
    }

    // establish connection
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        fmt.Println("DialTCP err = ", err.Error())
        return
    }

    // push cache
    body := GenTestBody()
    num, err := conn.Write([]byte(body))

    if err != nil {
        fmt.Println("conn Write err = ", err.Error())
        fmt.Println("num = ", num)
        return
    }

    // read push result
    buf := make([]byte, 1000)
    num ,err = conn.Read(buf)

    if err != nil {
        fmt.Println("conn Read err = ", err.Error())
        return
    }
    fmt.Print("read content = ", string(buf))
}

func main() {
    push()
}
