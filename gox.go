package gox

import (
    "net/http"
    "log"
    "strings"
    "io"
)

var (
    Certpem string
    Keypem string
)

func WriteString(w io.Writer, s string) {
    io.WriteString(w, s)
}

func Handle(uri string, args ...interface{}) {
    var handler interface{}
    var reqMethods []string
    for _, v := range args {
        switch x := v.(type) {
            case func(w http.ResponseWriter, req *http.Request):
                handler = x
            case string:
                reqMethods = append(reqMethods, strings.ToLower(x))
        }
    }
    http.HandleFunc(uri, func(w http.ResponseWriter, req *http.Request){
        if 0 < len(reqMethods) {
            isValid := false
            reqMethod := strings.ToLower(req.Method)
            for _, v := range reqMethods{
                if v == reqMethod {
                    isValid = true
                }
            }
            if !isValid {
                io.WriteString(w, "http.Request.Method error\n")
                return
            }
        }
        handlerFun := handler.(func(w http.ResponseWriter, req *http.Request))
        handlerFun(w, req)
    }) 
}

func Run(uri string) {
    var err error
    if "" != Certpem && "" != Keypem {
        err = http.ListenAndServeTLS(uri, Certpem, Keypem, nil)
    } else {
        err = http.ListenAndServe(uri, nil)
    }
    log.Fatal(err)
}
