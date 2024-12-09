# Slack Notifier

このリポジトリは、Slackに通知を送るためのGo言語によるCLIツールです。

## 必要な環境変数

- `SLACK_TOKEN`: SlackのBotトークン（例: `xoxb-...`）
- `SLACK_CHANNEL`: メッセージを送信したいチャンネルIDもしくはチャンネル名（`general`など）

これらが設定されていない場合、ツールはエラーを返します。

## インストール・ビルド

```bash
make build
```

## 実行例

export SLACK_TOKEN="xoxb-xxxxxxxxx-xxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxx"
export SLACK_CHANNEL="#general"

./slack-notifier "Hello Slack!"
