## How-to run a local test?

1. Start your application
```sh
go run ./example/local-test/main.go
```

2. Send a request
```sh
curl --location --request POST 'http://localhost:8000/bigquery' \
      --header 'Content-Type: application/json' \
      -d '{
      "requestId": "124ab1c",
      "caller": "//bigquery.googleapis.com/projects/myproject/jobs/myproject:US.bquxjob_5b4c112c_17961fafeaf",
      "sessionUser": "",
      "userDefinedContext": {},
      "calls": [["TEST REQUEST"]]
  }'
```
