package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/arquivei/bigqueryudf"
)

func init() {
	functions.HTTP("bigquery", bigqueryudf.NewHandler(transformationExample))
}

// This example creates a Cloud Function using the functions framework. To run it, follow all instriuctions from how-to.md
func main() {
	port := "8000"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}


func transformationExample(input []byte) (any, error) {
	return "hello world", nil
}
