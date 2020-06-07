module github.com/EwanValentine/shippy/shippy-service-consignment

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

// replace github.com/EwanValentine/shippy/shippy-service-vessel => ../shippy-service-vessel

require (
	github.com/EwanValentine/shippy-service-consignment v0.0.0-20200113004730-e48fe0dbef52
	github.com/EwanValentine/shippy-service-vessel v0.0.0-20200113232044-568e51dd7413
	github.com/EwanValentine/shippy/shippy-service-vessel v0.0.0-20200607004308-a5699169cc6d
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.8.1-0.20200603084508-7b379bf1f16e
	github.com/micro/micro/v2 v2.8.1-0.20200603100651-e57d42a20d26
	github.com/pkg/errors v0.9.1
	go.mongodb.org/mongo-driver v1.2.1
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2
	google.golang.org/grpc v1.29.1
)
