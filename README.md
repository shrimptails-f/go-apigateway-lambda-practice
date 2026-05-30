# リポジトリ概要

このリポジトリは、Go 製の Lambda を API Gateway 配下で動かし、それを AWS CDK でデプロイするための PoC です。

今回やりたいことは次の 4 点です。

- Go で書いた Lambda を API Gateway から呼び出す
- AWS CDK で API Gateway と Lambda をまとめてデプロイする
- Lambda は ECR のコンテナイメージを使って起動する
- Bearer トークン認証を API Gateway か Lambda のどちらで持つか判断する
- Lambda はプライベートサブネットに配置し、ECRからイメージのダウンロードができること

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

前提:

- envにAWS 認証情報を設定していること
  - AWS_ACCESS_KEY_ID
  - AWS_SECRET_ACCESS_KEY

```bash
task infra:deploy:aws
```

デプロイされるもの:

- VPC
- public subnet
- private isolated subnet
- API Gateway
- ECR ベースの Lambda

# デプロイ後動作確認

※トークンベースの認証をPoCとして実装しています。このリポジトリではトークンの運用は扱わないので別途検討してください。
```bash
curl -H 'Authorization: Bearer Poc_tokens' https://<API_ID>.execute-api.ap-northeast-1.amazonaws.com/prod/users
```
