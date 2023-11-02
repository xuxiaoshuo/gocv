# to build this docker image:
#   docker build .
FROM ghcr.io/hybridgroup/opencv:4.8.1

ENV GOPATH /go

COPY . /go/src/github.com/xuxiaoshuo/gocv/

WORKDIR /go/src/github.com/xuxiaoshuo/gocv
RUN go build -tags example -o /build/gocv_version -i ./cmd/version/

CMD ["/build/gocv_version"]
