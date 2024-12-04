package ai

import (
	"fmt"
	"image"
	"os"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
)

var model *onnx.Model
var backend *gorgonnx.Graph

//Ref: https://github.com/oramasearch/onnx-go/blob/v0.5.0/examples/tiny_yolov2/main.go#L258

// Initialize the model by loading it
func InitModel(modelPath string) error {
	// Read the model file using os.ReadFile since the other way is depricated
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Current working directory:", cwd)

	modelData, err := os.ReadFile(modelPath)
	if err != nil {
		return fmt.Errorf("failed to read ONNX model file: %v", err)
	}

	// Create a backend receiver
	backend = gorgonnx.NewGraph()

	// Create a model and set the execution backend
	model = onnx.NewModel(backend)

	// Ensure the model is loaded and ready
	if model == nil {
		return fmt.Errorf("model is not loaded")
	}

	// Decode the model data into the ONNX model
	err = model.UnmarshalBinary(modelData)
	if err != nil {
		return fmt.Errorf("failed to unmarshal ONNX model: %v", err)
	}

	// Return nil on successful initialization
	return nil
}

// PredictFood takes an image and returns the predicted food class
func PredictFood(img image.Image) (string, error) {

	// Preprocess the image into the required tensor format and resize it
	inputTensor, err := PreprocessImage(img)

	if err != nil {
		return "", fmt.Errorf("image preprocessing failed: %v", err)
	}

	// Ensure that inputTensor is not nil before using it
	if inputTensor == nil {
		return "", fmt.Errorf("the input tensor is nil")
	}

	// Ensure that the tensor is correctly formatted (batch_size, channels, height, width)
	if inputTensor.Shape()[0] != 1 || inputTensor.Shape()[1] != 3 || inputTensor.Shape()[2] != 224 || inputTensor.Shape()[3] != 224 {
		return "", fmt.Errorf("input tensor shape is invalid. Expected [1, 3, 224, 224], got: %v", inputTensor.Shape())
	}

	// fmt.Println("Model input shape:", model)
	// fmt.Printf("Input tensor shape: %+v, data type: %v\n", inputTensor.Shape(), inputTensor.Dtype())

	// Run prediction with the model
	model.SetInput(0, inputTensor)
	err = backend.Run()
	if err != nil {
		return "", fmt.Errorf("prediction failed: %v", err)
	}
	fmt.Println("Made it here")

	output, err := model.GetOutputTensors()
	if err != nil || len(output) == 0 {
		return "", fmt.Errorf("failed to retrieve model outputs")
	}

	fmt.Println("Made it here 2")
	// write the first output to stdout
	fmt.Println(output[0])

	// Decode the output into a class label
	var predictionLabel string = output[0].String()

	return predictionLabel, nil
}
