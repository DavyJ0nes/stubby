# Stubby

![Logo](./docs/logo.jpg)
 

## Description

Return stubbed HTTP responses from a config file

## Usage

Add the routes and the responses that you want in the config file

Then start the server and make requests

The [Makefile](./Makefile) provides some helpers for this

1. Run stubby with Docker and `config.yaml` file
```
$ make run_image
# Running Docker Image Locally...
2019/02/14 17:33:37 starting server on  :8080
```

2. Hit endpoint with curl
```
$ curl http://localhost:8080/foo
{
  "id": 123,
  "message": "foo"
}
```

## License

MIT
