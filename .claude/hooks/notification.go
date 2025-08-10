package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	botToken = os.Getenv("CC_HOOK_TELEGRAM_BOT_TOKEN")
	chatID   = os.Getenv("CC_HOOK_TELEGRAM_CHAT_ID")
)


func sendTelegramMessage(message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	
	data := url.Values{}
	data.Set("chat_id", chatID)
	data.Set("text", message)
	data.Set("parse_mode", "HTML")
	
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(apiURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram API returned status %d", resp.StatusCode)
	}
	
	return nil
}


func formatMessage(rawData map[string]interface{}) string {
	// 프로젝트 이름 추출
	projectName := "Unknown"
	if cwd, ok := rawData["cwd"].(string); ok && cwd != "" && cwd != "Unknown" {
		projectName = filepath.Base(cwd)
	}
	
	// 이벤트 이름 추출
	eventName := "Unknown"
	if event, ok := rawData["hook_event_name"].(string); ok {
		eventName = event
	}

	eventMessage := "Unknown"
	if msg, ok := rawData["message"].(string); ok {
		eventMessage = msg
	}

	needPermission := strings.HasPrefix(eventMessage, "Claude needs your permission")

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	
	// 메시지 작성 (HTML 형식 사용)
	message := fmt.Sprintf("🤖 <b>Project: %s</b>\n", projectName)
	message += fmt.Sprintf("⏰ %s\n", timestamp)
	message += fmt.Sprintf("✅ Event: <code>%s</code>\n", eventName)
	if needPermission {
		message += fmt.Sprintf("📌 <b>Need Permission:</b> <code>%v</code>\n", needPermission)
	}
	
	return message
}

func sendErrorNotification(errorMsg string) {
	// 현재 디렉토리에서 프로젝트 이름 가져오기
	cwd, _ := os.Getwd()
	projectName := filepath.Base(cwd)
	
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	
	// HTML 형식 사용
	errorMessage := "🚨 <b>Hook Error Alert</b>\n"
	errorMessage += fmt.Sprintf("📁 <b>Project:</b> <code>%s</code>\n", projectName)
	errorMessage += fmt.Sprintf("⏰ %s\n", timestamp)
	errorMessage += fmt.Sprintf("❌ <b>Error Details:</b>\n<pre>%s</pre>", errorMsg)
	
	_ = sendTelegramMessage(errorMessage)
}

func main() {
	// stdin에서 JSON 데이터 읽기
	inputData, err := io.ReadAll(os.Stdin)
	if err != nil {
		sendErrorNotification(fmt.Sprintf("Failed to read stdin: %v", err))
		os.Exit(1)
	}
	
	// 빈 입력 처리
	if len(bytes.TrimSpace(inputData)) == 0 {
		os.Exit(0)
	}
	
	// JSON 파싱 - 전체 데이터를 먼저 map으로 파싱
	var rawData map[string]interface{}
	if err := json.Unmarshal(inputData, &rawData); err != nil {
		sendErrorNotification(fmt.Sprintf("JSON Decode Error: %v\nInput: %s", err, string(inputData)))
		os.Exit(1)
	}
	
	// 메시지 포맷팅 및 전송
	message := formatMessage(rawData)
	if err := sendTelegramMessage(message); err != nil {
		sendErrorNotification(fmt.Sprintf("Failed to send telegram message: %v", err))
		os.Exit(1)
	}
	
	os.Exit(0)
}