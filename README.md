#Proyect Buffer

# Download protobuffer

## Install "protoc-24.4-win64.zip" from:

```
https://github.com/protocolbuffers/protobuf/releases/tag/v24.4
```

NOTE: Create a PATH for example "C:\protoc-24.4-win64\bin"

### Go plugins for the protocol compiler

$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

### Start Buffer

```
    make...
```