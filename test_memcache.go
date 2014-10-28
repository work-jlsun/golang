package main
 
import (
    "github.com/bradfitz/gomemcache/memcache"
    "fmt"
    "time"
    "io/ioutil"
)

func main() {
    mc := memcache.New("127.0.0.1:11211")
    now := time.Now().Unix()
    fmt.Println(now)
    
    // read file and then write to memcache
    bytes, err := ioutil.ReadFile("test.jpg")    

    mc.Set(&memcache.Item{Key: "foo", Value: bytes, Flags: uint32(now)})
    it, err := mc.Get("foo")
    if err != nil {
        fmt.Println(err.Error())
        return 
    }
    // read from memcache and the store to file    
    fmt.Println(string(it.Key))
    fmt.Println(it.Flags)
    err = ioutil.WriteFile("write.jpg", bytes, 0644)
    if err != nil {
        fmt.Printf("write file error  = %s", err.Error())
        return
    }
    time.Sleep(time.Second) 
}
