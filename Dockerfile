FROM golang:1.23.3

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o ./cmd/main ./cmd

CMD ["/app/cmd/main"]
