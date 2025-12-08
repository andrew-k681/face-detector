FROM gocv/opencv:4.8.0 AS backend-builder

RUN apt-get update && apt-get install -y golang-go && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o face-detection-server main.go

# Final image
FROM gocv/opencv:4.8.0

WORKDIR /app
COPY --from=backend-builder /app/face-detection-server .
COPY --from=backend-builder /app/internal/classifiers ./internal/classifiers

EXPOSE 8080
CMD ["./face-detection-server"]
