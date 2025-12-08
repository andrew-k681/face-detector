# Development Guide

Comprehensive guide for developing, testing, and contributing to the Face Detection App.

## Table of Contents

- [Development Setup](#development-setup)
- [Project Architecture](#project-architecture)
- [Development Workflow](#development-workflow)
- [Testing](#testing)
- [Code Structure](#code-structure)
- [Debugging](#debugging)
- [Best Practices](#best-practices)
- [Contributing](#contributing)

## Development Setup

### Prerequisites

- See [INSTALLATION.md](INSTALLATION.md) for initial setup
- Familiarity with Go and React
- Git for version control
- Your preferred code editor (VS Code recommended)

### Quick Setup

```bash
# Clone the repository
git clone <repository-url>
cd face-detector

# Install backend dependencies
cd backend
go mod download

# Install frontend dependencies
cd ../frontend
npm install

# Back to root
cd ..
```

### IDE Setup

#### VS Code Recommended Extensions

```json
{
  "extensions": [
    "golang.Go",
    "dbaeumer.vscode-eslint",
    "esbenp.prettier-vscode",
    "ms-vscode.live-server",
    "ms-azuretools.vscode-docker"
  ]
}
```

#### VS Code Settings (.vscode/settings.json)

```json
{
  "go.lintTool": "golangci-lint",
  "go.lintOnSave": "package",
  "[go]": {
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "[javascript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode",
    "editor.formatOnSave": true
  },
  "prettier.printWidth": 100,
  "prettier.tabWidth": 2
}
```

## Project Architecture

### Backend Architecture

```
backend/
├── main.go                          # Application entry point
├── go.mod                          # Module definition
├── go.sum                          # Dependency checksums
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── detect_face_handler.go    # Face detection endpoint
│   │   │   └── health_handler.go         # Health check endpoint
│   │   └── dto/
│   │       ├── face_detection_request.go
│   │       └── face_detection_response.go
│   ├── face_detect/
│   │   └── face_detect.go                # OpenCV face detection logic
│   └── classifiers/
│       └── haarcascade_frontalface_default.xml
└── Dockerfile
```

### Frontend Architecture

```
frontend/
├── src/
│   ├── App.js              # Main component
│   ├── App.css             # App styles
│   ├── index.js            # React entry point
│   ├── index.css           # Global styles
├── public/
│   └── index.html          # HTML template
├── package.json            # Dependencies
├── Dockerfile
└── nginx.conf
```

### Key Design Patterns

1. **Handler Pattern** (Backend)
   - HTTP request handlers process requests
   - DTO pattern for request/response serialization
   - Error handling with structured responses

2. **Component Pattern** (Frontend)
   - React functional components
   - State management with hooks
   - CSS modules for styling

3. **Service Architecture**
   - Frontend and backend as separate services
   - Communication via REST API
   - CORS enabled for cross-origin requests

## Development Workflow

### Starting Development Servers

#### Option 1: Using run.sh (Local)

```bash
./run.sh
```

- Frontend: `http://localhost:3000`
- Backend: `http://localhost:8080`

#### Option 2: Manual Start

Terminal 1 (Backend):
```bash
cd backend
go run main.go
```

Terminal 2 (Frontend):
```bash
cd frontend
npm start
```

#### Option 3: Docker Development

```bash
docker-compose up -d --build
```

### Making Code Changes

#### Backend Changes

1. Edit files in `backend/`
2. Backend will auto-reload if using air watch (optional)
3. Or manually restart: `docker-compose restart backend`

```bash
# Install air for auto-reload (optional)
go install github.com/cosmtrek/air@latest

# Run with air (from backend directory)
air
```

#### Frontend Changes

1. Edit files in `frontend/src/`
2. Changes auto-refresh in browser (hot reload enabled)
3. Check browser console for errors

### Git Workflow

```bash
# Create feature branch
git checkout -b feature/your-feature-name

# Make changes and commit
git add .
git commit -m "feat: description of changes"

# Push to remote
git push origin feature/your-feature-name

# Create Pull Request on GitHub
```

### Commit Message Convention

Follow conventional commits:

```
feat: add new feature
fix: fix a bug
docs: documentation changes
style: code formatting
refactor: code restructuring
test: add tests
chore: maintenance tasks
```

## Testing

### Backend Testing

#### Unit Tests

```bash
cd backend

# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run specific test
go test -v ./internal/api/handlers -run TestDetectFaceHandler

# Run with coverage
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

#### Example Test Structure

```go
package handlers

import (
    "testing"
)

func TestHealthHandler(t *testing.T) {
    // Test implementation
}
```

### Frontend Testing

#### Unit Tests

```bash
cd frontend

# Run tests in watch mode
npm test

# Run tests once
npm test -- --watchAll=false

# Run with coverage
npm test -- --coverage
```

#### Example Test Structure

```javascript
import { render, screen } from '@testing-library/react';
import App from './App';

test('renders face detection app', () => {
  render(<App />);
  expect(screen.getByText(/start camera/i)).toBeInTheDocument();
});
```

### Integration Testing

```bash
# Start both services
docker-compose up -d --build

# Test health endpoints
curl http://localhost:8080/api/health

# Test face detection API
curl -X POST http://localhost:8080/api/detect-face \
  -H "Content-Type: application/json" \
  -d '{"imageData":"base64_encoded_image_data"}'
```

## Code Structure

### Backend Code Organization

#### handlers/detect_face_handler.go

```go
func DetectFaceHandler(c *gin.Context) {
    // 1. Parse request
    // 2. Decode image
    // 3. Detect faces using OpenCV
    // 4. Return response
}
```

#### face_detect/face_detect.go

```go
func DetectFaces(imageData []byte) ([]Face, error) {
    // Core face detection logic
    // Uses OpenCV Haar Cascade classifier
}
```

#### dto/

```go
type FaceDetectionRequest struct {
    ImageData string `json:"imageData"`
}

type FaceDetectionResponse struct {
    Success   bool   `json:"success"`
    ImageData string `json:"imageData"`
    FaceCount int    `json:"faceCount"`
}
```

### Frontend Code Organization

#### App.js

```javascript
function App() {
  const [cameraActive, setCameraActive] = useState(false);
  const [photo, setPhoto] = useState(null);
  const [result, setResult] = useState(null);

  const startCamera = async () => { /* ... */ };
  const capturePhoto = () => { /* ... */ };
  const detectFace = async () => { /* ... */ };

  return (
    <div className="App">
      {/* JSX content */}
    </div>
  );
}
```

## Debugging

### Backend Debugging

#### VS Code Debugging

1. Install Go extension
2. Create `.vscode/launch.json`:

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Connect to Server",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "${workspaceFolder}/backend",
      "env": {},
      "args": []
    }
  ]
}
```

3. Set breakpoints and press F5

#### Console Logging

```go
import "log"

log.Printf("Debug: %v", variable)
log.Fatalf("Error: %v", err)
```

### Frontend Debugging

#### Browser DevTools

1. Press F12 to open DevTools
2. Use Console tab for logs:
   ```javascript
   console.log('Debug:', variable);
   console.error('Error:', error);
   ```

3. Network tab to inspect API calls
4. React DevTools extension for component inspection

#### VS Code Debugging

1. Install "Debugger for Chrome" extension
2. Add to `.vscode/launch.json`:

```json
{
  "type": "chrome",
  "request": "attach",
  "name": "Attach to Chrome",
  "port": 9222,
  "pathMapping": {
    "/": "${workspaceRoot}/",
    "/src": "${workspaceRoot}/src"
  }
}
```

## Best Practices

### Go/Backend Best Practices

1. **Error Handling**
   ```go
   if err != nil {
       log.Printf("Error: %v", err)
       return nil, err
   }
   ```

2. **Interface Design**
   ```go
   type FaceDetector interface {
       DetectFaces(imageData []byte) ([]Face, error)
   }
   ```

3. **Dependency Injection**
   ```go
   type Handler struct {
       detector FaceDetector
   }
   ```

4. **Testing**
   - Write tests alongside code
   - Use table-driven tests for multiple scenarios
   - Mock external dependencies

### React/Frontend Best Practices

1. **Functional Components**
   ```javascript
   const MyComponent = ({ prop1, prop2 }) => {
     return <div>{prop1}</div>;
   };
   ```

2. **Custom Hooks**
   ```javascript
   const useCamera = () => {
     const [stream, setStream] = useState(null);
     // Hook logic
     return { stream, startCamera, stopCamera };
   };
   ```

3. **Error Boundaries**
   ```javascript
   class ErrorBoundary extends React.Component {
     // Error handling logic
   }
   ```

4. **Prop Validation**
   ```javascript
   MyComponent.propTypes = {
     prop1: PropTypes.string.isRequired,
   };
   ```

### Code Quality

1. **Formatting**
   - Go: `gofmt` (automatic)
   - JavaScript: `prettier` (configured in VS Code)

2. **Linting**
   ```bash
   # Backend
   golangci-lint run ./...

   # Frontend
   npm run lint
   ```

3. **Type Safety**
   - Go: Strongly typed by default
   - JavaScript: Consider TypeScript for large projects

## Contributing

### Contribution Process

1. **Fork and Clone**
   ```bash
   git clone <your-fork-url>
   cd face-detector
   ```

2. **Create Feature Branch**
   ```bash
   git checkout -b feature/your-feature
   ```

3. **Make Changes**
   - Write code following project conventions
   - Add tests for new functionality
   - Update documentation

4. **Commit Changes**
   ```bash
   git add .
   git commit -m "feat: description"
   ```

5. **Push and Create PR**
   ```bash
   git push origin feature/your-feature
   ```

6. **Code Review**
   - Address review comments
   - Ensure CI passes
   - Merge when approved

### Code Review Checklist

- [ ] Code follows project style guide
- [ ] Changes are well documented
- [ ] New features have tests
- [ ] No breaking changes (or documented)
- [ ] Commit messages are clear
- [ ] All tests pass

### Documentation Updates

When adding features, update:

1. **Code comments** - Explain complex logic
2. **README.md** - Update if public API changes
3. **API.md** - Document new endpoints
4. **DEVELOPMENT.md** - Update if development process changes

## Common Tasks

### Adding a New API Endpoint

1. Create handler in `backend/internal/api/handlers/`
2. Define DTO in `backend/internal/api/dto/`
3. Register route in `backend/main.go`
4. Add tests
5. Update [API.md](API.md)

Example:
```go
// In main.go
api.POST("/new-endpoint", handlers.NewHandler)

// Create handlers/new_handler.go
func NewHandler(c *gin.Context) {
    // Implementation
}
```

### Adding a New React Component

1. Create component in `frontend/src/components/`
2. Add styling in `frontend/src/components/ComponentName.css`
3. Export from component file
4. Import and use in parent component
5. Add tests

Example:
```javascript
// frontend/src/components/MyComponent.js
export const MyComponent = ({ props }) => {
  return <div>{props}</div>;
};
```

### Database Integration (Future)

When adding database support:

1. Choose database (PostgreSQL recommended)
2. Create migrations
3. Define models/entities
4. Implement repositories
5. Update API endpoints to use persistence

## Performance Optimization

### Backend Optimization

- Cache OpenCV classifier in memory
- Batch process multiple images
- Use goroutines for concurrent requests
- Profile with pprof

```bash
# CPU profiling
go tool pprof http://localhost:8080/debug/pprof/profile
```

### Frontend Optimization

- Code splitting with React.lazy
- Memoization with React.memo
- Image optimization
- Lazy load components

```javascript
const DetectionResult = React.lazy(() => import('./DetectionResult'));
```

## Resources

- [Go Documentation](https://golang.org/doc/)
- [Gin Framework](https://gin-gonic.com/)
- [React Documentation](https://react.dev)
- [OpenCV Documentation](https://docs.opencv.org/)
- [Docker Documentation](https://docs.docker.com/)

## Getting Help

- Check existing issues and documentation
- Review commit history for similar changes
- Ask in project discussions
- Create detailed bug reports
