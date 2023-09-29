package example

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/arquivei/bigqueryudf"
)

// This example deploys a Cloud Function inside Big Query. To run it, follow all instriuctions from README.md 
func init() {
	functions.HTTP("bigquery", bigqueryudf.NewHandler(transformationExample))
}

func transformationExample(input []byte) (any, error) {
	return "hello world", nil
}
