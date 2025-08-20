FROM golang:1.23-bullseye

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o analyzer-api main.go

ENV ENV JWT_SECRET=${JWT_SECRET}

EXPOSE 8080

CMD ["./analyzer-api"]