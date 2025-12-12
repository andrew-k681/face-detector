# Face Detection App

A full-stack web application that captures photos from your camera, detects faces using OpenCV Haar Cascades, and displays results in a responsive React frontend.

## Features

- 📸 Real-time camera capture from browser
- 👤 Face detection using OpenCV Haar Cascades
- 🎨 Modern, responsive React UI
- 🔄 RESTful API with Gin framework
- 🐳 Docker & Docker Compose support
- ☸️ Kubernetes ready (Helm charts included)
- 🔒 CORS enabled for seamless frontend-backend communication

## Tech Stack

- **Backend**: Go 1.21 + Gin framework + GoCV (OpenCV bindings)
- **Frontend**: React 18 + CSS3
- **Container**: Docker & Docker Compose
- **Orchestration**: Kubernetes (Helm charts)
- **Package Manager**: npm (frontend), Go modules (backend)

## Quick Start

### Option 1: Docker Compose (Recommended)

```bash
docker-compose up -d --build
```

Then visit:
- **Frontend**: `http://localhost:3000`
- **Backend API**: `http://localhost:8080`

### Option 2: Local Development

See [INSTALLATION.md](INSTALLATION.md) for detailed local setup instructions.

### Option 3: Kubernetes (Docker Desktop)

1. **Enable Kubernetes in Docker Desktop**:
   - Open Docker Desktop
   - Go to Settings > Kubernetes
   - Check "Enable Kubernetes"
   - Click "Apply & Restart"

2. **Build images**:
   ```bash
   # Build backend image
   cd backend
   docker build -t face-detector-backend:latest .
   
   # Build frontend image
   cd ../frontend
   docker build -t face-detector-frontend:latest .
   ```

3. **Deploy using Helm**:
   ```bash
   cd ../charts/face-detector
   
   # Install the chart
   helm install face-detector .
   ```

4. **Access the application**:
   - **Port-forwarding**: `kubectl port-forward svc/face-detector-frontend 3000:80` then visit `http://localhost:3000`

5. **Cleanup**:
   ```bash
   helm uninstall face-detector
   ```

## Documentation

- [INSTALLATION.md](INSTALLATION.md) - Detailed setup and installation guide
- [DOCKER.md](DOCKER.md) - Docker and containerization guide
- [DEVELOPMENT.md](DEVELOPMENT.md) - Development workflow and contributing
- [API.md](API.md) - API endpoints and integration guide

## Project Structure

```
face-detector/
├── backend/                      # Go backend service
│   ├── main.go                  # Entry point
│   ├── go.mod                   # Go module definition
│   ├── Dockerfile               # Backend container image
│   └── internal/
│       ├── api/
│       │   ├── dto/             # Data transfer objects
│       │   └── handlers/        # HTTP request handlers
│       ├── classifiers/         # OpenCV cascade files
│       └── face_detect/         # Face detection logic
├── frontend/                     # React frontend
│   ├── public/
│   │   └── index.html           # HTML entry point
│   ├── src/
│   │   ├── App.js               # Main React component
│   │   ├── App.css              # Styling
│   │   ├── index.js             # React bootstrap
│   │   └── index.css            # Global styles
│   ├── Dockerfile               # Frontend container image
│   ├── nginx.conf               # Nginx configuration
│   └── package.json             # NPM dependencies
├── charts/                       # Helm chart for Kubernetes
│   └── face-detector/
│       ├── Chart.yaml
│       ├── values.yaml
│       └── templates/
├── docker-compose.yml            # Docker Compose orchestration
├── Makefile                      # Build automation
├── run.sh                        # Local startup script
├── README.md                     # This file
└── DOCKER.md                     # Docker guide
```

## Quick Commands

### Using Docker Compose

```bash
docker-compose up -d --build      # Build and start
docker-compose logs -f             # View logs
docker-compose down                # Stop and remove
docker-compose ps                  # Check status
```

### Using Makefile

```bash
make build                         # Build images
make up                           # Start containers
make down                         # Stop containers
make logs                         # View logs
make restart                      # Restart containers
make clean                        # Remove all containers/images
make rebuild                      # Rebuild and restart
```

### Local Development

```bash
./run.sh                          # Start both backend and frontend
```

## Architecture

### Deployment Architecture

The application follows a containerized architecture:

```
┌─────────────────────────────────────────┐
│         Frontend (React)                │
│      Port 3000 (Development)            │
│     Port 80 (Docker Container)          │
└──────────────┬──────────────────────────┘
               │
               │ HTTP Requests
               ▼
┌─────────────────────────────────────────┐
│    Backend (Go + Gin + OpenCV)          │
│     Port 8080                           │
│                                         │
│  ├─ POST /api/detect-face              │
│  └─ GET /api/health                    │
└─────────────────────────────────────────┘
```

### Service Communication

- **Frontend to Backend**: Direct HTTP requests (CORS enabled)
- **Docker Network**: Services communicate via `face-detector-network` bridge
- **Health Checks**: Both services include health check probes

## Usage

1. Open your browser to `http://localhost:3000` (or `http://localhost:8080` if using Docker single-port)
2. Allow camera permissions when prompted
3. Click "Start Camera" to activate your camera
4. Click "Capture Photo" to take a picture
5. Click "Detect Face" to analyze for faces
6. Results will show the photo with detected faces highlighted

## System Requirements

### Docker (Recommended)
- Docker 20.10+
- Docker Compose 2.0+
- 2GB RAM (minimum)
- 1GB disk space (for base images)

### Local Development
- Go 1.21 or higher
- Node.js 16+ with npm
- OpenCV 4.x
- 2GB RAM
- macOS / Linux / Windows with WSL2

## License

MIT License - See LICENSE file for details

## Support

For issues, questions, or contributions, please refer to [DEVELOPMENT.md](DEVELOPMENT.md).

