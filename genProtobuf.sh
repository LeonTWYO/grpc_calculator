#!/bin/bash

serv=${PWD##*/} 

if [ -d "./protobuf/" ];
then
     protoc -I./protobuf -I$GOPATH/include \
        --go_out=./protobuf --go_opt=paths=source_relative \
        --go-grpc_out=./protobuf --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=./protobuf --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=logtostderr=true:./protobuf --openapiv2_opt use_go_templates=true --openapiv2_opt enums_as_ints=true --openapiv2_opt omit_enum_default_value=true \
        ./protobuf/*.proto
else
    echo "protobuf directory was not found."
fi