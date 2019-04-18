// go run main.go -help
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Get _preferred_ outbound ip of this machine
func getOutboundIP() (*net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return &localAddr.IP, nil
}

func main() {
	ip := ""
	if outboundIP, err := getOutboundIP(); err == nil {
		ip = fmt.Sprintf("%s", outboundIP)
	}

	localHost := flag.String("local-host", "http://localhost", "local host to reverse proxy")
	localPort := flag.String("local-port", "4009", "local port to reverse proxy")
	publicHost := flag.String("public-host", ip, "address to serve on")
	publicPort := flag.String("public-port", "8080", "port to serve on")
	flag.Parse()

	local, err := url.Parse(*localHost + ":" + *localPort)
	if err != nil {
		log.Fatalf("error parsing local host/port: %v", err)
	}
	if *publicHost == "" {
		log.Fatal("missing -public-host flag", err)
	}

	public := *publicHost + ":" + *publicPort
	http.Handle("/", httputil.NewSingleHostReverseProxy(local))
	fmt.Printf("Listening...\n%s -> %s\n", local, public)
	log.Fatal(http.ListenAndServe(public, nil))
}
