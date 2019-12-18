package controller

import (
    "os"
    "time"
    "path"
    "crypto/md5"
    "encoding/hex"
    "github.com/gin-gonic/gin"
)

type File struct {}

func (*File) UpLoad(c *gin.Context) {
    if file, err := c.FormFile("file"); err != nil {
        data := err.Error()
        c.JSON(CODE_HTTP_FALIURE, gin.H{ "code": CODE_UPLOAD_FALIURE_B, "note": TEXT_UPLOAD_FALIURE_B, "data": data })
    } else {
        base := os.Getenv("LIFE_FILE_UPLOAD_PATH")
        date := time.Now().Format("20060102")
        temp := md5.Sum([]byte(file.Filename))
        name := "/static/upload/image/" + date + "/" + hex.EncodeToString(temp[:]) + path.Ext(file.Filename)

        os.MkdirAll(path.Dir(base + name), os.ModePerm)

        if err = c.SaveUploadedFile(file, base + name); err != nil {
            data := err.Error()
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_UPLOAD_FALIURE_A, "note": TEXT_UPLOAD_FALIURE_A, "data": data })
        } else {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_UPLOAD_SUCCESS_A, "note": TEXT_UPLOAD_SUCCESS_A, "data": name })
        }
    }
}
