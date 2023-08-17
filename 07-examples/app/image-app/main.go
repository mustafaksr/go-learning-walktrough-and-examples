package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"os"

	"github.com/fogleman/gg"
)

func rotateImage(imagePath string, degrees float64) (image.Image, error) {
	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	// Create a new context for drawing
	dc := gg.NewContextForImage(img)

	// Rotate the image
	dc.RotateAbout(gg.Radians(degrees), float64(img.Bounds().Dx())/2, float64(img.Bounds().Dy())/2)
	dc.DrawImage(img, 0, 0)

	return dc.Image(), nil
}

func handleRotateImage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, _, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Failed to read uploaded file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		degreesStr := r.FormValue("degrees")
		var degrees float64
		fmt.Sscanf(degreesStr, "%f", &degrees)

		rotatedImage, err := rotateImage("input.jpg", degrees)
		if err != nil {
			http.Error(w, "Failed to rotate image", http.StatusInternalServerError)
			return
		}

		// Set the content type as JPEG
		w.Header().Set("Content-Type", "image/jpeg")

		// Encode and write the rotated image to the response writer
		jpeg.Encode(w, rotatedImage, nil)
		return
	}

	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", handleRotateImage)

	fmt.Println("Server started on :9090")
	http.ListenAndServe(":9090", nil)
}
