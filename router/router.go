package router

import (
    "github.com/gin-gonic/gin"
    Middleware "life/plugin/middle"
    ControllerFore "life/app/fore/controller"
)

func Route(r *gin.Engine) (*gin.Engine) {
    r.Use(Middleware.Logger())
    r.Use(Middleware.CheckToken())
    r.Use(Middleware.CheckCrossDomain())

    groupApi := r.Group("/api")
    {
        groupApiFile := groupApi.Group("/file")
        {
            controller := &ControllerFore.File{}

            groupApiFile.POST("/upload", controller.UpLoad)
        }

        groupApiUser := groupApi.Group("/user")
        {
            controller := &ControllerFore.User{}

            groupApiUser.POST("/login", controller.Login)
            groupApiUser.POST("/status", controller.Status)
            groupApiUser.POST("/create", controller.Create)
            groupApiUser.POST("/detail", controller.Detail)
            groupApiUser.POST("/modify", controller.Modify)
            groupApiUser.POST("/delete", controller.Delete)
        }

        groupApiConf := groupApi.Group("/conf")
        {
            controller := &ControllerFore.Conf{}

            groupApiConf.POST("/tree", controller.Tree)
            groupApiConf.POST("/list", controller.List)
            groupApiConf.POST("/create", controller.Create)
            groupApiConf.POST("/detail", controller.Detail)
            groupApiConf.POST("/modify", controller.Modify)
            groupApiConf.POST("/delete", controller.Delete)
        }

        groupApiBill := groupApi.Group("/bill")
        {
            controller := &ControllerFore.Bill{}

            groupApiBill.POST("/search", controller.Search)
            groupApiBill.POST("/create", controller.Create)
            groupApiBill.POST("/detail", controller.Detail)
            groupApiBill.POST("/modify", controller.Modify)
            groupApiBill.POST("/delete", controller.Delete)
        }

        groupApiNote := groupApi.Group("/note")
        {
            controller := &ControllerFore.Note{}

            groupApiNote.POST("/search", controller.Search)
            groupApiNote.POST("/create", controller.Create)
            groupApiNote.POST("/detail", controller.Detail)
            groupApiNote.POST("/modify", controller.Modify)
            groupApiNote.POST("/delete", controller.Delete)
        }

        groupApiZsystem := groupApi.Group("/system")
        {
            controller := &ControllerFore.Zsystem{}

            groupApiZsystem.POST("/dbconninfo", controller.DbConnInfo)
        }
    }

    return r
}
