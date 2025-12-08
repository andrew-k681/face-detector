package dto

type FaceDetectionResponse struct {
	Success   bool   `json:"success"`
	ImageData string `json:"imageData"`
	FaceCount int    `json:"faceCount"`
	Message   string `json:"message,omitempty"`
}
