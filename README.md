# Go Products Catalog Service

Lightweight product catalog microservice written in Go (stdlib only — no external dependencies).
Part of the **GitOps Demo** multi-service application.

## Role in the Demo

```
Browser → Node.js Frontend → Go Catalog Service (this)
                           → Java Inventory Service
```

The Node.js frontend calls this service at `/products` to display the product catalog cards on the dashboard.

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/products` | Returns all products as JSON |
| GET | `/health` | Health check — returns `{"status":"ok","service":"products-go"}` |

### Sample Response — `GET /products`

```json
[
  {"id":1,"name":"Laptop Pro","category":"Electronics","price":1299.99,"stock":45},
  {"id":2,"name":"Wireless Mouse","category":"Accessories","price":29.99,"stock":200}
]
```

## Run Locally

**Prerequisites:** Go 1.21+

```bash
go run main.go
# Service starts on :8081

curl http://localhost:8081/health
curl http://localhost:8081/products
```

## Test

```bash
go test ./...
```

3 unit tests covering `/health` and `/products` handlers.

## Docker

```bash
docker build -t demo-go .
docker run -p 8081:8081 demo-go
```

Environment variable: `PORT` (default `8081`)

## Kubernetes

Manifests are in [`kubernetes/`](kubernetes/):

| File | Purpose |
|------|---------|
| `namespace.yaml` | Creates the `demo` namespace |
| `deployment.yaml` | 1 replica, ClusterIP, health probes |
| `service.yaml` | ClusterIP on port `8081` — internal only |
| `kustomization.yaml` | Kustomize entrypoint |

Apply:
```bash
kubectl apply -k kubernetes/
```

> **Image:** Update `DOCKERHUB_USER/demo-go:latest` in `deployment.yaml` with your registry image after the CI pipeline builds and pushes it.

## CI Pipeline (gitops-platform)

When onboarding in the gitops-platform, set:
- **Language:** Go
- **Dockerfile:** `Dockerfile`
- **Build context:** `.`
- **Compile command:** `go build ./...`
- **Test command:** `go test ./...`
