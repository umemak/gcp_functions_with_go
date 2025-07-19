# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## プロジェクト概要

このリポジトリは、Go言語を使用したGoogle Cloud Platform (GCP) Cloud Functionsのサンプル集です。各ディレクトリが独立したCloud Function実装で、以下のGCPサービスとの連携パターンを実演します：

- BigQuery（データ書き込み）
- Firestore（ドキュメント保存）
- Cloud Storage（ファイル書き込み・イベント処理）
- Cloud Pub/Sub（メッセージ処理）

## ディレクトリ構造

```
/bq_out/    - Pub/SubメッセージをBigQueryに書き込み
/fs_out/    - Pub/SubメッセージをFirestoreに保存
/gs_out/    - HTTPリクエストからCloud Storageにファイル書き込み
/bucket/    - Cloud Storageイベントをログ出力
/http/      - シンプルなHTTP Cloud Function
/pubsub/    - 基本的なPub/Subメッセージ処理
```

## 開発コマンド

### ビルド
各ディレクトリは独立したGoモジュールです：
```bash
cd <function-directory>
go build
```

### 依存関係管理
```bash
go mod tidy      # 依存関係の整理
go mod download  # 依存関係のダウンロード
```

### 実行・テスト
```bash
go run main.go   # ローカル実行（Functions Framework使用時）
```

## アーキテクチャ

### パッケージ構造
- 全ての関数で統一されたパッケージ名「functions」を使用
- 各ディレクトリに独立した`go.mod`ファイル
- Go 1.23.0を使用

### データ構造
- 共通の`Info`型構造体を使用（`Name`、`Value`フィールド）
- BigQueryスキーマは`bq_out/names.json`で定義

### トリガータイプ
- HTTP: `/http/`, `/gs_out/`
- Pub/Sub: `/bq_out/`, `/fs_out/`, `/pubsub/`
- Cloud Storage Events: `/bucket/`

### 環境変数
各関数で使用される環境変数：
- `GCP_PROJECT`: GCPプロジェクトID
- `BUCKET_NAME`: Cloud Storageバケット名
- `COLLECTION_NAME`: Firestoreコレクション名
- その他、各サービス固有の設定

## コード規約

### エラーハンドリング
- 統一されたエラーレスポンス形式
- 日本語エラーメッセージ
- 適切なHTTPステータスコード設定

### ログ出力
- 日本語ログメッセージ
- 構造化ログ（`log.Printf`使用）
- エラー時の詳細情報出力

## Cloud Functions固有の注意点

### デプロイメント
- 各ディレクトリを個別にデプロイ
- `main.go`にエントリーポイント関数定義
- 環境変数による設定外部化

### リソース設定
- メモリ、タイムアウト設定をGCPコンソールで調整
- IAM権限の適切な設定が必要（BigQuery、Firestore、Cloud Storage）

### 依存関係
- Dependabotによる自動更新が設定済み
- GCP公式ライブラリを使用
- セキュリティアップデートは定期的に適用