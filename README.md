# DNS over TLS Proxy

> Do not use this in production. This is for learning purposes

# Features
- *Support for Both TCP and UDP*: The proxy listens for DNS queries on both TCP and UDP on port 53.
- *DNS-over-TLS*: Queries are forwarded securely using DNS-over-TLS to Cloudflare's 1.1.1.1 service.
- *Concurrent Handling*: Multiple incoming DNS requests are handled concurrently.
- *Logging*: Basic logging of errors and important events for troubleshooting.

# Prerequisites
- Go (version 1.15 or later)
- Docker (if running using Docker)

# Usage

## Running Locally

1. build the proxy
```bash
go build
```

2. Run the Proxy
```bash
sudo ./dns_over_tls_proxy
```
Note: Root privileges are required to listen on port 53.

3. Test your proxy
```bash
dig @127.0.0.1 +tcp example.com.
```

```bash
dig @127.0.0.1 example.com.
```

## Docker
1. Build the Docker image:
```bash
docker build -t dns-over-tls-proxy .
```
2. Run the container
```bash
docker run -p 53:53/udp -p 53:53/tcp --name dns-proxy-instance dns-over-tls-proxy
```
This command maps port 53 on both TCP and UDP from the container to port 53 on the host.

# TODO

Currently, the proxy is pre-configured to use Cloudflare's 1.1.1.1 DNS-over-TLS service. You can modify the source code to change the DNS-over-TLS server or add additional configuration options. In the future this should be pass as a environment variable or some type of configuration for the DNS server that runs on TLS.

