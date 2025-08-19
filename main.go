package main

import (
    "github.com/gin-gonic/gin"
    "github.com/HomieWilliam/analyzer-api/config"
    "github.com/HomieWilliam/analyzer-api/routes"
)

func main() {
    r := gin.Default()
    routes.RegisterRoutes(r)
    port := config.GetPort()
    r.Run(":" + port)
}