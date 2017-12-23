#!/usr/bin/env bash

dest_dir="./tmp/protoc"

has() {
  type "$1" > /dev/null 2>&1
}

install_protoc() {
    local url=$1
    mkdir -p ${dest_dir}
    wget ${url} -O ${dest_dir}/protoc.zip
    cd ${dest_dir} && unzip ./protoc.zip && chmod +x ./bin/protoc
}

if [ ! -e ${dest_dir} ]
then
    if [ "$(uname)" == "Darwin" ]; then
        url="https://github.com/google/protobuf/releases/download/v3.5.0/protoc-3.5.0-osx-x86_64.zip"
        install_protoc ${url}
    elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
        url="https://github.com/google/protobuf/releases/download/v3.5.0/protoc-3.5.0-linux-x86_64.zip"
        install_protoc ${url}
    else
        echo 'Windows is not supported yet'
    fi
fi
