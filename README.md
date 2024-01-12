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

# Q&A

### Imagine this proxy being deployed in an infrastructure. What would be the security concerns you would raise?
- There are no Rate Limiting or access control so it can be used for DNS amplification attacks
- If there is a custom DNS server, ensure that the TLS connection to the DNS-over-TLS server is correctly configured.
### How would you integrate that solution in a distributed, microservices-oriented and containerized architecture?
*Kubernetes Service:* Define a Service in Kubernetes to expose the DNS proxy. Depending on your network policies, this could be an internal service or exposed externally. As we already have a dockerfile we can maintain lightweight and secure as the image of this service pods using a Kubernetes Deployment. 

With Kubernetes we can scale, load balance and monitor the proxy without adding on top of the complexity of the container.

Outside Kubernetes, it can be deployed as container in your private network (e.g. VPC in AWS). 

### What other improvements do you think would be interesting to add to the project?
- Implement checks to validate DNS queries and responses to prevent DNS spoofing or cache poisoning attacks.
- Proper error handling in the application is necessary to prevent leakage of sensitive information through error messages.
