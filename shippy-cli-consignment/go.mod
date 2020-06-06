module github.com/EwanValentine/shippy/shippy-cli-consignment

go 1.14

replace github.com/EwanValentine/shippy/shippy-service-consignment => ../shippy-service-consignment

require (
	github.com/EwanValentine/shippy/shippy-service-consignment v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.29.1
)
