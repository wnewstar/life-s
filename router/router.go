package router

import (
    "os"
    "github.com/gin-gonic/gin"
    Middleware "life/plugin/middle"
    ApiController "life/app/api/controller"
)

func Route(r *gin.Engine) (*gin.Engine) {
    base := os.Getenv("LIFE_FILE_UPLOAD_PATH")
    r.Static("/static", base + "/static")

    r.Use(gin.Recovery())
    r.Use(Middleware.Logger())
    r.Use(Middleware.CheckToken())
    r.Use(Middleware.CheckCrossDomain())

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

        groupApiBill := groupApi.Group("/bill")
        {
            controller := &ApiController.Bill{}

            groupApiBill.POST("/search", controller.Search)
            groupApiBill.POST("/create", controller.Create)
            groupApiBill.POST("/detail", controller.Detail)
            groupApiBill.POST("/modify", controller.Modify)
            groupApiBill.POST("/delete", controller.Delete)
        }

        groupApiNote := groupApi.Group("/note")
        {
            controller := &ApiController.Note{}

            groupApiNote.POST("/search", controller.Search)
            groupApiNote.POST("/create", controller.Create)
            groupApiNote.POST("/detail", controller.Detail)
            groupApiNote.POST("/modify", controller.Modify)
            groupApiNote.POST("/delete", controller.Delete)
        }
    }

    return r
}
