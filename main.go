package main
 
import (
    "github.com/gin-gonic/gin"
    "analyzer-api/config"
    "analyzer-api/routes"
)
 
func main() {
    r := gin.Default()
 
    // Registrar rotas
    routes.RegisterRoutes(r)
 
    // Pega a porta do config
    port := config.GetPort()
 
    // Rodar servidor
    r.Run(":" + port)
}