package main

import (
    tcpServer "github.com/dolfolife/dns_over_tls_proxy/tcp"
    udpServer "github.com/dolfolife/dns_over_tls_proxy/udp"
)

const (
    localAddress = ":53"         // Local endpoint for DNS queries over TCP and UDP
)

func main() {
    tcpListener := tcpServer.NewTCPServer()
    udpListener := udpServer.NewUDPServer()

    go tcpListener.ListenAndServe(localAddress)
    go udpListener.ListenAndServe(localAddress)

    // block main goroutine
    select {}
}

