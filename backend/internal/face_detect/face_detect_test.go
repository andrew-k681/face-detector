package face_detect

import (
	"bytes"
	"image"
	"image/png"
	"testing"
)

func TestDetectFaces_InvalidImage(t *testing.T) {
	invalidData := []byte("not an image")
	_, _, err := DetectFaces(invalidData)
	if err == nil {
		t.Error("Expected error for invalid image data")
	}
}

func TestDetectFaces_EmptyData(t *testing.T) {
	_, _, err := DetectFaces([]byte{})
	if err == nil {
		t.Error("Expected error for empty data")
	}
}

func TestDetectFaces_ValidImage(t *testing.T) {
	// Create a simple valid image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// Set some pixels to make it a valid image
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			img.Set(x, y, image.White)
		}
	}

	// Encode to PNG
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		t.Fatalf("Failed to encode test image: %v", err)
	}

	imageData := buf.Bytes()

	// Test DetectFaces
	resultImg, faceCount, err := DetectFaces(imageData)
	if err != nil {
		t.Errorf("Unexpected error for valid image: %v", err)
	}
	if resultImg == nil {
		t.Error("Expected non-nil result image")
	}
	if faceCount < 0 {
		t.Errorf("Face count should not be negative, got %d", faceCount)
	}
	// Since it's a simple image, expect no faces
	if faceCount != 0 {
		t.Errorf("Expected 0 faces in test image, got %d", faceCount)
	}
}
