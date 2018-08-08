# Build stage
# ========================================================
FROM golang:1.9.2-alpine3.6 AS builder

ENV GOPATH="$HOME/go"
ENV GOBIN="$GOPATH/bin"
ENV PATH="$PATH:$GOROOT/bin:$GOBIN"

COPY . /go/src/github.com/smilga/capture-go/

WORKDIR /go/src/github.com/smilga/capture-go/

RUN apk update && apk upgrade && \
    apk add --no-cache git && \
	go get -u github.com/golang/dep/cmd/dep  && \
	dep ensure -vendor-only


RUN go build -v -o capture-go cmd/capture-go-web/main.go

# Run stage
# ========================================================
FROM evpavel/slimerjs-alpine:57

RUN apk --update add --no-cache \
    git \
    autoconf \
    automake \
    build-base \
    libtool \
    nasm \
    pngquant

WORKDIR /root

RUN git clone git://github.com/mozilla/mozjpeg.git && \
    cd mozjpeg && \
    git checkout v3.1 && \
    autoreconf -fiv && ./configure --prefix=/opt/mozjpeg && make install

WORKDIR /root/capture-go

# copy compiled project
COPY --from=builder /go/src/github.com/smilga/capture-go/capture-go ./

# copy slimer scripts
COPY --from=builder /go/src/github.com/smilga/capture-go/slimer-script/ ./slimer-script/

CMD ["./capture-go"]
