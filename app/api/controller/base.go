package controller

import (
    "strconv"
    "github.com/gin-gonic/gin"
)

const (
    CODE_HTTP_SUCCESS = 200
    CODE_HTTP_FALIURE = 400

    CODE_SUCCESS_A = "0"
    TEXT_SUCCESS_A = "成功"

    CODE_LOGIN_SUCCESS_A = "0"
    CODE_LOGIN_FALIURE_A = "1001"
    CODE_LOGIN_FALIURE_B = "1002"
    TEXT_LOGIN_SUCCESS_A = "登录成功"
    TEXT_LOGIN_FALIURE_A = "登录失败"
    TEXT_LOGIN_FALIURE_B = "参数错误"

    CODE_UPLOAD_SUCCESS_A = "0"
    CODE_UPLOAD_FALIURE_A = "1101"
    CODE_UPLOAD_FALIURE_B = "1102"
    TEXT_UPLOAD_SUCCESS_A = "上传成功"
    TEXT_UPLOAD_FALIURE_A = "上传失败"
    TEXT_UPLOAD_FALIURE_B = "参数错误"

    CODE_SELECT_SUCCESS_A = "0"
    CODE_SELECT_FALIURE_A = "2001"
    CODE_SELECT_FALIURE_B = "2002"
    TEXT_SELECT_SUCCESS_A = "查询成功"
    TEXT_SELECT_FALIURE_A = "查询失败"
    TEXT_SELECT_FALIURE_B = "参数错误"

    CODE_CREATE_SUCCESS_A = "0"
    CODE_CREATE_FALIURE_A = "3001"
    CODE_CREATE_FALIURE_B = "3002"
    TEXT_CREATE_SUCCESS_A = "新增成功"
    TEXT_CREATE_FALIURE_A = "新增失败"
    TEXT_CREATE_FALIURE_B = "参数错误"

    CODE_DETAIL_SUCCESS_A = "0"
    CODE_DETAIL_FALIURE_A = "4001"
    CODE_DETAIL_FALIURE_B = "4002"
    TEXT_DETAIL_SUCCESS_A = "查询成功"
    TEXT_DETAIL_FALIURE_A = "查询失败"
    TEXT_DETAIL_FALIURE_B = "参数错误"

    CODE_MODIFY_SUCCESS_A = "0"
    CODE_MODIFY_FALIURE_A = "5001"
    CODE_MODIFY_FALIURE_B = "5002"
    TEXT_MODIFY_SUCCESS_A = "修改成功"
    TEXT_MODIFY_FALIURE_A = "修改失败"
    TEXT_MODIFY_FALIURE_B = "参数错误"

    CODE_DELETE_SUCCESS_A = "0"
    CODE_DELETE_FALIURE_A = "6001"
    CODE_DELETE_FALIURE_B = "6002"
    TEXT_DELETE_SUCCESS_A = "删除成功"
    TEXT_DELETE_FALIURE_A = "删除失败"
    TEXT_DELETE_FALIURE_B = "参数错误"
)

var Uid uint64

func base(c *gin.Context) {
    Uid, _ = strconv.ParseUint(c.Request.Header.Get("X-AUTH-USERID"), 10, 32)
}
