package controller

import (
    "time"
    "github.com/gin-gonic/gin"
    ApiModel "life/app/api/model"
)

type Bill struct {}

func (*Bill) Detail(c *gin.Context) {
    var Mbill ApiModel.Bill
    var Mbase ApiModel.Base

    if err := c.ShouldBindJSON(&Mbase); err != nil {
        var data = err.Error()
        c.JSON(CODE_HTTP_FALIURE, gin.H{ "code": CODE_DETAIL_FALIURE_B, "note": TEXT_DETAIL_FALIURE_B, "data": data })
    } else {
        base(c)
        Mbill.Id = Mbase.Id
        Mbill.Uid = uint(Uid)

        if data := Mbill.FindOneByIdUid(Mbill); data.Base.Id > 0 {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_DETAIL_SUCCESS_A, "note": TEXT_DETAIL_SUCCESS_A, "data": data })
        } else {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_DETAIL_FALIURE_A, "note": TEXT_DETAIL_FALIURE_A, "data": data })
        }
    }
}

func (*Bill) Modify(c *gin.Context) {
    var Mbill ApiModel.Bill

    if err := c.ShouldBindJSON(&Mbill); err != nil {
        var data = err.Error()
        c.JSON(CODE_HTTP_FALIURE, gin.H{ "code": CODE_MODIFY_FALIURE_B, "note": TEXT_MODIFY_FALIURE_B, "data": data })
    } else {
        base(c)
        Mbill.Uid = uint(Uid)
        Mbill.UserUpdate = uint(Uid)
        Mbill.TimeUpdate = uint64(time.Now().Unix())

        if data := Mbill.SaveOneByIdUid(Mbill); data > 0 {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_MODIFY_SUCCESS_A, "note": TEXT_MODIFY_SUCCESS_A, "data": data })
        } else {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_MODIFY_FALIURE_A, "note": TEXT_MODIFY_FALIURE_A, "data": data })
        }
    }
}

func (*Bill) Delete(c *gin.Context) {
    var Mbill ApiModel.Bill
    var Mbase ApiModel.Base

    if err := c.ShouldBindJSON(&Mbase); err != nil {
        var data = err.Error()
        c.JSON(CODE_HTTP_FALIURE, gin.H{ "code": CODE_DELETE_FALIURE_B, "note": TEXT_DELETE_FALIURE_B, "data": data })
    } else {
        base(c)
        Mbill = ApiModel.Bill{ Base: ApiModel.Base{ Id: Mbase.Id } }
        Mbill.Uid = uint(Uid)
        Mbill.Status = 1
        Mbill.UserDelete = uint(Uid)
        Mbill.TimeDelete = uint64(time.Now().Unix())

        if data := Mbill.SaveOneByIdUid(Mbill); data > 0 {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_DELETE_SUCCESS_A, "note": TEXT_DELETE_SUCCESS_A, "data": data })
        } else {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_DELETE_FALIURE_A, "note": TEXT_DELETE_FALIURE_A, "data": data })
        }
    }
}

func (*Bill) Create(c *gin.Context) {
    var Mbill ApiModel.Bill

    if err := c.ShouldBindJSON(&Mbill); err != nil {
        var data = err.Error()
        c.JSON(CODE_HTTP_FALIURE, gin.H{ "code": CODE_CREATE_FALIURE_B, "note": TEXT_CREATE_FALIURE_B, "data": data })
    } else {
        base(c)
        Mbill.Uid = uint(Uid)
        Mbill.UserCreate = uint(Uid)
        Mbill.TimeCreate = uint64(time.Now().Unix())

        if data := Mbill.Create(Mbill); data > 0 {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_CREATE_SUCCESS_A, "note": TEXT_CREATE_SUCCESS_A, "data": data })
        } else {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_CREATE_FALIURE_A, "note": TEXT_CREATE_FALIURE_A, "data": data })
        }
    }
}

func (*Bill) getTag(c *gin.Context) {

}

func (*Bill) Search(c *gin.Context) {
    var Mbill ApiModel.Bill

    base(c)
    data := Mbill.FindSetByUid(uint(Uid))
    c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_SELECT_SUCCESS_A, "note": TEXT_SELECT_SUCCESS_A, "data": data })
}
