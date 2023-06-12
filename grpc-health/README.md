## README

### Create proto classes
protoc --go_out=./generated --go_opt=paths=source_relative   --go-grpc_out=./generated --go-grpc_opt=paths=source_relative proto/hello-service.proto

### Using google GRPC health
https://pkg.go.dev/google.golang.org/grpc@v1.55.0/health/grpc_health_v1