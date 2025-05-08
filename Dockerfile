# ビルドステージ
FROM golang:1.24-alpine AS builder

WORKDIR /app

# 依存関係をコピー
COPY go.mod ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# 静的リンクでバイナリをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o slack-notifier .

# 実行ステージ
FROM alpine:3.19

# 証明書をインストール（HTTPS通信に必要）
RUN apk --no-cache add ca-certificates && \
    update-ca-certificates

WORKDIR /app

# ビルドステージからバイナリのみをコピー
COPY --from=builder /app/slack-notifier .

# 実行ユーザーを作成
RUN adduser -D appuser
USER appuser

# 環境変数の設定
ENV SLACK_TOKEN=""
ENV SLACK_CHANNEL=""

# コマンドライン引数を使ってスラック通知を送信
ENTRYPOINT ["/app/slack-notifier"]
