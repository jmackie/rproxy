# rproxy

A tiny command-line wrapper around Go's [`ReverseProxy`](https://golang.org/pkg/net/http/httputil/#ReverseProxy).
Useful for device testing and stuffs.

```
Usage of rproxy:
  -local-host string
    	local host to reverse proxy (default "http://localhost")
  -local-port string
    	local port to reverse proxy (default "4009")
  -public-host string
    	address to serve on (default "192.168.2.122")
  -public-port string
    	port to serve on (default "8080")
```
