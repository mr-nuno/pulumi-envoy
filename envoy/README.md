# Envoy Docker Compose Example with Dynamic File-based xDS Configuration

This project demonstrates how to run Envoy Proxy using Docker Compose, with dynamic configuration (xDS) loaded from the filesystem. The setup allows you to update listeners, clusters, and routes by editing YAML files, without rebuilding the container.

## Project Structure

```
.
├── docker-compose.yaml
├── config/
│   ├── bootstrap.yaml
│   └── dynamic_configs/
│       ├── listeners.yaml
│       ├── clusters.yaml
│       └── routes.yaml
├── backend-service/
│   ├── app.py
│   └── Dockerfile
```

## How It Works

- **Envoy** is started with a `bootstrap.yaml` that points to dynamic config files for LDS (listeners), CDS (clusters), and RDS (routes).
- **Dynamic config files** (`listeners.yaml`, `clusters.yaml`, `routes.yaml`) are mounted into the container and can be edited on the host. Envoy will pick up changes automatically.
- **Backend services** are simple HTTP servers, also managed by Docker Compose.

## Quick Start

1. **Clone the repository**

2. **Build and start the services:**
   ```sh
   docker-compose up --build
   ```

3. **Access Envoy:**
   - Proxy: [http://localhost:8080](http://localhost:8080)
   - Admin: [http://localhost:9901](http://localhost:9901)

4. **Test Routing:**
   - `curl http://localhost:8080/backend1` → routed to backend1
   - `curl http://localhost:8080/backend2` → routed to backend2
   - `curl http://localhost:8080/` → default route (backend1)

## Configuration Details

### docker-compose.yaml
- Mounts `config/bootstrap.yaml` and `config/dynamic_configs/` into the Envoy container at `/etc/envoy/`.
- Starts two backend services (`backend1`, `backend2`).

### config/bootstrap.yaml
- Configures Envoy to use file-based xDS for listeners and clusters.
- Example:
  ```yaml
  dynamic_resources:
    lds_config:
      path_config_source:
        path: /etc/envoy/dynamic_configs/listeners.yaml
        watched_directory:
          path: /etc/envoy/dynamic_configs/
      resource_api_version: V3
    cds_config:
      path_config_source:
        path: /etc/envoy/dynamic_configs/clusters.yaml
        watched_directory:
          path: /etc/envoy/dynamic_configs/
      resource_api_version: V3
  ```

### config/dynamic_configs/listeners.yaml
- Defines the HTTP listener and uses RDS to load routes from `routes.yaml`.

### config/dynamic_configs/routes.yaml
- Defines routing rules for `/backend1`, `/backend2`, and `/`.
- Uses `prefix_rewrite` to map `/backend1` and `/backend2` to `/` for the backend services.

### config/dynamic_configs/clusters.yaml
- Defines clusters for each backend service, matching the Docker Compose service names and ports.

## Hot Reloading
- Edit any of the YAML files in `config/dynamic_configs/`.
- Envoy will detect changes and reload configuration automatically.

## Troubleshooting
- Ensure the volume mount paths in `docker-compose.yaml` match the paths in `bootstrap.yaml`.
- Check Envoy logs (`docker-compose logs envoy`) for configuration errors.
- Use the Envoy admin interface (`localhost:9901`) to inspect config and stats.

## References
- [Envoy Proxy Documentation](https://www.envoyproxy.io/docs/envoy/latest/)
- [xDS API Overview](https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol)

---

Feel free to modify the backend services or add more clusters/routes as needed!
