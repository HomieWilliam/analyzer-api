package routes
 
import (
    "github.com/gin-gonic/gin"
    "github.com/HomieWilliam/analyzer-api/controllers"
    "github.com/HomieWilliam/analyzer-api/auth"
)
 
func RegisterRoutes(r *gin.Engine) {
    r.POST("/api/v1/login", controllers.Login)
    
    api := r.Group("/api/v1") 
    api.Use(auth.JWTAuthMiddleware())
    {
        api.POST("/analyze", controllers.AnalyzePhrase)
    }
}