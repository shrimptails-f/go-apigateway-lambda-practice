# Architecture

```mermaid
flowchart LR
    Client[Client] --> APIGW[API Gateway]
    APIGW --> Lambda[Lambda backend-user]

    subgraph VPC
        Lambda
        EcrApi[ECR API Endpoint]
        EcrDkr[ECR DKR Endpoint]
        S3[S3 Gateway Endpoint]
    end

    ECR[ECR backend-user] --> Lambda
```
