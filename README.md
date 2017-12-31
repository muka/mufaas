# μFaaS

golang micro FaaS framework

[![Build Status](https://travis-ci.org/muka/mufaas.svg?branch=master)](https://travis-ci.org/muka/mufaas) [![Coverage Status](https://coveralls.io/repos/github/muka/mufaas/badge.svg?branch=master)](https://coveralls.io/github/muka/mufaas?branch=master)

A minimal service to run container in a function as a service fashion. Easily embeddable and with a reasonably low resource consumption.

## Installation

Currently only `go get` is supported

`go get github.com/muka/mufaas`

## Usage

1.  Start the daemon
    `mufaas daemon`
2.  Point to a folder with a Dockerfile
    `mufaas add --name test1 ./somewhere`
3.  Once built you can run the function
    `mufaas run test1 arg1`
    The API has ongoing support for `stdin` as a io.Reader.
4.  Drop the function
    `mufaas remve test1`

Flags:

- `-v` enables debug logging
- `--url` set the daemon endpoint


## Exec mode benchmark

1. Create & run function container: 0.60s
2. Run function container: 0.45s
3. Exec on running container: **0.084s**

## To Do

- [ ] Move todo list to issue tracker
- [ ] Add test coverage to commands
- [ ] Add support for language specific deployment (from templates base image)
- [ ] Add stream support for stdin / stdout / stderr
- [ ] Add pipe-able command support
- [ ] Add docker release (`amd64`, `arm`)
- [ ] Add support for local registry
- [ ] Add support for pluggable authentication and authorization (`oauth2`, `jwt`)

## License

MIT License

Copyright (c) 2017 Luca Capra
