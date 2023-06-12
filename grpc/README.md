##Protobuff Golang repository

### Tutorial reference
https://www.practical-go-lessons.com/post/how-to-create-a-grpc-server-with-golang-ccdm795s4r5c70i1kacg

### How to execute

*Install protobuf*
For Linux :
`apt install -y protobuf-compiler`

For MacOs:
`brew install protobuf`

Check your installation with:
`protoc --version`

*Install protoc-gen-go-grpc*
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

*create directory for package*
`protos/test`

*generate protobuff class*

```
protoc --go_out=./proto --go_opt=paths=source_relative \
    --go-grpc_out=./proto --go-grpc_opt=paths=source_relative \
    invoicer.proto
```

```
protoc --go_out=./proto --go_opt=paths=source_relative \
    --go-grpc_out=./proto --go-grpc_opt=paths=source_relative \
    health.proto
```
