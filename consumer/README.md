# Pulumi Envoy Consumer Example

This folder contains an example Pulumi YAML program that uses the custom Envoy provider to generate dynamic configuration files for Envoy Proxy.

## Overview

- Demonstrates how to use the local Envoy provider to manage Envoy configuration as code.
- Generates LDS (listeners), CDS (clusters), and RDS (routes) YAML files in a specified output directory.
- Designed to work with the file-based dynamic xDS setup in the `envoy/` folder.

## Files

- `Pulumi.yaml`: Main Pulumi program definition.
- `Pulumi.<stack>.yaml`: Stack configuration file (e.g., `Pulumi.dev.yaml`) where you set the provider's `outputPath`.

## How It Works

1. **Provider Configuration**
   - The provider is configured with `outputPath`, which determines where the generated Envoy config files will be written.
   - Example stack config (`Pulumi.dev.yaml`):
     ```yaml
     config:
       envoy:outputPath: ../envoy/config/dynamic_configs
     ```

2. **Resources**
   - Define Envoy resources (e.g., clusters, listeners, routes) in `Pulumi.yaml`.
   - Example:
     ```yaml
     resources:
       cluster:
         type: envoy:provider:EnvoyCluster
         properties:
           name: demo
           type: demo
     ```

3. **Running the Program**
   - Run `pulumi up` to generate or update the Envoy config files in the specified output directory.
   - These files are then picked up by the running Envoy container (see `../envoy/docker-compose.yaml`).

## Usage

1. The provider will be built automatically when you run Pulumi commands, because the `Pulumi.yaml` in this folder specifies a relative path to the provider plugin (see the `plugins.providers` section).
2. Set the provider config in your stack config file (e.g., `Pulumi.dev.yaml`).
3. Run:
   ```sh
   pulumi up
   ```
4. Check the output directory (e.g., `../envoy/config/dynamic_configs/`) for the generated YAML files.
5. Envoy will automatically reload the configuration if running with file-based xDS.

## Customization

- You can add more resources (routes, clusters, listeners) to `Pulumi.yaml` as needed.
- The output directory can be changed by updating the `outputPath` config value.

## Troubleshooting

- Ensure the output directory exists and is writable.
- Check Pulumi logs for errors if resources are not created as expected.
- Make sure the Envoy container has the correct volume mount to the output directory.

---

For more details, see the main project [README](../README.md).
