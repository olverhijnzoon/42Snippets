# builder
FROM golang:1.19-alpine AS builder
WORKDIR /.
COPY rayleighscattering.go .
RUN go build -o rayleighscatteringX rayleighscattering.go

# rayleighscatterin
FROM alpine:latest AS rayleighscattering
WORKDIR /app
COPY --from=builder rayleighscatteringX /usr/bin
ENTRYPOINT ["/usr/bin/rayleighscatteringX"]