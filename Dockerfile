FROM golang:1.13-alpine as builder
WORKDIR /root
COPY ./  ./
RUN export GO111MODULE=on && CGO_ENABLED=0 GOOS=linux go build -o ./main ./main.go
COPY conf/  conf/

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /root/main ./
COPY --from=builder /root/conf conf/
ENTRYPOINT ["/root/main"]