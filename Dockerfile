FROM golang:1.12

WORKDIR /app
ADD . /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-w -s" -installsuffix cgo -o /app/directd .

FROM alpine:latest

WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=0 /app/directd /bin/directd

CMD /bin/directd -config /app/config.yml