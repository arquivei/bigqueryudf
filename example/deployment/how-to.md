## How to deploy a new version of the cloud function?

1. Install gcloud dependency following GCP instructions [here](https://cloud.google.com/sdk/docs/install).

2. Setup PROJECT_ID and REGION environment variables
```sh
export PROJECT_ID="my_project"
export REGION="us-central1"
``` 

3. Run gcloud functions deploy command
```sh
gcloud "--project=${PROJECT_ID}" \
        functions deploy "example-bigquery-cloud-function" \
        "--set-env-vars=PROJECT_ID=${PROJECT_ID}" \
        "--entry-point=bigquery" \
        "--region=${REGION}" \
        --gen2 \
        --runtime=go120 \
        --memory=256MiB \
        --source=. \
        --trigger-http \
        --allow-unauthenticated \
        --ingress-settings=internal-only
```

## How to create a UDF (user defined funcition) in Big Query

1. Create a connection to your Big Query project
```sh
bq mk --connection --display_name="Connection to remote Cloud Function" --connection_type=CLOUD_RESOURCE --project_id=${PROJECT_ID} --location=US remote-function-connection
```

2. Verify if your connection was created with success
```sh
bq show --location=US --connection  remote-function-connection
```

3. Get service account from your connection
```sh
export BQ_SVC_ACCOUNT=`bq show --format=json --location=US --connection  remote-function-connection  | jq -r '.cloudResource.serviceAccountId'`

echo ${BQ_SVC_ACCOUNT}
```

4. Allow service account from connect to call your cloud function
```sh
gcloud functions add-iam-policy-binding example-bigquery-cloud-function \
  --member="serviceAccount:${BQ_SVC_ACCOUNT}" \
  --role="roles/cloudfunctions.invoker"
```

5. Create your UDF inside Big Query to call your cloud function
```sh
bq --format=json query \
    --dataset_id=$PROJECT_ID:dataset  \
    --location=US \
    --nouse_legacy_sql  \
    "CREATE OR REPLACE FUNCTION example_function(x STRING)
    RETURNS STRING REMOTE WITH CONNECTION `${PROJECT_ID}.US.remote-function-connection` OPTIONS (endpoint = 'https://${REGION}-${PROJECT_ID}.cloudfunctions.net/example-bigquery-cloud-function', user_defined_context = [])"
```

6. Use your function inside any query
```sql
    SELECT `your_project.your_dataset.example_function`("I'M A TEST")
```