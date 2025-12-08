# Installation Guide

Detailed instructions for setting up the Face Detection App in various environments.

## Table of Contents

- [Docker Installation](#docker-installation) (Recommended)
- [Local Development Setup](#local-development-setup)
- [Platform-Specific Setup](#platform-specific-setup)
- [Troubleshooting](#troubleshooting)

## Docker Installation

### Prerequisites

- Docker 20.10 or higher
- Docker Compose 2.0 or higher
- 2GB RAM minimum
- 1GB disk space

### macOS Docker Installation

```bash
# Using Homebrew (recommended)
brew install --cask docker

# Or download Docker Desktop from:
# https://www.docker.com/products/docker-desktop

# Verify installation
docker --version
docker-compose --version
```

### Linux Docker Installation

#### Ubuntu/Debian

```bash
# Update package manager
sudo apt-get update

# Install Docker
sudo apt-get install -y docker.io docker-compose

# Add user to docker group (avoid sudo)
sudo usermod -aG docker $USER
newgrp docker

# Verify installation
docker --version
docker-compose --version
```

#### CentOS/RHEL

```bash
# Install Docker
sudo yum install -y docker docker-compose

# Start Docker daemon
sudo systemctl start docker
sudo systemctl enable docker

# Add user to docker group
sudo usermod -aG docker $USER
newgrp docker

# Verify installation
docker --version
docker-compose --version
```

### Windows Docker Installation

1. Download Docker Desktop from [docker.com](https://www.docker.com/products/docker-desktop)
2. Run the installer and follow the setup wizard
3. Ensure WSL2 (Windows Subsystem for Linux 2) is installed
4. Restart your computer if prompted
5. Open PowerShell and verify:
   ```powershell
   docker --version
   docker-compose --version
   ```

### Quick Start with Docker

```bash
# Navigate to project directory
cd face-detector

# Build and start all services
docker-compose up -d --build

# Check status
docker-compose ps

# View logs
docker-compose logs -f

# Access the application
# Frontend: http://localhost:3000
# Backend: http://localhost:8080
```

## Local Development Setup

For local development without Docker, you'll need to install dependencies on your machine.

### Prerequisites

- **Go 1.21 or higher**
- **Node.js 16 or higher** with npm
- **OpenCV 4.x**
- **Git** (for cloning the repository)
- **pkg-config** (for Go/OpenCV compilation)

### Step 1: Clone Repository

```bash
git clone <repository-url>
cd face-detector
```

### Step 2: Install OpenCV

#### macOS

```bash
# Using Homebrew (recommended)
brew install opencv

# Verify installation
pkg-config --cflags --libs opencv4

# If using OpenCV 3.x
pkg-config --cflags --libs opencv
```

#### Ubuntu/Debian

```bash
# Update package manager
sudo apt-get update

# Install OpenCV development libraries
sudo apt-get install -y libopencv-dev

# Install pkg-config
sudo apt-get install -y pkg-config

# Verify installation
pkg-config --cflags --libs opencv4
```

#### CentOS/RHEL

```bash
# Install OpenCV development libraries
sudo yum install -y opencv-devel

# Install pkg-config
sudo yum install -y pkgconfig

# Verify installation
pkg-config --cflags --libs opencv
```

#### Windows (WSL2)

```bash
# In WSL2 terminal
sudo apt-get update
sudo apt-get install -y libopencv-dev pkg-config

# Verify
pkg-config --cflags --libs opencv4
```

### Step 3: Backend Setup

```bash
# Navigate to backend directory
cd backend

# Download Go dependencies
go mod download

# Verify dependencies
go mod verify

# Test backend startup
go run main.go

# Backend should print: "Server starting on port 8080"
# Press Ctrl+C to stop
```

### Step 4: Frontend Setup

In a new terminal:

```bash
# Navigate to frontend directory
cd frontend

# Install npm dependencies
npm install

# Start development server
npm start

# Frontend will open at http://localhost:3000
```

### Step 5: Verify Installation

1. Open browser to `http://localhost:3000`
2. You should see the Face Detection App interface
3. Test backend health: `http://localhost:8080/api/health`
4. Should respond: `{"message":"ok"}`

## Platform-Specific Setup

### macOS Setup

#### Using Homebrew (Recommended)

```bash
# Install all dependencies
brew install go node opencv

# Install npm dependencies
cd frontend && npm install

# Verify versions
go version        # Should be 1.21+
node --version    # Should be 16+
npm --version     # Should be 8+
pkg-config --list-all | grep opencv
```

#### Using .zshrc

Add to your `~/.zshrc`:

```bash
# Go paths
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# OpenCV paths
export PKG_CONFIG_PATH="/usr/local/lib/pkgconfig"
```

Then reload:

```bash
source ~/.zshrc
```

### Linux (Ubuntu 22.04 LTS) Setup

```bash
#!/bin/bash
set -e

# Update system
sudo apt-get update && sudo apt-get upgrade -y

# Install Go
cd /tmp
wget https://go.dev/dl/go1.21.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.linux-amd64.tar.gz

# Install Node.js
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs

# Install OpenCV and dependencies
sudo apt-get install -y libopencv-dev pkg-config build-essential

# Add Go to PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Verify
go version
node --version
pkg-config --cflags --libs opencv4
```

### Windows (WSL2) Setup

```bash
# Open PowerShell as Administrator

# Install WSL2 if not already installed
wsl --install

# In WSL2 terminal:
curl -L https://golang.org/dl/go1.21.linux-amd64.tar.gz -o go1.21.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.linux-amd64.tar.gz

# Install Node.js
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs

# Install OpenCV
sudo apt-get update
sudo apt-get install -y libopencv-dev pkg-config

# Update PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

## Using run.sh Script

A convenience script is provided for local development:

```bash
# Make it executable
chmod +x run.sh

# Start both backend and frontend
./run.sh

# Both services will start in the background
# Frontend: http://localhost:3000
# Backend: http://localhost:8080

# To stop, press Ctrl+C
```

## Build from Source

### Frontend Build

```bash
cd frontend

# Install dependencies
npm install

# Build for production
npm run build

# Built files are in frontend/build/
```

### Backend Build

```bash
cd backend

# Build binary
go build -o face-detection-server main.go

# Run the binary
./face-detection-server

# Or build with specific output
go build -o dist/server main.go
```

## Environment Configuration

### Backend Environment Variables

```bash
# Port configuration (default: 8080)
PORT=8080

# Log level (optional)
LOG_LEVEL=info
```

### Frontend Configuration

Edit `frontend/package.json` proxy:

```json
"proxy": "http://localhost:8080"
```

## Verification Checklist

After installation, verify everything works:

- [ ] Backend starts without errors: `go run main.go`
- [ ] Backend health check: `curl http://localhost:8080/api/health`
- [ ] Frontend starts: `npm start` (from frontend directory)
- [ ] Frontend loads: `http://localhost:3000`
- [ ] Camera access is permitted in browser
- [ ] Face detection works with sample image

## Troubleshooting Installation

### OpenCV Not Found

```bash
# Verify installation
pkg-config --cflags --libs opencv4

# If not found, check installation:
# macOS: brew list opencv
# Linux: dpkg -l | grep opencv
# Windows/WSL2: apt list --installed | grep opencv
```

### Go Module Errors

```bash
# Clear cache and redownload
go clean -modcache
go mod download
go mod verify
```

### Port Already in Use

```bash
# macOS/Linux: Find process
lsof -i :8080
lsof -i :3000

# Kill process
kill -9 <PID>

# Or change ports in environment
PORT=8081 go run main.go
```

### npm Install Fails

```bash
# Clear npm cache
npm cache clean --force

# Try install again
npm install

# If still failing, try npm ci
npm ci
```

### Permission Denied

```bash
# Add execute permission
chmod +x run.sh

# For macOS/Linux binary
chmod +x face-detection-server
```

## Next Steps

After successful installation:

1. Read [README.md](README.md) for project overview
2. Read [DEVELOPMENT.md](DEVELOPMENT.md) for development workflow
3. Read [API.md](API.md) for API documentation
4. Read [DOCKER.md](DOCKER.md) for containerization details

## Getting Help

For issues not covered here:

1. Check application logs
2. Verify all prerequisites are installed
3. Consult [DEVELOPMENT.md](DEVELOPMENT.md) for known issues
4. Review error messages in detail
5. Search existing GitHub issues (if applicable)
