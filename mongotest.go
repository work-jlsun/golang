package main


import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    //. "gopkg.in/check.v1"
)

type KV struct {
    Key string "key"
    Value string "value"
}

func main() {
    session, err := mgo.Dial("127.0.0.1:27017")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("test").C("m1ymy12tkey1")

    index := mgo.Index{
        Key: []string{"key"},
        Unique: true,
        DropDups :true,
        Background: true, // See notes.
        Sparse: true,
    }
    err = c.EnsureIndex(index)
    if err != nil {
        panic(err)
    }

    // insert the keys
    keys := []string{"/ab","/ab/aaag","/ab/aaacdsds","/ab/aaacdsdsds"}
    for _, key := range keys{
        err  = c.Insert(&KV{key, key})
        if err != nil {
            log.Fatal(err)
        }
    }

   // range search the keys
   query := c.Find(bson.M{"key": bson.M{"$gte": "/ab/aaad"}}).Limit(2)

   iter := query.Iter()
    result  := KV{}
    for {
        ok := iter.Next(&result)
        if ok != true {
            break;
        }
        fmt.Println("--->", result.Key, result.Value)
    }

}
