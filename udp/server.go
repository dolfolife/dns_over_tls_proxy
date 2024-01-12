package udp

import(
    "net"
    "log"

    tlsResolver "github.com/dolfolife/dns_over_tls_proxy/tlsresolver"
    "github.com/miekg/dns"
)

type UDPServer struct {}

func handleDNSRequest(udpConn *net.UDPConn, clientAddr *net.UDPAddr, query []byte) {
    dnsMsg := &dns.Msg{}
    if err := dnsMsg.Unpack(query); err != nil {
        log.Printf("Failed to unpack DNS query: %v", err)
        return
    }

    response, err := tlsResolver.ResolveOverTLS(dnsMsg)
    
    responseBytes, err := response.Pack()
    if err != nil {
        log.Printf("Failed to pack DNS response: %v", err)
        return 
    }

    _, err = udpConn.WriteToUDP(responseBytes, clientAddr)
    if err != nil {
        log.Printf("Failed to write DNS response to UDP: %v", err)
        return
    }
}

func (s *UDPServer) ListenAndServe(address string) {
    udpAddress, _ := net.ResolveUDPAddr("udp", address)
    conn, err := net.ListenUDP("udp", udpAddress)
    if err != nil {
        log.Fatalf("Failed to set up UDP listener on %s: %v", address, err)
    }
    defer conn.Close()

    for {
        buffer := make([]byte, 512) // Standard DNS packet size, see: https://www.rfc-editor.org/rfc/rfc5966
        n, addr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            log.Printf("Error reading from UDP: %s", err)
            continue
        }

        go handleDNSRequest(conn, addr, buffer[:n])
    }
}

func NewUDPServer() UDPServer {
    return UDPServer{}
}
