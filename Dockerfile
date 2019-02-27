FROM golang

WORKDIR /app
ADD . /app
RUN go build github.com/enderian.directrd -o directrd

CMD directrd