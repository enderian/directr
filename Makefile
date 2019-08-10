.build:
	go build -o bin/directrd

test: .build

start: .build
	bin/directrd server