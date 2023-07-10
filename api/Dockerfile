FROM golang:1.18-alpine as builder

WORKDIR /app

COPY .. ./

RUN go mod download

RUN apk add gcc g++ make
RUN apk add git
RUN apk add curl
RUN apk add nodejs npm

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM alpine:latest

RUN apk add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 9999

CMD ["./main"]