FROM golang:1.14.7

RUN GO111MODULE=on go get github.com/golang/mock/mockgen@v1.4.3

WORKDIR /src
ENTRYPOINT ["go", "generate", "-v", "./..."]