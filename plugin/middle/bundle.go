package middle

import (
    "os"
    "time"
    "strings"
    "strconv"
    "net/http"
    "encoding/base64"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "github.com/rifflock/lfshook"
    "github.com/lestrrat/go-file-rotatelogs"
)

func Logger() (gin.HandlerFunc) {
    log := logrus.New()

    src, _ := os.OpenFile(os.DevNull, os.O_APPEND | os.O_WRONLY, os.ModeAppend)

    log.Out = src
    log.SetLevel(logrus.DebugLevel)

    writera, _ := rotatelogs.New(
        "data/log/api-" + "%Y-%m-%d.log",
        rotatelogs.WithLinkName("data/log/api-"),
        rotatelogs.WithMaxAge(30 * 24 * time.Hour),
        rotatelogs.WithRotationTime(24 * time.Hour),
    )
    writemap := lfshook.WriterMap{
        logrus.InfoLevel: writera,
        logrus.FatalLevel: writera,
    }
    log.AddHook(lfshook.NewHook(writemap, &logrus.JSONFormatter{}))

    return func (c *gin.Context) {
        s := time.Now()
        c.Next()  
        e := time.Now()

        diff := e.Sub(s)
        method := c.Request.Method
        status := c.Writer.Status()

        log.Infof("%s|%d|%v|%s|%s", method, status, diff, c.ClientIP(), c.Request.URL.Path)
    }
}

func CheckToken() (gin.HandlerFunc) {
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
                    c.JSON(http.StatusOK, gin.H{ "code": "0", "text": "无效账号" })
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
