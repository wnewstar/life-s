package main

import (
    "os"
    "log"
    "fmt"
    "net/http"
    _ "net/http/pprof"
    "life/router"
    "github.com/gin-gonic/gin"
)

func main() {
    go status()
    gin.SetMode(gin.ReleaseMode)
    obj := gin.Default()
    if len(os.Args) >= 2 {
        (router.Route(obj)).Run(os.Args[1])
        fmt.Println("The HTTP service start at host port %s", os.Args[1])
    } else {
        (router.Route(obj)).Run("0.0.0.0:3000")
        fmt.Println("The HTTP service start at default host port 0.0.0.0:3000")
    }
}

func status() {
    if len(os.Args) >= 3 {
        log.Println(http.ListenAndServe(os.Args[2], nil))
        fmt.Println("The HTTP status service start at host port %s", os.Args[2])
    }
}
