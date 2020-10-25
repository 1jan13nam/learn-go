FROM golang:1.8-alpine
ADD . /go/src/hello-app
RUN go install hello-app

FROM alpine:latest
COPY --from=0 /go/bin/hello-app .
ENV PORT 8080
ARG CONFIG_VERSION=0.0.1
ENV CONFIG_VER=$CONFIG_VERSION
CMD ["./hello-app"]
#docker build -t test-go:1.2.4 --build-arg CONFIG_VERSION=1.2.4  -f Dockerfile .
# kubectl set image deploy hello-app hello-app=test-go:1.2.3
