#!/bin/bash

BASEDIR=$PWD
PROTO_DEST=$BASEDIR/infrastructure/proto
TEMP_REPO=$BASEDIR/temp-proto

GO_MOD=github.com/rennanbadaro/discount-calculator
GO_DISCOUNT_PROTO_PACKAGE="github.com\/rennanbadaro\/discount-calculator\/infrastructure\/proto"

mkdir -p $TEMP_REPO $PROTO_DEST

cd $TEMP_REPO
git clone git@github.com:rennanbadaro/proto-graal.git --quiet

find ./ -name '*.proto' -exec cp {} $PROTO_DEST \;

# Set go package
cd $PROTO_DEST
sed -i "s/%GO_DISCOUNT_PROTO_PACKAGE%/${GO_DISCOUNT_PROTO_PACKAGE}/g" ./*.proto

rm -rf $TEMP_REPO
cd $BASEDIR

# Code generation
protoc --go_out=./infrastructure/proto \
    --go_opt=paths=source_relative \
    --go-grpc_out=./infrastructure/proto \
    --go-grpc_opt=paths=source_relative \
    -I $PROTO_DEST \
    $PROTO_DEST/*proto
