FROM golang:1.14 as build
WORKDIR /app
COPY . .
RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o socks5

FROM scratch
COPY --from=build /app/socks5 /socks5
EXPOSE 1080
ENTRYPOINT ["/socks5"]
