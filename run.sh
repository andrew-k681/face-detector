#!/bin/bash

# Face Detection App - Startup Script

echo "Starting Face Detection App..."
echo ""

# Check if OpenCV is available
if ! command -v pkg-config &> /dev/null || ! pkg-config --exists opencv4; then
    echo "Warning: OpenCV may not be installed. Please install OpenCV first."
    echo "macOS: brew install opencv"
    echo "Linux: sudo apt-get install libopencv-dev"
    echo ""
fi

# Start backend in background
echo "Starting Go backend server..."
cd backend
go run main.go &
BACKEND_PID=$!
cd ..

# Wait a moment for backend to start
sleep 2

# Start frontend
echo "Starting React frontend..."
cd frontend
npm start &
FRONTEND_PID=$!
cd ..

echo ""
echo "Servers started!"
echo "Backend PID: $BACKEND_PID"
echo "Frontend PID: $FRONTEND_PID"
echo ""
echo "Press Ctrl+C to stop all servers"

# Wait for user interrupt
trap "kill $BACKEND_PID $FRONTEND_PID 2>/dev/null; exit" INT TERM
wait

