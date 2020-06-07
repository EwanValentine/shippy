module github.com/EwanValentine/shippy/shippy-service-vessel

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/EwanValentine/shippy-service-vessel v0.0.0-20200113232044-568e51dd7413
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.8.0
	go.mongodb.org/mongo-driver v1.3.4
)
