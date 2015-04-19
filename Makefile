.PHONY: build fmt lint run test vendor.clean vendor.get vet

GOPATH := ${PWD}:${PWD}/_vendor:${GOPATH}
export GOPATH

default: build

build: vet
	go build -v -o ./bin/app ./src/main.go

fmt:
	go fmt ./src/...

lint:
	golint ./src

run: build
	./bin/app

test:
	go test -v ./src/...

vendor.clean:
	rm -dRf ./_vendor/src

vendor.get: vendor.clean
	GOPATH=${PWD}/_vendor cat ${PWD}/Dependencies | xargs go get -d -u -v
	if [ -d ./_vendor/src ]; then find ./_vendor/src -type d -name .git -exec echo rm -rf {} \; ; fi

vet:
	go vet ./src/payroll...


