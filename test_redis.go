package main

import (
    "fmt"
    "github.com/garyburd/redigo/redis"
    "io/ioutil"
)


func main() {
    // connect to  redis
    c, err := redis.Dial("tcp", ":6379")
    if err != nil {
        // handle error
        fmt.Printf("connect fail err = %s\n", err.Error())
    }
    defer c.Close()
   
    // read file and then write to memcache
    bytes, err := ioutil.ReadFile("test.jpg")    
    if err != nil {
        fmt.Printf("read file error,err = %s\n",err.Error())
        return
    } 
    _,err = c.Do("SET", "Key",bytes)  
    if err != nil {
        fmt.Printf("set redis error,err = %s\n",err.Error())
        return
    }

    // get the return & store the file
    s, err := redis.Bytes(c.Do("GET", "Key"))
    if err != nil {
        fmt.Printf("redis GET error = %s\n", err.Error())
        return
    } 
    fmt.Printf("Read bytes = %d\n", len(s)) 

    err = ioutil.WriteFile("write2.jpg", s, 0644)
    if err != nil {
        fmt.Printf("write file error  = %s", err.Error())
        return
    }
    
    // read part from redis & compact to one file
    fileLength := len(s)
    resultbytes := make([]byte, 0) 
    i := 0
    j := 0
    for {
        j = i + 100
        if j > fileLength {
            j = -1
        } 
        bytes, err := redis.Bytes(c.Do("GETRANGE", "Key", i, j))
        if err != nil {
            fmt.Printf("redis GET error = %s\n", err.Error())
            return
        }
        resultbytes = append(resultbytes, bytes...) 
        if j == -1 { break }
        i = j + 1
    }
   
    
    err = ioutil.WriteFile("write3.jpg", resultbytes, 0644)
    if err != nil {
        fmt.Printf("write file error  = %s", err.Error())
        return
    }

}
