package controllers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/HomieWilliam/analyzer-api/models"
    "github.com/HomieWilliam/analyzer-api/services"
)

func TestAnalyzePhrase_ValidRequest(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    r.POST("/api/v1/analyze", AnalyzePhrase)

    reqBody, _ := json.Marshal(models.AnalyzerRequest{Phrase: "Hello world"})
    req, _ := http.NewRequest(http.MethodPost, "/api/v1/analyze", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("Esperado status 200, recebido %d", w.Code)
    }

    var resp models.AnalyzerResponse
    if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
        t.Fatalf("Erro ao decodificar JSON: %v", err)
    }

    expected := services.Analyze("Hello world")
    if resp != expected {
        t.Errorf("Esperado %+v, recebido %+v", expected, resp)
    }
}

func TestAnalyzePhrase_InvalidRequest(t *testing.T) {
    gin.SetMode(gin.TestMode)
    r := gin.Default()
    r.POST("/api/v1/analyze", AnalyzePhrase)

    // JSON inválido
    req, _ := http.NewRequest(http.MethodPost, "/api/v1/analyze", bytes.NewBuffer([]byte(`{bad json}`)))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != http.StatusBadRequest {
        t.Fatalf("Esperado status 400, recebido %d", w.Code)
    }

    var resp map[string]string
    if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
        t.Fatalf("Erro ao decodificar JSON: %v", err)
    }

    if _, ok := resp["error"]; !ok {
        t.Errorf("Esperado campo 'error' na resposta")
    }
}
