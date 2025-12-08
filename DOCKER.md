# Docker Guide

Complete guide to building and running the Face Detection App using Docker and Docker Compose.

## Quick Start

```bash
# Build and start everything
docker-compose up -d --build

# View logs
docker-compose logs -f

# Stop and remove
docker-compose down
```

## Docker Compose Configuration

### Services

The application consists of two services:

#### Frontend Service
- **Image**: `face-detection-frontend:latest`
- **Container**: `face-detection-frontend`
- **Port Mapping**: `3000:80` (Nginx serves React build)
- **Base Image**: Node 18 Alpine (build), Nginx (runtime)
- **Purpose**: Serves the compiled React application
- **Health Check**: Checks `http://localhost:80` every 30s

#### Backend Service
- **Image**: `face-detection-backend:latest`
- **Container**: `face-detection-app`
- **Port Mapping**: `8080:8080`
- **Base Image**: Golang 1.21 (build), Debian Bullseye (runtime)
- **Purpose**: Go API server with OpenCV integration
- **Environment**: `PORT=8080`
- **Health Check**: Checks `/api/health` endpoint every 30s
- **Restart Policy**: `unless-stopped`

### Network

Both services are connected via the `face-detector-network` bridge network, enabling service-to-service communication.

## Using Makefile

For convenience, use the provided Makefile:

```bash
make help                         # Show all available commands
make build                        # Build Docker images
make up                          # Start containers in background
make down                        # Stop and remove containers
make logs                        # Show real-time logs
make restart                     # Restart running containers
make clean                       # Remove containers, networks, and images
make rebuild                     # Clean build and restart
```

## Building Images

### Manual Build

```bash
# Build all services
docker-compose build

# Build specific service
docker-compose build frontend
docker-compose build backend

# Force rebuild without cache
docker-compose build --no-cache
```

### Build Process

1. **Frontend** builds first:
   - Base: Node 18 Alpine
   - Installs npm dependencies
   - Runs `npm run build`
   - Creates optimized production build

2. **Backend** builds second:
   - Base: Golang 1.21 Bullseye (build stage)
   - Installs OpenCV dev libraries
   - Downloads Go modules
   - Compiles Go application
   - Final stage: Debian Bullseye slim
   - Copies compiled binary and frontend build

## Running Containers

### Start Services

```bash
# Start in background
docker-compose up -d

# Start and build if needed
docker-compose up -d --build

# Start with verbose output
docker-compose up
```

### Access Services

- **Frontend**: `http://localhost:3000`
- **Backend API**: `http://localhost:8080`
- **Health Check**: `curl http://localhost:8080/api/health`

### View Logs

```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f backend
docker-compose logs -f frontend

# Last 100 lines
docker-compose logs --tail=100

# No follow mode
docker-compose logs
```

### Check Status

```bash
# Container status
docker-compose ps

# Resource usage
docker stats

# Container processes
docker-compose top
```

## Container Management

### Stop Services

```bash
# Stop (keeps containers)
docker-compose stop

# Stop specific service
docker-compose stop backend

# Restart
docker-compose restart

# Force kill
docker-compose kill
```

### Remove Resources

```bash
# Remove containers and networks (keeps images)
docker-compose down

# Remove containers, networks, and volumes
docker-compose down -v

# Remove containers, networks, images, and volumes
docker-compose down -v --rmi all

# Only remove volumes
docker-compose down -v
```

## Troubleshooting

### Common Issues

#### Port Already in Use

```bash
# Find process using port
lsof -i :8080
lsof -i :3000

# Kill process
kill -9 <PID>

# Or change port in docker-compose.yml
```

**Solution**: Modify `docker-compose.yml`:
```yaml
services:
  backend:
    ports:
      - "8081:8080"  # Use 8081 instead
  frontend:
    ports:
      - "3001:80"    # Use 3001 instead
```

#### Build Failures

**OpenCV compilation errors:**
```bash
# Force rebuild without cache
docker-compose build --no-cache

# Check builder logs
docker-compose build --no-cache backend 2>&1 | head -50
```

**Dependency issues:**
```bash
# Verify images exist
docker images

# Clean and rebuild
docker-compose down -v --rmi all
docker-compose up -d --build
```

#### Container Exit/Crash

```bash
# Check logs
docker-compose logs backend
docker-compose logs frontend

# Check exit code
docker-compose ps

# Inspect container
docker inspect face-detection-app
```

#### Slow/Stuck Container

```bash
# Increase startup timeout
# (modify healthcheck start_period in docker-compose.yml)

# Force restart
docker-compose restart

# Rebuild
docker-compose up -d --build
```

### Debugging

#### Access Container Shell

```bash
# Backend
docker-compose exec backend /bin/bash

# Frontend (nginx)
docker-compose exec frontend /bin/sh
```

#### View Container Logs with Timestamps

```bash
docker-compose logs -f --timestamps

# With container names
docker-compose logs -f --timestamps frontend backend
```

#### Monitor Resource Usage

```bash
docker stats --no-stream
docker stats --no-stream face-detection-app
docker stats --no-stream face-detection-frontend
```

#### Network Debugging

```bash
# Check network
docker network inspect face-detector-network

# Test service connectivity from inside container
docker-compose exec backend ping frontend
docker-compose exec frontend ping backend
```

## Development Workflow

### Local Development with Docker

```bash
# Start services
docker-compose up -d --build

# Make code changes
# (edit files in your editor)

# Rebuild and restart (option 1)
docker-compose up -d --build

# Or restart just the changed service (option 2)
docker-compose restart backend
```

### With Volume Mounts (Faster Iteration)

For faster development, you can mount source directories:

```yaml
services:
  backend:
    volumes:
      - ./backend:/app/backend
      - /app/backend/.build  # Exclude build output

  frontend:
    volumes:
      - ./frontend:/app/frontend
      - /app/frontend/node_modules  # Exclude node_modules
```

Then use nodemon/air for auto-reload during development.

### Building Production Images

#### Multi-stage Builds

Both services use multi-stage builds to minimize image size:

- Frontend: Build in Node, serve from Nginx
- Backend: Build in Golang, minimal runtime in Debian slim

#### Image Sizes

```bash
# Check image sizes
docker images

# Example output:
# face-detection-backend:latest    ~450-550 MB (includes OpenCV)
# face-detection-frontend:latest   ~100-150 MB
```

#### Optimize Further

```dockerfile
# Use Alpine for smaller images
FROM golang:1.21-alpine AS builder

# Minimize final image
FROM alpine:latest
```

## Advanced Configuration

### Custom Environment Variables

```bash
# Create .env file
cat > .env << EOF
PORT=8080
ENVIRONMENT=production
LOG_LEVEL=info
EOF

# Reference in docker-compose.yml
environment:
  - PORT=${PORT}
```

### Scaling Services

Docker Compose doesn't scale the same service multiple times easily. For multiple instances:

```bash
# Using Kubernetes (see charts/face-detector/)
helm install face-detector ./charts/face-detector/
```

### Storage/Volumes

```yaml
services:
  backend:
    volumes:
      - backend-data:/data  # Named volume

volumes:
  backend-data:
    driver: local
```

## Security Considerations

### Best Practices

1. **Don't run as root** (already configured in Dockerfiles)
2. **Use specific image versions** (not `latest`)
3. **Scan images for vulnerabilities**:
   ```bash
   docker scan face-detection-backend:latest
   docker scan face-detection-frontend:latest
   ```

4. **Secrets management** (for production):
   - Use Docker secrets
   - Use environment variable files
   - Never commit `.env` files

5. **Network isolation**:
   - Services only communicate on internal network
   - Expose only necessary ports

## Performance Tuning

### Memory and CPU Limits

```yaml
services:
  backend:
    deploy:
      resources:
        limits:
          cpus: '0.5'
          memory: 512M
        reservations:
          cpus: '0.25'
          memory: 256M
```

### Optimization Tips

1. Use `.dockerignore` to exclude unnecessary files
2. Layer caching: put frequently changing commands near the end
3. Use Alpine images for smaller size
4. Clean up after installation: `apt-get clean`, `rm -rf /var/cache/apt/*`

## Kubernetes Deployment

For production deployment with Kubernetes:

```bash
# Install Helm chart
helm install face-detector ./charts/face-detector/

# Check status
kubectl get deployments
kubectl get services
kubectl logs -f deployment/face-detector
```

See [charts/face-detector/values.yaml](charts/face-detector/values.yaml) for configuration options.

## Additional Resources

- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [Docker Best Practices](https://docs.docker.com/develop/dev-best-practices/)
- [OpenCV Documentation](https://docs.opencv.org/)
- [Gin Web Framework](https://gin-gonic.com/)

