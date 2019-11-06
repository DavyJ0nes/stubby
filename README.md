# Stubby
[![Go Report Card](https://goreportcard.com/badge/github.com/davyj0nes/stubby)](https://goreportcard.com/report/github.com/davyj0nes/stubby)
[![Build Status](https://travis-ci.org/DavyJ0nes/stubby.svg?branch=master)](https://travis-ci.org/DavyJ0nes/stubby)

![Logo](./docs/logo.jpg)

<!-- vim-markdown-toc Redcarpet -->

* [Description](#description)
* [Usage](#usage)
  * [Configuration](#configuration)
  * [Docker](#docker)
* [License](#license)

<!-- vim-markdown-toc -->

## Description

Return stubbed HTTP responses defined in a config file

## Usage

### Configuration

### Basic

Add the routes and the responses that you want in the [config file](./comfig.yaml).

A basic route definition would look like:

```yaml
routes:
  - path: /foo
    status: 200
    response: >-
      {
        "id": 123,
        "message": "foo"
      }
```

### URL Query Params

If the response has URL parameters then these need to be defined as follows:

```yaml
routes:
  - path: /foo
    queries:
      - show_deleted # this is the param key
      - true         # this is the param value
    status: 200
    respose: >-
      {
        "id": 987,
        "message": "bar"
      }
```

The reason for having them defined in a list rather than as a key/value pair
is due to how the (Queries](https://www.gorillatoolkit.org/pkg/mux#Route.Queries)
method is defined in the router package used ([gorilla mux](https://www.gorillatoolkit.org)).

### Custom Response Headers

If you want the response to include a header then you can add it as such:

```yaml
routes:
  - path: /foo
    status: 200
    headers:
      X-Custom: Header
      X-Request-Id: ef835eaf-a658-458b-86ae-d2d771f5e745
    respose: >-
      {
        "id": 987,
        "message": "bar"
      }
```

### Docker

The artifact is stored as a docker image and is located on [Docker Hub](https://hub.docker.com/r/davyj0nes/stubby)

You can also build locally if required by running `make image`.

Run the docker container using `make run_image` in this directory.

```
$ make run_image
# Running Docker Image Locally...
2019/02/14 17:33:37 starting server on  :8080
```

Or you can run the following docker run command anywhere:

```
docker run --rm -v "$PWD/config.yaml:/bin/config.yaml" -p 8080:8080 davyj0nes/stubby
```

## License

MIT
