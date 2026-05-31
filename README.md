# リポジトリ概要

このリポジトリは、Go 製の Lambda を API Gateway 配下で動かし、それを AWS CDK でデプロイするための PoC です。

今回やりたいことは次の 5 点です。

- [Go で書いた Lambda を API Gateway から呼び出す](./docs/goals/01-call-go-lambda-from-api-gateway.md)
- [AWS CDK で API Gateway と Lambda をまとめてデプロイする](./docs/goals/02-deploy-api-gateway-and-lambda-with-cdk.md)
- [Lambda は ECR のコンテナイメージを使って起動する](./docs/goals/03-run-lambda-from-ecr-image.md)
- [Bearer トークン認証を API Gateway の Lambda authorizer で行う](./docs/goals/04-bearer-token-authentication-at-api-gateway.md)
- [Lambda はプライベートサブネットに配置し、ECRからイメージのダウンロードができること](./docs/goals/05-put-lambda-in-private-subnet-and-enable-ecr-image-download.md)

## 環境構築手順

1. ソースをクローン
2. .envをコピー
```bash
cp .devcontainer/.env.sample .devcontainer/.env
```
3. VsCodeでプロジェクトフォルダーを開く
4. Reopen in Containerを押下
  Ctrl Shift P → Reopen in containerと入力して実行

## ディレクトリ構成

[こちら](./docs/structure.md)を参照してください。

## アーキテクチャ

概要レベルですが、[こちら](./infra/README.md)を参照してください。

## デプロイ手順

- デプロイされるもの
  - VPC
  - public subnet
  - private isolated subnet
  - API Gateway
  - Lambda authorizer
  - ECR ベースの Lambda

### 実AWSの場合

#### 前提

- envにAWS 認証情報を設定していること
  - AWS_ACCESS_KEY_ID
  - AWS_SECRET_ACCESS_KEY
- ECR に `backend-user` リポジトリを作成済みであること

#### デプロイコマンド実行

```bash
# イメージのプッシュ
docker build -f build/user/Dockerfile -t backend-user:1 .
docker tag backend-user:1 <アカウントID>.dkr.ecr.ap-northeast-1.amazonaws.com/backend-user:1
docker push <アカウントID>.dkr.ecr.ap-northeast-1.amazonaws.com/backend-user:1

# デプロイ
task infra:deploy:aws
```

#### デプロイ後動作確認

```bash
curl -i -H 'Authorization: Bearer Poc_tokens' https://<API_ID>.execute-api.ap-northeast-1.amazonaws.com/prod/users
curl -i -H 'Authorization: Bearer invalid' https://<API_ID>.execute-api.ap-northeast-1.amazonaws.com/prod/users
```
