# Stubby
[![Go Report Card](https://goreportcard.com/badge/github.com/davyj0nes/stubby)](https://goreportcard.com/report/github.com/davyj0nes/stubby)
[![Build Status](https://travis-ci.org/DavyJ0nes/stubby.svg?branch=master)](https://travis-ci.org/DavyJ0nes/stubby)

![Logo](./docs/logo.jpg)

<!-- vim-markdown-toc Redcarpet -->

- [Description](#description)
- [Usage](#usage)
- [License](#license)

<!-- vim-markdown-toc -->

## Description

Return stubbed HTTP responses from a config file

## Usage

Add the routes and the responses that you want in the config file

Run the docker container using `make run_image` in this directory.

```
$ make run_image
# Running Docker Image Locally...
2019/02/14 17:33:37 starting server on  :8080
```

Or you can run the following docker run command anywhere:

```
docker run --rm -v "$PWD/config.yaml:/bin/config.yaml" -p 8080:8080 davyj0nes/stubby:0.2.0
```

## License

MIT
