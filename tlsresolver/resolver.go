package tlsresolver

import (
    "crypto/tls"
    "log"

    "github.com/miekg/dns"
)

const (
    tlsDNSAddress        = "1.1.1.1:853" // Cloudflare's DNS-over-TLS endpoint
)

func ResolveOverTLS(msg *dns.Msg) (*dns.Msg, error) {
    tlsConn, err := tls.Dial("tcp", tlsDNSAddress, &tls.Config{})
    if err != nil {
        log.Printf("Failed to establish a TLS connection: %v", err)
        return nil, err
    }
    defer tlsConn.Close()

    dnsConn := &dns.Conn{Conn: tlsConn}
    if err := dnsConn.WriteMsg(msg); err != nil {
        log.Printf("Failed to send message over TLS: %v", err)
        return nil, err
    }

    response, err := dnsConn.ReadMsg()
    if err != nil {
        log.Printf("Failed to read message from TLS: %v", err)
        return nil, err
    }

    return response, nil
}
