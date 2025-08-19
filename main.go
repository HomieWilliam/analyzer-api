package main

import (
    "github.com/gin-gonic/gin"
    "github.com/HomieWilliam/analyzer-api/config"
    "github.com/HomieWilliam/analyzer-api/routes"
)

func main() {
      // Carrega configuração
    cfg := config.LoadConfig() 

    // Cria servidor Gin
    r := gin.Default()

    // Registra rotas
    routes.RegisterRoutes(r)

    // Executa servidor na porta configurada
    r.Run(cfg.ServerPort)
}