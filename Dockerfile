FROM golang

WORKDIR /app
ADD . /app
RUN go build ender.gr/directrd -o directrd

CMD directrd