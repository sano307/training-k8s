FROM golang:1.12 as builder

WORKDIR /go/src/demo-telepresence/

COPY /boot.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o boot .

FROM alpine:3.9.2

WORKDIR /root/

RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/demo-telepresence/boot .

ENTRYPOINT ["./boot"]

EXPOSE 8080
