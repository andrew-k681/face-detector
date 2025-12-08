# Documentation Index

Quick navigation guide to the Face Detection App documentation.

## 📚 Documentation Overview

| Document | Purpose | Audience |
|----------|---------|----------|
| [README.md](README.md) | Project overview, quick start, features | Everyone |
| [INSTALLATION.md](INSTALLATION.md) | Detailed setup instructions for all platforms | Developers, DevOps |
| [DOCKER.md](DOCKER.md) | Docker & Docker Compose guide | DevOps, Operators |
| [API.md](API.md) | Complete API reference & integration guide | Frontend devs, API consumers |
| [DEVELOPMENT.md](DEVELOPMENT.md) | Development workflow, testing, contributing | Developers |

## 🚀 Getting Started

### I want to run the app quickly
→ See [README.md - Quick Start](README.md#quick-start)

### I need to set up locally
→ See [INSTALLATION.md - Local Development Setup](INSTALLATION.md#local-development-setup)

### I want to use Docker
→ See [DOCKER.md](DOCKER.md)

### I'm integrating with the API
→ See [API.md](API.md)

### I want to contribute/develop
→ See [DEVELOPMENT.md](DEVELOPMENT.md)

## 📋 Documentation Map

```
📖 README.md
   ├─ Features
   ├─ Tech Stack
   ├─ Quick Start (Docker)
   └─ System Requirements

📖 INSTALLATION.md
   ├─ Docker Installation
   │  ├─ macOS
   │  ├─ Linux (Ubuntu/Debian, CentOS/RHEL)
   │  └─ Windows (WSL2)
   ├─ Local Development Setup
   │  ├─ Install OpenCV
   │  ├─ Backend Setup
   │  └─ Frontend Setup
   ├─ Platform-Specific Guides
   └─ Troubleshooting

📖 DOCKER.md
   ├─ Quick Start
   ├─ Makefile Commands
   ├─ Service Configuration
   ├─ Building Images
   ├─ Running Containers
   ├─ Troubleshooting
   ├─ Development Workflow
   ├─ Advanced Configuration
   └─ Security & Performance

📖 API.md
   ├─ Base URL & Auth
   ├─ GET /api/health
   │  └─ Health check endpoint
   ├─ POST /api/detect-face
   │  ├─ Request/Response format
   │  ├─ Error scenarios
   │  ├─ Example usage (JS, Python, cURL)
   │  └─ Performance notes
   ├─ Error Handling
   ├─ Rate Limiting
   └─ CORS Configuration

📖 DEVELOPMENT.md
   ├─ Development Setup
   ├─ Project Architecture
   ├─ Development Workflow
   ├─ Testing
   ├─ Code Structure
   ├─ Debugging
   ├─ Best Practices
   ├─ Contributing
   └─ Common Tasks
```

## 🎯 Common Scenarios

### "I just want to try it out"
1. Install Docker (see [INSTALLATION.md - Docker Installation](INSTALLATION.md#docker-installation))
2. Run: `docker-compose up -d --build`
3. Open: `http://localhost:3000`

### "I want to develop locally"
1. Read [INSTALLATION.md - Local Development Setup](INSTALLATION.md#local-development-setup)
2. Run: `./run.sh`
3. Frontend: `http://localhost:3000`, Backend: `http://localhost:8080`
4. See [DEVELOPMENT.md](DEVELOPMENT.md) for workflow

### "I need to integrate the API"
1. See [API.md - Detect Faces](API.md#2-detect-faces)
2. Check example usage for your language (JS, Python, cURL)
3. Handle error responses as documented

### "I'm having deployment issues"
1. Check [DOCKER.md - Troubleshooting](DOCKER.md#troubleshooting)
2. Review [INSTALLATION.md - Troubleshooting](INSTALLATION.md#troubleshooting-installation)
3. Check logs: `docker-compose logs -f`

### "I want to contribute"
1. Read [DEVELOPMENT.md - Contributing](DEVELOPMENT.md#contributing)
2. Read [DEVELOPMENT.md - Code Structure](DEVELOPMENT.md#code-structure)
3. Check [DEVELOPMENT.md - Best Practices](DEVELOPMENT.md#best-practices)

## 📚 Topic Index

### Installation & Setup
- [Docker Setup](INSTALLATION.md#docker-installation)
- [Local Development](INSTALLATION.md#local-development-setup)
- [macOS Setup](INSTALLATION.md#macos-setup)
- [Linux Setup](INSTALLATION.md#linux-ubuntu-2204-lts-setup)
- [Windows Setup](INSTALLATION.md#windows-wsl2-setup)

### Deployment
- [Docker Compose](DOCKER.md#docker-compose-configuration)
- [Using Makefile](DOCKER.md#using-makefile)
- [Building Images](DOCKER.md#building-images)
- [Running Containers](DOCKER.md#running-containers)

### API Integration
- [Health Check](API.md#1-health-check)
- [Face Detection](API.md#2-detect-faces)
- [Error Handling](API.md#error-handling)
- [Code Examples](API.md#example-usage-in-javascript)

### Development
- [Development Setup](DEVELOPMENT.md#development-setup)
- [Architecture](DEVELOPMENT.md#project-architecture)
- [Testing](DEVELOPMENT.md#testing)
- [Debugging](DEVELOPMENT.md#debugging)
- [Contributing](DEVELOPMENT.md#contributing)

### Troubleshooting
- [Installation Issues](INSTALLATION.md#troubleshooting-installation)
- [Docker Issues](DOCKER.md#troubleshooting)
- [Common Tasks](DEVELOPMENT.md#common-tasks)

## 🔍 Quick Command Reference

### Docker
```bash
docker-compose up -d --build      # Start
docker-compose logs -f             # View logs
docker-compose down                # Stop
```

### Makefile
```bash
make build                         # Build
make up                           # Start
make logs                         # View logs
make down                         # Stop
```

### Local Development
```bash
./run.sh                          # Start both services
go run main.go                    # Start backend
npm start                         # Start frontend
```

### Testing
```bash
go test ./...                     # Run Go tests
npm test                          # Run React tests
```

## 🌐 External Resources

- [Go Documentation](https://golang.org/doc/)
- [Gin Framework](https://gin-gonic.com/)
- [React Documentation](https://react.dev)
- [OpenCV Docs](https://docs.opencv.org/)
- [Docker Docs](https://docs.docker.com/)
- [Kubernetes Docs](https://kubernetes.io/docs/)

## 📝 Documentation Updates

The documentation was last updated on: **December 8, 2025**

Included topics:
- ✅ Project overview and quick start
- ✅ Complete installation guide for all platforms
- ✅ Docker and containerization details
- ✅ API reference with examples
- ✅ Development workflow and contributing guide
- ✅ Troubleshooting guides

## ❓ Can't Find What You Need?

1. Check this index
2. Use Ctrl+F to search documentation
3. Check [DEVELOPMENT.md - Getting Help](DEVELOPMENT.md#getting-help)
4. Review commit history for context
