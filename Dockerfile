FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN go build -o tasks .

FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/tasks .
EXPOSE 8080
CMD ["./tasks"]