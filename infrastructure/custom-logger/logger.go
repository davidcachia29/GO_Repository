package custom_logger

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/gorm/logger"
)

type Logger interface {
	LogMode(level logger.LogLevel) logger.Interface
	Info(ctx context.Context, msg string, data ...interface{})
	Warn(ctx context.Context, msg string, data ...interface{})
	Error(ctx context.Context, msg string, data ...interface{})
	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
	Write(p []byte) (n int, err error)
}

type CustomLogger struct {
	consoleLogger *log.Logger
	fileLogger    *log.Logger
}

type LoggerProvider struct {
	consoleLogger *log.Logger
	fileLogger    *log.Logger
}

var loggerInstance *CustomLogger
var once sync.Once

func init() {
	// Initialize the loggerInstance when the package is loaded.
	once.Do(func() {
		loggerInstance = NewCustomLogger()
	})
}

func NewCustomLogger() *CustomLogger {
	// Open the log file
	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return &CustomLogger{
		consoleLogger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
		fileLogger:    log.New(file, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
	}
}

// Implement the Logger interface methods...

// GetInstance returns the global instance of the CustomLogger.
func GetInstance() *CustomLogger {
	once.Do(func() {
		loggerInstance = NewCustomLogger()
	})
	return loggerInstance
}

func (c *CustomLogger) LogMode(level logger.LogLevel) logger.Interface {
	// Implement the logic to change the log level if necessary
	// In this case, we simply return the same logger as we won't change the log level.
	return c
}

func (c *CustomLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	// Implement the logic to log the information message
	c.consoleLogger.Printf("[INFO] "+msg, data...)
	c.fileLogger.Printf("[INFO] "+msg, data...)
}

func (c *CustomLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	// Implement the logic to log the warning message
	c.consoleLogger.Printf("[WARN] "+msg, data...)
	c.fileLogger.Printf("[WARN] "+msg, data...)
}

func (c *CustomLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	// Implement the logic to log the error message
	c.consoleLogger.Printf("[ERROR] "+msg, data...)
	c.fileLogger.Printf("[ERROR] "+msg, data...)
}

func (c *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	// Implement the logic to log SQL traces
	sql, rows := fc()
	if err != nil {
		c.consoleLogger.Printf("[ERROR] %s\n[SQL] %s\n[ROWS AFFECTED] %d\n[ERROR] %s", time.Since(begin), sql, rows, err)
		c.fileLogger.Printf("[ERROR] %s\n[SQL] %s\n[ROWS AFFECTED] %d\n[ERROR] %s", time.Since(begin), sql, rows, err)
	} else {
		c.consoleLogger.Printf("[INFO] %s\n[SQL] %s\n[ROWS AFFECTED] %d", time.Since(begin), sql, rows)
		c.fileLogger.Printf("[INFO] %s\n[SQL] %s\n[ROWS AFFECTED] %d", time.Since(begin), sql, rows)
	}
}

func (c *CustomLogger) Write(p []byte) (n int, err error) {
	// Implement the logic to write to the log file
	c.consoleLogger.Print(string(p))
	c.fileLogger.Print(string(p))
	return len(p), nil
}
