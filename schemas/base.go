package schema

import (
	"time"

	"github.com/google/uuid"
)

// TranslateOut defines the structure for the translated messages (can be a map or a specific struct)
type TranslateOut struct {
	En string `json:"en"` // You can add more languages if needed
}

// IResponseBase defines the structure for the base response
type IResponseBase struct {
	LogID   string      `json:"log_id"`
	Success int         `json:"success"`
	Code    string      `json:"code"`
	Message TranslateOut `json:"message"`
	Data    interface{} `json:"data,omitempty"` // Can hold any type of data
}

// NewIResponseBase creates a new IResponseBase with a generated UUID for log_id and default values for success and message.
func NewIResponseBase(code, message string, data interface{}) *IResponseBase {
	return &IResponseBase{
		LogID:   generateUUID(),
		Success: 1,  // success set to 1 by default
		Code:    code,
		Message: TranslateOut{En: message}, // You can extend this with other languages if needed
		Data:    data,
	}
}

// generateUUID generates a new UUID
func generateUUID() string {
	return uuid.New().String()
}

// GetCurrentDate returns the current UTC date (equivalent to getCurrentDate() in Python)
func GetCurrentDate() string {
	return time.Now().UTC().Format(time.RFC3339)
}
