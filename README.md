# gox
它是一个很简洁的golang框架，为快速集成业务而写

# demo
package main

import (
    "net/http"
    "github.com/gox"
    "encoding/json"
    "fmt"
    "github.com/gomodule/redigo/redis"
)

func ping(w http.ResponseWriter, req *http.Request) {
    gox.WriteString(w, "pong\n")
}

type testDB struct {
    Item string `json:"item"`
    Price float64 `json:"price"`
}

func query(w http.ResponseWriter, req *http.Request) {
    db := gox.DBconn()
    defer db.Close()
    //查询多条
    select_sql := "select * from test where item != ?"
    select_rows,select_err := db.Query(select_sql,"16")
    if select_err != nil {
        gox.WriteString(w, "查询test表失败\n")
        return
    }
    defer select_rows.Close()
    var testSlice []testDB
    for select_rows.Next(){
        var test testDB
        if err := select_rows.Scan(&test.Item,&test.Price); err != nil {
            gox.WriteString(w, "获取test值失败\n")
            return
        }
        testSlice = append(testSlice, test)
    }
    b, err := json.Marshal(testSlice)
    if err != nil {
        gox.WriteString(w, "json encode 失败\n")
    }
    gox.WriteString(w, string(b))
}

func redisGet(w http.ResponseWriter, req *http.Request) {
    r, err := gox.RedisConn()
    if err != nil {

    }
    defer r.Close()
    s, err := redis.String(r.Do("GET", "rec_info"))
    fmt.Printf("%#v\n", s)  // "world"
    fmt.Printf(s)
}

func main() {
    gox.DBconfig= "root:chr@911225@/test?charset=utf8"
    gox.InitConn("tcp", "127.0.0.1:6379", "chr@911225")
    gox.Handle("/ping", ping)
    gox.Handle("/query", query)
    gox.Handle("/get", redisGet)
    gox.Run(":1324")
}
