package router

import (
    "github.com/gin-gonic/gin"
    "github.com/pp00x/hydrate/internal/handler"
    "github.com/pp00x/hydrate/internal/middleware"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.Use(middleware.CorsMiddleware())

    api := r.Group("/api/v1")
    {
        api.POST("/register", handler.Register)
        api.POST("/login", handler.Login)

        protected := api.Group("/")
        protected.Use(middleware.Authenticate())
        {
            protected.POST("/water-intake", handler.CreateWaterIntake)
            protected.GET("/water-intake", handler.GetWaterIntakes)
        }
    }

    return r
}