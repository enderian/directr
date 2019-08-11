.build:
	go build -o bin/directrd

start: .build
	bin/directrd server

start-agent: .build
	bin/directrd agent start --deamon=false

test: .build

generate-proto:
	protoc -I=./proto --go_out=./types proto/*.proto