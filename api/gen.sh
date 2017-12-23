#!/usr/bin/env bash

PROTOSRC="./api.proto"

PKGDIR=${GOPATH}/src/github.com/muka/mufaas

protoc_bin=${PKGDIR}/tmp/protoc/bin/protoc
protoc_include=${PKGDIR}/tmp/protoc/include
googleapis=${PKGDIR}/tmp/googleapis

# generate the gRPC code
${protoc_bin} \
    -I. \
    -I${protoc_include} \
    -I${googleapis} \
    --go_out=plugins=grpc:. \
    $PROTOSRC

# generate the JSON interface code
${protoc_bin} \
    -I. \
    -I${protoc_include} \
    -I$GOPATH/src \
    -I${googleapis} \
    --go_out=plugins=grpc:. \
    $PROTOSRC

# generate the reverse proxy
${protoc_bin} \
    -I. \
    -I${protoc_include} \
    -I$GOPATH/src \
    -I${googleapis} \
    --grpc-gateway_out=logtostderr=true:. \
    $PROTOSRC

# generate the swagger definitions
${protoc_bin} \
    -I. \
    -I${protoc_include} \
    -I$GOPATH/src \
    -I${googleapis} \
    --swagger_out=logtostderr=true:../swagger \
    $PROTOSRC
