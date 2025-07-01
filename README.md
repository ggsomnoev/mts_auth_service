# TLS Auth Service

## Description

This service exposes a secure `/auth` endpoint that performs mutual TLS (mTLS) authentication. The client must provide a valid certificate signed by a trusted Certificate Authority (CA), and the Common Name (CN) in the client certificate must match a list of preconfigured trusted values.

## Libraries & tools:

| Concern | Library |
|---------|---------|
| HTTP server | [Echo](https://echo.labstack.com/) |
| Testing | [Ginkgo](https://onsi.github.io/ginkgo/) + [Gomega](https://onsi.github.io/gomega/) |

## Prerequisites

You must generate the necessary certificates **before running** the application (whether locally or in Docker). Use the following `make` command to create:

```bash
make generate-all-certs
```

If you want to generate a client certificate with custom CN:

```bash
make client CN=my-custom-cn
```

## How to run the server

To run the server:
```bash
make run
```

To run the server using docker:
```bash
make run-docker
```

## How to run the tests

To run the unit tests:

```bash
make test
```

## Example Requests

### Valid certificates

```bash
make client

curl https://localhost:8443/auth \
  --cert ./certs/client/client.crt \
  --key ./certs/client/client.key \
  --cacert ./certs/ca/ca.crt
```

### Valid certificates and unknown CN

```bash
make client CN=my-custom-cn

curl https://localhost:8443/auth \
  --cert ./certs/client/client.crt \
  --key ./certs/client/client.key \
  --cacert ./certs/ca/ca.crt
```

### Certificates signed by untrusted CA

```bash
curl https://localhost:8443/auth \
  --cert ./certs/client-untrusted/client.crt \
  --key ./certs/client-untrusted/client.key \
  --cacert ./certs/ca-untrusted/ca.crt
```

### No certificate

```bash
curl https://localhost:8443/auth
```