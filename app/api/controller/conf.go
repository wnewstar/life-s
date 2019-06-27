package controller

import (
    "time"
    "github.com/gin-gonic/gin"
    ApiModel "life/app/api/model"
    ApiLogic "life/app/api/logic"
)

type Conf struct {}

func (*Conf) List(c *gin.Context) {
    var Lconf ApiLogic.Conf

    base(c)
    data := Lconf.GetListByUid(uint(Uid))
    c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_SELECT_SUCCESS_A, "note": TEXT_SELECT_SUCCESS_A, "data": data })
}

func (*Conf) Tree(c *gin.Context) {
    var Lconf ApiLogic.Conf

    base(c)
    data := Lconf.GetTreeByUid(uint(Uid))
    c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_SELECT_SUCCESS_A, "note": TEXT_SELECT_SUCCESS_A, "data": data })
}

func (*Conf) Detail(c *gin.Context) {
    var Mconf ApiModel.Conf
    var Mbase ApiModel.Base

    if err := c.ShouldBindJSON(&Mbase); err != nil {
        var data = err.Error()
        c.JSON(CODE_HTTP_FALIURE, gin.H{ "code": CODE_DETAIL_FALIURE_B, "note": TEXT_DETAIL_FALIURE_B, "data": data })
    } else {
        base(c)
        Mconf.Id = Mbase.Id
        Mconf.Uid = uint(Uid)

        if data := Mconf.FindOneByIdUid(Mconf); data.Id > 0 {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_DETAIL_SUCCESS_A, "note": TEXT_DETAIL_SUCCESS_A, "data": data })
        } else {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_DETAIL_FALIURE_A, "note": TEXT_DETAIL_FALIURE_A, "data": data })
        }
    }
}

func (*Conf) Modify(c *gin.Context) {
    var Mconf ApiModel.Conf

    if err := c.ShouldBindJSON(&Mconf); err != nil {
        var data = err.Error()
        c.JSON(CODE_HTTP_FALIURE, gin.H{ "code": CODE_MODIFY_FALIURE_B, "note": TEXT_MODIFY_FALIURE_B, "data": data })
    } else {
        base(c)
        Mconf.Uid = uint(Uid)
        Mconf.UserUpdate = uint(Uid)
        Mconf.TimeUpdate = uint64(time.Now().Unix())

        if data := Mconf.SaveOneByIdUid(Mconf); data > 0 {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_MODIFY_SUCCESS_A, "note": TEXT_MODIFY_SUCCESS_A, "data": data })
        } else {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_MODIFY_FALIURE_A, "note": TEXT_MODIFY_FALIURE_A, "data": data })
        }
    }
}

func (*Conf) Delete(c *gin.Context) {
    var Mconf ApiModel.Conf
    var Mbase ApiModel.Base

    if err := c.ShouldBindJSON(&Mbase); err != nil {
        var data = err.Error()
        c.JSON(CODE_HTTP_FALIURE, gin.H{ "code": CODE_DELETE_FALIURE_B, "note": TEXT_DELETE_FALIURE_B, "data": data })
    } else {
        base(c)
        Mconf = ApiModel.Conf{ Base: ApiModel.Base{ Id: Mbase.Id } }
        Mconf.Uid = uint(Uid)
        Mconf.Status = 1
        Mconf.UserDelete = uint(Uid)
        Mconf.TimeDelete = uint64(time.Now().Unix())

        if data := Mconf.SaveOneByIdUid(Mconf); data > 0 {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_DELETE_SUCCESS_A, "note": TEXT_DELETE_SUCCESS_A, "data": data })
        } else {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_DELETE_FALIURE_A, "note": TEXT_DELETE_FALIURE_A, "data": data })
        }
    }
}

func (*Conf) Create(c *gin.Context) {
    var Mconf ApiModel.Conf

    if err := c.ShouldBindJSON(&Mconf); err != nil {
        var data = err.Error()
        c.JSON(CODE_HTTP_FALIURE, gin.H{ "code": CODE_CREATE_FALIURE_B, "note": TEXT_CREATE_FALIURE_B, "data": data })
    } else {
        base(c)
        Mconf.Uid = uint(Uid)
        Mconf.UserCreate = uint(Uid)
        Mconf.TimeCreate = uint64(time.Now().Unix())

        if data := Mconf.Create(Mconf); data > 0 {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_CREATE_SUCCESS_A, "note": TEXT_CREATE_SUCCESS_A, "data": data })
        } else {
            c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_CREATE_FALIURE_A, "note": TEXT_CREATE_FALIURE_A, "data": data })
        }
    }
}
