package config

import "os"

type Config struct {
    ServerPort string
}

func LoadConfig() Config {
    port := os.Getenv("PORT")
    if port == "" {
        port = ":8080"
    }
    return Config{
        ServerPort: port,
    }
}
