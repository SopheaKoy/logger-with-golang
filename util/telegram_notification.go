package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

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

// SendTelegramMessage sends a message to the configured Telegram chat
func SendTelegramMessage(message string) error {
	// Log the start of the message sending
	log.Println("Attempting to send Telegram message...")

	// Get Telegram bot token and chat ID from environment variables
	botToken := viper.GetString("TELEGRAM_BOT_TOKEN")
	chatID := viper.GetString("TELEGRAM_CHAT_ID")

	// Validate configuration
	if botToken == "" || chatID == "" {
		// Print detailed debug information
		if botToken == "" {
			log.Println("Error: TELEGRAM_BOT_TOKEN not configured")
		}
		if chatID == "" {
			log.Println("Error: TELEGRAM_CHAT_ID not configured")
		}
		return fmt.Errorf("telegram configuration incomplete")
	}

	// Prepare the request payload
	msgReq := TelegramMessageRequest{
		ChatID:    chatID,
		Text:      message,
		ParseMode: "HTML", // Enable HTML formatting if needed, can be removed
	}

	// Marshal the payload to JSON
	jsonPayload, err := json.Marshal(msgReq)
	if err != nil {
		log.Printf("Error marshaling JSON for Telegram: %v", err)
		return err
	}

	// Prepare the request URL
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	// Create a new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("Error creating Telegram request: %v", err)
		return err
	}

	// Set content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending Telegram message: %v", err)
		return err
	}
	defer resp.Body.Close()

	// Read response body for logging
	bodyBytes, _ := io.ReadAll(resp.Body)
	respBody := string(bodyBytes)

	// Check if the response status code is 200 (OK)
	if resp.StatusCode != http.StatusOK {
		log.Printf("Non-200 response from Telegram: %s\nResponse body: %s", resp.Status, respBody)
		return fmt.Errorf("telegram API error: %s", resp.Status)
	}

	// Log success for successful requests
	log.Println("Message successfully sent to Telegram!")

	// Optionally log the response for debugging
	if viper.GetBool("DEBUG") {
		log.Printf("Telegram API response: %s", respBody)
	}

	return nil
}
