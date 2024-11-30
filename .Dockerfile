FROM golang:stable

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main .

CMD ["/app/cmd/main"]


