.build:
	go build -o bin/directrd

start: .build
	bin/directrd server

start-agent: .build
	bin/directrd agent

test: .build