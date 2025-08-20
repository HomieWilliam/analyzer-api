# Analyzer API

**Analyzer API** é uma aplicação web desenvolvida em Go que permite analisar frases, retornando a contagem de palavras, vogais e consoantes. O projeto é containerizado com Docker e pode ser facilmente implantado em instâncias AWS EC2 via Terraform e Ansible.

## Funcionalidades

- Receber uma frase via API REST e retornar:
  - Número de palavras
  - Número de vogais
  - Número de consoantes
- Autenticação via JWT
- Deploy automatizado em EC2
- Container Docker para fácil distribuição

## Tecnologias

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin) – Framework web
- [Docker](https://www.docker.com/)
- [AWS EC2](https://aws.amazon.com/ec2/)
- [Terraform](https://www.terraform.io/)
- [Ansible](https://www.ansible.com/)
- JWT para autenticação

## Como Rodar Localmente

1. Clone o repositório:

git clone https://github.com/HomieWilliam/analyzer-api.git
cd analyzer-api
go mod tidy
go run main.go


curl -X POST "http://localhost:8080/api/v1/login" -H "Content-Type: application/json" -d "{\"username\":\"admin\",\"password\":\"admin123\"}"

curl -X POST "http://localhost:8080/api/v1/analyze" -H "Content-Type: application/json" -H "Authorization: Bearer <JWT_TOKEN>" -d "{\"phrase\":\"Hello World\"}"

## Para Criar o docker 
docker build -t analyzer-api:latest .
docker run -d -p 8080:8080 analyzer-api:latest

## Para executar os testes 
go test ./...