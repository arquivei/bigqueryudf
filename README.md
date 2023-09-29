# bigqueryudf
An Cloud Function interface wrapper that could be used as UDF in Big Query

This implementation uses http interface and bigquery entities expected inside a Cloud Function

TL;DR

```go
func init() {
	functions.HTTP("bigquery", bigqueryudf.NewHandler(transformationExample))
}

func transformationExample(input []byte) (any, error) {
	return "hello world", nil
}
```

To understand more about how to use this project, there are some examples
- [Local run for tests](https://github.com/arquivei/bigqueryudf/tree/main/examples/local-test)
- [Cloud function deployment](https://github.com/arquivei/bigqueryudf/tree/main/examples/deployment)

Comments, discussions, issues and pull-requests are welcomed.
