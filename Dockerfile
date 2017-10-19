# build container
FROM golang:1.9 AS build

ENV GHQ_ROOT=/go/src
ENV CGO_ENABLED=0

RUN go get github.com/motemen/ghq && \
    ghq get nasa9084/salamander

WORKDIR /go/src/github.com/nasa9084
RUN curl https://glide.sh/get | sh

WORKDIR /go/src/github.com/nasa9084/salamander/salamander
RUN glide install && \
    go build -o /tmp/salamander cmd/salamander/main.go && \
    chmod +x /tmp/salamander

# application container
FROM alpine:latest
LABEL maintainer="nasa9084"
COPY --from=build /tmp/salamander salamander
CMD ["salamander"]
