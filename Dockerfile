ARG base_image=alpine:3.15
ARG build_image=golang:1.18-alpine3.15

FROM ${base_image} AS base_image

FROM ${build_image} AS build_image
ENV CGO_ENABLED=0
WORKDIR $GOPATH/src/github.com/frantjc/sequence
COPY go.mod go.sum ./
RUN go mod download

FROM build_image AS build
COPY . .
ARG version=0.0.0
ARG prerelease=
RUN go build -ldflags "-s -w" -o /usr/local/bin ./internal/cmd/sqnc-shim-source
RUN go build -ldflags "-s -w" -o /usr/local/bin ./internal/cmd/sqnc-shim-uses
RUN mv /usr/local/bin/sqnc-shim-source shim/
RUN mv /usr/local/bin/sqnc-shim-source shim/
RUN go build -ldflags "-s -w -X github.com/frantjc/sequence.Version=${version} -X github.com/frantjc/sequence.Prerelease=${prerelease}" -o /usr/local/bin ./cmd/sqnc
RUN go build -ldflags "-s -w -X github.com/frantjc/sequence.Version=${version} -X github.com/frantjc/sequence.Prerelease=${prerelease}" -o /usr/local/bin ./cmd/sqncd
RUN go build -ldflags "-s -w" -o /usr/local/bin ./internal/cmd/sqnc-runtime-docker

FROM base_image AS sequence
COPY --from=build /usr/local/bin /usr/local/bin
RUN ln -s /etc/sqnc/plugins/default-runtime /usr/local/bin/sqnc-runtime-docker
ENTRYPOINT ["sqncd"]

FROM sequence
