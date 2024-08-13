package main

import (
	"fmt"
	"sync"
)

type Logger interface {
	Log(msg string)
}

type Logger1 struct{}

func (l *Logger1) Log(msg string) {
	fmt.Println("This is a logger1 structure Log method.")
}

type Logger2 struct{}

func (l *Logger2) Log(msg string) {
	fmt.Println("This is a logger2 structure Log method.")
}

type CompanyLogger struct{}

func (c *CompanyLogger) createLogger(msg string) Logger {
	fmt.Println("Here we are using factory design pattern.")
	if msg == "logger1" {
		return &Logger1{}
	} else {
		return &Logger2{}
	}
}

type SingletonLogger struct {
	logger Logger
}

var instance *SingletonLogger
var once sync.Once

func GetInstance() *SingletonLogger {
	once.Do(func() {
		instance = &SingletonLogger{}
	})
	return instance
}

func (sl *SingletonLogger) SetLogger(logger Logger) {
	sl.logger = logger
}

func (sl *SingletonLogger) Log(message string) {
	if sl.logger != nil {
		sl.logger.Log(message)
	} else {
		fmt.Println("No logger set.")
	}
}

func main() {
	obj := CompanyLogger{}
	LoggingObject := obj.createLogger("logger1")
	singleton := GetInstance()
	singleton.SetLogger(LoggingObject)
	singleton.Log("This is a message for the LoggingObject.")

}
