package controllers
 
import (
    "github.com/gin-gonic/gin"
    "analyzer-api/models"
    "analyzer-api/services"
    "net/http"
)
 
func AnalyzePhrase(c *gin.Context) {
    var req models.AnalyzerRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
 
    result := services.Analyze(req.Phrase)
 
    c.JSON(http.StatusOK, result)
}