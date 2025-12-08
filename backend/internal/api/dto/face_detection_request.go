package dto

type FaceDetectionRequest struct {
	ImageData string `json:"imageData" binding:"required"`
}
