# Docker Deployment Guide

This guide explains how to build and run the Todo Management application using Docker containers.

## Quick Start

### Using Docker Compose (Recommended)

#### React CSR Frontend
1. **Build and start CSR services**:
   ```bash
   docker-compose up --build
   ```

2. **Access the application**:
   - React Frontend: http://localhost:3000
   - Backend API: http://localhost:8080

3. **Stop the application**:
   ```bash
   docker-compose down
   ```

#### Go SSR Frontend
1. **Build and start SSR services**:
   ```bash
   docker-compose -f docker-compose-ssr.yml up --build
   ```

2. **Access the application**:
   - Go SSR Frontend: http://localhost:3001
   - Backend API: http://localhost:8080

3. **Stop the application**:
   ```bash
   docker-compose -f docker-compose-ssr.yml down
   ```

### Using Individual Docker Commands

#### Backend

1. **Build the backend image**:
   ```bash
   cd backend
   docker build -t todo-backend .
   ```

2. **Run the backend container**:
   ```bash
   docker run -d \
     --name todo-backend \
     -p 8080:8080 \
     todo-backend
   ```

#### React CSR Frontend

1. **Build the React frontend image**:
   ```bash
   cd frontend
   docker build -t todo-frontend .
   ```

2. **Run the React frontend container**:
   ```bash
   docker run -d \
     --name todo-frontend \
     -p 3000:8080 \
     -e VITE_API_URL=http://localhost:8080 \
     todo-frontend
   ```

#### Go SSR Frontend

1. **Build the Go SSR frontend image**:
   ```bash
   cd frontend-ssr-go
   docker build -t todo-frontend-ssr .
   ```

2. **Run the Go SSR frontend container**:
   ```bash
   docker run -d \
     --name todo-frontend-ssr \
     -p 3001:3001 \
     -e API_URL=http://localhost:8080 \
     -e PORT=3001 \
     todo-frontend-ssr
   ```

## Docker Images Details

### Backend Image (`backend/Dockerfile`)

**Base Images**:
- Build stage: `golang:1.21-alpine`
- Runtime stage: `alpine:latest`

**Features**:
- Multi-stage build for optimized image size
- Non-root user execution for security
- Health check endpoint
- Static binary compilation
- Minimal attack surface

**Exposed Port**: 8080

**Health Check**: `wget http://localhost:8080/health`

### React CSR Frontend Image (`frontend/Dockerfile`)

**Base Images**:
- Build stage: `node:24-alpine`
- Runtime stage: `nginx:1.25-alpine`

**Features**:
- Multi-stage build with Node.js build and Nginx runtime
- Custom Nginx configuration with SPA routing support
- Gzip compression enabled
- Security headers configured
- Non-root user execution
- Static asset caching

**Exposed Port**: 8080 (mapped to 3000 in docker-compose)

**Health Check**: `wget http://localhost:8080/health`

### Go SSR Frontend Image (`frontend-ssr-go/Dockerfile`)

**Base Images**:
- Build stage: `golang:1.24-alpine`
- Runtime stage: `alpine:latest`

**Features**:
- Multi-stage build for optimized image size
- Server-side rendering with Go templates
- HTTP server with routing and API integration
- Non-root user execution for security
- Static files and templates bundled
- Minimal attack surface

**Exposed Port**: 3001

**Health Check**: `wget http://localhost:3001/health`

## Configuration

### Environment Variables

#### Backend
- `GIN_MODE`: Set to `release` for production (default in docker-compose)

#### React CSR Frontend
- `VITE_API_URL`: Backend API URL (default: `http://localhost:8080`)

#### Go SSR Frontend
- `API_URL`: Backend API URL (default: `http://localhost:8080`)
- `PORT`: Server port (default: `3001`)

### Configuration Files

#### React CSR Frontend
The React frontend uses a custom Nginx configuration (`frontend/nginx.conf`) with:
- SPA routing support (fallback to index.html)
- Gzip compression
- Security headers
- Static asset caching
- Health check endpoint

#### Go SSR Frontend
The Go SSR frontend uses:
- Go html/template for server-side rendering
- Custom CSS matching React styling
- Static file serving for CSS and assets
- Built-in HTTP server with health checks

## Production Deployment

### Environment-Specific Configurations

#### Development
```bash
# Use docker-compose.yml as-is
docker-compose up --build
```

#### Production
```bash
# Override environment variables
VITE_API_URL=https://api.yourdomain.com \
GIN_MODE=release \
docker-compose up --build -d
```

### Security Considerations

1. **Non-root execution**: Both containers run as non-root users
2. **Minimal base images**: Alpine Linux for smaller attack surface
3. **Security headers**: CSP, XSS protection, frame options
4. **Health checks**: Container health monitoring
5. **Resource limits**: Can be added to docker-compose.yml

### Example Production docker-compose.override.yml

```yaml
version: '3.8'
services:
  backend:
    environment:
      - GIN_MODE=release
    deploy:
      resources:
        limits:
          cpus: '0.5'
          memory: 512M
        reservations:
          cpus: '0.1'
          memory: 128M

  frontend:
    environment:
      - VITE_API_URL=https://api.yourdomain.com
    deploy:
      resources:
        limits:
          cpus: '0.2'
          memory: 256M
        reservations:
          cpus: '0.05'
          memory: 64M
```

## Build Optimization

### .dockerignore Files

Both services include `.dockerignore` files to exclude unnecessary files:

**Backend**:
- Binaries and test files
- IDE and OS files
- Documentation
- Git files

**Frontend**:
- node_modules and build artifacts
- Environment files
- IDE and OS files
- Documentation

### Multi-stage Builds

Both Dockerfiles use multi-stage builds:
- **Backend**: Go build stage + minimal Alpine runtime
- **Frontend**: Node.js build stage + Nginx runtime

This approach significantly reduces final image sizes.

## Monitoring and Logging

### Health Checks

Both containers include health checks:
- **Interval**: 30 seconds
- **Timeout**: 3-10 seconds
- **Retries**: 3
- **Start period**: 40 seconds

### Logs

View container logs:
```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f backend
docker-compose logs -f frontend

# Individual containers
docker logs todo-backend
docker logs todo-frontend
```

## Troubleshooting

### Common Issues

1. **Port conflicts**:
   ```bash
   # Check if ports are in use
   lsof -i :3000
   lsof -i :8080
   
   # Change ports in docker-compose.yml if needed
   ```

2. **Build failures**:
   ```bash
   # Clean build (no cache)
   docker-compose build --no-cache
   
   # Remove existing containers
   docker-compose down --volumes --remove-orphans
   ```

3. **Frontend can't reach backend**:
   - Check `VITE_API_URL` environment variable
   - Ensure backend container is healthy
   - Verify network connectivity

### Development vs Production

**Development**:
- Use docker-compose for easy setup
- Mount volumes for live reload (if needed)
- Use default environment variables

**Production**:
- Use production environment variables
- Add resource limits
- Configure proper domain/SSL
- Use orchestration tools (Kubernetes, Docker Swarm)

## Image Sizes

Approximate image sizes after optimization:
- **Backend**: ~15-20 MB
- **Frontend**: ~25-30 MB

These sizes are achieved through:
- Multi-stage builds
- Alpine base images
- Minimal runtime dependencies
- Effective .dockerignore files

## Multi-Architecture Docker Images with Buildx

### Prerequisites

1. **Enable Docker Buildx** (available in Docker Desktop by default):
   ```bash
   # Check if buildx is available
   docker buildx version
   
   # Create and use a new builder instance
   docker buildx create --name multiarch-builder --use
   
   # Bootstrap the builder
   docker buildx inspect --bootstrap
   ```

2. **Login to Docker Hub** (or your registry):
   ```bash
   docker login
   ```

### Building Multi-Architecture Images

#### Backend Multi-Arch Build
```bash
cd backend

# Build and push for multiple architectures
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  --tag your-dockerhub-username/todo-backend:latest \
  --tag your-dockerhub-username/todo-backend:v1.0.0 \
  --push \
  .

# Build locally for testing (single platform)
docker buildx build \
  --platform linux/amd64 \
  --tag todo-backend:latest \
  --load \
  .
```

#### Frontend Multi-Arch Build
```bash
cd frontend

# Build and push for multiple architectures
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  --tag your-dockerhub-username/todo-frontend:latest \
  --tag your-dockerhub-username/todo-frontend:v1.0.0 \
  --push \
  .

# Build locally for testing (single platform)
docker buildx build \
  --platform linux/amd64 \
  --tag todo-frontend:latest \
  --load \
  .
```

### Automated Multi-Arch Build Script

Create a `build-multiarch.sh` script in the project root:

```bash
#!/bin/bash

# Configuration
DOCKER_USERNAME="your-dockerhub-username"
VERSION="v1.0.0"
LATEST_TAG="latest"

# Platforms to build for
PLATFORMS="linux/amd64,linux/arm64"

echo "Building multi-architecture Docker images..."

# Ensure buildx builder exists and is active
docker buildx create --name multiarch-builder --use 2>/dev/null || docker buildx use multiarch-builder

# Build backend
echo "Building backend for platforms: $PLATFORMS"
docker buildx build \
  --platform $PLATFORMS \
  --tag $DOCKER_USERNAME/todo-backend:$VERSION \
  --tag $DOCKER_USERNAME/todo-backend:$LATEST_TAG \
  --push \
  ./backend

# Build frontend
echo "Building frontend for platforms: $PLATFORMS"
docker buildx build \
  --platform $PLATFORMS \
  --tag $DOCKER_USERNAME/todo-frontend:$VERSION \
  --tag $DOCKER_USERNAME/todo-frontend:$LATEST_TAG \
  --push \
  ./frontend

echo "Multi-architecture build complete!"
echo "Images pushed:"
echo "  $DOCKER_USERNAME/todo-backend:$VERSION"
echo "  $DOCKER_USERNAME/todo-backend:$LATEST_TAG"
echo "  $DOCKER_USERNAME/todo-frontend:$VERSION"
echo "  $DOCKER_USERNAME/todo-frontend:$LATEST_TAG"
```

Make the script executable and run it:
```bash
chmod +x build-multiarch.sh
./build-multiarch.sh
```

### Using Multi-Arch Images in Production

Update your `docker-compose.yml` to use the published images:

```yaml
version: '3.8'
services:
  backend:
    image: your-dockerhub-username/todo-backend:latest
    container_name: todo-backend
    ports:
      - "8080:8080"
    environment:
      - GIN_MODE=release
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "wget", "--no-verbose", "--tries=1", "--spider", "http://localhost:8080/health"]
      interval: 30s
      timeout: 3s
      retries: 3
      start_period: 40s

  frontend:
    image: your-dockerhub-username/todo-frontend:latest
    container_name: todo-frontend
    ports:
      - "3000:8080"
    environment:
      - VITE_API_URL=http://backend:8080
    restart: unless-stopped
    depends_on:
      backend:
        condition: service_healthy
    healthcheck:
      test: ["CMD", "wget", "--no-verbose", "--tries=1", "--spider", "http://localhost:8080/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s

networks:
  default:
    name: todo-network
```

### Supported Architectures

Both images support the following architectures:
- **linux/amd64**: Intel/AMD 64-bit (most common)
- **linux/arm64**: ARM 64-bit (Apple Silicon, AWS Graviton, etc.)

### Buildx Commands Reference

```bash
# List available builders
docker buildx ls

# Create new builder
docker buildx create --name mybuilder --use

# Remove builder
docker buildx rm mybuilder

# Inspect current builder
docker buildx inspect

# Build for specific platform only
docker buildx build --platform linux/arm64 -t myimage:latest .

# Build and load locally (single platform only)
docker buildx build --platform linux/amd64 -t myimage:latest --load .

# Build and push to registry (supports multi-platform)
docker buildx build --platform linux/amd64,linux/arm64 -t myimage:latest --push .
```

### CI/CD Integration Example (GitHub Actions)

Create `.github/workflows/docker-build.yml`:

```yaml
name: Build and Push Multi-Arch Docker Images

on:
  push:
    branches: [ main ]
    tags: [ 'v*' ]
  pull_request:
    branches: [ main ]

env:
  REGISTRY: docker.io
  BACKEND_IMAGE: your-username/todo-backend
  FRONTEND_IMAGE: your-username/todo-frontend

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - name: Checkout
      uses: actions/checkout@v4

    - name: Set up Docker Buildx
      uses: docker/setup-buildx-action@v3

    - name: Login to Docker Hub
      if: github.event_name != 'pull_request'
      uses: docker/login-action@v3
      with:
        registry: ${{ env.REGISTRY }}
        username: ${{ secrets.DOCKERHUB_USERNAME }}
        password: ${{ secrets.DOCKERHUB_TOKEN }}

    - name: Extract metadata for backend
      id: meta-backend
      uses: docker/metadata-action@v5
      with:
        images: ${{ env.REGISTRY }}/${{ env.BACKEND_IMAGE }}
        tags: |
          type=ref,event=branch
          type=ref,event=pr
          type=semver,pattern={{version}}
          type=semver,pattern={{major}}.{{minor}}

    - name: Extract metadata for frontend
      id: meta-frontend
      uses: docker/metadata-action@v5
      with:
        images: ${{ env.REGISTRY }}/${{ env.FRONTEND_IMAGE }}
        tags: |
          type=ref,event=branch
          type=ref,event=pr
          type=semver,pattern={{version}}
          type=semver,pattern={{major}}.{{minor}}

    - name: Build and push backend
      uses: docker/build-push-action@v5
      with:
        context: ./backend
        platforms: linux/amd64,linux/arm64
        push: ${{ github.event_name != 'pull_request' }}
        tags: ${{ steps.meta-backend.outputs.tags }}
        labels: ${{ steps.meta-backend.outputs.labels }}
        cache-from: type=gha
        cache-to: type=gha,mode=max

    - name: Build and push frontend
      uses: docker/build-push-action@v5
      with:
        context: ./frontend
        platforms: linux/amd64,linux/arm64
        push: ${{ github.event_name != 'pull_request' }}
        tags: ${{ steps.meta-frontend.outputs.tags }}
        labels: ${{ steps.meta-frontend.outputs.labels }}
        cache-from: type=gha
        cache-to: type=gha,mode=max
```

## Commands Reference

```bash
# Build and start
docker-compose up --build

# Start in background
docker-compose up -d

# Stop and remove containers
docker-compose down

# View logs
docker-compose logs -f

# Rebuild specific service
docker-compose build backend
docker-compose build frontend

# Scale services (if needed)
docker-compose up --scale backend=2

# Execute commands in running containers
docker-compose exec backend sh
docker-compose exec frontend sh

# Multi-architecture builds
docker buildx build --platform linux/amd64,linux/arm64 -t myimage:latest --push .
```