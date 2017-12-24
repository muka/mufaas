.PHONY: deps generate build clean docker/push docker/build test

build=./build
dist=./dist

dockerName := opny/mufaas

gittag := $(shell git describe --tag --always)
tag := $(shell echo ${gittag} | cut -d'-' -f 1)
basetag := $(shell echo ${gittag} | cut -d'.' -f 1)

deps:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	cd ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway && go get ./...
	rm -rf ./tmp/googleapis
	mkdir -p ./tmp
	git clone --depth 1 https://github.com/googleapis/googleapis.git tmp/googleapis
	./bin/install_protoc.sh
	@dep ensure

generate:
	go generate ./api/api.go

build: generate
	CGO_ENABLED=0 go build -o ${OUTPUT_DIR}/mufaas mufaas.go

clean:
	rm -rf ${OUTPUT_DIR}
	rm -rf ${DIST_DIR}

test:
	@go test ./...

docker/build: build
	echo "Building ${tag}"
	docker build . -t ${dockerName}:${tag}
	docker tag ${dockerName}:${tag} ${dockerName}:${basetag}

docker/push: docker/build
	docker push ${dockerName}:${tag}
	docker push ${dockerName}:${basetag}
