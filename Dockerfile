# -- Builder Image
FROM golang:1.11.4-alpine3.8 As Builder

RUN apk update && \
    apk add curl ca-certificates git && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/github.com/davyj0nes/stubby

COPY ./Gopkg.toml Gopkg.toml
COPY ./Gopkg.lock Gopkg.lock
RUN dep ensure -vendor-only

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo --installsuffix netgo -o stubby .

# -- Main Image
FROM alpine:3.8

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
