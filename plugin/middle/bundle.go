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

var log = logrus.New()

var host, _ = os.Hostname()

func init() {
    log.SetLevel(logrus.DebugLevel)

    pvisit := "/data/log/visit"
    writera, _ := rotatelogs.New(
        pvisit + ".%Y%m%d",
        rotatelogs.WithLinkName(pvisit),
        rotatelogs.WithMaxAge(99 * 24 * time.Hour),
        rotatelogs.WithRotationTime(24 * time.Hour),
    )
    pdebug := "/data/log/debug"
    writerb, _ := rotatelogs.New(
        pdebug + ".%Y%m%d",
        rotatelogs.WithLinkName(pdebug),
        rotatelogs.WithMaxAge(99 * 24 * time.Hour),
        rotatelogs.WithRotationTime(24 * time.Hour),
    )
    perror := "/data/log/error"
    writerc, _ := rotatelogs.New(
        perror + ".%Y%m%d",
        rotatelogs.WithLinkName(perror),
        rotatelogs.WithMaxAge(99 * 24 * time.Hour),
        rotatelogs.WithRotationTime(24 * time.Hour),
    )
    writemap := lfshook.WriterMap{
        logrus.InfoLevel:  writera,
        logrus.WarnLevel:  writerc,
        logrus.DebugLevel: writerb,
        logrus.ErrorLevel: writerc,
        logrus.FatalLevel: writerc,
        logrus.PanicLevel: writerc,
    }
    log.AddHook(lfshook.NewHook(writemap, &logrus.JSONFormatter{}))
    log.Out, _ = os.OpenFile(os.DevNull, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
}

/**
 * @note        日志记录
 * @date        2019-12-06
 * @author      wnewstar
 */
func Logger() (gin.HandlerFunc) {
    return func (c *gin.Context) {
        defer LogRecover(c)

        s := time.Now().UnixNano() / 1e6
        c.Next()
        e := time.Now().UnixNano() / 1e6

        t := e - s
        url := c.Request.URL.Path
        method := c.Request.Method
        status := c.Writer.Status()

        log.Infof("%s|%s|%d|%v|%s|%s", host, method, status, t, c.ClientIP(), url)
    }
}

func LogRecover(c *gin.Context) {
    if r := recover(); r != nil {
        t := time.Now().UnixNano() / 1e6

        url := c.Request.URL.Path
        method := c.Request.Method
        status := c.Writer.Status()

        log.Panicf("%s|%s|%d|%v|%s|%s|%s", host, method, status, t, c.ClientIP(), url, r)
    }
}

/**
 * @desc        登录检查
 * @date        2019-12-06
 * @author      wnewstar
 */
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

/**
 * @desc        跨域处理
 * @date        2019-12-06
 * @author      wnewstar
 */
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
