FROM golang:1.19 AS build-dist
ENV GOPROXY='https://mirrors.aliyun.com/goproxy'
WORKDIR /data/release
COPY . .
RUN go build

FROM centos:latest as prod
WORKDIR /data/db-go-gin
COPY --from=build-dist /data/release/cloud-svr ./
COPY --from=build-dist /data/release/conf /data/cloud-svr/conf

EXPOSE 10010

CMD ["/data/cloud-svr/db-go-gin"]