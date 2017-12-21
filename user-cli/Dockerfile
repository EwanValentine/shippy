FROM golang:1.9.0 as builder

RUN go get -u github.com/EwanValentine/shippy/user-cli

WORKDIR /go/src/github.com/EwanValentine/shippy/user-cli

RUN go get && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/EwanValentine/shippy/user-cli/user-cli .

ENTRYPOINT ["./user-cli"]

CMD ["./user-cli"]
