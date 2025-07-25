# ---------- Builder ----------
FROM golang:1.22 AS builder

# Enable Go modules and static linking
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /src

# Copy go.mod and go.sum first for dependency download caching
COPY go.mod ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy the rest of the source
COPY . .

# Build with optimizations (stripped symbols)
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags="-s -w" -o /bin/backend .

# ---------- Final (minimal) image ----------
FROM scratch

# Copy binary and any needed assets
COPY --from=builder /bin/backend /bin/backend

# If your app needs timezone or certs, switch to an alpine base.

EXPOSE 8080

ENTRYPOINT ["/bin/backend"]
