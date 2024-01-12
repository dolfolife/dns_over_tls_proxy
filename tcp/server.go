package tcp_server

import(
    "net"
    "log"

    tlsResolver "github.com/dolfolife/dns_over_tls_proxy/tlsresolver"
    "github.com/miekg/dns"
)

type TCPServer struct {}

func handleDNSRequest(conn net.Conn) {
    defer conn.Close()

    dnsConn := dns.Conn{Conn: conn}
    dnsMsg, err := dnsConn.ReadMsg()
    if err != nil {
        log.Printf("Failed to read the DNS query: %v", err)
        return
    }

    response, err := tlsResolver.ResolveOverTLS(dnsMsg)
    if err != nil {
        log.Println("Error resolving the query:", err)
        return
    }

    if err = dnsConn.WriteMsg(response); err != nil {
        log.Fatalf("Error forwarding the response into the client conn with: \n %v \n", err)
    }
}

func (s *TCPServer) ListenAndServe(address string) {
    listener, err := net.Listen("tcp", address)
    if err != nil {
        log.Fatalf("Failed to set up TCP listener on %s: %v", address, err)
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println("Failed to accept connection:", err)
            continue
        }
        go handleDNSRequest(conn)
    }
}

func NewTCPServer() TCPServer {
    return TCPServer{}
}
