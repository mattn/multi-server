# multi-server

multi-server is a simple Go-based virtual hosting static file server. It serves static files from different directories based on the HTTP Host header, allowing you to host multiple sites on a single port. Perfect for development, testing, or lightweight production setups.

## Features

- Virtual hosting: Serve files from site-specific directories based on the `Host` header.
- Automatic directory detection: Maps hostnames to subdirectories under a base sites directory.
- CORS support: Adds `Access-Control-Allow-Origin: *` header.
- Lightweight and fast: Built with the Go standard library.

## Installation

Install the binary using Go:

```bash
go install github.com/mattn/multi-server@latest
```

Requires Go 1.25 or later.

## Usage

Run the server with default settings (listen on `:8080`, sites in `/data`):

```bash
./multi-server
```

Access sites via different hostnames, e.g., `http://example.com:8080/` serves from `/data/example.com/`.

### Command-Line Flags

| Flag | Description | Default |
|------|-------------|---------|
| `-addr` | Address to listen on (host:port) | `:8080` |
| `-sites-dir` | Base directory containing site folders (e.g., `/data/example.com/`) | `/data` |

### Setup

1. Create site directories under the sites directory, e.g.:
   ```
   /data/
   ├── example.com/
   │   └── index.html
   └── test.local/
       └── index.html
   ```

2. Update your `/etc/hosts` for local testing:
   ```
   127.0.0.1 example.com test.local
   ```

3. Run the server:
   ```bash
   multi-server -sites-dir /data -addr :8080
   ```

4. Access: `http://example.com:8080/` or `http://test.local:8080/`

If a site directory doesn't exist, it returns a 404 "Site not found".

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
