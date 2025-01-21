FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o riskservice .

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/riskservice .

EXPOSE 8080

CMD ["./riskservice"]
