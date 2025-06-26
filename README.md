# pulumi-envoy

This repository provides a complete example of building, running, and consuming a custom Pulumi provider for Envoy Gateway configuration.

## Project Structure

- **envoy/**: Minimal working Envoy and Envoy Gateway setup using Docker Compose.
  - `docker-compose.yaml`: Runs Envoy and Envoy Gateway containers.
  - `envoy.yaml`: Minimal Envoy configuration.
- **provider/**: Pulumi provider for Envoy Gateway configuration (Go, REST API).
  - Implements resources for Envoy routes, clusters, and listeners.
  - Uses provider config for `endpoint` (Envoy Gateway API URL) and `token` (API bearer token).
  - Communicates with Envoy Gateway via REST endpoints (e.g., `/v1/routes`, `/v1/clusters`).
  - See `provider/README.md` for development and extension details.
- **consumer/**: Example Pulumi YAML program using the local Envoy provider.
  - `Pulumi.yaml`: Minimal Pulumi YAML using the provider.

---

## Quick Start

### 1. Envoy & Envoy Gateway Setup

```sh
cd envoy
# Start Envoy and Envoy Gateway
docker-compose up
```

- The admin interface is available at [http://localhost:9901](http://localhost:9901)
- The Envoy Gateway API is available at [http://localhost:8080](http://localhost:8080)

### 2. Provider Development

- The provider is implemented in Go using [`pulumi-go-provider`](https://github.com/pulumi/pulumi-go-provider).
- Resources supported: **Route**, **Cluster**, **Listener** (CRUD via REST).
- Provider config:
  - `endpoint`: Envoy Gateway API URL (default: `http://localhost:8080`)
  - `token`: Bearer token for Envoy Gateway (if required)

#### Build the Provider

```sh
cd provider
make
# or
# go build -o bin/pulumi-resource-envoy main.go
```

#### Use Locally

- Copy or link the built provider binary to your consumer project, or ensure it is in your `PATH`.

### 3. Consumer Example

- See `consumer/Pulumi.yaml` for a minimal Pulumi YAML program using the provider.
- Run the example:

```sh
cd consumer
pulumi up
```

---

## Provider Features

- Written in Go using [`pulumi-go-provider`](https://github.com/pulumi/pulumi-go-provider).
- Supports Envoy Gateway REST API for configuration.
- Resources: **Route**, **Cluster**, **Listener** (CRUD via REST).
- Easily extensible for more Envoy resources.
- Example resource configuration:

```yaml
resources:
  my-route:
    type: envoy:index:Route
    properties:
      name: example-route
      cluster: example-cluster
      prefix: "/api"
  my-cluster:
    type: envoy:index:Cluster
    properties:
      name: example-cluster
      type: logical_dns
```

---

## Development & Extending

- Add new resources by creating a new file in `provider/` and following the pattern in existing resources.
- Update the provider registration in `provider.go` to include your new resource.
- See the [Envoy Gateway API docs](https://gateway.envoyproxy.io/latest/api/) for available endpoints and resource types.

---

## References

- [Envoy Gateway](https://gateway.envoyproxy.io/)
- [Pulumi Go Provider](https://github.com/pulumi/pulumi-go-provider)
- [Pulumi Documentation](https://www.pulumi.com/docs/)

---

See the README in each folder for more details and advanced usage.