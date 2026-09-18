# WebSocket Chat

A small Go web application for experimenting with a browser-based chat interface and WebSocket connections.

## Requirements

- [Go](https://go.dev/dl/) 1.27.1 or newer
- [Air](https://github.com/air-verse/air) for automatic rebuilds during development

## Getting started

Clone the repository and move into the project directory:

```bash
git clone <repository-url>
cd websocket_project
```

Download the Go dependencies:

```bash
go mod download
```

## Run the server

### Development mode with automatic reload

Install Air once:

```bash
go install github.com/air-verse/air@latest
```

Start the development server:

```bash
air
```

If the Go binary directory is not already on your `PATH`, use:

```bash
PATH="$(go env GOPATH)/bin:$PATH" air
```

Air reads [.air.toml](./.air.toml), watches Go, HTML, JavaScript, and CSS files, and rebuilds/restarts the server after changes are saved.

Open [http://localhost:8080](http://localhost:8080) in a browser.

### Run without automatic reload

```bash
go run .
```

## Project structure

```text
.
├── frontend/
│   ├── app.js       # Browser-side JavaScript
│   ├── index.html   # Main page
│   └── styles.css   # Frontend styles
├── .air.toml        # Air development reloader configuration
├── go.mod           # Go module and dependency definitions
├── main.go          # HTTP server entrypoint
└── manager.go       # WebSocket connection handling
```

## Development commands

Format Go source files:

```bash
gofmt -w *.go
```

Run tests:

```bash
go test ./...
```

Build the application:

```bash
go build ./...
```

## Notes

- The server listens on port `8080`.
- Air writes its temporary binary to `tmp/`; this directory is generated and should not be committed.
- Application and frontend HTTP request logs are appended to `logs/app.log` and also printed in the terminal. The `logs/` directory is generated and should not be committed.
- The WebSocket implementation is under active development.
