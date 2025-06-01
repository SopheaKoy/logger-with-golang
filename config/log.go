package config

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/sirupsen/logrus"
)

// ANSI color codes
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[91m"
	ColorGreen  = "\033[92m"
	ColorYellow = "\033[93m"
	ColorBlue   = "\033[94m"
)

// Severity levels
type ErrorLevel string

const (
	LevelNone     ErrorLevel = "NONE"
	LevelLow      ErrorLevel = "LOW"
	LevelMedium   ErrorLevel = "MEDIUM"
	LevelHigh     ErrorLevel = "HIGH"
	LevelCritical ErrorLevel = "CRITICAL"
)

// File Log
const (
	logFolderName = "logs"
	logFileName   = "app.log"
)


type DetailLevel int

const (
    BasicLevel DetailLevel = iota
    StandardLevel
    VerboseLevel
)
type LoggerConfig struct {
    LogDir              string
    MaxSize             int64       
    TimeFormat          string
    RetentionDays       int64
    RotationFormat      string
    DetailLevel         DetailLevel 
}

var DefaultConfig    = LoggerConfig{
    LogDir           : "logs"                ,
    MaxSize          : 10                    , // MB
    TimeFormat       : "2006-01-02 15:04:05" ,
    RotationFormat   : "2006-01-02"          , 
    RetentionDays    : 7                     ,
    DetailLevel      : BasicLevel            ,
}

// CustomFormatter that formats logs
type CustomFormatter struct {
	IsTerminal bool
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Get the current timestamp
	timestamp := entry.Time.Format("2006-01-02 15:04:05,000")

	// Get the message
	message := entry.Message

	// Create the log message based on the format you want
	var logMessage string
	if f.IsTerminal {
		// For terminal, include colors
		logMessage = fmt.Sprintf("%s - %s - %s - %s\n", timestamp, entry.Level, entry.Level.String(), message)
	} else {
		// For file, strip colors and use clean format
		cleanMessage := stripColorCodes(message)
		logMessage = fmt.Sprintf("%s - %s - %s - %s\n", timestamp, entry.Level, entry.Level.String(), cleanMessage)
	}

	// Return the formatted log message as a byte slice
	return []byte(logMessage), nil
}

// Logger struct
type Logger struct {
	log        *logrus.Logger
	consoleLog *logrus.Logger
}

// NewLogger creates and configures the logger
func NewLogger() *Logger {
	l := logrus.New()

	// Ensure the log directory exists
	if err := os.MkdirAll(logFolderName, os.ModePerm); err != nil {
		panic(fmt.Sprintf("Could not create log directory: %v", err))
	}

	// Define the log file path
	logFilePath := fmt.Sprintf("%s/%s", logFolderName, logFileName)

	// Open log file in append mode
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		l.Fatalf("Failed to open log file: %v", err)
	}

	// Create a multi-writer to write to both file and console
	multiWriter := io.MultiWriter(file, os.Stdout)

	// Set up the logger output
	l.SetOutput(multiWriter)
	l.SetLevel(logrus.InfoLevel)

	// Use different formatters for file and console
	l.SetFormatter(&CustomFormatter{IsTerminal: false})  // File formatter (no colors)
	
	// Create a separate logger for console output with colors
	consoleLogger := logrus.New()
	consoleLogger.SetOutput(os.Stdout)
	consoleLogger.SetLevel(logrus.InfoLevel)
	consoleLogger.SetFormatter(&CustomFormatter{IsTerminal: true})  // Console formatter (with colors)

	return &Logger{
		log: l,
		consoleLog: consoleLogger,
	}
}

// Utility function to check if the output is a terminal
func isTerminal(f *os.File) bool {
	stat, err := f.Stat()
	if err != nil {
		return false
	}
	// Check if it's a terminal (usually a TTY device)
	return (stat.Mode() & os.ModeCharDevice) != 0
}

// Info prints a green info message
func (l *Logger) Info(args ...interface{}) {
	msg := fmt.Sprint(args...)
	l.log.Infof("%s", msg)  // File log (no colors)
	l.consoleLog.Infof("%s%s%s", ColorGreen, msg, ColorReset)  // Console log (with colors)
}

// Warn prints a yellow warning
func (l *Logger) Warn(args ...interface{}) {
	msg := fmt.Sprint(args...)
	l.log.Warnf("%s", msg)  // File log (no colors)
	l.consoleLog.Warnf("%s%s%s", ColorYellow, msg, ColorReset)  // Console log (with colors)
}

// Track prints a green message (custom track)
func (l *Logger) Track(args ...interface{}) {
	msg := fmt.Sprint(args...)
	l.log.Infof("%s", msg)  // File log (no colors)
	l.consoleLog.Infof("%s%s%s", ColorGreen, msg, ColorReset)  // Console log (with colors)
}

// ErrorHandler handles errors with levels
type ErrorHandler struct {
	logger *Logger
}

// Error prints red error
func (eh *ErrorHandler) Error(args ...interface{}) string {
	msg := fmt.Sprint(args...)
	eh.logger.log.Errorf("%s", msg)  // File log (no colors)
	eh.logger.consoleLog.Errorf("%s%s%s", ColorRed, msg, ColorReset)  // Console log (with colors)
	return msg
}

func (eh *ErrorHandler) Low(args ...interface{}) {
	msg := eh.Error(args...)
	sendError(LevelLow, msg)
}

func (eh *ErrorHandler) Medium(args ...interface{}) {
	msg := eh.Error(args...)
	sendError(LevelMedium, msg)
}

func (eh *ErrorHandler) High(args ...interface{}) {
	msg := eh.Error(args...)
	sendError(LevelHigh, msg)
}

func (eh *ErrorHandler) Critical(args ...interface{}) {
	msg := eh.Error(args...)
	sendError(LevelCritical, msg)
}

// Expose handlers
func (l *Logger) Error() *ErrorHandler {
	return &ErrorHandler{logger: l}
}

// Dummy external error sender
func sendError(level ErrorLevel, msg string) {
	fmt.Printf("[SendError] Level: %s | Message: %s\n", level, msg)
}

// Helper function to strip ANSI color codes from strings
func stripColorCodes(input string) string {
	// This regex will match ANSI escape codes (including color codes)
	colorPattern := regexp.MustCompile("\033\\[[0-9;]*m")
	return colorPattern.ReplaceAllString(input, "")
}