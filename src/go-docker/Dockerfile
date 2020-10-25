FROM golang:1.14.4
WORKDIR /go/src/
COPY websrv-demo.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
WORKDIR /root/goapp/
COPY --from=0 /go/src/app .
RUN chown -R appuser:appgroup /root/goapp
RUN chmod 777 -R /root/goapp
USER appuser
CMD ["./app"]