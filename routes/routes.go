package routes
 
import (
    "github.com/gin-gonic/gin"
    "analyzer-api/controllers"
)
 
func RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api/v1") // prefixo de versão
    {
        api.POST("/analyze", controllers.AnalyzePhrase)
    }
}