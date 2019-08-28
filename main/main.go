package main

import (
    "fmt"
    "flag"
    "net/http"
    _ "net/http/pprof"
    "life/router"
    "github.com/gin-gonic/gin"
)

var (
    hpweb string
    hpsta string
)

func init() {
    flag.StringVar(&hpweb, "hpweb", "127.0.0.1:3000", "nil")
    flag.StringVar(&hpsta, "hpsta", "127.0.0.1:3100", "nil")
}

func main() {
    flag.Parse()

    if (len(hpweb) <= 0) {
        fmt.Println("Please set host and port")
    } else {
        go httpstatus()

        gin.DisableConsoleColor()
        gin.SetMode(gin.ReleaseMode)

        obj := gin.New()
        obj.Use(gin.Recovery())

        router.Route(obj).Run(hpweb)
        fmt.Printf("The HTTP service start at host port %s", hpweb)
    }
}

func httpstatus() {
    if (len(hpsta) > 0) {
        http.ListenAndServe(hpsta, nil)
        fmt.Printf("The HTTP status service start at host port %s", hpsta)
    }
}


