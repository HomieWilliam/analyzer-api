package config

import (
    "os"
    "testing"
)

func TestLoadConfig_DefaultPort(t *testing.T) {
    os.Unsetenv("PORT") // garante que PORT não está definida
    cfg := LoadConfig()
    if cfg.ServerPort != ":8080" {
        t.Errorf("Esperado :8080, recebido %s", cfg.ServerPort)
    }
}

func TestLoadConfig_CustomPort(t *testing.T) {
    os.Setenv("PORT", ":9090")
    defer os.Unsetenv("PORT") // limpa após o teste
    cfg := LoadConfig()
    if cfg.ServerPort != ":9090" {
        t.Errorf("Esperado :9090, recebido %s", cfg.ServerPort)
    }
}
