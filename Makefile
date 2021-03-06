.PHONY: build deps generate clean test coverage docker/push docker/build bindata build-idle

build_dir=./build
bin_dir=./bin

dockerName := opny/mufaas

gittag := $(shell git describe --tag --always)
tag := $(shell echo ${gittag} | cut -d'-' -f 1)
basetag := $(shell echo ${gittag} | cut -d'.' -f 1)

all: deps test coverage

deps:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/mattn/goveralls
	cd ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway && go get ./...
	rm -rf ./tmp/googleapis
	mkdir -p ./tmp
	git clone --depth 1 https://github.com/googleapis/googleapis.git tmp/googleapis
	./bin/install_protoc.sh
	@dep ensure

generate: bindata
	go generate ./api/api.go

bindata: build-idle
	go-bindata -pkg asset -o asset/bindata.go -ignore=idle.go ./idle/bin

build:
	mkdir -p ${build_dir}
	CGO_ENABLED=0 go build -o ${build_dir}/mufaas main.go

build-idle:
	mkdir -p ./idle/bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./idle/bin/idle-amd64 idle/idle.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ./idle/bin/idle-arm idle/idle.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./idle/bin/idle-arm64 idle/idle.go

clean:
	rm -rf ${build_dir}

test:
	@go test ./... -v

coverage:
	echo "mode: count" > coverage.out

	go test -covermode="count" -coverprofile="coverage.tmp" ./service
	cat coverage.tmp | grep -v "mode: count" >> coverage.out

	go test -covermode="count" -coverprofile="coverage.tmp" ./docker
	cat coverage.tmp | grep -v "mode: count" >> coverage.out

	go test -covermode="count" -coverprofile="coverage.tmp" ./util
	cat coverage.tmp | grep -v "mode: count" >> coverage.out

	rm coverage.tmp
	goveralls -service=travis-ci -coverprofile=./coverage.out

docker/build: build
	echo "Building ${tag}"
	docker build . -t ${dockerName}:${tag}
	docker tag ${dockerName}:${tag} ${dockerName}:${basetag}

docker/push: docker/build
	docker push ${dockerName}:${tag}
	docker push ${dockerName}:${basetag}
