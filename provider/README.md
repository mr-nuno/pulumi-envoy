# Pulumi Envoy Provider Skeleton

This folder contains a skeleton for a Pulumi provider for Envoy, based on the structure of [mr-nuno/pulumi-loopia-dns](https://github.com/mr-nuno/pulumi-loopia-dns).

## Getting Started

- Implement the provider logic in Go.
- Update the provider schema and build the provider binary.
- See the referenced repo for detailed implementation guidance.

## Structure
- `cmd/` - Provider entrypoint
- `internal/` - Provider logic
- `pkg/` - Go SDK
- `provider/` - Pulumi provider plugin
- `Makefile` - Build commands
- `go.mod`/`go.sum` - Go dependencies

> **Note:** This is a skeleton. You must implement the actual Envoy resource logic.
