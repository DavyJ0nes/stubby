# -- Builder Image
FROM golang:1.12rc1-stretch As Builder

ENV GO111MODULE=on
WORKDIR /go/src/github.com/davyj0nes/stubby

COPY ./go.mod go.mod
COPY ./go.sum go.sum
RUN go mod vendor

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo --installsuffix netgo -o stubby .

# -- Main Image
FROM alpine:3.9

LABEL MAINTAINER=DavyJ0nes
LABEL OWNER=DavyJ0nes

# Copy binary from builder image
COPY --from=Builder /go/src/github.com/davyj0nes/stubby/stubby /bin/stubby
RUN chmod +x /bin/stubby

# Ensure not runnning as root user
RUN adduser -D -s /bin/sh app
USER app

EXPOSE 8080
CMD ["stubby"]
