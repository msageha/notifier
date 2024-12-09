package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var slackAPIURL = "https://slack.com/api/chat.postMessage"

type slackMessage struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func main() {
	slackToken := os.Getenv("SLACK_TOKEN")
	slackChannel := os.Getenv("SLACK_CHANNEL")

	if slackToken == "" {
		log.Fatal("環境変数 SLACK_TOKEN が設定されていません。")
	}
	if slackChannel == "" {
		log.Fatal("環境変数 SLACK_CHANNEL が設定されていません。")
	}

	if len(os.Args) < 2 {
		log.Fatalf("使用方法: %s [メッセージ本文]", os.Args[0])
	}

	message := os.Args[1]

	err := postToSlack(slackToken, slackChannel, message)
	if err != nil {
		log.Fatalf("Slackへのメッセージ送信に失敗しました: %v", err)
	}

	fmt.Println("Slackへのメッセージ送信に成功しました")
}

func postToSlack(token, channel, text string) error {
	msg := slackMessage{
		Channel: channel,
		Text:    text,
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", slackAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Slack API returned status code %d", resp.StatusCode)
	}

	var result struct {
		OK bool `json:"ok"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if !result.OK {
		return fmt.Errorf("Slack API returned ok=false")
	}

	return nil
}
