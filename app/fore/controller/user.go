package controller

import (
    "time"
    "golang.org/x/crypto/bcrypt"
    "github.com/gin-gonic/gin"
    ModelFore "life/app/fore/model"
    LogicFore "life/app/fore/logic"
)

type User struct {}

func (*User) Login(c *gin.Context) {
    var Luser LogicFore.User
    var Muser ModelFore.User

    if err := c.ShouldBindJSON(&Muser); err != nil {
        var data = err.Error()
        c.JSON(CODE_HTTP_FALIURE, gin.H{ "code": CODE_LOGIN_FALIURE_B, "note": TEXT_LOGIN_FALIURE_B, "data": data })
    } else {
        if data, err := Luser.Login(Muser); err == nil {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_LOGIN_SUCCESS_A, "note": TEXT_LOGIN_SUCCESS_A, "data": data })
        } else {
            var data = err.Error()
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_LOGIN_FALIURE_A, "note": TEXT_LOGIN_FALIURE_A, "data": data })
        }
    }
}

func (*User) Status(c *gin.Context) {

    c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_SUCCESS_A, "note": TEXT_SUCCESS_A, "data": ModelFore.Musers })
}

func (*User) Detail(c *gin.Context) {

}

func (*User) Modify(c *gin.Context) {

}

func (*User) Delete(c *gin.Context) {

}

func (*User) Create(c *gin.Context) {
    var Muser ModelFore.User

    if err := c.ShouldBindJSON(&Muser); err != nil {
        var data = err.Error()
        c.JSON(CODE_HTTP_FALIURE, gin.H{ "code": CODE_CREATE_FALIURE_B, "note": TEXT_CREATE_FALIURE_B, "data": data })
    } else {
        base(c)

        password, _ := bcrypt.GenerateFromPassword([]byte(Muser.Password), bcrypt.DefaultCost)

        Muser.Password = string(password)
        Muser.UserCreate = uint(Uid)
        Muser.TimeCreate = uint64(time.Now().Unix())

        if data := Muser.Create(Muser); data > 0 {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_CREATE_SUCCESS_A, "note": TEXT_CREATE_SUCCESS_A, "data": data })
        } else {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_CREATE_FALIURE_A, "note": TEXT_CREATE_FALIURE_A, "data": data })
        }
    }
}
