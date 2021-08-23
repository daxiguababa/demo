FROM golang:1.13-alpine as builder
WORKDIR /root
COPY ./  ./
RUN export GO111MODULE=on && CGO_ENABLED=0 GOOS=linux go build -o ./main ./main.go


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /root/main ./

ENTRYPOINT ["/root/main"]
#
#FROM golang:1.13-alpine as builder
#WORKDIR /root
#COPY ./  ./
#
#RUN export GO111MODULE=on && CGO_ENABLED=0 GOOS=linux go build -o ./main ./main.go
#
#FROM node:12-alpine as node-builder
#WORKDIR /root
#COPY dashboard .
#RUN npm rebuild node-sass && npm install && npm run build
#
#
#FROM alpine:latest
#RUN apk --no-cache add ca-certificates
#WORKDIR /root
#COPY --from=builder /root/build/* ./
#COPY --from=node-builder /root/build dist
#COPY deployments/docker/all-in-one/start.sh .
#ENV DATABASE_URL="root:12345@(127.0.0.1:3306)/xconf?charset=utf8&parseTime=true&loc=Local"
#
#EXPOSE 8080
#ENTRYPOINT ["/root/start.sh"]
