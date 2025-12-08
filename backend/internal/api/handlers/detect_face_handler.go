package handlers

import (
	"bytes"
	"encoding/base64"
	"face-detection-app/internal/api/dto"
	"face-detection-app/internal/face_detect"
	"fmt"
	"image/jpeg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DetectFaceHandler(c *gin.Context) {
	var request dto.FaceDetectionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.FaceDetectionResponse{
			Success: false,
			Message: "Invalid request: " + err.Error(),
		})
		return
	}

	// Decode base64 image
	imageData, err := base64.StdEncoding.DecodeString(request.ImageData)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.FaceDetectionResponse{
			Success: false,
			Message: "Failed to decode image: " + err.Error(),
		})
		return
	}

	// Detect faces
	resultImg, faceCount, err := face_detect.DetectFaces(imageData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.FaceDetectionResponse{
			Success: false,
			Message: "Face detection failed: " + err.Error(),
		})
		return
	}

	// Encode to base64
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, resultImg, nil); err != nil {
		c.JSON(http.StatusInternalServerError, dto.FaceDetectionResponse{
			Success: false,
			Message: "Failed to encode result image: " + err.Error(),
		})
		return
	}

	resultBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	c.JSON(http.StatusOK, dto.FaceDetectionResponse{
		Success:   true,
		ImageData: "data:image/jpeg;base64," + resultBase64,
		FaceCount: faceCount,
		Message:   fmt.Sprintf("Detected %d face(s)", faceCount),
	})
}
