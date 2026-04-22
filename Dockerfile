# Multistage Build

# -- Set ARGS
ARG APP_NAME=stubby
ARG MAIN_PATH=cmd/main.go

# -- Builder Image
FROM golang:1.26 AS builder

ARG APP_NAME
ARG MAIN_PATH

WORKDIR /app

# Set up dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the package code
COPY . .

# Build the statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags="-s -w" -o $APP_NAME $MAIN_PATH

# -- Main Image
FROM gcr.io/distroless/static-debian12:nonroot

ARG APP_NAME

COPY --from=builder /app/${APP_NAME} /app

EXPOSE 8080
ENTRYPOINT ["/app"]
