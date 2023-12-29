# Gunakan image resmi Golang sebagai base image
FROM golang:latest

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /go/src/app/main .

EXPOSE 8080

CMD ["./main"]
