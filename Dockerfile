FROM golang:1.12.4-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache \
    git \
    autoconf \
    automake \
    build-base \
    libtool \
    nasm \
    pngquant

RUN git clone git://github.com/mozilla/mozjpeg.git && \
    cd mozjpeg && \
    git checkout v3.1 && \
    autoreconf -fiv && ./configure --prefix=/opt/mozjpeg && make install

WORKDIR capture-go

COPY . .

#RUN install node deps from puppeteer

RUN go build -v -o capture-go cmd/capture-go-web/main.go

CMD ["./capture-go"]
