package controller

import (
    "time"
    "github.com/gin-gonic/gin"
    ModelFore "life/app/fore/model"
)

type Note struct {}

func (*Note) Detail(c *gin.Context) {
    var Mnote ModelFore.Note
    var Mbase ModelFore.Base

    if err := c.ShouldBindJSON(&Mbase); err != nil {
        var data = err.Error()
        c.JSON(CODE_HTTP_FALIURE, gin.H{ "code": CODE_DETAIL_FALIURE_B, "note": TEXT_DETAIL_FALIURE_B, "data": data })
    } else {
        base(c)
        Mnote.Id = Mbase.Id
        Mnote.Uid = uint(Uid)

        if data := Mnote.FindOneByIdUid(Mnote); data.Id > 0 {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_DETAIL_SUCCESS_A, "note": TEXT_DETAIL_SUCCESS_A, "data": data })
        } else {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_DETAIL_FALIURE_A, "note": TEXT_DETAIL_FALIURE_A, "data": data })
        }
    }
}

func (*Note) Modify(c *gin.Context) {
    var Mnote ModelFore.Note

    if err := c.ShouldBindJSON(&Mnote); err != nil {
        var data = err.Error()
        c.JSON(CODE_HTTP_FALIURE, gin.H{ "code": CODE_MODIFY_FALIURE_B, "note": TEXT_MODIFY_FALIURE_B, "data": data })
    } else {
        base(c)
        Mnote.Uid = uint(Uid)
        Mnote.UserUpdate = uint(Uid)
        Mnote.TimeUpdate = uint64(time.Now().Unix())

        if data := Mnote.SaveOneByIdUid(Mnote); data > 0 {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_MODIFY_SUCCESS_A, "note": TEXT_MODIFY_SUCCESS_A, "data": data })
        } else {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_MODIFY_FALIURE_A, "note": TEXT_MODIFY_FALIURE_A, "data": data })
        }
    }
}

func (*Note) Delete(c *gin.Context) {
    var Mnote ModelFore.Note
    var Mbase ModelFore.Base

    if err := c.ShouldBindJSON(&Mbase); err != nil {
        var data = err.Error()
        c.JSON(CODE_HTTP_FALIURE, gin.H{ "code": CODE_DELETE_FALIURE_B, "note": TEXT_DELETE_FALIURE_B, "data": data })
    } else {
        base(c)
        Mnote = ModelFore.Note{ Base: ModelFore.Base{ Id: Mbase.Id } }
        Mnote.Uid = uint(Uid)
        Mnote.Status = 1
        Mnote.UserDelete = uint(Uid)
        Mnote.TimeDelete = uint64(time.Now().Unix())

        if data := Mnote.SaveOneByIdUid(Mnote); data > 0 {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_DELETE_SUCCESS_A, "note": TEXT_DELETE_SUCCESS_A, "data": data })
        } else {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_DELETE_FALIURE_A, "note": TEXT_DELETE_FALIURE_A, "data": data })
        }
    }
}

func (*Note) Create(c *gin.Context) {
    var Mnote ModelFore.Note

    if err := c.ShouldBindJSON(&Mnote); err != nil {
        var data = err.Error()
        c.JSON(CODE_HTTP_FALIURE, gin.H{ "code": CODE_CREATE_FALIURE_B, "note": TEXT_CREATE_FALIURE_B, "data": data })
    } else {
        base(c)
        Mnote.Uid = uint(Uid)
        Mnote.UserCreate = uint(Uid)
        Mnote.TimeCreate = uint64(time.Now().Unix())

        if data := Mnote.Create(Mnote); data > 0 {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_CREATE_SUCCESS_A, "note": TEXT_CREATE_SUCCESS_A, "data": data })
        } else {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_CREATE_FALIURE_A, "note": TEXT_CREATE_FALIURE_A, "data": data })
        }
    }
}

func (*Note) Search(c *gin.Context) {
    var Mnote ModelFore.Note

    base(c)
    data := Mnote.FindSetByUid(uint(Uid))

    c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_SELECT_SUCCESS_A, "note": TEXT_SELECT_SUCCESS_A, "data": data })
}
