```
export GO111MODULE="on"
protoc --go_out=plugins=grpc:login login.proto
```