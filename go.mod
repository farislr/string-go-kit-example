module github.com/string-service

go 1.14

require (
	github.com/farislr/core-proto v0.0.0-00010101000000-000000000000
	github.com/go-kit/kit v0.10.0
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/gorilla/mux v1.7.4
	github.com/prometheus/client_golang v1.3.0 // indirect
	google.golang.org/grpc v1.29.1
)

replace github.com/farislr/core-proto => ../acite/core-proto
