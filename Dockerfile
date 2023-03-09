FROM golang:alpine3.17

RUN set -ex \
    && apk add --no-cache --virtual .build-deps \
        curl \
        gcc \
        g++ \
        make \
        libtool \
        autoconf \
        automake \
        git \
    && mkdir -p /src \
    && mkdir -p /data \
    && cd /src \
    && git clone https://github.com/openvenues/libpostal.git \
    && cd libpostal \
    && ./bootstrap.sh \
    && ./configure --datadir=/data MODEL=senzing \
    && make -j "$(nproc)" \
    && make install \
    && apk del .build-deps \
    && rm -rf /src

RUN apk add --no-cache gcc musl-dev pkgconfig

WORKDIR /app

ENV GO111MODULE=on
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /usr/bin/address-parser main.go

ENTRYPOINT ["/usr/bin/address-parser"]
