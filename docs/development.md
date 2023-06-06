# Development Guide

## Development Environment

The development environment is a [dev container](https://code.visualstudio.com/docs/remote/containers) that is configured to run the application in a containerized environment. The container environment will have the runtime dependencies installed and configured.

## Dependencies

For the runtime environment, the following dependencies are required:

* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/)
* [dev container](https://code.visualstudio.com/docs/remote/containers)

For the development environment, the following dependencies are required:
* [Golang](https://golang.org/)
* [CockroachDB](https://www.cockroachlabs.com/)
* [NATS](https://nats.io/)

These are all contained in the dev container. The `makefile` is configured to run in the dev container.

## Getting Started

Checkout the [Infratographer guide](https://infratographer.com/docs/development/local-setup/) on setting up `dev containers`.
