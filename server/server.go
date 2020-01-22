package server

import (
    "github.com/gin-gonic/gin"

    "github.com/kondo1018008/golangPractice/controller"
)

// Init is initialize server
func Init() {
    r := router()
    r.Run()
}

func router() *gin.Engine {
    r := gin.Default()

    u := r.Group("/users")
    {
        ctrl := user.Controller{}
        u.GET("", ctrl.Index)
        u.GET("/:id", ctrl.Show)
        u.POST("", ctrl.Create)
        u.PUT("/:id", ctrl.Update)
        u.DELETE("/:id", ctrl.Delete)
    }

    return r
}