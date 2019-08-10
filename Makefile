.build:
	go build -o bin/directrd

start: .build
	bin/directrd server