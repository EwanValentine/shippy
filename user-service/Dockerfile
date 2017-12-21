FROM golang:1.9.0 as builder

RUN go get -u github.com/EwanValentine/shippy/vessel-service

WORKDIR /go/src/github.com/EwanValentine/shippy/vessel-service

RUN go get && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/EwanValentine/shippy/vessel-service/vessel-service .

CMD ["./vessel-service"]
