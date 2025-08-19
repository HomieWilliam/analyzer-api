FROM golang:1.25-bullseye

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o analyzer-api main.go

EXPOSE 8080

CMD ["./analyzer-api"]