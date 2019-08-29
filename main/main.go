package main

import (
    "fmt"
    "flag"
    "life/router"
    "github.com/gin-gonic/gin"
)

var (
    hpweb string
)

func init() {
    flag.StringVar(&hpweb, "hpweb", ":3000", "nil")
}

func main() {
    flag.Parse()

    if (len(hpweb) <= 0) {
        fmt.Printf("Please set host and port\n")
    } else {
        gin.SetMode(gin.ReleaseMode)

        router.Route(gin.New()).Run(hpweb)
        fmt.Printf("The HTTP service start at host port %s\n", hpweb)
    }
}
