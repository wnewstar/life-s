package controller

import (
    "github.com/gin-gonic/gin"
    ModelFore "life/app/fore/model"
)

type Zsystem struct {}

func (*Zsystem) DbConnInfo(c *gin.Context) {
    data := []interface{}{ ModelFore.DbQuery.DB().Stats(), ModelFore.DbWrite.DB().Stats() }

    c.JSON(CODE_HTTP_SUCCESS, gin.H{ "code": CODE_SELECT_SUCCESS_A, "note": TEXT_SELECT_SUCCESS_A, "data": data })
}
