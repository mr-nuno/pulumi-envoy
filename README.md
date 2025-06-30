# pulumi-envoy

This repository provides a complete example of building, running, and consuming a custom Pulumi provider for Envoy configuration using file-based dynamic xDS (LDS, CDS, RDS) and Docker Compose.

## Project Structure

- **envoy/**: Minimal working Envoy setup using Docker Compose and file-based dynamic configuration.
  - `docker-compose.yaml`: Runs Envoy and backend containers.
  - `config/bootstrap.yaml`: Envoy bootstrap config, points to dynamic config files.
  - `config/dynamic_configs/`: Folder for dynamic xDS config files (listeners.yaml, clusters.yaml, routes.yaml).
- **provider/**: Pulumi provider for Envoy configuration (Go, file-based xDS).
  - Implements resources for Envoy routes, clusters, and listeners.
  - Uses provider config for `outputPath` (directory for generated config files).
  - See `provider/README.md` for development and extension details.
- **consumer/**: Example Pulumi YAML program using the local Envoy provider.
  - `Pulumi.yaml`: Minimal Pulumi YAML using the provider and file-based config.

---

## Quick Start

### 1. Envoy Setup

```sh
cd envoy
# Start Envoy and backend services
docker-compose up
```

- The admin interface is available at [http://localhost:9901](http://localhost:9901)
- The HTTP proxy is available at [http://localhost:8080](http://localhost:8080)

### 2. Provider Development

- The provider is implemented in Go using [`pulumi-go-provider`](https://github.com/pulumi/pulumi-go-provider).
- Resources supported: **Route**, **Cluster**, **Listener** (CRUD via file-based xDS).
- Provider config:
  - `outputPath`: Directory for generated Envoy config files (e.g., `./demo`)

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
- Set the provider config for the stack (e.g., in `Pulumi.dev.yaml`):

```yaml
config:
  envoy:outputPath: ./demo
```

- Run the example:

```sh
cd consumer
pulumi up
```

---

## Provider Features

- Written in Go using [`pulumi-go-provider`](https://github.com/pulumi/pulumi-go-provider`).
- Supports file-based dynamic xDS configuration for Envoy (LDS, CDS, RDS).
- Resources: **Route**, **Cluster**, **Listener** (CRUD via file generation).
- Easily extensible for more Envoy resources.
- Example resource configuration:

```yaml
resources:
  cluster:
    type: envoy:provider:EnvoyCluster
    properties:
      name: demo
      type: demo
```

---

## Development & Extending

- Add new resources by creating a new file in `provider/` and following the pattern in existing resources.
- Update the provider registration in `provider.go` to include your new resource.

---

## Generating a .NET SDK for the Provider

You can generate a .NET SDK for your Pulumi provider using the `pulumi-gen-dotnet` tool. This allows you to consume your provider from C#, F#, or VB.NET Pulumi programs.

### Prerequisites
- Ensure you have the provider schema (`schema.json`) generated and available in your provider directory.
- Install the Pulumi .NET SDK generator if you haven't already:
  ```sh
  go install github.com/pulumi/pulumi/pkg/v3/cmd/pulumi-gen-dotnet@latest
  ```

### Generate the SDK
From the root of your provider directory, run:

```sh
pulumi-gen-dotnet --out sdk/dotnet/
```

- This will generate the .NET SDK in the `sdk/dotnet/` directory.
- The generator reads your `schema.json` and produces C# classes and types for all your provider resources and functions.

### Using the SDK
- You can reference the generated SDK in your .NET Pulumi projects by adding it as a project reference or packaging it as a NuGet package.
- Example usage in C#:
  ```csharp
  using Pulumi;
  using Pulumi.Envoy;
  // ...
  ```

---

## References

- [Envoy Proxy](https://www.envoyproxy.io/)
- [Pulumi Go Provider](https://github.com/pulumi/pulumi-go-provider)
- [Pulumi Documentation](https://www.pulumi.com/docs/)

---

See the README in each folder for more details and advanced usage.