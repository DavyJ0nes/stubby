# Stubby
[![Go Report Card](https://goreportcard.com/badge/github.com/davyj0nes/stubby)](https://goreportcard.com/report/github.com/davyj0nes/stubby)
[![Lint and Test](https://github.com/DavyJ0nes/stubby/actions/workflows/lint-and-test.yml/badge.svg)](https://github.com/DavyJ0nes/stubby/actions/workflows/lint-and-test.yml)
[![Docker Info](https://img.shields.io/badge/docker-info-blue?style=flat&logo=docker)](https://hub.docker.com/repository/docker/davyj0nes/stubby/general)

![Logo](./docs/logo.jpg)

- [Description](#description)
- [Usage](#usage)
  - [Configuration](#configuration)
    - [Basic](#basic)
    - [URL Query Parameters](#url-query-parameters)
    - [Custom Response Headers](#custom-response-headers)
  - [Docker](#docker)
- [Contributing](#contributing)
- [License](#license)

## Description

Return stubbed HTTP responses defined in a config file.

## Usage

### Configuration

#### Basic

Add the routes and the responses that you want in the [config file](./config.yaml).

A basic route definition would look like:

```yaml
port: 8080
routes:
  - path: /foo
    status: 200
    response: >-
      {
        "id": 123,
        "message": "foo"
      }
```

#### URL Query Parameters

If the response has URL query parameters then define them as a flat list of key/value pairs:

```yaml
routes:
  - path: /foo
    queries:
      - show_deleted # param key
      - true         # param value
    status: 200
    response: >-
      {
        "id": 987,
        "message": "bar"
      }
```

The reason for using a flat list rather than a map is how the [Queries](https://www.gorillatoolkit.org/pkg/mux#Route.Queries) method is defined in gorilla/mux.

#### Custom Response Headers

```yaml
routes:
  - path: /foo
    status: 200
    headers:
      X-Custom: Header
      X-Request-Id: ef835eaf-a658-458b-86ae-d2d771f5e745
    response: >-
      {
        "id": 987,
        "message": "bar"
      }
```

### Docker

The image is available on [Docker Hub](https://hub.docker.com/r/davyj0nes/stubby). Build locally with:

```bash
mise run image
```

Run the container, mounting your config file:

```bash
docker run --rm -v "$PWD/config.yaml:/config.yaml" -p 8080:8080 davyj0nes/stubby -config /config.yaml
```

Stubby logs in JSON format:

```json
{"time":"2026-01-01T00:00:00Z","level":"INFO","msg":"starting stubby","addr":":8080"}
{"time":"2026-01-01T00:00:00Z","level":"INFO","msg":"route configured","path":"/foo"}
```

## Contributing

If you would like to contribute to this project then please check out the guidance
within: [CONTRIBUTING.md](./CONTRIBUTING.md)

## License

[MIT](./LICENSE)
