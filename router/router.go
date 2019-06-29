package router

import (
    "os"
    "github.com/gin-gonic/gin"
    middle "life/plugin/middle"
    ApiController "life/app/api/controller"
)

func Route(r *gin.Engine) (*gin.Engine) {
    base := os.Getenv("LIFE_FILE_UPLOAD_PATH")
    r.Static("/static", base + "/static")

    r.Use(middle.Logger())
    r.Use(middle.CheckToken())
    r.Use(middle.CheckCrossDomain())

    groupApi := r.Group("/api")
    {
        groupApiUser := groupApi.Group("/user")
        {
            controller := &ApiController.User{}

            groupApiUser.POST("/login", controller.Login)
            groupApiUser.POST("/status", controller.Status)
            groupApiUser.POST("/create", controller.Create)
            groupApiUser.POST("/detail", controller.Detail)
            groupApiUser.POST("/modify", controller.Modify)
            groupApiUser.POST("/delete", controller.Delete)
        }

        groupApiFile := groupApi.Group("/file")
        {
            controller := &ApiController.File{}

            groupApiFile.POST("/upload", controller.UpLoad)
        }

        groupApiConf := groupApi.Group("/conf")
        {
            controller := &ApiController.Conf{}

            groupApiConf.GET("/list", controller.List)
            groupApiConf.GET("/tree", controller.Tree)
            groupApiConf.GET("/search", controller.List)
            groupApiConf.POST("/create", controller.Create)
            groupApiConf.POST("/detail", controller.Detail)
            groupApiConf.POST("/modify", controller.Modify)
            groupApiConf.POST("/delete", controller.Delete)
        }

        groupApiBill := groupApi.Group("/bill")
        {
            controller := &ApiController.Bill{}

            groupApiBill.GET("/list", controller.List)
            groupApiBill.GET("/search", controller.List)
            groupApiBill.POST("/create", controller.Create)
            groupApiBill.POST("/detail", controller.Detail)
            groupApiBill.POST("/modify", controller.Modify)
            groupApiBill.POST("/delete", controller.Delete)
        }

        groupApiNote := groupApi.Group("/note")
        {
            controller := &ApiController.Note{}

            groupApiNote.GET("/list", controller.List)
            groupApiNote.GET("/search", controller.List)
            groupApiNote.POST("/create", controller.Create)
            groupApiNote.POST("/detail", controller.Detail)
            groupApiNote.POST("/modify", controller.Modify)
            groupApiNote.POST("/delete", controller.Delete)
        }
    }

    return r
}
