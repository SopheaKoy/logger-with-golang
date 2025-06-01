package middlewares

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// ErrorLevel represents the severity level of an error
type ErrorLevel string

const (
	None     ErrorLevel = "NONE"
	Low      ErrorLevel = "LOW"
	Medium   ErrorLevel = "MEDIUM"
	High     ErrorLevel = "HIGH"
	Critical ErrorLevel = "CRITICAL"
)

// StyleModifier contains ANSI color codes for terminal output
type StyleModifier struct {
	Blue          string
	Green         string
	Red           string
	Yellow        string
	TextBold      string
	TextUnderline string
	EndC          string
}

// NewStyleModifier creates a new StyleModifier instance
func NewStyleModifier() *StyleModifier {
	return &StyleModifier{
		Blue:          "\033[94m",
		Green:         "\033[92m",
		Red:           "\033[91m",
		Yellow:        "\033[93m",
		TextBold:      "\033[1m",
		TextUnderline: "\033[4m",
		EndC:          "\033[0m",
	}
}

// Logger is the main logger struct
type Logger struct {
	errorLogger *ErrorLogger
	style       *StyleModifier
	consoleLog  *log.Logger
	fileLog     *log.Logger
}

// ErrorLogger handles error logging with different severity levels
type ErrorLogger struct {
	logger *Logger
}

// New creates a new Logger instance
func New() *Logger {
	style := NewStyleModifier()
	
	// Create console logger
	consoleLog := log.New(os.Stdout, "", log.LstdFlags)
	
	// Ensure logs directory exists
	logsDir := "logs"
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		log.Fatal("Failed to create logs directory:", err)
	}
	
	// Create file logger
	logFile := filepath.Join(logsDir, "app.log")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	
	// Create file logger without colors
	fileLog := log.New(file, "", log.LstdFlags)
	
	return &Logger{
		errorLogger: &ErrorLogger{},
		style:       style,
		consoleLog:  consoleLog,
		fileLog:     fileLog,
	}
}

// Error returns the error logger
func (l *Logger) Error() *ErrorLogger {
	l.errorLogger.logger = l
	return l.errorLogger
}

// logError is a private method to handle error logging
func (e *ErrorLogger) logError(args ...interface{}) string {
	msg := fmt.Sprint(args...)
	// Log to console with colors
	e.logger.consoleLog.Printf("%s%s%s", e.logger.style.Red, msg, e.logger.style.EndC)
	// Log to file without colors
	e.logger.fileLog.Printf("ERROR: %s", msg)
	return msg
}

// Low logs a low severity error
func (e *ErrorLogger) Low(args ...interface{}) {
	msg := e.logError(args...)
	sendError(Low, msg)
}

// Medium logs a medium severity error
func (e *ErrorLogger) Medium(args ...interface{}) {
	msg := e.logError(args...)
	sendError(Medium, msg)
}

// High logs a high severity error
func (e *ErrorLogger) High(args ...interface{}) {
	msg := e.logError(args...)
	sendError(High, msg)
}

// Critical logs a critical severity error
func (e *ErrorLogger) Critical(args ...interface{}) {
	msg := e.logError(args...)
	sendError(Critical, msg)
}

// Info logs an informational message
func (l *Logger) Info(args ...interface{}) {
	msg := fmt.Sprint(args...)
	// Log to console with colors
	l.consoleLog.Printf("%s%s%s", l.style.Blue, msg, l.style.EndC)
	// Log to file without colors
	l.fileLog.Printf("INFO: %s", msg)
}

// Warn logs a warning message
func (l *Logger) Warn(args ...interface{}) {
	msg := fmt.Sprint(args...)
	// Log to console with colors
	l.consoleLog.Printf("%s%s%s", l.style.Yellow, msg, l.style.EndC)
	// Log to file without colors
	l.fileLog.Printf("WARN: %s", msg)
}

// Track logs a tracking message
func (l *Logger) Track(args ...interface{}) {
	msg := fmt.Sprint(args...)
	// Log to console with colors
	l.consoleLog.Printf("%s%s%s", l.style.Green, msg, l.style.EndC)
	// Log to file without colors
	l.fileLog.Printf("TRACK: %s", msg)
}

// sendError is a helper function to send errors (implementation depends on your needs)
func sendError(level ErrorLevel, msg string) {
	// Implement your error sending logic here
	log.Printf("Sending error - Level: %s, Message: %s", level, msg)
}

// Global logger instance
var DefaultLogger = New() 