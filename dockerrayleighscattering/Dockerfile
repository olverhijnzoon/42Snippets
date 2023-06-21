# builder
FROM golang:1.19-alpine AS builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
# download in a separate layer so Docker can cache it
RUN go mod download
COPY rayleighscattering.go .
RUN go build -o rayleighscatteringX rayleighscattering.go

# rayleighscattering
FROM alpine:latest AS rayleighscattering
WORKDIR /app
# app runs as appuser (no-login non-root no-pass and home is /app)
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/rayleighscatteringX /usr/bin
ENTRYPOINT ["/usr/bin/rayleighscatteringX"]
