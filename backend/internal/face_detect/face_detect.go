package face_detect

import (
	"bytes"
	"errors"
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

const ClassifierPath = "./internal/classifiers/haarcascade_frontalface_default.xml"

func DetectFaces(imageData []byte) (image.Image, int, error) {
	// Convert to image.Image
	img, _, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		// c.JSON(http.StatusBadRequest, dto.FaceDetectionResponse{
		// 	Success: false,
		// 	Message: "Failed to decode image format: " + err.Error(),
		// })
		return nil, 0, err
	}

	// Convert image.Image to gocv.Mat
	mat, err := imageToMat(img)
	if err != nil {
		return nil, 0, err
	}
	defer mat.Close()

	// Load face classifier
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(ClassifierPath) {
		return nil, 0, errors.New("failed to load face classifier")
	}

	// Detect faces
	rects := classifier.DetectMultiScale(mat)
	faceCount := len(rects)
	// Draw rectangles around detected faces
	green := color.RGBA{R: 0, G: 255, B: 0, A: 255} // Green color
	for _, r := range rects {
		gocv.RectangleWithParams(&mat, r, green, 3, gocv.LineAA, 0)
	}

	// Convert back to image
	resultImg, err := matToImage(mat)
	if err != nil {
		return nil, 0, err
	}

	return resultImg, faceCount, nil
}

func imageToMat(img image.Image) (gocv.Mat, error) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Create RGBA image for efficient conversion
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	// Create Mat from RGBA data
	mat, err := gocv.NewMatFromBytes(height, width, gocv.MatTypeCV8UC4, rgba.Pix)
	if err != nil {
		return gocv.Mat{}, err
	}

	// Convert RGBA to BGR (OpenCV default)
	bgrMat := gocv.NewMat()
	gocv.CvtColor(mat, &bgrMat, gocv.ColorRGBAToBGR)
	mat.Close()

	return bgrMat, nil
}

func matToImage(mat gocv.Mat) (image.Image, error) {
	// Convert BGR back to RGBA
	rgbaMat := gocv.NewMat()
	gocv.CvtColor(mat, &rgbaMat, gocv.ColorBGRToRGBA)
	defer rgbaMat.Close()

	// Create image from Mat data
	img := image.NewRGBA(image.Rect(0, 0, rgbaMat.Cols(), rgbaMat.Rows()))
	img.Pix = rgbaMat.ToBytes()

	return img, nil
}
