FROM golang:1.7.3 as builder

WORKDIR /go/src/github.com/EwanValentine/shippy/consignment-service

COPY . .

RUN go get github.com/EwanValentine/shippy/vessel-service/proto/vessel
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/EwanValentine/shippy/consignment-service/consignment-service .
ADD consignment-service /app/consignment-service

CMD ["./consignment-service"]
