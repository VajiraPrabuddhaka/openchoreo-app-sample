# Docker Deployment Guide

This guide explains how to build and run the Todo Management application using Docker containers.

## Quick Start

### Using Docker Compose (Recommended)

1. **Build and start all services**:
   ```bash
   docker-compose up --build
   ```

2. **Access the application**:
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080

3. **Stop the application**:
   ```bash
   docker-compose down
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

#### Frontend

1. **Build the frontend image**:
   ```bash
   cd frontend
   docker build -t todo-frontend .
   ```

2. **Run the frontend container**:
   ```bash
   docker run -d \
     --name todo-frontend \
     -p 3000:8080 \
     -e VITE_API_URL=http://localhost:8080 \
     todo-frontend
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

### Frontend Image (`frontend/Dockerfile`)

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

## Configuration

### Environment Variables

#### Backend
- `GIN_MODE`: Set to `release` for production (default in docker-compose)

#### Frontend
- `VITE_API_URL`: Backend API URL (default: `http://localhost:8080`)

### Nginx Configuration

The frontend uses a custom Nginx configuration (`frontend/nginx.conf`) with:
- SPA routing support (fallback to index.html)
- Gzip compression
- Security headers
- Static asset caching
- Health check endpoint

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
```