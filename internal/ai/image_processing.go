package ai

import (
	// "bytes"
	"fmt"
	"image"

	// "image/jpeg"
	"golang.org/x/image/draw"
	"gorgonia.org/tensor"
)

// PreprocessImage resizes and normalizes the image for the model
func PreprocessImage(img image.Image) (tensor.Tensor, error) {
	// Resize the image to 224x224 pixels
	resizedImg := image.NewRGBA(image.Rect(0, 0, 224, 224))

	// Use `draw.CatmullRom` for high-quality resizing
	draw.CatmullRom.Scale(resizedImg, resizedImg.Bounds(), img, img.Bounds(), draw.Over, nil)

	// Normalize the image and create a tensor
	inputTensor, err := ConvertImageToTensor(resizedImg)
	if err != nil {
		return nil, fmt.Errorf("failed to convert image to tensor: %v", err)
	}

	return inputTensor, nil
}

// ConvertImageToTensor normalizes image pixels and creates a tensor
func ConvertImageToTensor(img image.Image) (tensor.Tensor, error) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Ensure the image is the correct size (224x224)
	if width != 224 || height != 224 {
		return nil, fmt.Errorf("image must be 224x224, but got %dx%d", width, height)
	}

	// Create a tensor for the model input (CHW format)
	data := make([]float32, 3*width*height)
	index := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			// Normalize pixel values to [0, 1] range by dividing by 255.0
			data[index] = float32(r>>8) / 255.0
			data[index+1] = float32(g>>8) / 255.0
			data[index+2] = float32(b>>8) / 255.0
			index += 3
		}
	}

	// Normalize the tensor - same as the pytorch model used
	mean := []float32{0.485, 0.456, 0.406}
	std := []float32{0.229, 0.224, 0.225}
	NormalizeTensor(data, mean, std)

	// Create a 4D tensor: (Batch, Channels, Height, Width)
	inputDense := tensor.New(
		tensor.WithShape(1, 3, height, width),
		tensor.Of(tensor.Float32),
		tensor.WithBacking(data),
	)

	// Return the tensor.Dense (which implements tensor.Tensor)
	return inputDense, nil
}

// NormalizeTensor normalizes the tensor data with the given mean and std
func NormalizeTensor(tensorData []float32, mean, std []float32) {
	channelSize := len(tensorData) / 3 // Each channel's size
	for c := 0; c < 3; c++ {           // Iterate over channels (RGB)
		for i := 0; i < channelSize; i++ {
			index := c*channelSize + i
			tensorData[index] = (tensorData[index] - mean[c]) / std[c]
		}
	}
}
