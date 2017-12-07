# build container
FROM golang:1.9 AS build

ENV CGO_ENABLED=0

RUN go get -u github.com/golang/dep/... &&\
    mkdir -p /go/src/github.com/nasa9084/salamander/salamander

WORKDIR /go/src/github.com/nasa9084/salamander/salamander

COPY salamander/ ./

WORKDIR /go/src/github.com/nasa9084/salamander/salamander
RUN dep ensure && \
    go build -o /tmp/salamander cmd/salamander/main.go && \
    chmod +x /tmp/salamander

# application container
FROM scratch
LABEL maintainer="nasa9084"
COPY --from=build /tmp/salamander salamander
CMD ["./salamander"]
