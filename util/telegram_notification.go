package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// init automatically reads environment variables when the package is imported
func init() {
	// Automatically read environment variables and replace '.' with '_'
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

// TelegramMessageRequest represents the request body for sending a message
type TelegramMessageRequest struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

type BodyParams struct {
	System 			string
	Method          string
	Status          int
	Endpoint        string
	ResponseMessage string
	LogID           string
	RequestID       string
	IP              string
	UserAgent       string
	Duration        float64
	RequestBody     string
	ResponseBody    string
}

func SendTelegramMessage(message BodyParams) error {
	botToken := viper.GetString("TELEGRAM_BOT_TOKEN")
	chatID   := viper.GetString("TELEGRAM_CHAT_ID")

	// Validate required configuration
	if botToken == "" {
		return fmt.Errorf("TELEGRAM_BOT_TOKEN not configured")
	}
	if chatID == "" {
		return fmt.Errorf("TELEGRAM_CHAT_ID not configured")
	}

	// Format the message text
	messageText := formatNotificationText(message)

	// Prepare the request payload
	messageSend := TelegramMessageRequest{
		ChatID	  : chatID,
		Text	  : messageText,
		ParseMode : "HTML",
	}

	jsonPayload, err := json.Marshal(messageSend)
	if err != nil {
		return fmt.Errorf("error marshaling JSON for Telegram: %v", err)
	}

	// Create a new request
	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken), bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("error creating Telegram request: %v", err)
	}

	// Set content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Send the request with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending Telegram message: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}
	respBody := string(bodyBytes)

	// Check if the response status code is 200 (OK)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram API error: %s, response: %s", resp.Status, respBody)
	}

	return nil
}

// SendTelegramMessageText sends a simple text message (renamed to avoid conflict)
func SendTelegramMessageText(text string) error {
	botToken := viper.GetString("TELEGRAM_BOT_TOKEN")
	chatID := viper.GetString("TELEGRAM_CHAT_ID")

	if botToken == "" || chatID == "" {
		return fmt.Errorf("telegram configuration missing")
	}

	messageSend := TelegramMessageRequest{
		ChatID: chatID,
		Text:   text,
	}

	jsonPayload, err := json.Marshal(messageSend)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending message: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("telegram API error: %s", resp.Status)
		}
		return fmt.Errorf("telegram API error: %s, response: %s", resp.Status, string(bodyBytes))
	}

	return nil
}

// formatNotificationText formats the BodyParams into a readable message
func formatNotificationText(params BodyParams) string {
	// Get current time in UTC+7
	timestamp := time.Now().In(time.FixedZone("UTC+7", 7*60*60)).Format("02/01/2006 15:04:05")
	
	// Get system and environment from viper
	system := strings.ToUpper(viper.GetString("PROJECT_NAME"))
	if system == "" {
		system = "N/A"
	}
	env := strings.ToUpper(viper.GetString("ENV"))
	if env == "" {
		env = "N/A"
	}

	// Format the message with HTML tags for better formatting
	message := fmt.Sprintf(
		"<code>DATE         : 🕒 %s</code>\n"+
			"<code>SYSTEM       : %s</code>\n"+
			"<code>ENVIRONMENT  : %s</code>\n"+
			"<code>METHOD       : %s</code>\n"+
			"<code>STATUS       : %s %d</code>\n"+
			"<code>ENDPOINT     : %s</code>\n"+
			"<code>MESSAGE      : %s</code>",
		timestamp,
		system,
		env,
		params.Method,
		getColorForStatus(params.Status), params.Status,
		params.Endpoint,
		params.ResponseMessage,
	)

	return message
}

// getStatusEmoji returns emoji based on HTTP status
func getColorForStatus(status int) string {
	switch {
	case status >= 200 && status < 300:
		return "🟢"
	case status >= 300 && status < 400:
		return "🔵"
	case status >= 400 && status < 600:
		return "🔴"
	default:
		return "🟡"
	}
}