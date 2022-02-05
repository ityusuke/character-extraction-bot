package main

import (
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"fmt"
	"log"
	"os"
)

const (
	sampleImagePath = "./image/sample.jpeg"
)

func main() {
	ocr(sampleImagePath)
}

func ocr(fileName string) {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	detection, err := client.DetectWeb(ctx, image, nil)
	if err != nil {
		log.Fatalf("Failed to detect labels: %v", err)
	}
	matchImages := detection.GetPagesWithMatchingImages()
	if len(matchImages) == 0 {
		log.Println("failed to find image in web")
	}
	for _, matchImage := range matchImages {
		fmt.Println(matchImage.PageTitle)
		fmt.Println(matchImage.Url)
	}
}
