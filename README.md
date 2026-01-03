# Go HTTP Server

Minimal HTTP server written in Go.

This is a **development and learning project**.  
Its purpose is to explore HTTP, networking, and basic service design in Go.


## Description

The server provides:
- a simple HTTP endpoint
- HTML rendering using Go templates
- system information display (hostname)
- client IP address display

The project uses only the Go standard library.


## Requirements

- Go (>= 1.20)
- Linux operating system


## Directory Structure
```
.
├── server.go
└── static/
└── home.html
```

# Build
```
go build server.go
```

## Run

By default, the server listens on port `8090`.
```
./server
```

## Testing
```
curl http://localhost:8090
```

In a web browser:
http://localhost:8090

## Notes

- Ports < 1024 require administrative privileges.
- When accessing via `localhost`, the client IP address may appear as `::1` (IPv6 loopback).

## Development Direction

- HTTP request logging
- `/health` endpoint
- reverse proxy support (X-Forwarded-For)
- systemd service integration
- Nginx reverse proxy configuration
- basic security mechanisms
