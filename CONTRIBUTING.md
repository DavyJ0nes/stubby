# Contributing

<!-- vim-markdown-toc GFM -->

- [Pull Requests](#pull-requests)
- [Testing](#testing)
  - [Unit Tests](#unit-tests)
  - [Integration Tests](#integration-tests)

<!-- vim-markdown-toc -->

Thank you so much for wanting to contribute to the [stubby]() project.

I built the tool to help fulfil a need that I had with simplifying
testing and I hope that it helps you as well.

Below I have put some requests and tips for how to request and add features
and updates to the tool.

## Pull Requests

If you want to contribute an update or a change, could you fork the repository
and then open a pull request merging your fork into the upstream. This helps
to keep the history of this upstream tidy :).

## Testing

### Unit Tests

As stubby just returns configured HTTP responses, most of the testing is handled
within the [router](./router) package.

Please follow the conventions used within the [router_test.go](./router/router_test.go)
file and ensure that any new functionality is covered by a test, existing or
newly created.

### Integration Tests

A Docker Image is built for stubby to make using it within other repositories easier.

To help ensure the image that is provided works as expected, some integration tests
are run whenever a pull request is opened. These utilise [docker-compose](https://docs.docker.com/compose)
to build the stubby Docker Image and then run tests that interact with a running
version of stubby.

The integration tests are located in the [testing/](./testing) package. The associated Compose
file can be found [here](docker-compose.test.yml).
