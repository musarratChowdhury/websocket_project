# Project instructions

## Overview

- This is a Go web application with static frontend files in `frontend/`.
- The Go module requires Go 1.27.1 or newer.
- The application serves the frontend on port `8080`.

## Development workflow

- Run the application with Air during development so changes are rebuilt automatically:

  ```bash
  air
  ```

- If Air is not on `PATH`, install it with:

  ```bash
  go install github.com/air-verse/air@latest
  ```

  Then run it with:

  ```bash
  PATH="$(go env GOPATH)/bin:$PATH" air
  ```

- Air is configured in `.air.toml` and watches Go, HTML, JavaScript, and CSS files.
- The generated `tmp/` directory is build output and should not be edited or committed.

## Validation

- Format Go changes with `gofmt`.
- Run the test suite with:

  ```bash
  go test ./...
  ```

- Build the application with:

  ```bash
  go build ./...
  ```

## Change guidelines

- Keep backend code in the project root unless a new package is justified.
- Keep browser assets under `frontend/`.
- Preserve the existing HTTP entrypoint on port `8080` unless the configuration and documentation are updated together.
