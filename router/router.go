package router

import (
    "os"
    "time"
    "strings"
    "strconv"
    "net/http"
    "encoding/base64"
    "github.com/gin-gonic/gin"
    ApiController "life-b/app/api/controller"
)

func Route(r *gin.Engine) (*gin.Engine) {
    base := os.Getenv("LIFE_FILE_UPLOAD_PATH")
    r.Static("/static", base + "/static")

    r.Use(CheckLogin())
    r.Use(CheckCrossDomain())

    groupApi := r.Group("/api")
    {
        groupApiUser := groupApi.Group("/user")
        {
            controller := &ApiController.User{}

            groupApiUser.POST("/login", controller.Login)
            groupApiUser.POST("/status", controller.Status)
            groupApiUser.POST("/create", controller.Create)
            groupApiUser.POST("/detail", controller.Detail)
            groupApiUser.POST("/modify", controller.Modify)
            groupApiUser.POST("/delete", controller.Delete)
        }

        groupApiFile := groupApi.Group("/file")
        {
            controller := &ApiController.File{}

            groupApiFile.POST("/upload", controller.UpLoad)
        }

        groupApiConf := groupApi.Group("/conf")
        {
            controller := &ApiController.Conf{}

            groupApiConf.GET("/list", controller.List)
            groupApiConf.GET("/tree", controller.Tree)
            groupApiConf.POST("/create", controller.Create)
            groupApiConf.POST("/detail", controller.Detail)
            groupApiConf.POST("/modify", controller.Modify)
            groupApiConf.POST("/delete", controller.Delete)
        }

        groupApiBill := groupApi.Group("/bill")
        {
            controller := &ApiController.Bill{}

            groupApiBill.GET("/list", controller.List)
            groupApiBill.POST("/create", controller.Create)
            groupApiBill.POST("/detail", controller.Detail)
            groupApiBill.POST("/modify", controller.Modify)
            groupApiBill.POST("/delete", controller.Delete)
        }

        groupApiNote := groupApi.Group("/note")
        {
            controller := &ApiController.Note{}

            groupApiNote.GET("/list", controller.List)
            groupApiNote.POST("/create", controller.Create)
            groupApiNote.POST("/detail", controller.Detail)
            groupApiNote.POST("/modify", controller.Modify)
            groupApiNote.POST("/delete", controller.Delete)
        }
    }

    return r
}

func CheckLogin() (gin.HandlerFunc) {
    return func(c *gin.Context) {
        if c.Request.Method != "OPTIONS" &&
            c.Request.URL.Path != "/api/user/login" {
            token := c.Request.Header.Get("X-AUTH-TOKEN")
            chars, err := base64.StdEncoding.DecodeString(token)
            if err != nil || len(chars) == 0 {
                defer c.Abort()
                c.JSON(http.StatusOK, gin.H{ "code": "0", "text": "没有登录" })
            } else {
                if strs := strings.Split(string(chars), "|"); len(strs) != 2 {
                    defer c.Abort()
                    c.JSON(http.StatusOK, gin.H{ "code": "0", "text": "无效登陆" })
                } else {
                    ntime := uint64(time.Now().Unix())
                    etime, err := strconv.ParseUint(strs[1], 10, 64)
                    if err == nil && etime > ntime {
                        c.Request.Header.Set("X-AUTH-USERID", strs[0])
                    } else {
                        defer c.Abort()
                        c.JSON(http.StatusOK, gin.H{ "code": "0", "text": "登录超时" })
                    }
                }
            }
        }
    }
}

func CheckCrossDomain() (gin.HandlerFunc) {
    return func(c *gin.Context) {
        if c.Request.Method == "OPTIONS" {
            defer c.JSON(http.StatusOK, gin.H{ "code": "0" })
            defer c.Abort()
        } else {
            defer c.Next()
        }

        w := c.Writer
        w.Header().Set("Access-Control-Allow-Origin", "*") 
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.Header().Add("Access-Control-Allow-Headers", "X-AUTH-TOKEN")
        w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, HEAD, GET, PUT, POST, PATCH, DELETE")
    }
}
