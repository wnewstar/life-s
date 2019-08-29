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
    "github.com/lestrrat-go/file-rotatelogs"
)

func Logger() (gin.HandlerFunc) {
    log := logrus.New()

    h, _ := os.Hostname()
    src, _ := os.OpenFile(os.DevNull, os.O_APPEND | os.O_WRONLY, os.ModeAppend)

    log.Out = src
    log.SetLevel(logrus.DebugLevel)

    pacc := "/data/log/acc-" + h
    writera, _ := rotatelogs.New(
        pacc + ".%Y%m%d",
        rotatelogs.WithLinkName(pacc),
        rotatelogs.WithMaxAge(99 * 24 * time.Hour),
        rotatelogs.WithRotationTime(24 * time.Hour),
    )
    perr := "/data/log/err-" + h
    writerb, _ := rotatelogs.New(
        perr + ".%Y%m%d",
        rotatelogs.WithLinkName(perr),
        rotatelogs.WithMaxAge(99 * 24 * time.Hour),
        rotatelogs.WithRotationTime(24 * time.Hour),
    )
    writemap := lfshook.WriterMap{
        logrus.InfoLevel: writera,
        logrus.FatalLevel: writerb,
    }
    log.AddHook(lfshook.NewHook(writemap, &logrus.JSONFormatter{}))

    return func (c *gin.Context) {
        s := time.Now().UnixNano() / 1e6
        c.Next()  
        e := time.Now().UnixNano() / 1e6

        diff := e - s
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
                c.JSON(http.StatusOK, gin.H{ "code": "0", "text": "没有登录" })
                c.Abort()
                return
            }
            infos := strings.Split(string(chars), "|")
            if len(infos) != 2 {
                c.JSON(http.StatusOK, gin.H{ "code": "0", "text": "无效账号" })
                c.Abort()
                return
            }
            ntime := uint64(time.Now().Unix())
            etime, err := strconv.ParseUint(infos[1], 10, 64)
            if err == nil && etime > ntime {
                c.Request.Header.Set("X-AUTH-USERID", infos[0])
            } else {
                c.Abort()
                c.JSON(http.StatusOK, gin.H{ "code": "0", "text": "登录超时" })
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
