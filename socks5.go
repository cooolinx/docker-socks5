package main

import (
    "os"
    "net"
    "fmt"
    "strconv"

    "github.com/haxii/socks5"
)

func getenv(key, def string) string {
    val := os.Getenv(key)
    if len(val) == 0 {
        return def
    }
    return val
}

func main() {
    host := getenv("PROXY_HOST", "0.0.0.0")
    port := getenv("PROXY_PORT", "1081")
    user := getenv("PROXY_USER", "")
    pass := getenv("PROXY_PASS", "")

    authRequired := len(user) != 0 && len(pass) != 0

    iport, err := strconv.Atoi(port)
    if err != nil {
        panic(err)
    }

    conf := &socks5.Config{
        BindIP:   net.ParseIP(host),
        BindPort: iport,
    }
    if authRequired {
        conf.Credentials = socks5.StaticCredentials{
            user: pass,
        }
    }
    server, err := socks5.New(conf)
    if err != nil {
        panic(err)
    }

    // Create SOCKS5 proxy
    if authRequired {
        fmt.Printf("Listening on %v:%v with auth required (%v)...\n", host, port, user)
    } else {
        fmt.Printf("Listening on %v:%v...\n", host, port)
    }
    
    if err := server.ListenAndServe("tcp", host + ":" + port); err != nil {
        panic(err)
    }
}
