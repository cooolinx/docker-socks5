# Docker Socks5

A docker image with socks5 proxy server inside.

Image already on docker hub [cooolin/socks5](https://hub.docker.com/r/cooolin/socks5).

## Features

By using https://github.com/haxii/socks5, we support:

- "No Auth" mode
- User/Password authentication
- Support for the CONNECT command
- Support for the ASSOCIATE command
- Rules to do granular filtering of commands
- Custom DNS resolution
- Unit tests

NOT yet support:

- BIND command

## Usage

1-line command:

```sh
docker run --rm --name socks5 -p 1080:1080 -e PROXY_HOST=0.0.0.0 -e PROXY_PORT=1080 -e PROXY_USER=cooolin -e PROXY_PASS=password cooolin/socks5
```

with `docker-compose.yml`:

```yaml
version: '3.4'

services:
  squid:
    image: cooolin/socks5
    container_name: socks5
    ports:
      - "1080:1080"
    environment:
      - PROXY_HOST=0.0.0.0
      - PROXY_PORT=1080
      - PROXY_USER=cooolin
      - PROXY_PASS=password
    restart: always
```

then up it

```sh
docker-compose up -d
```
