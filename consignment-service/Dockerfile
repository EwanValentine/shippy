FROM golang:1.9.0 as builder

RUN go get -u github.com/EwanValentine/shippy/consignment-service

WORKDIR /go/src/github.com/EwanValentine/shippy/consignment-service

RUN go get && CGO_ENABLED=0 GOOS=linux go build  -o consignment-service -a -installsuffix cgo main.go repository.go handler.go datastore.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/EwanValentine/shippy/consignment-service/consignment-service .

CMD ["./consignment-service"]
