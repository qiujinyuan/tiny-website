FROM golang:1.14-alpine

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/yrjkqq/tiny-website
COPY . $GOPATH/src/github.com/yrjkqq/tiny-website
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./tiny-website"]