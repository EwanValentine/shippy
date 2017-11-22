build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/ewanvalentine/shippy/consignment-service \
		proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .

run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-service
