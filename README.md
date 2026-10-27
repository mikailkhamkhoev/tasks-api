# Tasks API

## Main idea

A small REST API for a task list, containerized. In-memory storage (no
database), so this project exists for one reason: **practicing Docker**.

## Endpoints

- `GET    /tasks`      - list all tasks
- `POST   /tasks`      - create a task (`{"title":"..."}`)
- `GET    /tasks/{id}` - get one task
- `DELETE /tasks/{id}` - delete a task

## Run it locally

```bash
go test ./...
go run .
```

Then from another terminal:

```bash
curl -X POST localhost:8080/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title":"buy milk"}'

curl localhost:8080/tasks
```

## Run it in Docker

```bash
docker build -t tasks .
docker run -p 8080:8080 tasks
```

Then from another terminal:

```bash
curl localhost:8080/tasks
```

## What's inside the image

The Dockerfile has two `FROM` lines - a multi-stage build:

- Stage 1 (`golang:1.22-alpine`) - has the Go toolchain, compiles the
  binary. ~350MB, thrown away.
- Stage 2 (`alpine:3.20`) - fresh, minimal Linux. Copies only the
  compiled binary from stage 1. ~10MB.

Final image: ~24MB instead of ~350MB. The Go toolchain is not in the
shipped image - only the binary is. Check with:

```bash
docker run --rm tasks ls /usr/local/go   # No such file or directory
docker run --rm tasks ls /app            # tasks
```

## Files

- `main.go` - startup and routes
- `go.mod` - module definition
- `store.go` - the in-memory task store
- `handlers.go` - the HTTP handlers
- `main_test.go` - a small test
- `Dockerfile` - multi-stage build
- `.dockerignore` - keeps README and .git out of the build
- `.gitignore` - keeps binaries and editor files out of git
- `LICENSE` - MIT
